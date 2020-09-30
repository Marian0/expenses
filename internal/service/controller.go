package service

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/zgoldy/expenses/internal/db"
)

// Controller -
type Controller struct {
	Repo	*db.Repository
}

// CreateController -
func createController(r *db.Repository) *Controller {
	return &Controller{
		Repo: r,
	}
}

// GetAllExpenses -
func (c *Controller) getAllExpenses(ctx *gin.Context) {
	expenses, err := c.Repo.FindAll()
	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":	http.StatusNotFound,
				"data":		expenses,
				"message":	"Data Not Found",
			})
	} else {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"status":	http.StatusOK,
				"data":		expenses,
				"message":	"Data Retrieved Successfully",
			})
	}
}

// GetExpenseByID -
func (c *Controller) getExpenseByID(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id"))
	
	expense, err := c.Repo.FindByID(ID)
	if err != nil {
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":	http.StatusNotFound,
				"data":		expense,
				"message":	"Data Not Found",
			})
	} else {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"status":	http.StatusOK,
				"data":		expense,
				"message":	"Data Retrieved Successfully",
			})
	}
}