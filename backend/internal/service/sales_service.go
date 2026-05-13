package service

import (
	"errors"
	"fmt"
	"time"

	"erp-backend/internal/model"
	"erp-backend/internal/repository"
)

type salesService struct {
	repo    repository.SalesRepository
	invRepo repository.InventoryRepository
}

func NewSalesService(repo repository.SalesRepository, invRepo repository.InventoryRepository) SalesService {
	return &salesService{repo: repo, invRepo: invRepo}
}

func (s *salesService) GetAllCustomers() ([]model.Customer, error) {
	return s.repo.FindAllCustomers()
}

func (s *salesService) GetCustomerByID(id uint) (*model.Customer, error) {
	c, err := s.repo.FindCustomerByID(id)
	if err != nil {
		return nil, model.ErrNotFound
	}
	return c, nil
}

func (s *salesService) CreateCustomer(req model.CreateCustomerRequest) (*model.Customer, error) {
	c := &model.Customer{
		Code:        req.Code,
		Name:        req.Name,
		Email:       req.Email,
		Phone:       req.Phone,
		Address:     req.Address,
		TaxID:       req.TaxID,
		CreditLimit: req.CreditLimit,
	}
	return c, s.repo.CreateCustomer(c)
}

func (s *salesService) UpdateCustomer(id uint, req model.UpdateCustomerRequest) (*model.Customer, error) {
	c, err := s.repo.FindCustomerByID(id)
	if err != nil {
		return nil, model.ErrNotFound
	}
	if req.Name != "" {
		c.Name = req.Name
	}
	if req.Email != "" {
		c.Email = req.Email
	}
	if req.Phone != "" {
		c.Phone = req.Phone
	}
	if req.Address != "" {
		c.Address = req.Address
	}
	if req.TaxID != "" {
		c.TaxID = req.TaxID
	}
	if req.CreditLimit != 0 {
		c.CreditLimit = req.CreditLimit
	}
	return c, s.repo.UpdateCustomer(c)
}

func (s *salesService) DeleteCustomer(id uint) error {
	if _, err := s.repo.FindCustomerByID(id); err != nil {
		return model.ErrNotFound
	}
	return s.repo.DeleteCustomer(id)
}

func (s *salesService) GetAllOrders() ([]model.SalesOrder, error) {
	return s.repo.FindAllOrders()
}

func (s *salesService) GetOrderByID(id uint) (*model.SalesOrder, error) {
	o, err := s.repo.FindOrderByID(id)
	if err != nil {
		return nil, errors.New("sales order not found")
	}
	return o, nil
}

func (s *salesService) CreateOrder(req model.CreateSalesOrderRequest, createdBy uint) (*model.SalesOrder, error) {
	orderDate := req.OrderDate
	if orderDate.IsZero() {
		orderDate = time.Now()
	}

	var total float64
	items := make([]model.SalesOrderItem, 0, len(req.Items))
	for _, item := range req.Items {
		subtotal := (item.UnitPrice * item.Quantity) - item.Discount
		total += subtotal
		items = append(items, model.SalesOrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
			Discount:  item.Discount,
			Subtotal:  subtotal,
		})
	}

	order := &model.SalesOrder{
		OrderNo:     fmt.Sprintf("SO-%d", time.Now().UnixNano()),
		CustomerID:  req.CustomerID,
		Status:      model.SODraft,
		OrderDate:   orderDate,
		TotalAmount: total,
		Note:        req.Note,
		CreatedBy:   createdBy,
		Items:       items,
	}
	return order, s.repo.CreateOrder(order)
}

func (s *salesService) UpdateOrderStatus(id uint, status model.SalesOrderStatus) (*model.SalesOrder, error) {
	order, err := s.repo.FindOrderByID(id)
	if err != nil {
		return nil, errors.New("sales order not found")
	}

	// ตรวจและตัด stock เมื่อ confirm
	if status == model.SOConfirmed && order.Status == model.SODraft {
		for _, item := range order.Items {
			p, err := s.invRepo.FindProductByID(item.ProductID)
			if err != nil {
				return nil, fmt.Errorf("product id %d not found", item.ProductID)
			}
			if p.StockQuantity < item.Quantity {
				return nil, fmt.Errorf("insufficient stock for %s (available: %.2f)", p.Name, p.StockQuantity)
			}
		}
		for _, item := range order.Items {
			mv := &model.StockMovement{
				ProductID:     item.ProductID,
				Type:          model.StockOut,
				Quantity:      item.Quantity,
				ReferenceType: "sales_order",
				ReferenceID:   order.ID,
				Note:          fmt.Sprintf("SO: %s", order.OrderNo),
				CreatedBy:     order.CreatedBy,
			}
			if err := s.invRepo.CreateStockMovement(mv); err != nil {
				return nil, err
			}
			if err := s.invRepo.UpdateProductStock(item.ProductID, -item.Quantity); err != nil {
				return nil, err
			}
		}
	}

	order.Status = status
	return order, s.repo.UpdateOrder(order)
}

func (s *salesService) DeleteOrder(id uint) error {
	o, err := s.repo.FindOrderByID(id)
	if err != nil {
		return errors.New("sales order not found")
	}
	if o.Status != model.SODraft {
		return errors.New("only draft orders can be deleted")
	}
	return s.repo.DeleteOrder(id)
}
