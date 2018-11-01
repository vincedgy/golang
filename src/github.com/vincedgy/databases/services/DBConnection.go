package services

import (
	"database/sql"
	"fmt"
)

func db() *sql.DB {
	return conn
}

func InitConnection(confFile string) {
	/**
	GetConnection serves a singleton db connector
	It initializes it the first time and serves it
	*/

	config = InitializeConfFromFile(confFile)

	if conn == nil {
		log.Debugf("Create the db connection")

		// Formating the database connection string
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			config.Datasource.Main.Host,
			config.Datasource.Main.Port,
			config.Datasource.Main.User,
			config.Datasource.Main.Password,
			config.Datasource.Main.Dbname)

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
