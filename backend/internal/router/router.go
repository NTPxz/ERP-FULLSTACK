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

	// Middleware shortcuts
	jwt     := middleware.JWTProtected()
	admin   := middleware.RequireRole("admin")
	hr      := middleware.RequireRole("admin", "hr")
	inv     := middleware.RequireRole("admin", "inventory")
	prodR   := middleware.RequireRole("admin", "inventory", "sales", "purchase")
	sales   := middleware.RequireRole("admin", "sales")
	pur     := middleware.RequireRole("admin", "purchase")

	// ── Any authenticated user ──────────────────────────────────────────────
	api.Get("/me", jwt, userH.Me)

	// ── Admin only ──────────────────────────────────────────────────────────
	api.Get("/users",          jwt, admin, userH.GetAll)
	api.Get("/users/:id",      jwt, admin, userH.GetByID)
	api.Post("/users",         jwt, admin, userH.AdminCreate)
	api.Patch("/users/:id/role", jwt, admin, userH.AdminUpdate)
	api.Delete("/users/:id",   jwt, admin, userH.Delete)

	// ── HR (admin | hr) ────────────────────────────────────────────────────
	api.Get("/departments",        jwt, hr, empH.GetAllDepartments)
	api.Get("/departments/:id",    jwt, hr, empH.GetDepartmentByID)
	api.Post("/departments",       jwt, hr, empH.CreateDepartment)
	api.Put("/departments/:id",    jwt, hr, empH.UpdateDepartment)
	api.Delete("/departments/:id", jwt, hr, empH.DeleteDepartment)
	api.Get("/positions",          jwt, hr, empH.GetAllPositions)
	api.Post("/positions",         jwt, hr, empH.CreatePosition)
	api.Get("/employees",          jwt, hr, empH.GetAllEmployees)
	api.Get("/employees/:id",      jwt, hr, empH.GetEmployeeByID)
	api.Post("/employees",         jwt, hr, empH.CreateEmployee)
	api.Put("/employees/:id",      jwt, hr, empH.UpdateEmployee)
	api.Delete("/employees/:id",   jwt, hr, empH.DeleteEmployee)

	// ── Inventory (admin | inventory) ──────────────────────────────────────
	api.Get("/categories",         jwt, inv, invH.GetAllCategories)
	api.Post("/categories",        jwt, inv, invH.CreateCategory)
	api.Post("/products",          jwt, inv, invH.CreateProduct)
	api.Put("/products/:id",       jwt, inv, invH.UpdateProduct)
	api.Delete("/products/:id",    jwt, inv, invH.DeleteProduct)
	api.Post("/stock/adjust",      jwt, inv, invH.AdjustStock)

	// Products readable by inventory + sales + purchase
	api.Get("/products",                     jwt, prodR, invH.GetAllProducts)
	api.Get("/products/:id",                 jwt, prodR, invH.GetProductByID)
	api.Get("/products/:id/movements",       jwt, prodR, invH.GetStockMovements)

	// ── Sales (admin | sales) ──────────────────────────────────────────────
	api.Get("/customers",              jwt, sales, salesH.GetAllCustomers)
	api.Get("/customers/:id",          jwt, sales, salesH.GetCustomerByID)
	api.Post("/customers",             jwt, sales, salesH.CreateCustomer)
	api.Put("/customers/:id",          jwt, sales, salesH.UpdateCustomer)
	api.Delete("/customers/:id",       jwt, sales, salesH.DeleteCustomer)
	api.Get("/sales-orders",           jwt, sales, salesH.GetAllOrders)
	api.Get("/sales-orders/:id",       jwt, sales, salesH.GetOrderByID)
	api.Post("/sales-orders",          jwt, sales, salesH.CreateOrder)
	api.Patch("/sales-orders/:id/status", jwt, sales, salesH.UpdateOrderStatus)
	api.Delete("/sales-orders/:id",    jwt, sales, salesH.DeleteOrder)

	// ── Purchase (admin | purchase) ────────────────────────────────────────
	api.Get("/suppliers",                     jwt, pur, purchaseH.GetAllSuppliers)
	api.Get("/suppliers/:id",                 jwt, pur, purchaseH.GetSupplierByID)
	api.Post("/suppliers",                    jwt, pur, purchaseH.CreateSupplier)
	api.Put("/suppliers/:id",                 jwt, pur, purchaseH.UpdateSupplier)
	api.Delete("/suppliers/:id",              jwt, pur, purchaseH.DeleteSupplier)
	api.Get("/purchase-orders",               jwt, pur, purchaseH.GetAllOrders)
	api.Get("/purchase-orders/:id",           jwt, pur, purchaseH.GetOrderByID)
	api.Post("/purchase-orders",              jwt, pur, purchaseH.CreateOrder)
	api.Patch("/purchase-orders/:id/status",  jwt, pur, purchaseH.UpdateOrderStatus)
	api.Post("/purchase-orders/:id/receive",  jwt, pur, purchaseH.ReceiveOrder)
	api.Delete("/purchase-orders/:id",        jwt, pur, purchaseH.DeleteOrder)
}
