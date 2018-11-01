///bin/true; exec /usr/bin/env go run "$0" "$@"
package main

import (
	"github.com/juju/loggo"
	_ "github.com/lib/pq"
	m "github.com/vincedgy/databases/models"
	s "github.com/vincedgy/databases/services"
)

func main() {
	/*
		The main function of our local module
	*/

	loggo.ConfigureLoggers("<root>=DEBUG")
	var log = loggo.GetLogger("main")

	log.Infof("Starting...")
	defer s.CloseConnection()

	// Create a collection of Users a fetch for each id from 0 to the length of the collection
	var users = new([10]m.User)

	for id := 0; id < len(users); id++ {
		users[id] = s.GetUserById(id)
		log.Infof("The user's with id %d (%d) has email : %s", id, users[id].Id, users[id].Email)
	}

	// Count the total numbers of Users
	var usersCount int = s.GetUsersCount()
	log.Infof("The number of counted users in db is : %d", usersCount)

	// Retrieve all the users
	var usersCollection []m.User = s.GetUsers()
	log.Infof("The number of users fetched from db is : %d", len(usersCollection))

	log.Infof("Done")
}
