package handlers

import (
	"github.com/gnotnek/fiber-task-list-backend/internal/database"
	"github.com/gnotnek/fiber-task-list-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	database.DB.Find(&todos)
	return c.JSON(todos)
}

func GetTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	database.DB.Where("id = ?", id).First(&todo)
	return c.JSON(todo)
}

func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Create(&todo)
	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DB.Delete(&models.Todo{}, id)
	return c.SendString("Todo deleted successfully")
}

func UpdateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Save(&todo)
	return c.JSON(todo)
}

func UpdateTodoByID(c *fiber.Ctx) error {
	id := c.Params("id")
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Model(&models.Todo{}).Where("id = ?", id).Updates(todo)
	return c.JSON(todo)
}

func CompleteAllTodos(c *fiber.Ctx) error {
	database.DB.Model(&models.Todo{}).Update("completed", true)
	return c.SendString("All todos completed successfully")
}

func GetCompletedTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	database.DB.Where("completed = ?", true).Find(&todos)
	return c.JSON(todos)
}
