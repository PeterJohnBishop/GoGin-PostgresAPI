package postgres

import (
	"database/sql"
	"fmt"
	"time"
)

type Message struct {
	ID           int       `json:"id"`
	Sender       string    `json:"sender"` // the author user.id
	TextContent  string    `json:"text_content"`
	MediaContent []string  `json:"media_content"` // stores urls
	Likes        []string  `json:"likes"`         // stores user.id's
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func CreateMessage(db *sql.DB, message Message) error {
	query := "INSERT INTO messages (sender, text_content, media_content) VALUES ($1, $3, $4) RETURNING id, created_at"
	queryErr := db.QueryRow(query, message.Sender, message.TextContent, message.MediaContent).Scan(&message.ID, &message.CreatedAt)
	if queryErr != nil {
		fmt.Println(queryErr)
		return queryErr
	}

	return nil
}

func GetMessages(db *sql.DB) ([]Message, error) {
	rows, err := db.Query("SELECT id, sender, text_content, media_content, likes, created_at, updated_at FROM messages;")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var message Message
		if err := rows.Scan(&message.ID, &message.Sender, &message.TextContent, &message.MediaContent, &message.Likes, &message.CreatedAt, &message.UpdatedAt); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}
		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return nil, err
	}

	return messages, nil
}

// Disabling for now as I don't see a use case for this
//
// func GetMessageById(db *sql.DB, id int) (Message, error) {
// 	var message Message
// 	query := "SELECT id, sender, text_content, media_content, likes, created_at, updated_at FROM messages WHERE id = $1"
// 	err := db.QueryRow(query, id).Scan(&message.ID, &message.Sender, &message.TextContent, &message.MediaContent, &message.Likes, &message.CreatedAt, &message.UpdatedAt)
// 	if err != nil {
// 		fmt.Println("Error executing query:", err)
// 		return message, err
// 	}

// 	return message, nil
// }

func DeleteMessage(db *sql.DB, id int) error {

	query := "DELETE FROM messages WHERE id = $1"
	res, err := db.Exec(query, id)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return err
	}

	return nil
}
