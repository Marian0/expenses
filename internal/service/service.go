package service

import (
	"os"
	
	"github.com/gin-gonic/gin"
)

// Engine -
type Engine struct {
	Router 		*gin.Engine
	Controller	*controller
}

// CreateEngine -
func createEngine(e *gin.Engine, c *controller) *Engine {
	return &Engine{
		Router:		e,
		Controller:	c,
	}
}

// SetupEndpoints -
func (eng *Engine) setupEndpoints() {
	apiEndpoints := eng.Router.Group("/api/")
	
	apiEndpoints.GET("expenses", eng.Controller.getAllExpenses)
	apiEndpoints.GET("expenses/:id", eng.Controller.getExpenseByID)
}

// Run -
func (eng *Engine) Run() {
	eng.Router.Run(":" + os.Getenv("APP_PORT"))
}
