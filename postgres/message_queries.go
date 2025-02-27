package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/lib/pq"
)

type Message struct {
	MessageID    string    `json:"message_id"`
	Sender       string    `json:"sender"` // the author user.id
	TextContent  string    `json:"text_content"`
	MediaContent []string  `json:"media_content"` // stores urls
	Likes        []string  `json:"likes"`         // stores user.id's
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// CREATE TABLE messages (message_id TEXT UNIQUE NOT NULL PRIMARY KEY, sender TEXT NOT NULL, text_content TEXT NOT NULL, media_content TEXT[] DEFAULT '{}', likes TEXT[] DEFAULT '{}', created_at TIMESTAMP DEFAULT NOW(), updated_at TIMESTAMP DEFAULT NOW());
// {
// 	"message_id": "message_e0de5e50-f3ac-11ef-a249-0a400849f31f",
// 	"sender": "cce3855a-c206-4bc6-bbae-106c8f73892a",
// 	"text_content": "Hello, world!",
// 	"media_content": [],
// 	"likes": [],
// 	"created_at": "2025-02-25T12:15:30.439203Z",
// 	"updated_at": "2025-02-25T12:15:30.439203Z"
// }

func CreateMessage(db *sql.DB, message *Message) error {
	id, err := uuid.NewV1()
	if err != nil {
		log.Fatal(err)
	}
	messageID := "message_" + id.String()

	query := `
		INSERT INTO messages (message_id, sender, text_content, media_content)
		VALUES ($1, $2, $3, $4)
		RETURNING created_at, likes
	`

	queryErr := db.QueryRow(query, messageID, message.Sender, message.TextContent, pq.Array(message.MediaContent)).
		Scan(&message.CreatedAt, pq.Array(&message.Likes))

	if queryErr != nil {
		fmt.Println("Error inserting message:", queryErr)
		return queryErr
	}

	message.MessageID = messageID

	return nil
}

func GetMessages(db *sql.DB) ([]Message, error) {
	rows, err := db.Query("SELECT message_id, sender, text_content, media_content, likes, created_at, updated_at FROM messages;")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return nil, err
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var message Message

		if err := rows.Scan(
			&message.MessageID,
			&message.Sender,
			&message.TextContent,
			pq.Array(&message.MediaContent),
			pq.Array(&message.Likes),
			&message.CreatedAt,
			&message.UpdatedAt,
		); err != nil {
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

func DeleteMessage(db *sql.DB, id string) error {

	query := "DELETE FROM messages WHERE message_id = $1"
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
