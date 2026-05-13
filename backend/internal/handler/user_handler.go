package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"erp-backend/internal/model"
	"erp-backend/internal/service"
)

type UserHandler struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

// GetAll godoc
// @Summary      List all users
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  map[string][]model.User
// @Router       /users [get]
func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	users, err := h.svc.GetAll()
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, users)
}

// GetByID godoc
// @Summary      Get user by ID
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  map[string]model.User
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [get]
func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	user, err := h.svc.GetByID(uint(id))
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, user)
}

// Register godoc
// @Summary      Register new user
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      model.RegisterRequest  true  "Register payload"
// @Success      201      {object}  map[string]model.User
// @Failure      400      {object}  map[string]string
// @Router       /auth/register [post]
func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req model.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	user, err := h.svc.Register(req)
	if err != nil {
		return respondError(c, err)
	}
	return respondCreated(c, user)
}

// Login godoc
// @Summary      Login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request  body      model.LoginRequest  true  "Login payload"
// @Success      200      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Router       /auth/login [post]
func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	token, err := h.svc.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	return respondOK(c, fiber.Map{"token": token})
}

// Me godoc
// @Summary      Get current user
// @Tags         auth
// @Security     BearerAuth
// @Produce      json
// @Success      200  {object}  map[string]model.User
// @Failure      401  {object}  map[string]string
// @Router       /me [get]
func (h *UserHandler) Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	user, err := h.svc.GetByID(userID)
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, user)
}

// Update godoc
// @Summary      Update own profile
// @Tags         users
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id       path      int                      true  "User ID"
// @Param        request  body      model.UpdateUserRequest  true  "Update payload"
// @Success      200      {object}  map[string]model.User
// @Failure      400      {object}  map[string]string
// @Router       /users/{id} [put]
func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var req model.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	user, err := h.svc.Update(uint(id), req)
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, user)
}

// Delete godoc
// @Summary      Delete user
// @Tags         users
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [delete]
func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.svc.Delete(uint(id)); err != nil {
		return respondError(c, err)
	}
	return c.JSON(fiber.Map{"message": "deleted successfully"})
}

// AdminCreate godoc
// @Summary      Create user with role (admin only)
// @Tags         users
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request  body      model.AdminCreateUserRequest  true  "User data with role"
// @Success      201      {object}  map[string]model.User
// @Failure      400      {object}  map[string]string
// @Failure      403      {object}  map[string]string
// @Router       /users [post]
func (h *UserHandler) AdminCreate(c *fiber.Ctx) error {
	var req model.AdminCreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	user, err := h.svc.AdminCreate(req)
	if err != nil {
		return respondError(c, err)
	}
	return respondCreated(c, user)
}

// AdminUpdate godoc
// @Summary      Update user role/info (admin only)
// @Tags         users
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id       path      int                           true  "User ID"
// @Param        request  body      model.AdminUpdateUserRequest  true  "Update payload"
// @Success      200      {object}  map[string]model.User
// @Failure      400      {object}  map[string]string
// @Failure      403      {object}  map[string]string
// @Router       /users/{id}/role [patch]
func (h *UserHandler) AdminUpdate(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var req model.AdminUpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}
	user, err := h.svc.AdminUpdate(uint(id), req)
	if err != nil {
		return respondError(c, err)
	}
	return respondOK(c, user)
}
