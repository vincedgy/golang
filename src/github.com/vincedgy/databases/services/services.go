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

var config Conf

const DefaultConfigFile = "resources/application.yaml"
