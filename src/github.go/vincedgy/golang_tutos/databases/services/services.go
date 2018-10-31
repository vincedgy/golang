package services

import (
	"database/sql"
	"github.com/juju/loggo"
)

// Modules logging
var log = loggo.GetLogger("services")

// conn : the main db connection object
var conn *sql.DB = nil

// dbConfig : the database configuration
var dbConfig Conf