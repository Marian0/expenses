package service

import (
	"os"
	
	"github.com/gin-gonic/gin"
)

// Engine -
type Engine struct {
	router 		*gin.Engine
	controller	*controller
}

// CreateEngine -
func createEngine(e *gin.Engine, c *controller) *Engine {
	return &Engine{
		router:		e,
		controller:	c,
	}
}

// SetupEndpoints -
func (eng *Engine) setupEndpoints() {
	apiEndpoints := eng.router.Group("/api/")
	
	apiEndpoints.GET("expenses", eng.controller.getAllExpenses)
	apiEndpoints.GET("expenses/:id", eng.controller.getExpenseByID)
}

// Run -
func (eng *Engine) Run() {
	eng.router.Run(":" + os.Getenv("APP_PORT"))
}
