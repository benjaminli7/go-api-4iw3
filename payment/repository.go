package payment

import (
	// "errors"
	"gorm.io/gorm"
)

type Repository interface {
	GetAll() ([]Payment, error)
	GetByID(id int) (*Payment, error)
	Create(payment Payment) (Payment, error)
	Store(payment Payment) (Payment, error)
	Update(payment Payment) error
	Delete(id int) error
	Subscribe() <-chan Payment
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

func (r *repository) GetByID(id int) (*Payment, error) {
	var payment Payment
	err := r.db.Find(&payment, id).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *repository) Create(payment Payment) (int, error) {
	err := r.db.Create(&payment).Error
	if err != nil {
		return 0, err
	}
	return payment.ID, nil
}

func (r *repository) Update(payment Payment) error {
	return r.db.Save(&payment).Error
}

func (r *repository) Delete(id int) error {
	return r.db.Delete(&Payment{ID: id}).Error
}

func (r *repository) Subscribe() <-chan Payment {
	ch := make(chan Payment)
	go func() {
		for {
			var payment Payment
			r.db.First(&payment)
			ch <- payment
		}
	}()
	return ch
}
