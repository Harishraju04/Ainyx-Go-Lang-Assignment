package handler

import (
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateUserHandler(c *fiber.Ctx) error {
	var user CreateUserRequest

	if err := c.BodyParser(&user); err != nil {
		return err
	}
	res, err := h.svc.CreateUser(c.Context(), &service.CreateUserRequest{
		Name: user.Name,
		Dob:  user.Dob,
	})
	if err != nil {
		return err
	}

	createUserResponse := &User{
		Id:   res.Id,
		Name: res.Name,
		Dob:  res.Dob,
	}

	return c.JSON(createUserResponse)
}
