package handlers

import (
	"github.com/gnotnek/fiber-task-list-backend/internal/database"
	"github.com/gnotnek/fiber-task-list-backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// GetTodos godoc
// @Summary Get all todos
// @Description Get all todos
// @Tags todos
// @Produce json
// @Success 200 {array} models.Todo
// @Router /todos [get]
func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	database.DB.Where("user_id = ?", c.Locals("userID").(uuid.UUID)).Find(&todos)
	return c.JSON(todos)
}

// GetTodoByID godoc
// @Summary Get todo by ID
// @Description Get todo by ID
// @Tags todos
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} models.Todo
// @Failure 404 {string} string "Todo not found"
// @Router /todos/{id} [get]
func GetTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	database.DB.Where("id = ?", id).First(&todo)
	return c.JSON(todo)
}

// CreateTodo godoc
// @Summary Create a new todo
// @Description Create a new todo
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Todo object"
// @Success 200 {object} models.Todo
// @Router /todos [post]
func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	todo.UserID = c.Locals("userID").(uuid.UUID)
	database.DB.Create(&todo)
	return c.JSON(todo)
}

// DeleteTodo godoc
// @Summary Delete todo by ID
// @Description Delete todo by ID
// @Tags todos
// @Produce plain
// @Param id path string true "Todo ID"
// @Success 200 {string} string "Todo deleted successfully"
// @Router /todos/{id} [delete]
func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Where("id = ?", id).Delete(&models.Todo{})
	return c.SendString("Todo deleted successfully")
}

// UpdateTodoByID godoc
// @Summary Update a todo by ID
// @Description Update a todo by ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param todo body models.Todo true "Todo object"
// @Success 200 {object} models.Todo
// @Router /todos/{id} [put]
func UpdateTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Model(&models.Todo{}).Where("id = ?", id).Updates(todo)
	return c.JSON(todo)
}

// CompleteAllTodos godoc
// @Summary Complete all todos
// @Description Complete all todos
// @Tags todos
// @Produce plain
// @Success 200 {string} string "All todos completed successfully"
// @Router /todos/complete [put]
func CompleteAllTodos(c *fiber.Ctx) error {
	database.DB.Model(&models.Todo{}).Where("user_id = ?", c.Locals("userID").(uuid.UUID)).Update("completed", true)
	return c.SendString("All todos completed successfully")
}

// GetCompletedTodos godoc
// @Summary Get all completed todos
// @Description Get all completed todos
// @Tags todos
// @Produce json
// @Success 200 {array} models.Todo
// @Router /todos/completed [get]
func GetCompletedTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	database.DB.Where("user_id = ? AND completed = ?", c.Locals("userID").(uuid.UUID), true).Find(&todos)
	return c.JSON(todos)
}
