package server

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"symetrical-fishstick-go/main.go/routes"

	"github.com/gin-gonic/gin"
)

func isGitRepo(path string) bool {
	_, err := os.Stat(path + "/HEAD") // Git repositories contain a HEAD file
	return err == nil
}

func initGitRepo(path string) error {
	if isGitRepo(path + "/.git") {
		fmt.Println("Git repository already initialized at", path)
		return nil
	}

	fmt.Println("Initializing new Git repository at", path)
	cmd := exec.Command("git", "init", "--bare", path)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to initialize Git repo: %w", err)
	}
	fmt.Println("Git repository initialized successfully.")
	return nil
}

func Gin_Server() {
	repoPath := "./repo"

	if err := initGitRepo(repoPath); err != nil {
		log.Fatalf("Error initializing Git repository: %v", err)
	}

	router := gin.Default()

	router.POST("/upload", routes.HandleUpload)
	router.GET("/pull", routes.HandlePull)
	router.POST("/commit", routes.HandleCommit)

	log.Println("Server running on :8888")
	router.Run(":8080")
}
