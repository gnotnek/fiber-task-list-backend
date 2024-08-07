package handlers

import (
	"github.com/gnotnek/fiber-task-list-backend/internal/database"
	"github.com/gnotnek/fiber-task-list-backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

// SignUp godoc
// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User object"
// @Success 200 {object} models.User
// @Router /users [post]
func SignUp(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	hash, err := HashPassword(user.Password)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	user.Password = hash
	database.DB.Create(&user)
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	var dbUser models.User
	database.DB.Where("email = ?", user.Email).First(&dbUser)
	if dbUser.ID == uuid.Nil {
		return c.Status(400).SendString("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		return c.Status(400).SendString("invalid credentials")
	}

	return c.JSON(fiber.Map{
		"message": "success",
		"token":   "token", // not implemented yet
	})
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
