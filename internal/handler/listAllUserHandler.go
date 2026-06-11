package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) ListAllUsers(c *fiber.Ctx) error {
	res, err := h.svc.ListAllUsers(c.Context())
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(http.StatusOK).JSON(res)
}
