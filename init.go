package main

import (
	"github.com/gin-gonic/gin"
)

func initializeService() *service {
	dbConn:= createDBConnection()
	
	migrate(dbConn)
	
	expensesRepository := createRepository(dbConn)
	
	engine := gin.Default()
	expensesController := createController(expensesRepository)
	
	svc := createService(engine, expensesController)

	svc.setupEndpoints()

	return svc
}