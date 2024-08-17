package routes

import (
	"github.com/gnotnek/fiber-task-list-backend/internal/handlers"
	"github.com/gnotnek/fiber-task-list-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	//default route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Todo list API")
	})

	// Todos
	todos := app.Group("/todos")
	todos.Get("/", middleware.AuthRequired, handlers.GetTodos)
	todos.Get("/:id", middleware.AuthRequired, handlers.GetTodoByID)
	todos.Post("/", middleware.AuthRequired, handlers.CreateTodo)
	todos.Delete("/:id", middleware.AuthRequired, handlers.DeleteTodo)
	todos.Put("/:id", middleware.AuthRequired, handlers.UpdateTodoByID)
	todos.Put("/complete", middleware.AuthRequired, handlers.CompleteAllTodos)

	// Register
	app.Post("/register", handlers.SignUp)

	// Login
	app.Post("/login", handlers.Login)

	app.Use(func(c *fiber.Ctx) error { return c.SendStatus(404) })
}
