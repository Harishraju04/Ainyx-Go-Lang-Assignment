package handler

import (
	"log"
	"strconv"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	var updateUser UpdateUserRequest

	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 32)

	if err := c.BodyParser(&updateUser); err != nil {
		log.Println(updateUser)
		return fiber.ErrBadRequest
	}

	res, err := h.svc.UpdateUser(c.Context(), &service.UpdateUserRequest{
		Id:   int32(id),
		Name: updateUser.Name,
		Dob:  updateUser.Dob,
	})

	if err != nil {
		return fiber.ErrInternalServerError
	}

	updateUserResponse := &User{
		Id:   res.Id,
		Name: res.Name,
		Dob:  res.Dob,
	}

	return c.JSON(updateUserResponse)
}
