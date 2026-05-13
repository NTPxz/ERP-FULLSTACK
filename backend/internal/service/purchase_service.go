package service

import (
	"errors"
	"fmt"
	"time"

	"erp-backend/internal/model"
	"erp-backend/internal/repository"
)

type purchaseService struct {
	repo    repository.PurchaseRepository
	invRepo repository.InventoryRepository
}

func NewPurchaseService(repo repository.PurchaseRepository, invRepo repository.InventoryRepository) PurchaseService {
	return &purchaseService{repo: repo, invRepo: invRepo}
}

func (s *purchaseService) GetAllSuppliers() ([]model.Supplier, error) {
	return s.repo.FindAllSuppliers()
}

func (s *purchaseService) GetSupplierByID(id uint) (*model.Supplier, error) {
	sup, err := s.repo.FindSupplierByID(id)
	if err != nil {
		return nil, model.ErrNotFound
	}
	return sup, nil
}

func (s *purchaseService) CreateSupplier(req model.CreateSupplierRequest) (*model.Supplier, error) {
	sup := &model.Supplier{
		Code:         req.Code,
		Name:         req.Name,
		Email:        req.Email,
		Phone:        req.Phone,
		Address:      req.Address,
		TaxID:        req.TaxID,
		PaymentTerms: req.PaymentTerms,
	}
	return sup, s.repo.CreateSupplier(sup)
}

func (s *purchaseService) UpdateSupplier(id uint, req model.UpdateSupplierRequest) (*model.Supplier, error) {
	sup, err := s.repo.FindSupplierByID(id)
	if err != nil {
		return nil, model.ErrNotFound
	}
	if req.Name != "" {
		sup.Name = req.Name
	}
	if req.Email != "" {
		sup.Email = req.Email
	}
	if req.Phone != "" {
		sup.Phone = req.Phone
	}
	if req.Address != "" {
		sup.Address = req.Address
	}
	if req.TaxID != "" {
		sup.TaxID = req.TaxID
	}
	if req.PaymentTerms != "" {
		sup.PaymentTerms = req.PaymentTerms
	}
	return sup, s.repo.UpdateSupplier(sup)
}

func (s *purchaseService) DeleteSupplier(id uint) error {
	if _, err := s.repo.FindSupplierByID(id); err != nil {
		return model.ErrNotFound
	}
	return s.repo.DeleteSupplier(id)
}

func (s *purchaseService) GetAllOrders() ([]model.PurchaseOrder, error) {
	return s.repo.FindAllOrders()
}

func (s *purchaseService) GetOrderByID(id uint) (*model.PurchaseOrder, error) {
	o, err := s.repo.FindOrderByID(id)
	if err != nil {
		return nil, errors.New("purchase order not found")
	}
	return o, nil
}

func (s *purchaseService) CreateOrder(req model.CreatePurchaseOrderRequest, createdBy uint) (*model.PurchaseOrder, error) {
	orderDate := req.OrderDate
	if orderDate.IsZero() {
		orderDate = time.Now()
	}

	var total float64
	items := make([]model.PurchaseOrderItem, 0, len(req.Items))
	for _, item := range req.Items {
		subtotal := item.UnitPrice * item.Quantity
		total += subtotal
		items = append(items, model.PurchaseOrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
			Subtotal:  subtotal,
		})
	}

	order := &model.PurchaseOrder{
		PONo:         fmt.Sprintf("PO-%d", time.Now().UnixNano()),
		SupplierID:   req.SupplierID,
		Status:       model.PODraft,
		OrderDate:    orderDate,
		ExpectedDate: req.ExpectedDate,
		TotalAmount:  total,
		Note:         req.Note,
		CreatedBy:    createdBy,
		Items:        items,
	}
	return order, s.repo.CreateOrder(order)
}

func (s *purchaseService) UpdateOrderStatus(id uint, status model.PurchaseOrderStatus) (*model.PurchaseOrder, error) {
	order, err := s.repo.FindOrderByID(id)
	if err != nil {
		return nil, errors.New("purchase order not found")
	}
	if status == model.POReceived {
		return nil, errors.New("use POST /purchase-orders/:id/receive to receive goods")
	}
	order.Status = status
	return order, s.repo.UpdateOrder(order)
}

// ReceiveOrder รับสินค้าตาม PO → เพิ่ม stock อัตโนมัติ
func (s *purchaseService) ReceiveOrder(id uint, createdBy uint) (*model.PurchaseOrder, error) {
	order, err := s.repo.FindOrderByID(id)
	if err != nil {
		return nil, errors.New("purchase order not found")
	}
	if order.Status == model.POReceived {
		return nil, errors.New("order already received")
	}
	if order.Status == model.POCancelled {
		return nil, errors.New("cannot receive a cancelled order")
	}

	for _, item := range order.Items {
		mv := &model.StockMovement{
			ProductID:     item.ProductID,
			Type:          model.StockIn,
			Quantity:      item.Quantity,
			ReferenceType: "purchase_order",
			ReferenceID:   order.ID,
			Note:          fmt.Sprintf("PO: %s", order.PONo),
			CreatedBy:     createdBy,
		}
		if err := s.invRepo.CreateStockMovement(mv); err != nil {
			return nil, err
		}
		if err := s.invRepo.UpdateProductStock(item.ProductID, item.Quantity); err != nil {
			return nil, err
		}
	}

	order.Status = model.POReceived
	return order, s.repo.UpdateOrder(order)
}

func (s *purchaseService) DeleteOrder(id uint) error {
	o, err := s.repo.FindOrderByID(id)
	if err != nil {
		return errors.New("purchase order not found")
	}
	if o.Status != model.PODraft {
		return errors.New("only draft orders can be deleted")
	}
	return s.repo.DeleteOrder(id)
}
