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

// CREATE TABLE product_id TEXT UNIQUE NOT NULL PRIMARY KEY, name TEXT UNIQUE NOT NULL, description TEXT NOT NULL, media_content TEXT[] DEFAULT '{}', price MONEY DEFAULT 0.00, inventory INTEGER, likes UUID[] DEFAULT '{}', created_at TIMESTAMP DEFAULT NOW(), updated_at TIMESTAMP DEFAULT NOW());

func CreateProduct(db *sql.DB, product *Product) error {
	id, err := uuid.NewV1()
	if err != nil {
		log.Fatal(err)
	}
	productID := "product_" + id.String()

	query := `
		INSERT INTO messages (product_id, name, description, media_content, price, inventory)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING created_at, likes
	`

	queryErr := db.QueryRow(query, productID, product.Name, product.Description, pq.Array(product.MediaContent), product.Price, product.Inventory).
		Scan(&product.CreatedAt, pq.Array(&product.Likes))

	if queryErr != nil {
		fmt.Println("Error inserting product:", queryErr)
		return queryErr
	}

	product.ProductID = productID

	return nil
}
