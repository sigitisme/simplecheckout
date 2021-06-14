package product

import (
	"simplecheckout/entity"
	"time"
)

// Service of product
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateProduct create a Product
func (s *Service) CreateProduct(sku string, name string, price float64, quantity int) (*entity.Product, error) {
	p, err := entity.NewProduct(sku, name, price, quantity)
	if err != nil {
		return p, err
	}
	return s.repo.Create(p)
}

//GetProduct Get a product
func (s *Service) GetProduct(id string) (*entity.Product, error) {
	return s.repo.Get(id)
}

// GetProductByName get product by name
func (s *Service) GetProductByName(name string) (*entity.Product, error) {
	p, err := s.repo.GetByName(name)
	if err != nil {
		return p, err
	}

	return p, nil
}

// GetProductBySKU get product by sku
func (s *Service) GetProductBySKU(sku string) (*entity.Product, error) {
	p, err := s.repo.GetBySKU(sku)
	if err != nil {
		return p, err
	}

	return p, nil
}

//ListProducts List Products
func (s *Service) ListProducts() ([]*entity.Product, error) {
	return s.repo.List()
}

//DeleteProduct Delete an Product
func (s *Service) DeleteProduct(id string) error {
	u, err := s.GetProduct(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

//UpdateProduct Update an Product
func (s *Service) UpdateProduct(e *entity.Product) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
