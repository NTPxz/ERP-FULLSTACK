package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"erp-backend/internal/model"
	"erp-backend/internal/service"
)

type InventoryHandler struct {
	svc service.InventoryService
}

func NewInventoryHandler(svc service.InventoryService) *InventoryHandler {
	return &InventoryHandler{svc: svc}
}

// GetAllCategories godoc
// @Summary      List categories
// @Tags         inventory
// @Produce      json
// @Success      200  {object}  map[string][]model.Category
// @Router       /categories [get]
func (h *InventoryHandler) GetAllCategories(c *fiber.Ctx) error {
	cats, err := h.svc.GetAllCategories()
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, cats)
}

// CreateCategory godoc
// @Summary      Create category
// @Tags         inventory
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      model.CreateCategoryRequest  true  "Category data"
// @Success      201      {object}  map[string]model.Category
// @Failure      400      {object}  map[string]string
// @Router       /categories [post]
func (h *InventoryHandler) CreateCategory(c *fiber.Ctx) error {
	var req model.CreateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	cat, err := h.svc.CreateCategory(req)
	if err != nil {
		return respondError(c, err)
	}
	return respondCreated(c, cat)
}

// GetAllProducts godoc
// @Summary      List products
// @Tags         inventory
// @Produce      json
// @Success      200  {object}  map[string][]model.Product
// @Router       /products [get]
func (h *InventoryHandler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.svc.GetAllProducts()
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, products)
}

// GetProductByID godoc
// @Summary      Get product
// @Tags         inventory
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  map[string]model.Product
// @Failure      404  {object}  map[string]string
// @Router       /products/{id} [get]
func (h *InventoryHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	p, err := h.svc.GetProductByID(uint(id))
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, p)
}

// CreateProduct godoc
// @Summary      Create product
// @Tags         inventory
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      model.CreateProductRequest  true  "Product data"
// @Success      201      {object}  map[string]model.Product
// @Failure      400      {object}  map[string]string
// @Router       /products [post]
func (h *InventoryHandler) CreateProduct(c *fiber.Ctx) error {
	var req model.CreateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	p, err := h.svc.CreateProduct(req)
	if err != nil {
		return respondError(c, err)
	}
	return respondCreated(c, p)
}

// UpdateProduct godoc
// @Summary      Update product
// @Tags         inventory
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id       path      int                         true  "Product ID"
// @Param        request  body      model.UpdateProductRequest  true  "Product data"
// @Success      200      {object}  map[string]model.Product
// @Failure      400      {object}  map[string]string
// @Router       /products/{id} [put]
func (h *InventoryHandler) UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var req model.UpdateProductRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	p, err := h.svc.UpdateProduct(uint(id), req)
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, p)
}

// DeleteProduct godoc
// @Summary      Delete product
// @Tags         inventory
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /products/{id} [delete]
func (h *InventoryHandler) DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.svc.DeleteProduct(uint(id)); err != nil {
		return respondError(c, err)
	}
	return c.JSON(fiber.Map{"message": "deleted successfully"})
}

// AdjustStock godoc
// @Summary      Manual stock adjustment
// @Description  Adjust stock quantity with an audit trail entry
// @Tags         inventory
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      model.StockAdjustRequest  true  "Adjustment data (positive=add, negative=remove)"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  map[string]string
// @Router       /stock/adjust [post]
func (h *InventoryHandler) AdjustStock(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var req model.StockAdjustRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	if err := h.svc.AdjustStock(req, userID); err != nil {
		return respondError(c, err)
	}
	return c.JSON(fiber.Map{"message": "stock adjusted"})
}

// GetStockMovements godoc
// @Summary      Get stock movements for a product
// @Tags         inventory
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  map[string][]model.StockMovement
// @Router       /products/{id}/movements [get]
func (h *InventoryHandler) GetStockMovements(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	movements, err := h.svc.GetStockMovements(uint(id))
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, movements)
}
