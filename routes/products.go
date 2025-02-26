package routes

import (
	"database/sql"
	"net/http"
	"strings"

	"symetrical-fishstick-go/main.go/authentication"
	"symetrical-fishstick-go/main.go/postgres"

	"github.com/gin-gonic/gin"
)

func CreateProductHandler(db *sql.DB, c *gin.Context) {

	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token missing!"})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == authHeader {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token format!"})
		return
	}

	err := authentication.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify token!"})
		return
	}

	var product postgres.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	msgErr := postgres.CreateProduct(db, &product)
	if msgErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create message"})
		return
	}

	c.JSON(http.StatusCreated, product)
}
