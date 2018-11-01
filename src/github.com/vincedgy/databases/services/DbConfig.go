package services

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

// This is the structure of the yaml file we expect
type Conf struct {
	Datasource struct {
		Main struct {
			Host     string
			Port     int
			User     string
			Password string
			Dbname   string
		}
	}
}

// GetConf reads a yaml file and returns a configuration set for database connection
func InitializeConfFromFile(fileName string) Conf {

	var c Conf

	// Build up the default config file
	absPath, _ := filepath.Abs(fileName)
	_, err := os.Stat(absPath)
	if err != nil {
		panic(err)
	} else {
		log.Debugf("DB config file is %s", absPath)
	}

	yamlFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Debugf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Debugf("Unmarshal: %v", err)
	}

	return c
}
