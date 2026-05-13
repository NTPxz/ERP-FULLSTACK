package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"erp-backend/internal/model"
	"erp-backend/internal/service"
)

type PurchaseHandler struct {
	svc service.PurchaseService
}

func NewPurchaseHandler(svc service.PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{svc: svc}
}

// GetAllSuppliers godoc
// @Summary      List suppliers
// @Tags         purchase
// @Produce      json
// @Success      200  {object}  map[string][]model.Supplier
// @Router       /suppliers [get]
func (h *PurchaseHandler) GetAllSuppliers(c *fiber.Ctx) error {
	suppliers, err := h.svc.GetAllSuppliers()
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, suppliers)
}

// GetSupplierByID godoc
// @Summary      Get supplier
// @Tags         purchase
// @Produce      json
// @Param        id   path      int  true  "Supplier ID"
// @Success      200  {object}  map[string]model.Supplier
// @Failure      404  {object}  map[string]string
// @Router       /suppliers/{id} [get]
func (h *PurchaseHandler) GetSupplierByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	sup, err := h.svc.GetSupplierByID(uint(id))
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, sup)
}

// CreateSupplier godoc
// @Summary      Create supplier
// @Tags         purchase
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      model.CreateSupplierRequest  true  "Supplier data"
// @Success      201      {object}  map[string]model.Supplier
// @Failure      400      {object}  map[string]string
// @Router       /suppliers [post]
func (h *PurchaseHandler) CreateSupplier(c *fiber.Ctx) error {
	var req model.CreateSupplierRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	sup, err := h.svc.CreateSupplier(req)
	if err != nil {
		return respondError(c, err)
	}
	return respondCreated(c, sup)
}

// UpdateSupplier godoc
// @Summary      Update supplier
// @Tags         purchase
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id       path      int                          true  "Supplier ID"
// @Param        request  body      model.UpdateSupplierRequest  true  "Supplier data"
// @Success      200      {object}  map[string]model.Supplier
// @Failure      400      {object}  map[string]string
// @Router       /suppliers/{id} [put]
func (h *PurchaseHandler) UpdateSupplier(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var req model.UpdateSupplierRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	sup, err := h.svc.UpdateSupplier(uint(id), req)
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, sup)
}

// DeleteSupplier godoc
// @Summary      Delete supplier
// @Tags         purchase
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Supplier ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /suppliers/{id} [delete]
func (h *PurchaseHandler) DeleteSupplier(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.svc.DeleteSupplier(uint(id)); err != nil {
		return respondError(c, err)
	}
	return c.JSON(fiber.Map{"message": "deleted successfully"})
}

// GetAllOrders godoc
// @Summary      List purchase orders
// @Tags         purchase
// @Produce      json
// @Success      200  {object}  map[string][]model.PurchaseOrder
// @Router       /purchase-orders [get]
func (h *PurchaseHandler) GetAllOrders(c *fiber.Ctx) error {
	orders, err := h.svc.GetAllOrders()
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, orders)
}

// GetOrderByID godoc
// @Summary      Get purchase order
// @Tags         purchase
// @Produce      json
// @Param        id   path      int  true  "PO ID"
// @Success      200  {object}  map[string]model.PurchaseOrder
// @Failure      404  {object}  map[string]string
// @Router       /purchase-orders/{id} [get]
func (h *PurchaseHandler) GetOrderByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	o, err := h.svc.GetOrderByID(uint(id))
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, o)
}

// CreateOrder godoc
// @Summary      Create purchase order
// @Description  Creates a draft PO. Use /receive to receive goods and update stock.
// @Tags         purchase
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      model.CreatePurchaseOrderRequest  true  "PO data"
// @Success      201      {object}  map[string]model.PurchaseOrder
// @Failure      400      {object}  map[string]string
// @Router       /purchase-orders [post]
func (h *PurchaseHandler) CreateOrder(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var req model.CreatePurchaseOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	o, err := h.svc.CreateOrder(req, userID)
	if err != nil {
		return respondError(c, err)
	}
	return respondCreated(c, o)
}

// UpdateOrderStatus godoc
// @Summary      Update PO status
// @Description  Valid values: sent, cancelled (use /receive endpoint to mark as received)
// @Tags         purchase
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id       path      int                                    true  "PO ID"
// @Param        request  body      model.UpdatePurchaseOrderStatusRequest  true  "New status"
// @Success      200      {object}  map[string]model.PurchaseOrder
// @Failure      400      {object}  map[string]string
// @Router       /purchase-orders/{id}/status [patch]
func (h *PurchaseHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var req model.UpdatePurchaseOrderStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	o, err := h.svc.UpdateOrderStatus(uint(id), req.Status)
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, o)
}

// ReceiveOrder godoc
// @Summary      Receive goods from PO
// @Description  Marks PO as received and automatically increases stock for all items
// @Tags         purchase
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "PO ID"
// @Success      200  {object}  map[string]model.PurchaseOrder
// @Failure      400  {object}  map[string]string
// @Router       /purchase-orders/{id}/receive [post]
func (h *PurchaseHandler) ReceiveOrder(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	o, err := h.svc.ReceiveOrder(uint(id), userID)
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, o)
}

// DeleteOrder godoc
// @Summary      Delete purchase order
// @Description  Only draft orders can be deleted
// @Tags         purchase
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "PO ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /purchase-orders/{id} [delete]
func (h *PurchaseHandler) DeleteOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.svc.DeleteOrder(uint(id)); err != nil {
		return respondError(c, err)
	}
	return c.JSON(fiber.Map{"message": "deleted successfully"})
}
