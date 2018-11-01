package services

import (
	"database/sql"
	"github.com/juju/loggo"
)

/**
Global variables for the module services
*/

// Modules logging
var log = loggo.GetLogger("services")

// conn : the main db connection object
var conn *sql.DB = nil

// dbConfig : the database configuration
var dbConfig Conf

const DefaultConfigFile = "/Users/vdagoury/Projects/golang/workspace/src/github.go/vincedgy/golang_tutos/databases/services/application.yaml"