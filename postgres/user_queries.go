package postgres

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func GetUsers(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, email, created_at FROM users;") // Include created_at
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt); err != nil {
			fmt.Println("Error scanning row:", err)
			return
		}
		users = append(users, user)
		fmt.Println(users)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return
	}
}
