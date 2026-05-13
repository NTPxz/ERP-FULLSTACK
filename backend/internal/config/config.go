package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"erp-backend/internal/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file, reading from OS environment")
	}
}

func ConnectDB() *gorm.DB {
	var db *gorm.DB
	var err error

	driver := os.Getenv("DB_DRIVER")

	switch driver {
	case "mysql":
		log.Fatal("mysql driver: uncomment mysql import in config.go")
	default:
		dbFile := os.Getenv("DB_FILE")
		if dbFile == "" {
			dbFile = "app.db"
		}
		db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
	}

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(
		&model.User{},
		&model.Department{},
		&model.Position{},
		&model.Employee{},
		&model.Category{},
		&model.Product{},
		&model.StockMovement{},
		&model.Customer{},
		&model.SalesOrder{},
		&model.SalesOrderItem{},
		&model.Supplier{},
		&model.PurchaseOrder{},
		&model.PurchaseOrderItem{},
	)

	log.Println("database connected")
	return db
}

// SeedDefaultUsers creates one user per role if not present, and ensures
// the admin account always has the admin role.
func SeedDefaultUsers(db *gorm.DB) {
	defaults := []struct {
		name, email, password string
		role                  model.UserRole
	}{
		{"Administrator", "admin@erp.com", "admin1234", model.RoleAdmin},
		{"HR Manager", "hr@erp.com", "pass1234", model.RoleHR},
		{"Inventory Manager", "inventory@erp.com", "pass1234", model.RoleInventory},
		{"Sales Representative", "sales@erp.com", "pass1234", model.RoleSales},
		{"Purchasing Manager", "purchase@erp.com", "pass1234", model.RolePurchase},
	}

	for _, u := range defaults {
		var existing model.User
		if db.Where("email = ?", u.email).First(&existing).Error == nil {
			// Ensure role is correct for existing accounts
			db.Model(&existing).Update("role", string(u.role))
			continue
		}
		hashed, err := bcrypt.GenerateFromPassword([]byte(u.password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("failed to hash password for %s: %v", u.email, err)
			continue
		}
		db.Create(&model.User{
			Name:     u.name,
			Email:    u.email,
			Password: string(hashed),
			Role:     u.role,
		})
	}
	log.Println("default users seeded")
}
