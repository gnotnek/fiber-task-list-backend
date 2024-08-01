package routes

import (
	"github.com/gnotnek/fiber-task-list-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Todos
	todos := app.Group("/todos")
	todos.Get("/", handlers.GetTodos)
	todos.Get("/:id", handlers.GetTodoByID)
	todos.Post("/", handlers.CreateTodo)
	todos.Delete("/:id", handlers.DeleteTodo)
	todos.Put("/:id", handlers.UpdateTodoByID)
	todos.Put("/complete", handlers.CompleteAllTodos)

	// Users
	users := app.Group("/users")
	users.Get("/", handlers.GetUsers)
	users.Get("/:id", handlers.GetUserByID)
	users.Post("/", handlers.CreateUser)

	app.Use(func(c *fiber.Ctx) error { return c.SendStatus(404) })
}
