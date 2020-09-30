package service

import (
	"github.com/gin-gonic/gin"
	"github.com/zgoldy/expenses/internal/db"
)

// InitializeService -
func InitializeService() *Engine {
	dbConn:= db.CreateDBConnection()
	
	db.Migrate(dbConn)
	
	expensesRepository := db.CreateRepository(dbConn)
	
	router := gin.Default()
	expensesController := createController(expensesRepository)
	
	svc := createEngine(router, expensesController)

	svc.setupEndpoints()

	return svc
}