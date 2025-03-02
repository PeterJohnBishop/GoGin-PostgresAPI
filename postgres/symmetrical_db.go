package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func ConnectPSQL(db *sql.DB) *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("PSQL_HOST")
	portStr := os.Getenv("PSQL_PORT")

	port, err := strconv.Atoi(portStr) // Convert to int
	if err != nil {
		fmt.Println("Invalid port number:", err)
	}
	user := os.Getenv("PSQL_USER")
	fmt.Println("User:", user)
	password := os.Getenv("PSQL_PASSWORD")
	fmt.Println("Password:",
		password)
	dbname := os.Getenv("PSQL_DBNAME")
	fmt.Println("DB Name:", dbname)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println("PSQL Info:", psqlInfo)
	mydb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = mydb.Ping()
	if err != nil {
		panic(err)
	}

	return mydb
}
