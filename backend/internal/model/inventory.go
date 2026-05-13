package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string    `json:"name"               gorm:"not null;unique"`
	Description string    `json:"description"`
	Products    []Product `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
}

type Product struct {
	gorm.Model
	SKU           string   `json:"sku"            gorm:"unique;not null"`
	Name          string   `json:"name"           gorm:"not null"`
	Description   string   `json:"description"`
	CategoryID    uint     `json:"category_id"`
	Category      Category `json:"category,omitempty"`
	Unit          string   `json:"unit"`
	CostPrice     float64  `json:"cost_price"`
	SalePrice     float64  `json:"sale_price"`
	StockQuantity float64  `json:"stock_quantity" gorm:"default:0"`
	ReorderLevel  float64  `json:"reorder_level"  gorm:"default:0"`
}

type StockMovementType string

const (
	StockIn     StockMovementType = "in"
	StockOut    StockMovementType = "out"
	StockAdjust StockMovementType = "adjust"
)

type StockMovement struct {
	gorm.Model
	ProductID     uint              `json:"product_id"`
	Product       Product           `json:"product,omitempty"`
	Type          StockMovementType `json:"type"           gorm:"not null"`
	Quantity      float64           `json:"quantity"       gorm:"not null"`
	ReferenceType string            `json:"reference_type"`
	ReferenceID   uint              `json:"reference_id"`
	Note          string            `json:"note"`
	CreatedBy     uint              `json:"created_by"`
}

type CreateCategoryRequest struct {
	Name        string `json:"name"        validate:"required"`
	Description string `json:"description"`
}

type CreateProductRequest struct {
	SKU          string  `json:"sku"          validate:"required"`
	Name         string  `json:"name"         validate:"required"`
	Description  string  `json:"description"`
	CategoryID   uint    `json:"category_id"  validate:"required"`
	Unit         string  `json:"unit"`
	CostPrice    float64 `json:"cost_price"`
	SalePrice    float64 `json:"sale_price"   validate:"required"`
	ReorderLevel float64 `json:"reorder_level"`
}

type UpdateProductRequest struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	CategoryID   uint    `json:"category_id"`
	Unit         string  `json:"unit"`
	CostPrice    float64 `json:"cost_price"`
	SalePrice    float64 `json:"sale_price"`
	ReorderLevel float64 `json:"reorder_level"`
}

type StockAdjustRequest struct {
	ProductID uint    `json:"product_id" validate:"required"`
	Quantity  float64 `json:"quantity"   validate:"required"`
	Note      string  `json:"note"`
}
