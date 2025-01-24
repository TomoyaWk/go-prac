package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id     int    `json:"id"`
	UserId string `json:"name" unique:"true"`
	Email  string `json:"email"`
}
