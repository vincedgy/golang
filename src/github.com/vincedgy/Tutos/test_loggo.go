package main

import (
	"github.com/vincedgy/databases/vendor/github.com/juju/loggo"
)

var first = loggo.GetLogger("first")

func FirstCritical(message string) {
	first.Criticalf(message)
}

func FirstError(message string) {
	first.Errorf(message)
}

func FirstWarning(message string) {
	first.Warningf(message)
}

func FirstInfo(message string) {
	first.Infof(message)
}

func FirstDebug(message string) {
	first.Debugf(message)
}

func FirstTrace(message string) {
	first.Tracef(message)
}

func main() {
	loggo.ConfigureLoggers("first=DEBUG")
	FirstDebug("Test")
}
