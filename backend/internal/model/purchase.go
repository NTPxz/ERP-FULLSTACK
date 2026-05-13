package model

import (
	"time"

	"gorm.io/gorm"
)

type Supplier struct {
	gorm.Model
	Code         string `json:"code"          gorm:"unique;not null"`
	Name         string `json:"name"          gorm:"not null"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	TaxID        string `json:"tax_id"`
	PaymentTerms string `json:"payment_terms"`
}

type PurchaseOrderStatus string

const (
	PODraft     PurchaseOrderStatus = "draft"
	POSent      PurchaseOrderStatus = "sent"
	POReceived  PurchaseOrderStatus = "received"
	POCancelled PurchaseOrderStatus = "cancelled"
)

type PurchaseOrder struct {
	gorm.Model
	PONo         string              `json:"po_no"         gorm:"unique;not null"`
	SupplierID   uint                `json:"supplier_id"   gorm:"not null"`
	Supplier     Supplier            `json:"supplier,omitempty"`
	Status       PurchaseOrderStatus `json:"status"        gorm:"default:draft"`
	OrderDate    time.Time           `json:"order_date"`
	ExpectedDate time.Time           `json:"expected_date"`
	TotalAmount  float64             `json:"total_amount"  gorm:"default:0"`
	Note         string              `json:"note"`
	CreatedBy    uint                `json:"created_by"`
	Items        []PurchaseOrderItem `json:"items,omitempty" gorm:"foreignKey:PurchaseOrderID"`
}

type PurchaseOrderItem struct {
	gorm.Model
	PurchaseOrderID uint    `json:"purchase_order_id" gorm:"not null"`
	ProductID       uint    `json:"product_id"        gorm:"not null"`
	Product         Product `json:"product,omitempty"`
	Quantity        float64 `json:"quantity"          gorm:"not null"`
	UnitPrice       float64 `json:"unit_price"        gorm:"not null"`
	ReceivedQty     float64 `json:"received_qty"      gorm:"default:0"`
	Subtotal        float64 `json:"subtotal"          gorm:"not null"`
}

type CreateSupplierRequest struct {
	Code         string `json:"code"          validate:"required"`
	Name         string `json:"name"          validate:"required"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	TaxID        string `json:"tax_id"`
	PaymentTerms string `json:"payment_terms"`
}

type UpdateSupplierRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	TaxID        string `json:"tax_id"`
	PaymentTerms string `json:"payment_terms"`
}

type PurchaseOrderItemRequest struct {
	ProductID uint    `json:"product_id" validate:"required"`
	Quantity  float64 `json:"quantity"   validate:"required"`
	UnitPrice float64 `json:"unit_price" validate:"required"`
}

type CreatePurchaseOrderRequest struct {
	SupplierID   uint                       `json:"supplier_id"   validate:"required"`
	OrderDate    time.Time                  `json:"order_date"`
	ExpectedDate time.Time                  `json:"expected_date"`
	Note         string                     `json:"note"`
	Items        []PurchaseOrderItemRequest `json:"items"         validate:"required,min=1"`
}

type UpdatePurchaseOrderStatusRequest struct {
	Status PurchaseOrderStatus `json:"status" validate:"required"`
}
