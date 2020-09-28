package main

import (
	"github.com/joho/godotenv"
)

func init() { 
	err := godotenv.Load()
	logFatal(err)
}

func main() {
	expensesSVC := initializeService()
	
	expensesSVC.run()
}