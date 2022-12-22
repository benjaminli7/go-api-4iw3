package payment

type Service interface {
	GetAll() ([]Payment, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) GetAll() ([]Payment, error) {
	payments, err := s.repository.GetAll()
	if err != nil {
		return payments, err
	}

	return payments, nil
}