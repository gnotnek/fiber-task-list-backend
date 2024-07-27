package routes

import (
	"github.com/gnotnek/fiber-task-list-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h *handlers.TodoHandler) {
	app.Get("/todos", h.GetTodos)
	app.Post("/todos", h.CreateTodo)
	app.Put("/todos", h.UpdateTodo)
	app.Delete("/todos/:id", h.DeleteTodo)
	app.Put("/todos/:id", h.UpdateTodoByID)
	app.Put("/todos/complete", h.CompleteAllTodos)
}
