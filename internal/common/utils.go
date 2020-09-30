package common

import "log"

func LogFatal(err error) {
	if err != nil {
		log.Fatal("Fatal error: " + err.Error())
	}
}