package main

import (
	"log"

	"github.com/gnotnek/fiber-task-list-backend/internal/database"
	"github.com/gnotnek/fiber-task-list-backend/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.InitDB()

	app := fiber.New()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3500"))
}
