package main

import (
	"log"

	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/config"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/db"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/db/sqlc"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/handler"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/repository"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/routes"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/service"
	"github.com/Harishraju04/Ainyx-Go-Lang-Assignment/internal/validator"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()

	pool := db.InitDB(cfg.DBurl)
	defer pool.Close()

	validator.Init()

	app := fiber.New()

	queires := sqlc.New(pool)
	repo := repository.NewRepository(queires)

	svc := service.NewService(repo)

	handler := handler.NewHandler(svc)

	routes.SetUpRoutes(app, handler)

	log.Fatal(app.Listen(cfg.Port))

}
