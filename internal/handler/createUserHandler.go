package handler

import (
	"log"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/service"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateUserHandler(c *fiber.Ctx) error {
	var user CreateUserRequest

	if err := c.BodyParser(&user); err != nil {
		return err
	}
	log.Println("handler", user)
	res, err := h.svc.CreateUser(c.Context(), &service.CreateUserRequest{
		Name: user.Name,
		Dob:  user.Dob,
	})
	if err != nil {
		return err
	}

	return c.JSON(res)
}
