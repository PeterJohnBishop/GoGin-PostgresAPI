package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/lib/pq"
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

func CreateMessage(db *sql.DB, message *Message) error {
	query := "INSERT INTO messages (sender, text_content, media_content) VALUES ($1, $2, $3) RETURNING id, created_at, likes"

	queryErr := db.QueryRow(query, message.Sender, message.TextContent, pq.Array(message.MediaContent)).Scan(&message.ID, &message.CreatedAt, pq.Array(&message.Likes))
	if queryErr != nil {
		fmt.Println("Error inserting message:", queryErr)
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
		var updatedAt sql.NullTime

		if err := rows.Scan(
			&message.ID,
			&message.Sender,
			&message.TextContent,
			pq.Array(&message.MediaContent),
			pq.Array(&message.Likes),
			&message.CreatedAt,
			&updatedAt,
		); err != nil {
			fmt.Println("Error scanning row:", err)
			return nil, err
		}

		if updatedAt.Valid {
			message.UpdatedAt = updatedAt.Time
		} else {
			message.UpdatedAt = time.Time{}
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return nil, err
	}

	return messages, nil
}

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
