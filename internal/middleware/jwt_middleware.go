package middleware

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// doesnt work yet, because i dont know how to implement it
func JWTMiddleware() string {
	err := godotenv.Load(filepath.Join("..", ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		log.Fatal("Error signing token")
	}

	return tokenString
}
