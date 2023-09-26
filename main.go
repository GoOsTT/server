package main

import (
	GET "main/Handlers/GET"
	POST_ALBUM "main/Handlers/POST/ALBUM"
	POST_AUTH "main/Handlers/POST/AUTH"
	AUTHMIDDLEWARE "main/Middleware"

	"github.com/gin-gonic/gin"
)


func main() {
    router := gin.Default()

    //GET Routes
    router.GET("/albums", GET.GetAlbums)
    router.GET("/items", GET.GetItems)
    router.GET("/albums/:id", GET.GetSingleAlbum)

    //Authenticated Routes
    //POST routes
    router.POST("/albums", AUTHMIDDLEWARE.AuthMiddleware, POST_ALBUM.AddAlbum)

    //Authentication routes
    router.POST("/login", POST_AUTH.LoginHandler)

    router.Run("localhost:8080")
}



