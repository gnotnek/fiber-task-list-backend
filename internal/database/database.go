package database

import (
	"fmt"
	"os"

	"github.com/gnotnek/fiber-task-list-backend/internal/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func InitDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		return nil, fmt.Errorf("DATABASE_URL is not set")
	}

	db, err := gorm.Open("postgres", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	err = db.AutoMigrate(models.Todo{}).Error
	if err != nil {
		return nil, fmt.Errorf("error migrating database: %w", err)
	}

	return db, nil
}

type DB struct {
	*gorm.DB
}
