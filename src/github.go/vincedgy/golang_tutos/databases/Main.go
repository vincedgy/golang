package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	s "github.go/vincedgy/golang_tutos/databases/services"
	"log"
)

func main() {
	/*
	The main function of our local module
	 */

	log.Println("Starting...")

	var conn *sql.DB = s.GetConnection("/Users/vdagoury/Projects/golang/workspace/src/github.go/vincedgy/golang_tutos/databases/application.yaml")
	var user = s.GetUserById(conn,1)

	log.Printf("The user's with id %d has email : %s", user.Id, user.Email)

	user = s.GetUserById(conn,100)
	log.Printf("The user's with id %d has email : %s", user.Id, user.Email)


	s.CloseConnection()
	log.Println("Done")
}
