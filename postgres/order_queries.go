package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/lib/pq"
)

type Order struct {
	OrderID   string    `json:"order_id"`
	UserID    string    `json:"user_id"`
	Cart      []string  `json:"cart"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CREATE TABLE orders (order_id TEXT UNIQUE NOT NULL PRIMARY KEY, user_id TEXT NOT NULL, cart TEXT[] DEFAULT '{}', status TEXT NOT NULL, created_at TIMESTAMP DEFAULT NOW(), updated_at TIMESTAMP DEFAULT NOW());

func CreateOrder(db *sql.DB, order *Order) error {
	id, err := uuid.NewV1()
	if err != nil {
		log.Fatal(err)
	}
	orderID := "order_" + id.String()
	query := `INSERT INTO orders (order_id, user_id, cart, status) VALUES ($1, $2, $3, $4) RETURNING created_at`
	queryErr := db.QueryRow(query, orderID, order.UserID, pq.Array(order.Cart), order.Status).Scan(&order.CreatedAt)
	if queryErr != nil {
		fmt.Println("Error inserting order:", queryErr)
		return queryErr
	}
	order.OrderID = orderID
	return nil
}

func GetOrderById(db *sql.DB, order_id string) (Order, error) {
	var order Order
	query := "SELECT order_id, user_id, cart, status, created_at, update_at FROM orders WHERE order_id = $1"
	err := db.QueryRow(query, order_id).Scan(&order.OrderID, &order.UserID, pq.Array(&order.Cart), &order.Status, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return order, err
	}
	return order, nil
}

func GetOrders(db *sql.DB) ([]Order, error) {
	rows, err := db.Query("SELECT order_id, user_id, cart, status, created_at, updated_at FROM orders")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()
	var orders []Order
	for rows.Next() {
		var order Order
		if err := rows.Scan(&order.OrderID, &order.UserID, pq.Array(&order.Cart), &order.Status, &order.CreatedAt, &order.UpdatedAt); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		orders = append(orders, order)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return nil, err
	}
	return orders, nil
}

func UpdateOrder(db *sql.DB, orderID string, order Order) (Order, error) {
	query := `UPDATE orders SET cart = $2, status = $3, updated_at = NOW() WHERE order_id = $1 RETURNING order_id, user_id, cart, status, created_at, updated_at`
	var updatedOrder Order
	err := db.QueryRow(query, orderID, updatedOrder.UserID, pq.Array(updatedOrder.Cart), updatedOrder.Status).
		Scan(&updatedOrder.OrderID, &updatedOrder.UserID, pq.Array(&updatedOrder.Cart), &updatedOrder.Status, &updatedOrder.CreatedAt, &updatedOrder.UpdatedAt)
	if err != nil {
		fmt.Println("Error updating product:", err)
		return Order{}, err
	}
	return updatedOrder, nil
}

func DeleteOrder(db *sql.DB, order_id string) error {
	query := "DELETE FROM orders WHERE order_id = $1"
	res, err := db.Exec(query, order_id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return err
	}
	return nil
}
