package services

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
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
func GetConf(fileName string) Conf {

	var c Conf

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Debugf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Debugf("Unmarshal: %v", err)
	}

	return c
}
