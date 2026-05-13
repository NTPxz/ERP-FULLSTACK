package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"erp-backend/internal/model"
	"erp-backend/internal/service"
)

type SalesHandler struct {
	svc service.SalesService
}

func NewSalesHandler(svc service.SalesService) *SalesHandler {
	return &SalesHandler{svc: svc}
}

// GetAllCustomers godoc
// @Summary      List customers
// @Tags         sales
// @Produce      json
// @Success      200  {object}  map[string][]model.Customer
// @Router       /customers [get]
func (h *SalesHandler) GetAllCustomers(c *fiber.Ctx) error {
	customers, err := h.svc.GetAllCustomers()
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, customers)
}

// GetCustomerByID godoc
// @Summary      Get customer
// @Tags         sales
// @Produce      json
// @Param        id   path      int  true  "Customer ID"
// @Success      200  {object}  map[string]model.Customer
// @Failure      404  {object}  map[string]string
// @Router       /customers/{id} [get]
func (h *SalesHandler) GetCustomerByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	cust, err := h.svc.GetCustomerByID(uint(id))
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, cust)
}

// CreateCustomer godoc
// @Summary      Create customer
// @Tags         sales
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      model.CreateCustomerRequest  true  "Customer data"
// @Success      201      {object}  map[string]model.Customer
// @Failure      400      {object}  map[string]string
// @Router       /customers [post]
func (h *SalesHandler) CreateCustomer(c *fiber.Ctx) error {
	var req model.CreateCustomerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	cust, err := h.svc.CreateCustomer(req)
	if err != nil {
		return respondError(c, err)
	}
	return respondCreated(c, cust)
}

// UpdateCustomer godoc
// @Summary      Update customer
// @Tags         sales
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id       path      int                          true  "Customer ID"
// @Param        request  body      model.UpdateCustomerRequest  true  "Customer data"
// @Success      200      {object}  map[string]model.Customer
// @Failure      400      {object}  map[string]string
// @Router       /customers/{id} [put]
func (h *SalesHandler) UpdateCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var req model.UpdateCustomerRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	cust, err := h.svc.UpdateCustomer(uint(id), req)
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, cust)
}

// DeleteCustomer godoc
// @Summary      Delete customer
// @Tags         sales
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Customer ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /customers/{id} [delete]
func (h *SalesHandler) DeleteCustomer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.svc.DeleteCustomer(uint(id)); err != nil {
		return respondError(c, err)
	}
	return c.JSON(fiber.Map{"message": "deleted successfully"})
}

// GetAllOrders godoc
// @Summary      List sales orders
// @Tags         sales
// @Produce      json
// @Success      200  {object}  map[string][]model.SalesOrder
// @Router       /sales-orders [get]
func (h *SalesHandler) GetAllOrders(c *fiber.Ctx) error {
	orders, err := h.svc.GetAllOrders()
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, orders)
}

// GetOrderByID godoc
// @Summary      Get sales order
// @Tags         sales
// @Produce      json
// @Param        id   path      int  true  "Order ID"
// @Success      200  {object}  map[string]model.SalesOrder
// @Failure      404  {object}  map[string]string
// @Router       /sales-orders/{id} [get]
func (h *SalesHandler) GetOrderByID(c *fiber.Ctx) error {
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
// @Summary      Create sales order
// @Description  Creates a draft sales order with line items. Confirm to deduct stock.
// @Tags         sales
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      model.CreateSalesOrderRequest  true  "Order data"
// @Success      201      {object}  map[string]model.SalesOrder
// @Failure      400      {object}  map[string]string
// @Router       /sales-orders [post]
func (h *SalesHandler) CreateOrder(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var req model.CreateSalesOrderRequest
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
// @Summary      Update sales order status
// @Description  Valid transitions: draft→confirmed (deducts stock), confirmed→completed, any→cancelled
// @Tags         sales
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id       path      int                                  true  "Order ID"
// @Param        request  body      model.UpdateSalesOrderStatusRequest  true  "New status"
// @Success      200      {object}  map[string]model.SalesOrder
// @Failure      400      {object}  map[string]string
// @Router       /sales-orders/{id}/status [patch]
func (h *SalesHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var req model.UpdateSalesOrderStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	o, err := h.svc.UpdateOrderStatus(uint(id), req.Status)
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, o)
}

// DeleteOrder godoc
// @Summary      Delete sales order
// @Description  Only draft orders can be deleted
// @Tags         sales
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Order ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Router       /sales-orders/{id} [delete]
func (h *SalesHandler) DeleteOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.svc.DeleteOrder(uint(id)); err != nil {
		return respondError(c, err)
	}
	return c.JSON(fiber.Map{"message": "deleted successfully"})
}
