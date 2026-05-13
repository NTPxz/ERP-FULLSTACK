package repository

import (
	"erp-backend/internal/model"
	"gorm.io/gorm"
)

type purchaseRepository struct {
	db *gorm.DB
}

func NewPurchaseRepository(db *gorm.DB) PurchaseRepository {
	return &purchaseRepository{db: db}
}

func (r *purchaseRepository) FindAllSuppliers() ([]model.Supplier, error) {
	var suppliers []model.Supplier
	return suppliers, r.db.Find(&suppliers).Error
}

func (r *purchaseRepository) FindSupplierByID(id uint) (*model.Supplier, error) {
	var s model.Supplier
	if err := r.db.First(&s, id).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *purchaseRepository) CreateSupplier(s *model.Supplier) error {
	return r.db.Create(s).Error
}

func (r *purchaseRepository) UpdateSupplier(s *model.Supplier) error {
	return r.db.Save(s).Error
}

func (r *purchaseRepository) DeleteSupplier(id uint) error {
	return r.db.Delete(&model.Supplier{}, id).Error
}

func (r *purchaseRepository) FindAllOrders() ([]model.PurchaseOrder, error) {
	var orders []model.PurchaseOrder
	return orders, r.db.Preload("Supplier").Preload("Items.Product").Find(&orders).Error
}

func (r *purchaseRepository) FindOrderByID(id uint) (*model.PurchaseOrder, error) {
	var o model.PurchaseOrder
	if err := r.db.Preload("Supplier").Preload("Items.Product").First(&o, id).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *purchaseRepository) CreateOrder(o *model.PurchaseOrder) error {
	return r.db.Create(o).Error
}

func (r *purchaseRepository) UpdateOrder(o *model.PurchaseOrder) error {
	return r.db.Save(o).Error
}

func (r *purchaseRepository) DeleteOrder(id uint) error {
	return r.db.Delete(&model.PurchaseOrder{}, id).Error
}
