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
	password := os.Getenv("PSQL_PASSWORD")
	dbname := os.Getenv("PSQL_DBNAME")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
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
