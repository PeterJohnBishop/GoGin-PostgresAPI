package server

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"symetrical-fishstick-go/main.go/postgres"
	"symetrical-fishstick-go/main.go/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Gin_Server(db *sql.DB) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("GIN_PORT")

	postgres.GetUsers(db)

	router := gin.Default()

	// USER ROUTES
	router.POST("/users/new", func(c *gin.Context) {
		routes.CreateUserHandler(db, c)
	})
	router.POST("/login", func(c *gin.Context) {
		email := c.Request.FormValue("email")
		password := c.Request.FormValue("password")
		routes.Login(db, email, password, c)
	})
	router.GET("/users/", func(c *gin.Context) {
		routes.GetUsersHandler(db, c)
	})
	router.GET("/users/email/:email", func(c *gin.Context) {
		email := c.Param("email")
		routes.GetUserByEmailHandler(db, email, c)
	})
	router.GET("/users/uuid/:uuid", func(c *gin.Context) {
		uuid := c.Param("uuid")
		routes.GetUserByUUIDHandler(db, uuid, c)
	})
	router.PUT("/users/:uuid", func(c *gin.Context) {
		uuid := c.Param("uuid")
		routes.UpdateUserHandler(db, uuid, c)
	})
	router.DELETE("/users/:uuid", func(c *gin.Context) {
		uuid := c.Param("uuid")
		routes.DeleteUserHandler(db, uuid, c)
	})

	// MESSAGE ROUTES
	router.POST("/messages/new", func(c *gin.Context) {
		routes.CreateMessageHandler(db, c)
	})
	router.GET("/messages/", func(c *gin.Context) {
		routes.GetMessagesHandler(db, c)
	})
	router.DELETE("/messages/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid message ID"})
			return
		}
		routes.DeleteMessageHandler(db, idInt, c)
	})

	// PRODUCT ROUTES
	router.POST("/products/new", func(c *gin.Context) {
		routes.CreateProductHandler(db, c)
	})
	router.GET("/products/", func(c *gin.Context) {
		routes.GetProductsHandler(db, c)
	})
	router.GET("/products/:productID", func(c *gin.Context) {
		id := c.Param("productID")
		routes.GetProductByIdHandler(db, id, c)
	})
	router.PUT("/products/:productID", func(c *gin.Context) {
		id := c.Param("productID")
		routes.UpdateProductHandler(db, id, c)
	})
	router.DELETE("/products/:productID", func(c *gin.Context) {
		id := c.Param("productID")
		routes.DeleteProductHandler(db, id, c)
	})

	log.Println("Server running on :8888")
	router.Run(port)
}
