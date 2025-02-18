package server

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"symetrical-fishstick-go/main.go/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func isGitRepo(path string) bool {
	_, err := os.Stat(path + "/HEAD") // Git repositories contain a HEAD file
	return err == nil
}

func initGitRepo(path string) error {
	if isGitRepo(path + ".git") {
		fmt.Println("Git repository already initialized at", path)
		return nil
	}

	fmt.Println("Initializing new Git repository at", path)
	cmd := exec.Command("git", "init", path)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to initialize Git repo: %w", err)
	}
	fmt.Println("Git repository initialized successfully.")
	return nil
}

func connectPSQL() {
	host := os.Getenv("PSQL_HOST")
	port := os.Getenv("PSQL_PORT")
	user := os.Getenv("PSQL_USER")
	password := os.Getenv("PSQL_PASSWORD")
	dbname := os.Getenv("PSQL_DBNAME")
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func Gin_Server() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("GIN_PORT")

	repoPath := "/Users/peterbishop/Development/local/"

	if err := initGitRepo(repoPath); err != nil {
		log.Fatalf("Error initializing Git repository: %v", err)
	}

	connectPSQL()

	router := gin.Default()

	router.POST("/upload", routes.HandleUpload)
	router.GET("/pull", routes.HandlePull)
	router.POST("/commit", routes.HandleCommit)

	log.Println("Server running on :8888")
	router.Run(port)
}
