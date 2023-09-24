package GET

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Define a struct for your data
type Item struct {
    ID   int `json:"ID"`
    Name string `json:"name"`
}

// Create a slice to store your data (for simplicity, we're using an in-memory store)
var items []Item


// GetItems returns a JSON list of all items
func GetItems(c *gin.Context) {
	if len(items) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusOK, items)
	}
	
	// CreateItem adds a new item to the list
	func CreateItem(w http.ResponseWriter, r *http.Request) {
		var newItem Item
		err := json.NewDecoder(r.Body).Decode(&newItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		// Add the new item to the list
		items = append(items, newItem)
	
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newItem)
	}
	
	// album represents data about a record album.
	type album struct {
		ID     string  `json:"id"`
		Title  string  `json:"title"`
		Artist string  `json:"artist"`
		Price  float64 `json:"price"`
	}
	
	// albums slice to seed record album data.
	var albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	
	// getAlbums responds with the list of all albums as JSON.
	func GetAlbums(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, albums)
	}
	
	
	
	func GetSingleAlbum(c *gin.Context){
		id := c.Param("id")
	
		for _, a := range albums {
			if a.ID == id {
				c.IndentedJSON(http.StatusOK, a)
				return
			}
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
	}