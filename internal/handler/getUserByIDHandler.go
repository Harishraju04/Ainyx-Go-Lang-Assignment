package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) GetUserByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)

	if err != nil {
		return fiber.ErrBadRequest
	}

	res, err := h.svc.GetUserByID(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	getUserByIDResponse := &User{
		Id:   res.Id,
		Name: res.Name,
		Dob:  res.Dob,
		Age:  res.Age,
	}

	return c.JSON(getUserByIDResponse)
}
