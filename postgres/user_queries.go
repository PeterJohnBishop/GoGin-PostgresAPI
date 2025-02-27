package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"symetrical-fishstick-go/main.go/authentication"
	"time"

	"github.com/gofrs/uuid/v5"
)

type User struct {
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CREATE TABLE users (user_id TEXT UNIQUE NOT NULL PRIMARY KEY, name TEXT UNIQUE NOT NULL, email TEXT UNIQUE NOT NULL, password TEXT NOT NULL, created_at TIMESTAMP DEFAULT NOW(), updated_at TIMESTAMP DEFAULT NOW());

func CreateUser(db *sql.DB, user User) error {
	id, err := uuid.NewV1()
	if err != nil {
		log.Fatal(err)
	}
	userID := "user_" + id.String()
	hashedPassword, err := authentication.HashedPassword(user.Password)
	if err != nil {
		fmt.Println(err)
		return err
	}
	query := "INSERT INTO users (user_id, name, email, password) VALUES ($1, $2, $3, $4) RETURNING created_at"
	queryErr := db.QueryRow(query, userID, user.Name, user.Email, hashedPassword).Scan(&user.CreatedAt)
	if queryErr != nil {
		fmt.Println(queryErr)
		return queryErr
	}
	return nil
}

func GetUserByEmail(db *sql.DB, email string) (User, error) {
	var user User
	query := "SELECT user_id, name, email, password, created_at, updated_at FROM users WHERE email = $1"
	err := db.QueryRow(query, email).Scan(&user.UserID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

func GetUserByUUID(db *sql.DB, user_id string) (User, error) {
	var user User
	query := "SELECT user_id, name, email, password, created_at, updated_at FROM users WHERE user_id = $1"
	err := db.QueryRow(query, user_id).Scan(&user.UserID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

func GetUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT user_id, name, email, password, created_at, updated_at FROM users;")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			log.Fatal(err)
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users, nil
}

func UpdateUser(db *sql.DB, user_id string, user User) (User, error) {
	query := `
	UPDATE users 
	SET name = $1, email = $2, password = $3, updated_at = NOW() 
	WHERE user_id = $4 
	RETURNING user_id, name, email, password, created_at, updated_at`
	var updatedUser User
	err := db.QueryRow(query, user.Name, user.Email, user.Password, user_id).
		Scan(&updatedUser.UserID, &updatedUser.Name, &updatedUser.Email, &updatedUser.Password, &updatedUser.CreatedAt, &updatedUser.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	return updatedUser, nil
}

func DeleteUser(db *sql.DB, user_id string) error {
	query := "DELETE FROM users WHERE user_id = $1"
	res, err := db.Exec(query, user_id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		fmt.Println(err)
		return err
	}
	return nil

}
