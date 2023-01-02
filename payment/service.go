package payment

import (
	"github.com/benjaminli7/go-api-4iw3/product"

)

type Service interface {
	Store(payment Payment) (Payment, error)
	GetAll() ([]Payment, error)
	// GetByID(id int) (Payment, error)
	// Update(payment Payment) error
	// Delete(id int) error
	// StreamPayments(c chan<- Payment)
}

type service struct {
	repository  Repository
	broadcaster *Broadcaster
}

func NewService(r Repository, b *Broadcaster) *service {
	return &service{r, b}
}


func (s *service) Store(input InputPayment) (Payment, error) {
	var payment Payment
	payment.ProductId = input.ProductId

	test, err := product.Service.GetById(input.ProductId)

	
	newPayment, err := s.repository.Store(payment)
	if err != nil {
		return newPayment, err
	}

	return newPayment, nil
}

func (s *service) GetAll() ([]Payment, error) {
	payments, err := s.repository.GetAll()
	if err != nil {
		return payments, err
	}

	return payments, nil
}

// func (s *service) GetByID(id int) (*Payment, error) {
// 	return s.repository.GetByID(id)
// }

// func (s *service) Store(payment Payment) (uint32, error) {
// 	p, err := s.repository.Store(payment)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return p.ID, nil
// }

// func (s *service) Update(payment Payment) error {
// 	return s.repository.Update(payment)
// }

// func (s *service) Delete(id int) error {
// 	return s.repository.Delete(id)
// }

// func (s *service) broadcastPayment(p Payment) {
// 	s.broadcaster.mu.Lock()
// 	defer s.broadcaster.mu.Unlock()

// 	for c := range s.broadcaster.clients {
// 		c <- p
// 	}
// }
