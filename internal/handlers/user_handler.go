package handlers

import (
	"github.com/gnotnek/fiber-task-list-backend/internal/database"
	"github.com/gnotnek/fiber-task-list-backend/internal/middleware"
	"github.com/gnotnek/fiber-task-list-backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// SignUp godoc
// @Summary Sign up
// @Description Sign up
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User info"
// @Router /register [post]
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

// Login godoc
// @Summary Login
// @Description Login
// @Tags users
// @Accept json
// @Produce json
// @Param user body struct{email string, password string} true "User credentials"
// @Success 200
// @Router /login [post]
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

	token, err := middleware.CreateToken(dbUser.ID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:  "jwt",
		Value: token,
	})

	return c.Next()
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
