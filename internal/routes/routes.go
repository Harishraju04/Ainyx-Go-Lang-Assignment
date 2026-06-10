package routes

import (
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App, h *handler.Handler) {

	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Server is healthy")
	})

	api.Post("/users", h.CreateUserHandler)
}
