package product

import (
	// "errors"
	"gorm.io/gorm"
)

type Repository interface {
	Store(product Product) (Product, error)
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) GetById(id int) (Product, error) {
	var product Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return product, err
	}

	return product, nil
}