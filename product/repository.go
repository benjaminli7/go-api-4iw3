package product

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Store(product Product) (Product, error)
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Update(id int, inputProduct InputProduct) (Product, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(product Product) (Product, error) {
	errFind := r.db.First(&Product{}, "name = ?", product.Name).Error
	if errFind == nil {
		return product, errors.New("Product with this name already exists")
	}
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

func (r *repository) Update(id int, inputProduct InputProduct) (Product, error) {
	errFind := r.db.First(&Product{}, "name = ?", inputProduct.Name).Error
	if errFind == nil {
		return Product{}, errors.New("Product with this name already exists")
	}
	product, err := r.GetById(id)
	if err != nil {
		return product, err
	}

	product.Name = inputProduct.Name
	product.Price = inputProduct.Price

	err = r.db.Save(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *repository) Delete(id int) error {
	product, err := r.GetById(id)
	if err != nil {
		return err
	}

	tx := r.db.Delete(product)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("Product not found")
	}

	return nil
}
