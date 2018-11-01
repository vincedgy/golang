package services

import (
	"database/sql"
	m "github.com/vincedgy/databases/models"

	_ "github.com/lib/pq"
)

func GetUserById(id int) m.User {

	log.Debugf("Starting GetUserById")

	// Fetch some data
	sqlStatement := `SELECT user_oid, login, email FROM sch_services.ccuser WHERE user_oid = $1;`

	var user m.User
	row := db().QueryRow(sqlStatement, id)

	switch err := row.Scan(&user.Id, &user.Login, &user.Email); err {
	case sql.ErrNoRows:
		log.Debugf("No rows were returned!")
	default:
		if err != nil {
			// handle this error
			panic(err)
		}
	}

	return user
}

func GetUsersCount() int {

	log.Debugf("Starting GetUserCount")
	var usersCount int = 0

	// Fetch some data
	sqlStatement := `SELECT count(user_oid) as cnt FROM sch_services.ccuser;`

	row := db().QueryRow(sqlStatement)

	switch err := row.Scan(&usersCount); err {
	case sql.ErrNoRows:
		log.Debugf("No rows were returned!")
	default:
		if err != nil {
			// handle this error
			panic(err)
		}
	}

	return usersCount
}

func GetUsers() []m.User {

	log.Debugf("Starting GetUsers")

	// Fetch some data
	sqlStatement := `SELECT user_oid, login, email FROM sch_services.ccuser;`

	rows, err := db().Query(sqlStatement)
	defer rows.Close()

	var users []m.User

	for rows.Next() {
		var user m.User
		switch err := rows.Scan(&user.Id, &user.Login, &user.Email); err {
		case sql.ErrNoRows:
			log.Debugf("No rows were returned!")
		default:
			if err != nil {
				// handle this error
				panic(err)
			}
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
