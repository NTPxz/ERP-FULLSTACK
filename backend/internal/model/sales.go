package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Code        string  `json:"code"         gorm:"unique;not null"`
	Name        string  `json:"name"         gorm:"not null"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Address     string  `json:"address"`
	TaxID       string  `json:"tax_id"`
	CreditLimit float64 `json:"credit_limit" gorm:"default:0"`
}

type SalesOrderStatus string

const (
	SODraft     SalesOrderStatus = "draft"
	SOConfirmed SalesOrderStatus = "confirmed"
	SOCompleted SalesOrderStatus = "completed"
	SOCancelled SalesOrderStatus = "cancelled"
)

type SalesOrder struct {
	gorm.Model
	OrderNo     string           `json:"order_no"     gorm:"unique;not null"`
	CustomerID  uint             `json:"customer_id"  gorm:"not null"`
	Customer    Customer         `json:"customer,omitempty"`
	Status      SalesOrderStatus `json:"status"       gorm:"default:draft"`
	OrderDate   time.Time        `json:"order_date"`
	TotalAmount float64          `json:"total_amount" gorm:"default:0"`
	Note        string           `json:"note"`
	CreatedBy   uint             `json:"created_by"`
	Items       []SalesOrderItem `json:"items,omitempty" gorm:"foreignKey:SalesOrderID"`
}

type SalesOrderItem struct {
	gorm.Model
	SalesOrderID uint    `json:"sales_order_id" gorm:"not null"`
	ProductID    uint    `json:"product_id"     gorm:"not null"`
	Product      Product `json:"product,omitempty"`
	Quantity     float64 `json:"quantity"       gorm:"not null"`
	UnitPrice    float64 `json:"unit_price"     gorm:"not null"`
	Discount     float64 `json:"discount"       gorm:"default:0"`
	Subtotal     float64 `json:"subtotal"       gorm:"not null"`
}

type CreateCustomerRequest struct {
	Code        string  `json:"code"         validate:"required"`
	Name        string  `json:"name"         validate:"required"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Address     string  `json:"address"`
	TaxID       string  `json:"tax_id"`
	CreditLimit float64 `json:"credit_limit"`
}

type UpdateCustomerRequest struct {
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Address     string  `json:"address"`
	TaxID       string  `json:"tax_id"`
	CreditLimit float64 `json:"credit_limit"`
}

type SalesOrderItemRequest struct {
	ProductID uint    `json:"product_id" validate:"required"`
	Quantity  float64 `json:"quantity"   validate:"required"`
	UnitPrice float64 `json:"unit_price" validate:"required"`
	Discount  float64 `json:"discount"`
}

type CreateSalesOrderRequest struct {
	CustomerID uint                    `json:"customer_id" validate:"required"`
	OrderDate  time.Time               `json:"order_date"`
	Note       string                  `json:"note"`
	Items      []SalesOrderItemRequest `json:"items"       validate:"required,min=1"`
}

type UpdateSalesOrderStatusRequest struct {
	Status SalesOrderStatus `json:"status" validate:"required"`
}
