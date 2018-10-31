package services

import (
	"database/sql"
	m "github.go/vincedgy/golang_tutos/databases/models"

	_ "github.com/lib/pq"
)

func GetUserById(db *sql.DB, id int64) m.User {

	log.Debugf("Starting")

	 // Fetch some data
	sqlStatement := `SELECT user_oid, login, email FROM sch_services.ccuser WHERE user_oid = $1;`

	var user m.User

	row := db.QueryRow(sqlStatement, id)

	switch err := row.Scan(&user.Id, &user.Login, &user.Email); err {
	case sql.ErrNoRows:
		log.Errorf("No rows were returned!")
	case nil:
		log.Infof("Data retrieved...")
	default:
		panic(err)
	}

	return user
}
