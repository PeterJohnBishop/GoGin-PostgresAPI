package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/lib/pq"
)

type Product struct {
	ProductID    string    `json:"product_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	MediaContent []string  `json:"media_content"`
	Price        float64   `json:"price"`
	Inventory    int       `json:"inventory"`
	Likes        []string  `json:"likes"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// CREATE TABLE products (product_id TEXT UNIQUE NOT NULL PRIMARY KEY, name TEXT UNIQUE NOT NULL, description TEXT NOT NULL, media_content TEXT[] DEFAULT '{}', price MONEY DEFAULT 0.00, inventory INTEGER, likes TEXT[] DEFAULT '{}', created_at TIMESTAMP DEFAULT NOW(), updated_at TIMESTAMP DEFAULT NOW());

func CreateProduct(db *sql.DB, product *Product) error {
	id, err := uuid.NewV1()
	if err != nil {
		log.Fatal(err)
	}
	productID := "product_" + id.String()
	query := `INSERT INTO products (product_id, name, description, media_content, price, inventory) VALUES ($1, $2, $3, $4, $5, $6) RETURNING created_at, likes`
	queryErr := db.QueryRow(query, productID, product.Name, product.Description, pq.Array(product.MediaContent), product.Price, product.Inventory).
		Scan(&product.CreatedAt, pq.Array(&product.Likes))
	if queryErr != nil {
		fmt.Println("Error inserting product:", queryErr)
		return queryErr
	}
	product.ProductID = productID
	return nil
}

func GetProductById(db *sql.DB, product_id string) (Product, error) {
	var prod Product
	query := "SELECT product_id, name, description, media_content, price::NUMERIC, inventory, likes, created_at, updated_at FROM products WHERE product_id = $1"
	err := db.QueryRow(query, product_id).Scan(&prod.ProductID, &prod.Name, &prod.Description, pq.Array(&prod.MediaContent), &prod.Price, &prod.Inventory, pq.Array(&prod.Likes), &prod.CreatedAt, &prod.UpdatedAt)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return prod, err
	}
	return prod, nil
}

func GetProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT product_id, name, description, media_content, price::NUMERIC, inventory, likes, created_at, updated_at FROM products")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()
	var prods []Product
	for rows.Next() {
		var prod Product
		if err := rows.Scan(&prod.ProductID, &prod.Name, &prod.Description, pq.Array(&prod.MediaContent), &prod.Price, &prod.Inventory, pq.Array(&prod.Likes), &prod.CreatedAt, &prod.UpdatedAt); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		prods = append(prods, prod)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return nil, err
	}
	return prods, nil
}

func UpdateProduct(db *sql.DB, productID string, prod Product) (Product, error) {
	query := `UPDATE products SET name = $2, description = $3, media_content = $4, price = $5, inventory = $6, likes = $7, updated_at = NOW() WHERE product_id = $1 RETURNING product_id, name, description, media_content, price::NUMERIC, inventory, likes, created_at, updated_at`
	var updatedProduct Product
	err := db.QueryRow(query, productID, prod.Name, prod.Description, pq.Array(prod.MediaContent), prod.Price, prod.Inventory, pq.Array(prod.Likes)).
		Scan(&updatedProduct.ProductID, &updatedProduct.Name, &updatedProduct.Description, pq.Array(&updatedProduct.MediaContent), &updatedProduct.Price, &updatedProduct.Inventory, pq.Array(&updatedProduct.Likes), &updatedProduct.CreatedAt, &updatedProduct.UpdatedAt)
	if err != nil {
		fmt.Println("Error updating product:", err)
		return Product{}, err
	}
	return updatedProduct, nil
}

func DeleteProduct(db *sql.DB, product_id string) error {
	query := "DELETE FROM products WHERE product_id = $1"
	res, err := db.Exec(query, product_id)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return err
	}
	return nil
}
