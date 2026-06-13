package routes

import (
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/handler"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App, h *handler.Handler) {
	// Global middleware
	app.Use(middleware.RequestIDMiddleware)
	app.Use(middleware.RequestDurationMiddleware)

	api := app.Group("/api")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Server is healthy")
	})

	api.Post("/users", h.CreateUserHandler)
	api.Get("/users/:id", h.GetUserByID)
	api.Put("/users/:id", h.UpdateUser)
	api.Delete("/users/:id", h.DeleteUser)
	api.Get("/users", h.ListAllUsers)
}
