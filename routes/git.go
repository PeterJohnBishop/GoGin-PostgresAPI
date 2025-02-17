package routes

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func HandleUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload failed"})
		return
	}

	savePath := "./repo/" + file.Filename
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	cmd := exec.Command("git", "-C", "./repo", "add", ".")
	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to stage file."})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File staged."})
}

func HandlePull(c *gin.Context) {
	cmd := exec.Command("git", "-C", "./repo", "pull")
	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to pull changes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Repo synced"})
}

func HandleCommit(c *gin.Context) {
	cmd := exec.Command("git", "-C", "./repo", "commit", "-m", "Backup commit")
	if err := cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit changes"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Repo synced"})
}
