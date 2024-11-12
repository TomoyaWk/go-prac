package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// tutorial in https://go.dev/doc/tutorial/web-service-gin
// albumの構造体
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albumのスライス
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// albumsのリスト取得
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// ID指定取得
func getAlbumsById(c *gin.Context) {
	id := c.Param("id")
	for _, al := range albums {
		if al.ID == id {
			c.IndentedJSON(http.StatusOK, al)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "album not Found."})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusOK, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsById)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
