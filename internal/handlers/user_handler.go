package handlers

import (
	"time"

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
// @Param {object} body string true "User email and password"
// @Success 200 {object} map[string]string "user created"
// @Router /register [post]
func SignUp(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	user.Password = string(hashedPassword)

	database.DB.Create(&user)
	return c.JSON(fiber.Map{"message": "user created"})
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags users
// @Accept json
// @Produce json
// @Param {object} body string true "User email and password"
// @Success 200 {object} map[string]string "login success!"
// @Router /login [post]
func Login(c *fiber.Ctx) error {
	var user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	var dbUser models.User
	database.DB.Where("email = ?", user.Email).First(&dbUser)
	if dbUser.ID == uuid.Nil {
		return c.Status(400).SendString("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		return c.Status(401).SendString("invalid password")
	}

	token, err := middleware.CreateToken(dbUser.ID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{"message": "login success!"})
}
