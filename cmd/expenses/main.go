package main

import (
	"github.com/joho/godotenv"
	"github.com/zgoldy/expenses/internal/common"
	"github.com/zgoldy/expenses/internal/service"
)

func init() { 
	err := godotenv.Load()
	common.LogFatal(err)
}

func main() {
	expensesSVC := service.InitializeService()
	
	expensesSVC.Run()
}