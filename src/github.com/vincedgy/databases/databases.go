/*

Package main implements a simple call of libraries developed for database interactions


 */
package main

import (
	"fmt"
	"github.com/juju/loggo"
	_ "github.com/lib/pq"
	"github.com/pborman/getopt/v2"
	m "github.com/vincedgy/databases/models"
	s "github.com/vincedgy/databases/services"
	"os"
)

// This is the main function
func main() {
	/*
		The main function of our local module
	*/
	handleArgs()

	/**
	------------------
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

// Displays help message and Usage
func help() {
	fmt.Print(`
This is a tool for managing db connection and get some data.

`)

	getopt.Usage()

	fmt.Print(`
The commands are:

	- There's no command yet
`)

	fmt.Println(`
Thank you.`)
}

// Handles arguments from the command line
func handleArgs() {
	// Declare the flags and parameters to use
	optHelp := getopt.BoolLong("help", 0, "Display this help message")
	configFileParam := getopt.StringLong("file", 'f', "resources/application.yaml", "The YAML file with datasource configuration")

	// Parse the program arguments
	getopt.Parse()

	// Get the remaining positional parameters
	//args := getopt.Args()

	// Handle the Help flag
	if *optHelp {
		//getopt.Usage()
		help()
		os.Exit(0)
	}

	// Handling the database config file parameter and assume a failover value
	if *configFileParam != "" {
		s.InitConnection(*configFileParam)
	} else {
		s.InitConnection(s.DefaultConfigFile)
	}

}
