package main

import (
	"main/Handlers/GET"
	"main/Handlers/POST"

	"github.com/gin-gonic/gin"
)


func main() {
    router := gin.Default()

    //GET Routes
    router.GET("/albums", GET.GetAlbums)
    router.GET("/items", GET.GetItems)
    router.GET("/albums/:id", GET.GetSingleAlbum)

    //POST routes
    router.POST("/albums", POST.AddAlbum)

    router.Run("localhost:8080")
}



