package routes

import (
	"database/sql"
	"net/http"
	"strings"

	"symetrical-fishstick-go/main.go/authentication"
	"symetrical-fishstick-go/main.go/postgres"

	"github.com/gin-gonic/gin"
)

func CreateOrderHandler(db *sql.DB, c *gin.Context) {
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
	var order postgres.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	msgErr := postgres.CreateOrder(db, &order)
	if msgErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create message"})
		return
	}
	c.JSON(http.StatusCreated, order)
}

func GetOrderByIdHandler(db *sql.DB, order_id string, c *gin.Context) {
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
	var order postgres.Order
	foundOrder, err := postgres.GetOrderById(db, order_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product by that id"})
		return
	}
	order = foundOrder
	c.JSON(http.StatusOK, order)
}

func GetOrdersHandler(db *sql.DB, c *gin.Context) {
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
	var orders []postgres.Order
	allOrders, err := postgres.GetOrders(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get all orders!"})
		return
	}
	orders = allOrders
	c.JSON(http.StatusOK, orders)
}

func UpdateOrderHandler(db *sql.DB, order_id string, c *gin.Context) {
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
	var order postgres.Order
	if err := c.ShouldBindBodyWithJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	updatedOrder, err := postgres.UpdateOrder(db, order_id, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Order updated successfully",
		"order":   updatedOrder,
	})
}

func DeleteOrderHandler(db *sql.DB, order_id string, c *gin.Context) {
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
	authErr := authentication.VerifyToken(token)
	if authErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify token!"})
		return
	}
	err := postgres.DeleteOrder(db, order_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
