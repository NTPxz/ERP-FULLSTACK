package service

import "erp-backend/internal/model"

type UserService interface {
	GetAll() ([]model.User, error)
	GetByID(id uint) (*model.User, error)
	Register(req model.RegisterRequest) (*model.User, error)
	Login(req model.LoginRequest) (string, error)
	Update(id uint, req model.UpdateUserRequest) (*model.User, error)
	Delete(id uint) error
	AdminCreate(req model.AdminCreateUserRequest) (*model.User, error)
	AdminUpdate(id uint, req model.AdminUpdateUserRequest) (*model.User, error)
}

type EmployeeService interface {
	GetAllDepartments() ([]model.Department, error)
	GetDepartmentByID(id uint) (*model.Department, error)
	CreateDepartment(req model.CreateDepartmentRequest) (*model.Department, error)
	UpdateDepartment(id uint, req model.CreateDepartmentRequest) (*model.Department, error)
	DeleteDepartment(id uint) error

	GetAllPositions() ([]model.Position, error)
	CreatePosition(req model.CreatePositionRequest) (*model.Position, error)

	GetAllEmployees() ([]model.Employee, error)
	GetEmployeeByID(id uint) (*model.Employee, error)
	CreateEmployee(req model.CreateEmployeeRequest) (*model.Employee, error)
	UpdateEmployee(id uint, req model.UpdateEmployeeRequest) (*model.Employee, error)
	DeleteEmployee(id uint) error
}

type InventoryService interface {
	GetAllCategories() ([]model.Category, error)
	CreateCategory(req model.CreateCategoryRequest) (*model.Category, error)

	GetAllProducts() ([]model.Product, error)
	GetProductByID(id uint) (*model.Product, error)
	CreateProduct(req model.CreateProductRequest) (*model.Product, error)
	UpdateProduct(id uint, req model.UpdateProductRequest) (*model.Product, error)
	DeleteProduct(id uint) error

	AdjustStock(req model.StockAdjustRequest, createdBy uint) error
	GetStockMovements(productID uint) ([]model.StockMovement, error)
}

type SalesService interface {
	GetAllCustomers() ([]model.Customer, error)
	GetCustomerByID(id uint) (*model.Customer, error)
	CreateCustomer(req model.CreateCustomerRequest) (*model.Customer, error)
	UpdateCustomer(id uint, req model.UpdateCustomerRequest) (*model.Customer, error)
	DeleteCustomer(id uint) error

	GetAllOrders() ([]model.SalesOrder, error)
	GetOrderByID(id uint) (*model.SalesOrder, error)
	CreateOrder(req model.CreateSalesOrderRequest, createdBy uint) (*model.SalesOrder, error)
	UpdateOrderStatus(id uint, status model.SalesOrderStatus) (*model.SalesOrder, error)
	DeleteOrder(id uint) error
}

type PurchaseService interface {
	GetAllSuppliers() ([]model.Supplier, error)
	GetSupplierByID(id uint) (*model.Supplier, error)
	CreateSupplier(req model.CreateSupplierRequest) (*model.Supplier, error)
	UpdateSupplier(id uint, req model.UpdateSupplierRequest) (*model.Supplier, error)
	DeleteSupplier(id uint) error

	GetAllOrders() ([]model.PurchaseOrder, error)
	GetOrderByID(id uint) (*model.PurchaseOrder, error)
	CreateOrder(req model.CreatePurchaseOrderRequest, createdBy uint) (*model.PurchaseOrder, error)
	UpdateOrderStatus(id uint, status model.PurchaseOrderStatus) (*model.PurchaseOrder, error)
	ReceiveOrder(id uint, createdBy uint) (*model.PurchaseOrder, error)
	DeleteOrder(id uint) error
}
