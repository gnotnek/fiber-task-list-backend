package routes

import (
	"github.com/gnotnek/fiber-task-list-backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, h *handlers.TodoHandler) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Todos API, Build with Fiber 🚀")
	})
	app.Get("/todos", h.GetTodos)
	app.Get("/todos/:id", h.GetTodoByID)
	app.Post("/todos", h.CreateTodo)
	app.Delete("/todos/:id", h.DeleteTodo)
	app.Put("/todos", h.UpdateTodo)
	app.Put("/todos/:id", h.UpdateTodoByID)
	app.Put("/todos/complete", h.CompleteAllTodos)
}
