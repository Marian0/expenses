package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/zgoldy/expenses/internal/common"
	"github.com/zgoldy/expenses/internal/db"
)

func init() { 
	err := godotenv.Load()
	common.LogFatal(err)
}
func main() {
	dbConnection := db.CreateDBConnection()

	fmt.Println("Starting the migration process...")
	db.Migrate(dbConnection)
	fmt.Println("Migration process has been finished successfully... :D")
}