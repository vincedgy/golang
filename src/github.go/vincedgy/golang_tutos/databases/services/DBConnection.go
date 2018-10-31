package services

import (
	"database/sql"
	"fmt"
)

func GetConnection(confFile string) *sql.DB {
	/**
	GetConnection serves a singleton db connector
	It initializes it the first time and serves it
	*/
	if conn == nil {

		log.Debugf("Create the db connection")

		// Retrieve information for connection
		dbConfig = GetConf(confFile)

		// Formating the database connection string
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			dbConfig.Datasource.Main.Host,
			dbConfig.Datasource.Main.Port,
			dbConfig.Datasource.Main.User,
			dbConfig.Datasource.Main.Password,
			dbConfig.Datasource.Main.Dbname)

		// Open the connection to the database
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			panic(err)
			defer db.Close()
		} else {
			conn = db
		}
	} else {
		log.Debugf("Existing connection")
	}

	log.Debugf("Returning the connection")
	return conn
}

func CloseConnection() {
	/**
	Close the connection to the database
	*/
	if conn != nil {
		log.Debugf("Closing db connection")
		defer conn.Close()
	} else {
		log.Debugf("Asking for closing db connection but not connection available")

	}
}
