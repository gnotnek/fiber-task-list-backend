package handlers

import (
	"github.com/gnotnek/fiber-task-list-backend/internal/database"
	"github.com/gnotnek/fiber-task-list-backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /users [get]
func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.JSON(users)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	database.DB.Where("id = ?", id).First(&user)
	return c.JSON(user)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} models.User
// @Router /users [post]
func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	database.DB.Create(&user)
	return c.JSON(user)
}
