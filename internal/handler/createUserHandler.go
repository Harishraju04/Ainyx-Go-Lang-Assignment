package handler

import (
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/service"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/validator"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateUserHandler(c *fiber.Ctx) error {
	var user CreateUserRequest

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON("invalid body")
	}

	validationErr := validator.Validate.Struct(&user)
	if validationErr != nil {
		errors := validator.FormatValidationErrors(validationErr)
		return c.Status(400).JSON(errors)
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
