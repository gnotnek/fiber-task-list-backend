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

	// Register
	app.Post("/register", handlers.SignUp)

	// Login
	app.Post("/login", handlers.Login)

	app.Use(func(c *fiber.Ctx) error { return c.SendStatus(404) })
}
