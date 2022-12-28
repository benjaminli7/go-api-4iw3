package payment

type Service interface {
	GetAll() ([]Payment, error)
	GetByID(id int) (*Payment, error)
	Create(payment Payment) error
	Update(payment Payment) error
	Delete(id int) error
	StreamPayments(c chan<- Payment)
}

type service struct {
	repository  Repository
	broadcaster *Broadcaster
}

func NewService(r Repository, b *Broadcaster) *service {
	return &service{r, b}
}

func (s *service) StreamPayments(c chan<- Payment) {
	s.broadcaster.Add(c)
	defer s.broadcaster.Remove(c)
}

func (s *service) GetAll() ([]Payment, error) {
	payments, err := s.repository.GetAll()
	if err != nil {
		return payments, err
	}

	return payments, nil
}

func (s *service) GetByID(id int) (*Payment, error) {
	return s.repository.GetByID(id)
}

func (s *service) Create(payment Payment) (int, error) {
	// Vérifiez ici que le "productId" existe et que le "pricePaid" est supérieur ou égal au prix du produit
	// Si les règles ne sont pas respectées, renvoyez une erreur appropriée
	// Si les règles sont respectées, créez le paiement en appelant la fonction "Create" de la struct "repository"
	id := 0
	id, err := s.repository.Create(payment)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *service) Update(payment Payment) error {
	return s.repository.Update(payment)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) broadcastPayment(p Payment) {
	s.broadcaster.mu.Lock()
	defer s.broadcaster.mu.Unlock()

	for c := range s.broadcaster.clients {
		c <- p
	}
}
