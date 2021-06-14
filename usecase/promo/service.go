package promo

import (
	"simplecheckout/entity"
	"time"
)

// Service of Promo
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreatePromo create a Promo
func (s *Service) CreatePromo(sku string, scheme string, freebieSKU string, minQty int, payOnly int, percentageOff int) (*entity.Promo, error) {
	p, err := entity.NewPromo(sku, scheme, freebieSKU, minQty, payOnly, percentageOff)
	if err != nil {
		return p, err
	}
	return s.repo.Create(p)
}

//CreatePromoBuy1Get1 create a Promo
func (s *Service) CreatePromoBuy1Get1(sku string, freebies string) (*entity.Promo, error) {
	p, err := entity.NewPromoBuy1Get1(sku, freebies)
	if err != nil {
		return p, err
	}
	return s.repo.Create(p)
}

//CreatePromoBuyXPayY create a Promo
func (s *Service) CreatePromoBuyXPayY(sku string, minqty int, payonly int) (*entity.Promo, error) {
	p, err := entity.NewPromoBuyXPayY(sku, minqty, payonly)
	if err != nil {
		return p, err
	}
	return s.repo.Create(p)
}

//CreatePromoBuyMinXOffY create a Promo
func (s *Service) CreatePromoBuyMinXOffY(sku string, minqty int, percentageoff int) (*entity.Promo, error) {
	p, err := entity.NewPromoBuyMinXOffY(sku, minqty, percentageoff)
	if err != nil {
		return p, err
	}
	return s.repo.Create(p)
}

//GetPromo Get a Promo
func (s *Service) GetPromo(id string) (*entity.Promo, error) {
	return s.repo.Get(id)
}

// GetPromoBySKU get promo by sku
func (s *Service) GetPromoBySKU(sku string) (*entity.Promo, error) {
	p, err := s.repo.GetBySKU(sku)
	if err != nil {
		return p, err
	}

	return p, nil
}

//ListPromos List Promos
func (s *Service) ListPromos() ([]*entity.Promo, error) {
	return s.repo.List()
}

//DeletePromo Delete an Promo
func (s *Service) DeletePromo(id string) error {
	u, err := s.GetPromo(id)
	if u == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

//UpdatePromo Update an Promo
func (s *Service) UpdatePromo(e *entity.Promo) error {
	err := e.Validate()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}
