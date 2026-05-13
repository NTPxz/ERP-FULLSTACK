package repository

import (
	"erp-backend/internal/model"
	"gorm.io/gorm"
)

type salesRepository struct {
	db *gorm.DB
}

func NewSalesRepository(db *gorm.DB) SalesRepository {
	return &salesRepository{db: db}
}

func (r *salesRepository) FindAllCustomers() ([]model.Customer, error) {
	var customers []model.Customer
	return customers, r.db.Find(&customers).Error
}

func (r *salesRepository) FindCustomerByID(id uint) (*model.Customer, error) {
	var c model.Customer
	if err := r.db.First(&c, id).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *salesRepository) CreateCustomer(c *model.Customer) error {
	return r.db.Create(c).Error
}

func (r *salesRepository) UpdateCustomer(c *model.Customer) error {
	return r.db.Save(c).Error
}

func (r *salesRepository) DeleteCustomer(id uint) error {
	return r.db.Delete(&model.Customer{}, id).Error
}

func (r *salesRepository) FindAllOrders() ([]model.SalesOrder, error) {
	var orders []model.SalesOrder
	return orders, r.db.Preload("Customer").Preload("Items.Product").Find(&orders).Error
}

func (r *salesRepository) FindOrderByID(id uint) (*model.SalesOrder, error) {
	var o model.SalesOrder
	if err := r.db.Preload("Customer").Preload("Items.Product").First(&o, id).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *salesRepository) CreateOrder(o *model.SalesOrder) error {
	return r.db.Create(o).Error
}

func (r *salesRepository) UpdateOrder(o *model.SalesOrder) error {
	return r.db.Save(o).Error
}

func (r *salesRepository) DeleteOrder(id uint) error {
	return r.db.Delete(&model.SalesOrder{}, id).Error
}
