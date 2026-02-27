package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    err := godotenv.Load()
    if err != nil {
        log.Println("⚠️ Warning: Error loading .env file")
    }
    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        log.Fatal("❌ DATABASE_URL tidak ditemukan di environment atau file .env")
    }
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("❌ Failed to connect to database: ", err)
    }

    DB = database
    log.Println("✅ Database connected successfully")
}