package routes

import (
	"database/sql"
	"net/http"

	"symetrical-fishstick-go/main.go/authentication"
	"symetrical-fishstick-go/main.go/postgres"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(db *sql.DB, c *gin.Context) {
	var user postgres.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := postgres.CreateUser(db, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)

}

func Login(db *sql.DB, email string, password string, c *gin.Context) {
	var user postgres.User
	foundUser, err := postgres.GetUserByEmail(db, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user by that email"})
		return
	}

	user = foundUser

	pass := authentication.CheckPasswordHash(password, user.Password)
	if !pass {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password Verfication Failed"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUserByEmailHandler(db *sql.DB, email string, c *gin.Context) {
	var user postgres.User
	foundUser, err := postgres.GetUserByEmail(db, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user by that email"})
		return
	}
	user = foundUser
	c.JSON(http.StatusOK, user)
}

func GetUserByIdHandler(db *sql.DB, id int, c *gin.Context) {
	var user postgres.User
	foundUser, err := postgres.GetUserById(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user by that id"})
		return
	}
	user = foundUser
	c.JSON(http.StatusOK, user)

}

func GetUsersHandler(db *sql.DB, c *gin.Context) {
	var users []postgres.User
	allUsers, err := postgres.GetUsers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all users"})
		return
	}

	users = allUsers
	c.JSON(http.StatusOK, users)
}

func UpdateUserHandler(db *sql.DB, id int, c *gin.Context) {
	var user postgres.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedUser, err := postgres.UpdateUser(db, id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    updatedUser,
	})
}

func DeleteUserHandler(db *sql.DB, id int, c *gin.Context) {
	err := postgres.DeleteUser(db, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}
