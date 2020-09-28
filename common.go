package main

import "log"

func logFatal(err error) {
	if err != nil {
		log.Fatal("Fatal error: " + err.Error())
	}
}