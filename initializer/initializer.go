package initializer

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/model"
)

var DB *gorm.DB

func SetupDatabase() {
	dsn := os.Getenv("DSN")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	if err := DB.AutoMigrate(&model.UserModel{}); err != nil {
		fmt.Printf("Error migrating database %v", err)
	}
}
func Envload() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v", err)
	}
}
