package main

import (
	"log"

	"github.com/gnotnek/fiber-task-list-backend/internal/database"
	"github.com/gnotnek/fiber-task-list-backend/internal/routes"
	"github.com/gofiber/fiber/v2"

	_ "github.com/gnotnek/fiber-task-list-backend/cmd/docs"
	"github.com/gofiber/swagger"
)

// @title Fiber Task List API
// @version 1.0
// @description This is a simple task list API
// @host localhost:3500
// @BasePath /
// @schemes http
// @produce json
// @consumes json
// @router /todos [get]
// @router /todos/{id} [get]
// @router /todos [post]
// @router /todos/{id} [delete]
// @router /todos [put]
// @router /todos/{id} [put]
// @router /todos/complete [put]
// @router /users [get]
// @router /users/{id} [get]
// @router /users [post]
func main() {
	database.InitDB()

	app := fiber.New()
	// Register swagger handler
	app.Get("/swagger/*", swagger.HandlerDefault)

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3500"))
}
