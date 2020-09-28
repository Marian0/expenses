package main

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
)

type controller struct {
	repo	*repository
}

func createController(r *repository) *controller {
	return &controller{
		repo: r,
	}
}

func (c *controller) getAllExpenses(ctx *gin.Context) {
	expenses, err := c.repo.findAll()
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

func (c *controller) getExpenseByID(ctx *gin.Context) {
	ID, _ := strconv.Atoi(ctx.Param("id"))
	
	expense, err := c.repo.findByID(ID)
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