package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"erp-backend/internal/model"
	"erp-backend/internal/service"
)

type EmployeeHandler struct {
	svc service.EmployeeService
}

func NewEmployeeHandler(svc service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{svc: svc}
}

// --- Department ---

// GetAllDepartments godoc
// @Summary      List departments
// @Tags         hr
// @Produce      json
// @Success      200  {object}  map[string][]model.Department
// @Router       /departments [get]
func (h *EmployeeHandler) GetAllDepartments(c *fiber.Ctx) error {
	deps, err := h.svc.GetAllDepartments()
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, deps)
}

// GetDepartmentByID godoc
// @Summary      Get department
// @Tags         hr
// @Produce      json
// @Param        id   path      int  true  "Department ID"
// @Success      200  {object}  map[string]model.Department
// @Failure      404  {object}  map[string]string
// @Router       /departments/{id} [get]
func (h *EmployeeHandler) GetDepartmentByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	dep, err := h.svc.GetDepartmentByID(uint(id))
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, dep)
}

// CreateDepartment godoc
// @Summary      Create department
// @Tags         hr
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      model.CreateDepartmentRequest  true  "Department data"
// @Success      201      {object}  map[string]model.Department
// @Failure      400      {object}  map[string]string
// @Router       /departments [post]
func (h *EmployeeHandler) CreateDepartment(c *fiber.Ctx) error {
	var req model.CreateDepartmentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	dep, err := h.svc.CreateDepartment(req)
	if err != nil {
		return respondError(c, err)
	}
	return respondCreated(c, dep)
}

// UpdateDepartment godoc
// @Summary      Update department
// @Tags         hr
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id       path      int                            true  "Department ID"
// @Param        request  body      model.CreateDepartmentRequest  true  "Department data"
// @Success      200      {object}  map[string]model.Department
// @Failure      400      {object}  map[string]string
// @Router       /departments/{id} [put]
func (h *EmployeeHandler) UpdateDepartment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var req model.CreateDepartmentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	dep, err := h.svc.UpdateDepartment(uint(id), req)
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, dep)
}

// DeleteDepartment godoc
// @Summary      Delete department
// @Tags         hr
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Department ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /departments/{id} [delete]
func (h *EmployeeHandler) DeleteDepartment(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.svc.DeleteDepartment(uint(id)); err != nil {
		return respondError(c, err)
	}
	return c.JSON(fiber.Map{"message": "deleted successfully"})
}

// --- Position ---

// GetAllPositions godoc
// @Summary      List positions
// @Tags         hr
// @Produce      json
// @Success      200  {object}  map[string][]model.Position
// @Router       /positions [get]
func (h *EmployeeHandler) GetAllPositions(c *fiber.Ctx) error {
	pos, err := h.svc.GetAllPositions()
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, pos)
}

// CreatePosition godoc
// @Summary      Create position
// @Tags         hr
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      model.CreatePositionRequest  true  "Position data"
// @Success      201      {object}  map[string]model.Position
// @Failure      400      {object}  map[string]string
// @Router       /positions [post]
func (h *EmployeeHandler) CreatePosition(c *fiber.Ctx) error {
	var req model.CreatePositionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	pos, err := h.svc.CreatePosition(req)
	if err != nil {
		return respondError(c, err)
	}
	return respondCreated(c, pos)
}

// --- Employee ---

// GetAllEmployees godoc
// @Summary      List employees
// @Tags         hr
// @Produce      json
// @Success      200  {object}  map[string][]model.Employee
// @Router       /employees [get]
func (h *EmployeeHandler) GetAllEmployees(c *fiber.Ctx) error {
	emps, err := h.svc.GetAllEmployees()
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, emps)
}

// GetEmployeeByID godoc
// @Summary      Get employee
// @Tags         hr
// @Produce      json
// @Param        id   path      int  true  "Employee ID"
// @Success      200  {object}  map[string]model.Employee
// @Failure      404  {object}  map[string]string
// @Router       /employees/{id} [get]
func (h *EmployeeHandler) GetEmployeeByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	emp, err := h.svc.GetEmployeeByID(uint(id))
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, emp)
}

// CreateEmployee godoc
// @Summary      Create employee
// @Tags         hr
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      model.CreateEmployeeRequest  true  "Employee data"
// @Success      201      {object}  map[string]model.Employee
// @Failure      400      {object}  map[string]string
// @Router       /employees [post]
func (h *EmployeeHandler) CreateEmployee(c *fiber.Ctx) error {
	var req model.CreateEmployeeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	emp, err := h.svc.CreateEmployee(req)
	if err != nil {
		return respondError(c, err)
	}
	return respondCreated(c, emp)
}

// UpdateEmployee godoc
// @Summary      Update employee
// @Tags         hr
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id       path      int                          true  "Employee ID"
// @Param        request  body      model.UpdateEmployeeRequest  true  "Employee data"
// @Success      200      {object}  map[string]model.Employee
// @Failure      400      {object}  map[string]string
// @Router       /employees/{id} [put]
func (h *EmployeeHandler) UpdateEmployee(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var req model.UpdateEmployeeRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	emp, err := h.svc.UpdateEmployee(uint(id), req)
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, emp)
}

// DeleteEmployee godoc
// @Summary      Delete employee
// @Tags         hr
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Employee ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /employees/{id} [delete]
func (h *EmployeeHandler) DeleteEmployee(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.svc.DeleteEmployee(uint(id)); err != nil {
		return respondError(c, err)
	}
	return c.JSON(fiber.Map{"message": "deleted successfully"})
}
