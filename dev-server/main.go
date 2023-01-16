package main

import (
	"net/http"

	"github.com/gabr0236/pokedle/dev-server/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/guess", getGuessResult)

	router.Run("localhost:3000")
}

/*
curl http://localhost:3000/guess --include --header "Content-Type: application/json" --request "GET" --data '"squirtle"'
*/

// getAlbums responds with the list of all albums as JSON.
func getGuessResult(c *gin.Context) {
	var name string
	println("Call enter")
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&name); err != nil {
		return
	}
	//fmt.Println("Guess object:", guess)
	//TODO: call service
	services.SubmitGuess(name)

	c.IndentedJSON(http.StatusOK, name)
}

// postAlbums adds an album from JSON received in the request body.
// func postAlbums(c *gin.Context) {
// 	var newAlbum album

// 	// Call BindJSON to bind the received JSON to
// 	// newAlbum.
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	// Add the new album to the slice.
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
// func getAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	// Loop over the list of albums, looking for
// 	// an album whose ID value matches the parameter.
// 	for _, a := range albums {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }

/*
curl http://localhost:3000/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

	curl http://localhost:3000/albums \
    --header "Content-Type: application/json" \
    --request "GET"
*/
