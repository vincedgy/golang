package services

import (
	"database/sql"
	m "github.com/vincedgy/databases/models"

	_ "github.com/lib/pq"
)



func GetUserById(id int) m.User {

	log.Debugf("Starting GetUserById")

	var db *sql.DB = GetConnection("")

	// Fetch some data
	sqlStatement := `SELECT user_oid, login, email FROM sch_services.ccuser WHERE user_oid = $1;`

	var user m.User
	row := db.QueryRow(sqlStatement, id)

	switch err := row.Scan(&user.Id, &user.Login, &user.Email); err {
	case sql.ErrNoRows:
		log.Debugf("No rows were returned!")
	case nil:
		log.Debugf("Data retrieved in %0.6dns", db.Stats().WaitCount * db.Stats().WaitDuration.Nanoseconds())
	default:
		panic(err)
	}

	return user
}

func GetUsersCount() int {

	log.Debugf("Starting GetUserCount")
	var db *sql.DB = GetConnection("")

	var usersCount int = 0

	// Fetch some data
	sqlStatement := `SELECT count(user_oid) as cnt FROM sch_services.ccuser;`

	row := db.QueryRow(sqlStatement)

	switch err := row.Scan(&usersCount); err {
	case sql.ErrNoRows:
		log.Debugf("No rows were returned!")
	case nil:
		log.Debugf("Data retrieved in %0.6dns", db.Stats().WaitCount * db.Stats().WaitDuration.Nanoseconds())
	default:
		panic(err)
	}

	return usersCount
}

func GetUsers() []m.User {

	log.Debugf("Starting GetUsers")

	var db *sql.DB = GetConnection("")

	// Fetch some data
	sqlStatement := `SELECT user_oid, login, email FROM sch_services.ccuser;`

	rows, err := db.Query(sqlStatement)
	defer rows.Close()

	var users []m.User

	for rows.Next() {
		var user m.User
		err = rows.Scan(&user.Id, &user.Login, &user.Email)
		if err != nil {
			// handle this error
			panic(err)
		}
		users = append(users, user)

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return users
}