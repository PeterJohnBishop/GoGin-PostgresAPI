package main

import (
	"database/sql"
	"symetrical-fishstick-go/main.go/postgres"
	"symetrical-fishstick-go/main.go/server"
)

var db *sql.DB

func main() {
	db := postgres.ConnectPSQL(db)
	//server.Http_Server()
	server.Gin_Server(db)
	defer db.Close()

}
