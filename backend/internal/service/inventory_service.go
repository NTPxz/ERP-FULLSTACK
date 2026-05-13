package service

import (
	"erp-backend/internal/model"
	"erp-backend/internal/repository"
)

type inventoryService struct {
	repo repository.InventoryRepository
}

func NewInventoryService(repo repository.InventoryRepository) InventoryService {
	return &inventoryService{repo: repo}
}

func (s *inventoryService) GetAllCategories() ([]model.Category, error) {
	return s.repo.FindAllCategories()
}

func (s *inventoryService) CreateCategory(req model.CreateCategoryRequest) (*model.Category, error) {
	cat := &model.Category{Name: req.Name, Description: req.Description}
	return cat, s.repo.CreateCategory(cat)
}

func (s *inventoryService) GetAllProducts() ([]model.Product, error) {
	return s.repo.FindAllProducts()
}

func (s *inventoryService) GetProductByID(id uint) (*model.Product, error) {
	p, err := s.repo.FindProductByID(id)
	if err != nil {
		return nil, model.ErrNotFound
	}
	return p, nil
}

func (s *inventoryService) CreateProduct(req model.CreateProductRequest) (*model.Product, error) {
	p := &model.Product{
		SKU:          req.SKU,
		Name:         req.Name,
		Description:  req.Description,
		CategoryID:   req.CategoryID,
		Unit:         req.Unit,
		CostPrice:    req.CostPrice,
		SalePrice:    req.SalePrice,
		ReorderLevel: req.ReorderLevel,
	}
	return p, s.repo.CreateProduct(p)
}

func (s *inventoryService) UpdateProduct(id uint, req model.UpdateProductRequest) (*model.Product, error) {
	p, err := s.repo.FindProductByID(id)
	if err != nil {
		return nil, model.ErrNotFound
	}
	if req.Name != "" {
		p.Name = req.Name
	}
	if req.Description != "" {
		p.Description = req.Description
	}
	if req.CategoryID != 0 {
		p.CategoryID = req.CategoryID
	}
	if req.Unit != "" {
		p.Unit = req.Unit
	}
	if req.CostPrice != 0 {
		p.CostPrice = req.CostPrice
	}
	if req.SalePrice != 0 {
		p.SalePrice = req.SalePrice
	}
	if req.ReorderLevel != 0 {
		p.ReorderLevel = req.ReorderLevel
	}
	return p, s.repo.UpdateProduct(p)
}

func (s *inventoryService) DeleteProduct(id uint) error {
	if _, err := s.repo.FindProductByID(id); err != nil {
		return model.ErrNotFound
	}
	return s.repo.DeleteProduct(id)
}

func (s *inventoryService) AdjustStock(req model.StockAdjustRequest, createdBy uint) error {
	if _, err := s.repo.FindProductByID(req.ProductID); err != nil {
		return model.ErrNotFound
	}
	movement := &model.StockMovement{
		ProductID:     req.ProductID,
		Type:          model.StockAdjust,
		Quantity:      req.Quantity,
		ReferenceType: "manual",
		Note:          req.Note,
		CreatedBy:     createdBy,
	}
	if err := s.repo.CreateStockMovement(movement); err != nil {
		return err
	}
	return s.repo.UpdateProductStock(req.ProductID, req.Quantity)
}

func (s *inventoryService) GetStockMovements(productID uint) ([]model.StockMovement, error) {
	return s.repo.FindMovementsByProductID(productID)
}
