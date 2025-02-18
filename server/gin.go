package server

import (
	"database/sql"
	"log"
	"os"
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

	router.POST("/users/new", func(c *gin.Context) {
		routes.CreateUser(db, c)
	})

	log.Println("Server running on :8888")
	router.Run(port)
}
