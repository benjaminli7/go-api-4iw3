package product

type Service interface {
	Store(input InputProduct) (Product, error)
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) Store(input InputProduct) (Product, error) {
	var product Product
	product.Name = input.Name
	product.Price = input.Price
	newProduct, err := s.repository.Store(product)
	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) GetById(id int) (Product, error) {
	product, err := s.repository.GetById(id)
	if err != nil {
		return product, err
	}

	return product, nil
}