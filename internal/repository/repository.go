package repository

import (
	"context"
	"gorm.io/gorm"
	"alexdenkk/test-task/internal/model"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (repository *Repository) CreateNumber(ctx context.Context, number model.Number) error {
	return repository.DB.Create(&number).Error
}

func (repository *Repository) GetAllNumbers(ctx context.Context) ([]model.Number, error) {
	var numbers []model.Number
	result := repository.DB.Find(&numbers)
	return numbers, result.Error
}
