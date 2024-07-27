package handlers

import (
	"github.com/gnotnek/fiber-task-list-backend/internal/database"
	"github.com/gnotnek/fiber-task-list-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	db *database.DB
}

func NewTodoHandler(db *database.DB) *TodoHandler {
	return &TodoHandler{db: db}
}

func (h *TodoHandler) GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	h.db.Find(&todos)

	return c.JSON(todos)
}

func (h *TodoHandler) GetTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	h.db.First(&todo, id)

	return c.JSON(todo)
}

func (h *TodoHandler) CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	h.db.Create(todo)

	return c.JSON(todo)
}

func (h *TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	h.db.Save(todo)

	return c.JSON(todo)
}

func (h *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	h.db.First(&todo, id)
	h.db.Delete(&todo)

	return c.SendString("Todo deleted")
}

func (h *TodoHandler) UpdateTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	h.db.First(&todo, id)

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	h.db.Save(&todo)

	return c.JSON(todo)
}

func (h *TodoHandler) CompleteAllTodos(c *fiber.Ctx) error {
	h.db.Model(&models.Todo{}).Update("completed", true)

	return c.SendString("All todos completed")
}

func (h *TodoHandler) DeleteCompletedTodos(c *fiber.Ctx) error {
	h.db.Where("completed = ?", true).Delete(&models.Todo{})

	return c.SendString("Completed todos deleted")
}
