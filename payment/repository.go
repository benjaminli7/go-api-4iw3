package payment

import (
	// "errors"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]Payment, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(payment Payment) (Payment, error) {
	err := r.db.Create(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (r *repository) GetAll() ([]Payment, error) {
	var payments []Payment
	err := r.db.Find(&payments).Error
	if err != nil {
		return payments, err
	}

	return payments, nil
}