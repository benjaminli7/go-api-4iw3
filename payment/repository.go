package payment

import (
	"errors"

	"github.com/benjaminli7/go-api-4iw3/product"
	"gorm.io/gorm"
)

type Repository interface {
	Store(payment Payment) (Payment, error)
	GetAll() ([]Payment, error)
	GetById(id int) (Payment, error)
	Update(id int, inputPayment InputPayment) (Payment, error)
	Delete(id int) error

	// Subscribe() <-chan Payment
}

type repository struct {
	db          *gorm.DB
	broadcaster Broadcaster
}

func NewRepository(db *gorm.DB, broadcaster Broadcaster) *repository {
	return &repository{db, broadcaster}
}

func (r *repository) Store(payment Payment) (Payment, error) {
	var product product.Product

	productErr := r.db.First(&product, payment.ProductId).Error
	if productErr != nil {
		return payment, productErr
	}

	payment.PricePaid = product.Price

	err := r.db.Create(&payment).Error
	if err != nil {
		return payment, err
	}
	r.broadcaster.Submit(payment)
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

func (r *repository) GetById(id int) (Payment, error) {
	var payment Payment
	err := r.db.First(&payment, id).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (r *repository) Update(id int, inputPayment InputPayment) (Payment, error) {
	var product product.Product

	payment, err := r.GetById(id)
	if err != nil {
		return payment, err
	}

	productErr := r.db.First(&product, inputPayment.ProductId).Error
	if productErr != nil {
		return payment, productErr
	}

	payment.ProductId = inputPayment.ProductId
	payment.PricePaid = product.Price

	err = r.db.Save(&payment).Error
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (r *repository) Delete(id int) error {
	payment, err := r.GetById(id)
	if err != nil {
		return err
	}

	tx := r.db.Delete(payment)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("Payment not found")
	}

	return nil
}

// func (r *repository) Subscribe() <-chan Payment {
// 	ch := make(chan Payment)
// 	go func() {
// 		for {
// 			var payment Payment
// 			r.db.First(&payment)
// 			ch <- payment
// 		}
// 	}()
// 	return ch
// }
