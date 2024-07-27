package main

import (
	"log"

	"github.com/gnotnek/fiber-task-list-backend/internal/database"
	"github.com/gnotnek/fiber-task-list-backend/internal/handlers"
	"github.com/gnotnek/fiber-task-list-backend/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	database.InitDB()

	handlers.NewTodoHandler(&database.DB{})
	routes.SetupRoutes(app, handlers.NewTodoHandler(&database.DB{}))

	log.Fatal(app.Listen(":3500"))
}
