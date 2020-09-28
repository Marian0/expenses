package main

import (
	"os"
	
	"github.com/gin-gonic/gin"
)

type service struct {
	engine 		*gin.Engine
	controller	*controller
}

func createService(e *gin.Engine, c *controller) *service {
	return &service{
		engine:		e,
		controller:	c,
	}
}

func (svc *service) setupEndpoints() {
	apiEndpoints := svc.engine.Group("/api/")
	
	apiEndpoints.GET("expenses", svc.controller.getAllExpenses)
	apiEndpoints.GET("expenses/:id", svc.controller.getExpenseByID)
}

func (svc *service) run() {
	svc.engine.Run(":" + os.Getenv("APP_PORT"))
}
