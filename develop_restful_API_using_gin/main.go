package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(context *gin.Context) {
	var newAlbum album

	// BindJSON binds the request body with the newAlbum created
	// if error is not nil i.e if error exists then return empty
	if error := context.BindJSON(&newAlbum); error != nil {
		context.IndentedJSON(http.StatusBadRequest, "invalid request body recieved")
		return
	}

	albums = append(albums, newAlbum)

	context.IndentedJSON(http.StatusCreated, albums)
}

func getAlbumById(c *gin.Context) {
	var id string = c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusFound, album)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	fmt.Printf("Creating a server")

	router := gin.Default()

	router.GET("/albums", getAlbums)

	router.POST("albums", postAlbums)

	router.GET("album/:id", getAlbumById)

	router.Run("localhost:8080")

}
