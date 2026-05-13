package router

import (
	"github.com/gofiber/fiber/v2"
	"erp-backend/internal/handler"
	"erp-backend/internal/middleware"
)

func SetupRoutes(
	app       *fiber.App,
	userH     *handler.UserHandler,
	empH      *handler.EmployeeHandler,
	invH      *handler.InventoryHandler,
	salesH    *handler.SalesHandler,
	purchaseH *handler.PurchaseHandler,
) {
	api := app.Group("/api")

	// ── Public ─────────────────────────────────────────────────────────────
	api.Post("/auth/register", userH.Register)
	api.Post("/auth/login", userH.Login)

	// Base protected group — any valid JWT
	p := api.Group("/", middleware.JWTProtected())
	p.Get("/me", userH.Me)

	// ── Admin only ─────────────────────────────────────────────────────────
	adm := p.Group("/", middleware.RequireRole("admin"))
	adm.Get("/users", userH.GetAll)
	adm.Get("/users/:id", userH.GetByID)
	adm.Post("/users", userH.AdminCreate)
	adm.Patch("/users/:id/role", userH.AdminUpdate)
	adm.Delete("/users/:id", userH.Delete)

	// ── HR (admin | hr) ────────────────────────────────────────────────────
	hr := p.Group("/", middleware.RequireRole("admin", "hr"))
	hr.Get("/departments", empH.GetAllDepartments)
	hr.Get("/departments/:id", empH.GetDepartmentByID)
	hr.Post("/departments", empH.CreateDepartment)
	hr.Put("/departments/:id", empH.UpdateDepartment)
	hr.Delete("/departments/:id", empH.DeleteDepartment)
	hr.Get("/positions", empH.GetAllPositions)
	hr.Post("/positions", empH.CreatePosition)
	hr.Get("/employees", empH.GetAllEmployees)
	hr.Get("/employees/:id", empH.GetEmployeeByID)
	hr.Post("/employees", empH.CreateEmployee)
	hr.Put("/employees/:id", empH.UpdateEmployee)
	hr.Delete("/employees/:id", empH.DeleteEmployee)

	// ── Inventory (admin | inventory) ─────────────────────────────────────
	inv := p.Group("/", middleware.RequireRole("admin", "inventory"))
	inv.Get("/categories", invH.GetAllCategories)
	inv.Post("/categories", invH.CreateCategory)
	inv.Post("/products", invH.CreateProduct)
	inv.Put("/products/:id", invH.UpdateProduct)
	inv.Delete("/products/:id", invH.DeleteProduct)
	inv.Post("/stock/adjust", invH.AdjustStock)

	// Products readable by inventory + sales + purchase (for order creation dropdowns)
	prod := p.Group("/", middleware.RequireRole("admin", "inventory", "sales", "purchase"))
	prod.Get("/products", invH.GetAllProducts)
	prod.Get("/products/:id", invH.GetProductByID)
	prod.Get("/products/:id/movements", invH.GetStockMovements)

	// ── Sales (admin | sales) ──────────────────────────────────────────────
	sales := p.Group("/", middleware.RequireRole("admin", "sales"))
	sales.Get("/customers", salesH.GetAllCustomers)
	sales.Get("/customers/:id", salesH.GetCustomerByID)
	sales.Post("/customers", salesH.CreateCustomer)
	sales.Put("/customers/:id", salesH.UpdateCustomer)
	sales.Delete("/customers/:id", salesH.DeleteCustomer)
	sales.Get("/sales-orders", salesH.GetAllOrders)
	sales.Get("/sales-orders/:id", salesH.GetOrderByID)
	sales.Post("/sales-orders", salesH.CreateOrder)
	sales.Patch("/sales-orders/:id/status", salesH.UpdateOrderStatus)
	sales.Delete("/sales-orders/:id", salesH.DeleteOrder)

	// ── Purchase (admin | purchase) ────────────────────────────────────────
	pur := p.Group("/", middleware.RequireRole("admin", "purchase"))
	pur.Get("/suppliers", purchaseH.GetAllSuppliers)
	pur.Get("/suppliers/:id", purchaseH.GetSupplierByID)
	pur.Post("/suppliers", purchaseH.CreateSupplier)
	pur.Put("/suppliers/:id", purchaseH.UpdateSupplier)
	pur.Delete("/suppliers/:id", purchaseH.DeleteSupplier)
	pur.Get("/purchase-orders", purchaseH.GetAllOrders)
	pur.Get("/purchase-orders/:id", purchaseH.GetOrderByID)
	pur.Post("/purchase-orders", purchaseH.CreateOrder)
	pur.Patch("/purchase-orders/:id/status", purchaseH.UpdateOrderStatus)
	pur.Post("/purchase-orders/:id/receive", purchaseH.ReceiveOrder)
	pur.Delete("/purchase-orders/:id", purchaseH.DeleteOrder)
}
