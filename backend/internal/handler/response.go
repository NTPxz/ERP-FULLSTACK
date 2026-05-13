package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"erp-backend/internal/model"
)

func respondError(c *fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, model.ErrNotFound):
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	case errors.Is(err, model.ErrBadRequest):
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	case errors.Is(err, model.ErrConflict):
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	case errors.Is(err, model.ErrForbidden):
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": err.Error()})
	default:
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}
}

func respondOK(c *fiber.Ctx, data any) error {
	return c.JSON(fiber.Map{"data": data})
}

func respondCreated(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": data})
}
