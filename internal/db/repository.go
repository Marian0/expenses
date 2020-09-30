package db

import (
	"errors"
	"gorm.io/gorm/clause"
	"gorm.io/gorm"
	"github.com/zgoldy/expenses/internal/models"
)

// Repository -
type Repository struct {
	DBConnection	*gorm.DB
}

// CreateRepository -
func CreateRepository(dbConn *gorm.DB) *Repository {
	return &Repository{
		DBConnection: dbConn,
	}
}

// FindAll -
func (r *Repository) FindAll() (*[]models.Expense, error) {
	var expenses []models.Expense
	
	result := r.DBConnection.Preload(clause.Associations).Find(&expenses)
	if result.Error != nil {
		return nil, result.Error
	}

	return &expenses, nil
}

// FindByID -
func (r *Repository) FindByID(ID int) (*models.Expense, error) {
	var expense models.Expense
	
	result := r.DBConnection.Preload(clause.Associations).First(&expense, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return &expense, nil
}