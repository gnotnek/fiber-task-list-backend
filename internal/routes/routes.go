package routes

import (
	"github.com/gnotnek/fiber-task-list-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/todos", handlers.GetTodos)
	app.Get("/todos/:id", handlers.GetTodoByID)
	app.Post("/todos", handlers.CreateTodo)
	app.Delete("/todos/:id", handlers.DeleteTodo)
	app.Put("/todos", handlers.UpdateTodo)
	app.Put("/todos/:id", handlers.UpdateTodoByID)
	app.Put("/todos/complete", handlers.CompleteAllTodos)
}
