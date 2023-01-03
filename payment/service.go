package payment

type Service interface {
	Store(input InputPayment) (Payment, error)
	GetAll() ([]Payment, error)
	GetById(id int) (Payment, error)
	Update(id int, inputPayment InputPayment) (Payment, error)
	Delete(id int) error

	// StreamPayments(c chan<- Payment)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) Store(input InputPayment) (Payment, error) {
	var payment Payment

	payment.ProductId = input.ProductId

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

func (s *service) GetById(id int) (Payment, error) {
	payment, err := s.repository.GetById(id)
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (s *service) Update(id int, input InputPayment) (Payment, error) {
	payment, err := s.repository.Update(id, input)
	if err != nil {
		return payment, err
	}

	return payment, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

// func (s *service) broadcastPayment(p Payment) {
// 	s.broadcaster.mu.Lock()
// 	defer s.broadcaster.mu.Unlock()

// 	for c := range s.broadcaster.clients {
// 		c <- p
// 	}
// }
