package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	e := echo.New()

	db := ConnectDb()
	defer db.Close()
	// 初期テーブルを作成
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		password TEXT NOT NULL UNIQUE
	);`
	if _, err := db.Exec(createTableQuery); err != nil {
		log.Fatal("Failed to create table:", err)
	}
	//初期レコードの投入
	initUserTable(db)

	//user list
	e.GET("/users", GetUsers())
	//create user
	e.POST("/users", CreateUser())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	e.GET("/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(http.StatusOK, fmt.Sprintf("Hello,"+name+"!"))
	})

	//start in "http://localhost:1323"
	e.Logger.Fatal(e.Start(":1323"))
}

func ConnectDb() *sql.DB {
	// SQLite接続
	db, err := sql.Open("sqlite3", "./db/practice.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// 初期データを挿入する関数
func initUserTable(db *sql.DB) {
	// テーブル内のレコード数を確認
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		log.Fatal("Failed to count users:", err)
	}

	// レコードが0件の場合 初期データを挿入
	if count == 0 {
		log.Println("Inserting initial data into users table...")

		initialUsers := []struct {
			Name     string
			Password string
		}{
			{"hanako", "password123"},
			{"taro", "password1234"},
			{"jiro", "password12345"},
		}

		for _, user := range initialUsers {
			_, err := db.Exec("INSERT INTO users (name, password) VALUES (?, ?)", user.Name, user.Password)
			if err != nil {
				log.Printf("Failed to insert user (%s, %s): %v", user.Name, user.Password, err)
			}
		}
	} else {
		log.Println("Users table already has data. Skipping initial data insertion.")
	}

}
