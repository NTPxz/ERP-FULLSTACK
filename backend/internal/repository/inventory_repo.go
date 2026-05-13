package repository

import (
	"erp-backend/internal/model"
	"gorm.io/gorm"
)

type inventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) FindAllCategories() ([]model.Category, error) {
	var cats []model.Category
	return cats, r.db.Find(&cats).Error
}

func (r *inventoryRepository) FindCategoryByID(id uint) (*model.Category, error) {
	var cat model.Category
	if err := r.db.First(&cat, id).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *inventoryRepository) CreateCategory(c *model.Category) error {
	return r.db.Create(c).Error
}

func (r *inventoryRepository) FindAllProducts() ([]model.Product, error) {
	var products []model.Product
	return products, r.db.Preload("Category").Find(&products).Error
}

func (r *inventoryRepository) FindProductByID(id uint) (*model.Product, error) {
	var product model.Product
	if err := r.db.Preload("Category").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *inventoryRepository) CreateProduct(p *model.Product) error {
	return r.db.Create(p).Error
}

func (r *inventoryRepository) UpdateProduct(p *model.Product) error {
	return r.db.Save(p).Error
}

func (r *inventoryRepository) DeleteProduct(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}

func (r *inventoryRepository) CreateStockMovement(m *model.StockMovement) error {
	return r.db.Create(m).Error
}

func (r *inventoryRepository) FindMovementsByProductID(productID uint) ([]model.StockMovement, error) {
	var movements []model.StockMovement
	return movements, r.db.Where("product_id = ?", productID).Order("created_at desc").Find(&movements).Error
}

func (r *inventoryRepository) UpdateProductStock(productID uint, delta float64) error {
	return r.db.Model(&model.Product{}).Where("id = ?", productID).
		Update("stock_quantity", gorm.Expr("stock_quantity + ?", delta)).Error
}
