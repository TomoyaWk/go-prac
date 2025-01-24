package main

import (
	"log"
	"sample-api/internal/config"

	"sample-api/internal/models"
	"sample-api/internal/routes"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.GetConfig()
	dbConfig := "host=" + cfg.DBHost + " user=" + cfg.DBUser + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " port=" + cfg.DBPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.User{})

	e := echo.New()
	routes.RegisterRoutes(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}
