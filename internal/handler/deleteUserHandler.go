package handler

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		return c.Status(400).JSON(fiber.Map{"id": "invalid id"})
	}

	if err := h.svc.DeleteUser(c.Context(), int32(id)); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(http.StatusOK)
}
