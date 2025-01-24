package routes

import (
	"sample-api/internal/handlers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	userHandler := handlers.NewUserHandler(db)
	e.GET("/users", userHandler.GetUsers)
	e.POST("/users", userHandler.CreateUser)
}
