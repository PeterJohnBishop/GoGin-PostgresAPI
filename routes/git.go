package routes

import (
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

func HandleUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Error retrieving file: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}

	savePath := "/Users/peterbishop/Development/local/" + file.Filename
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		log.Printf("Error saving file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	cmd := exec.Command("git", "-C", "/Users/peterbishop/Development/local/", "add", ".")
	output, _ := cmd.CombinedOutput()
	if err := cmd.Run(); err != nil {
		log.Printf("Error staging file: %v\nOutput: %s", err, string(output))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to stage file."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File staged."})
}

func HandlePull(c *gin.Context) {
	cmd := exec.Command("git", "-C", "/Users/peterbishop/Development/local/", "pull")
	output, err := cmd.CombinedOutput()

	if err != nil {
		if strings.Contains(string(output), "not found") {
			setUpstreamCmd := exec.Command("git", "-C", "/Users/peterbishop/Development/local/", "push", "--set-upstream", "origin", "main")
			upstreamOutput, upstreamErr := setUpstreamCmd.CombinedOutput()
			if upstreamErr != nil {
				log.Printf("Error setting upstream: %v\nOutput: %s", upstreamErr, string(upstreamOutput))
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set upstream"})
				return
			}
			cmd = exec.Command("git", "-C", "/Users/peterbishop/Development/local/", "pull")
			output, err = cmd.CombinedOutput()
		}

		if err != nil {
			log.Printf("Error pulling changes: %v\nOutput: %s", err, string(output))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to pull changes"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Repo synced"})
}

func HandleCommit(c *gin.Context) {
	cmd := exec.Command("git", "-C", "/Users/peterbishop/Development/local/", "commit", "-m", "Backup commit")
	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit changes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Repo synced"})
}
