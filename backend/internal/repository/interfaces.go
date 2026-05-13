package repository

import "erp-backend/internal/model"

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindByID(id uint) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id uint) error
}

type EmployeeRepository interface {
	FindAllDepartments() ([]model.Department, error)
	FindDepartmentByID(id uint) (*model.Department, error)
	CreateDepartment(d *model.Department) error
	UpdateDepartment(d *model.Department) error
	DeleteDepartment(id uint) error

	FindAllPositions() ([]model.Position, error)
	FindPositionByID(id uint) (*model.Position, error)
	CreatePosition(p *model.Position) error

	FindAllEmployees() ([]model.Employee, error)
	FindEmployeeByID(id uint) (*model.Employee, error)
	CreateEmployee(e *model.Employee) error
	UpdateEmployee(e *model.Employee) error
	DeleteEmployee(id uint) error
}

type InventoryRepository interface {
	FindAllCategories() ([]model.Category, error)
	FindCategoryByID(id uint) (*model.Category, error)
	CreateCategory(c *model.Category) error

	FindAllProducts() ([]model.Product, error)
	FindProductByID(id uint) (*model.Product, error)
	CreateProduct(p *model.Product) error
	UpdateProduct(p *model.Product) error
	DeleteProduct(id uint) error

	CreateStockMovement(m *model.StockMovement) error
	FindMovementsByProductID(productID uint) ([]model.StockMovement, error)
	UpdateProductStock(productID uint, delta float64) error
}

type SalesRepository interface {
	FindAllCustomers() ([]model.Customer, error)
	FindCustomerByID(id uint) (*model.Customer, error)
	CreateCustomer(c *model.Customer) error
	UpdateCustomer(c *model.Customer) error
	DeleteCustomer(id uint) error

	FindAllOrders() ([]model.SalesOrder, error)
	FindOrderByID(id uint) (*model.SalesOrder, error)
	CreateOrder(o *model.SalesOrder) error
	UpdateOrder(o *model.SalesOrder) error
	DeleteOrder(id uint) error
}

type PurchaseRepository interface {
	FindAllSuppliers() ([]model.Supplier, error)
	FindSupplierByID(id uint) (*model.Supplier, error)
	CreateSupplier(s *model.Supplier) error
	UpdateSupplier(s *model.Supplier) error
	DeleteSupplier(id uint) error

	FindAllOrders() ([]model.PurchaseOrder, error)
	FindOrderByID(id uint) (*model.PurchaseOrder, error)
	CreateOrder(o *model.PurchaseOrder) error
	UpdateOrder(o *model.PurchaseOrder) error
	DeleteOrder(id uint) error
}
