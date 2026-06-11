package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return fiber.ErrBadRequest
	}
	Err := h.svc.DeleteUser(context.Background(), int32(id))
	if Err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(http.StatusOK)
}
