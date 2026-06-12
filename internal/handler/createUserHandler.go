package handler

import (
	"net/http"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/service"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/validator"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateUserHandler(c *fiber.Ctx) error {
	var user CreateUserRequest

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := validator.Validate.Struct(&user); err != nil {
		return handleError(c, err)
	}

	res, err := h.svc.CreateUser(c.Context(), &service.CreateUserRequest{
		Name: user.Name,
		Dob:  user.Dob,
	})
	if err != nil {
		return handleError(c, err)
	}

	createUserResponse := &User{
		Id:   res.Id,
		Name: res.Name,
		Dob:  res.Dob,
	}

	return c.Status(http.StatusCreated).JSON(createUserResponse)
}
