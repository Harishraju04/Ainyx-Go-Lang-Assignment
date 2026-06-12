package handler

import (
	"strconv"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/service"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/validator"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	var updateUser UpdateUserRequest

	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil || id <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := validator.Validate.Struct(&updateUser); err != nil {
		return handleError(c, err)
	}

	if updateUser.Name == nil && updateUser.Dob == nil {
		return c.Status(400).JSON(fiber.Map{"error": "at least one field must be provided"})
	}

	res, err := h.svc.UpdateUser(c.Context(), &service.UpdateUserRequest{
		Id:   int32(id),
		Name: updateUser.Name,
		Dob:  updateUser.Dob,
	})

	if err != nil {
		return handleError(c, err)
	}

	updateUserResponse := &User{
		Id:   res.Id,
		Name: res.Name,
		Dob:  res.Dob,
	}

	return c.JSON(updateUserResponse)
}
