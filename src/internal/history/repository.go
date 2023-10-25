package history

import (
	"go-base-template/src/model"

	"gorm.io/gorm"
)

type Repository interface {
	// example repo
	Create(history model.History) (model.History, error)
}

type repository struct {
	collection *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{}
}

// example repo
func (r *repository) Create(history model.History) (model.History, error) {
	return history, nil
}
