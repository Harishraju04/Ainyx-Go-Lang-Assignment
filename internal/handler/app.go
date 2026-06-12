package handler

import (
	"errors"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/service"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/validator"
	validatorpkg "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

type Handler struct {
	svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func handleError(c *fiber.Ctx, err error) error {
	if err == nil {
		return nil
	}

	var validationErr validatorpkg.ValidationErrors
	if errors.As(err, &validationErr) {
		return c.Status(400).JSON(fiber.Map{
			"error":   "validation failed",
			"details": validator.FormatValidationErrors(validationErr),
		})
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}

	if errors.Is(err, fiber.ErrBadRequest) {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}

	if err.Error() == "at least one field must be provided" || err.Error() == "invalid date format" {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(500).JSON(fiber.Map{"error": "internal server error"})
}
