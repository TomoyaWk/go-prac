package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// UserList
func GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := ConnectDb()
		rows, err := db.Query("SELECT * FROM users;")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
		}
		defer rows.Close()

		var userList []User
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.ID, &user.Name, &user.Password); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
			}
			userList = append(userList, user)
		}
		return c.JSONPretty(http.StatusOK, userList, "	")
	}
}

// create User
// curl -X POST --data-urlencode 'name=testman' -d 'password=1234' http://localhost:1323/users
func CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		db := ConnectDb()
		name := c.FormValue("name")
		password := c.FormValue("password")
		insertQuery := `INSERT INTO users (name, password) VALUES (?, ?)`
		if _, err := db.Exec(insertQuery, name, password); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to Insert users"})
		}
		return c.NoContent(http.StatusOK)
	}
}
