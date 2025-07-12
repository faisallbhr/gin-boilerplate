package database

import (
	"fmt"
	"log"

	"github.com/faisallbhr/gin-boilerplate/config"
	"github.com/faisallbhr/gin-boilerplate/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbUser := config.GetEnv("DB_USER", "root")
	dbPassword := config.GetEnv("DB_PASSWORD", "")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "3306")
	dbName := config.GetEnv("DB_NAME", "golang")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Connected to database")

	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated")
}
