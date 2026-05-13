// @title           Go ERP API
// @version         2.0
// @description     ERP System — HR, Inventory, Sales, Purchase with JWT Authentication
// @host            localhost:8080
// @BasePath        /api

// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 Enter: Bearer <your-token>

package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fiberSwagger "github.com/gofiber/swagger"
	"erp-backend/internal/config"
	_ "erp-backend/docs"
	"erp-backend/internal/handler"
	"erp-backend/internal/repository"
	"erp-backend/internal/router"
	"erp-backend/internal/service"
)

func main() {
	config.LoadEnv()
	db := config.ConnectDB()
	config.SeedDefaultUsers(db)

	// Repositories
	userRepo     := repository.NewUserRepository(db)
	empRepo      := repository.NewEmployeeRepository(db)
	invRepo      := repository.NewInventoryRepository(db)
	salesRepo    := repository.NewSalesRepository(db)
	purchaseRepo := repository.NewPurchaseRepository(db)

	// Services (SalesService และ PurchaseService ใช้ invRepo ร่วมกันเพื่อจัดการ stock)
	userSvc     := service.NewUserService(userRepo)
	empSvc      := service.NewEmployeeService(empRepo)
	invSvc      := service.NewInventoryService(invRepo)
	salesSvc    := service.NewSalesService(salesRepo, invRepo)
	purchaseSvc := service.NewPurchaseService(purchaseRepo, invRepo)

	// Handlers
	userH     := handler.NewUserHandler(userSvc)
	empH      := handler.NewEmployeeHandler(empSvc)
	invH      := handler.NewInventoryHandler(invSvc)
	salesH    := handler.NewSalesHandler(salesSvc)
	purchaseH := handler.NewPurchaseHandler(purchaseSvc)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		},
	})

	app.Use(logger.New())
	app.Use(cors.New())

	app.Get("/swagger/*", fiberSwagger.HandlerDefault)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	router.SetupRoutes(app, userH, empH, invH, salesH, purchaseH)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("server  : http://localhost:%s", port)
	log.Printf("swagger : http://localhost:%s/swagger/index.html", port)
	log.Fatal(app.Listen(":" + port))
}
