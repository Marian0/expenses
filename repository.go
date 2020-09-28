package main

import (
	"errors"
	"gorm.io/gorm/clause"
	"gorm.io/gorm"
)

type repository struct {
	dbConnection	*gorm.DB
}

func createRepository(dbConn *gorm.DB) *repository {
	return &repository{
		dbConnection: dbConn,
	}
}

func (r *repository) findAll() (*[]expense, error) {
	var expenses []expense
	
	result := r.dbConnection.Preload(clause.Associations).Find(&expenses)
	if result.Error != nil {
		return nil, result.Error
	}

	return &expenses, nil
}

func (r *repository) findByID(ID int) (*expense, error) {
	var expense expense
	
	result := r.dbConnection.Preload(clause.Associations).First(&expense, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return &expense, nil
}