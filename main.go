package main

import "net/http"

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Kinda Blue", Artist: "Miles Davis", Price: 56.99},
	{ID: "2", Title: "Hail to the Thief", Artist: "Radiohead", Price: 20.99},
	{ID: "3", Title: "GNX", Artist: "Kendrick Lamar", Price: 35.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
