package LOGIN

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type user struct {
    Username string `json:"username"`
    Password string `json:"password"`
}


func LoginHandler(c *gin.Context) {
  	var user user

	if err := c.BindJSON(&user); err != nil {
        c.AbortWithStatus(http.StatusNetworkAuthenticationRequired)
    }
	fmt.Printf("Username: %s, Password: %s\n", user.Username, user.Password)
    if user.Username == "user" && user.Password == "password" {
        // Create a new token
        token := jwt.New(jwt.SigningMethodHS256)
        claims := token.Claims.(jwt.MapClaims)
        claims["username"] = user.Username
        // Set expiration time if needed
        claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

        // Sign the token with the secret key
        tokenString, err := token.SignedString(jwtKey)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
            return
        }

        // Send the token in the response
        c.JSON(http.StatusOK, gin.H{"token": tokenString})
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    }
}
