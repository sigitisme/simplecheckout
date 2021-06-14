package promo

import (
	"simplecheckout/entity"
)

//Inmem in memory repo
type Inmem struct {
	m map[string]*entity.Promo
}

//NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*entity.Promo{}
	return &Inmem{
		m: m,
	}
}

//Create a Promo
func (r *Inmem) Create(e *entity.Promo) (*entity.Promo, error) {
	r.m[e.ID] = e
	return e, nil
}

//Get a Promo
func (r *Inmem) Get(id string) (*entity.Promo, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//GetBySKU Get a Promo By SKU
func (r *Inmem) GetBySKU(sku string) (*entity.Promo, error) {
	for _, v := range r.m {
		if v.SKU == sku {
			return v, nil
		}
	}
	return nil, entity.ErrNotFound
}

//List Promos
func (r *Inmem) List() ([]*entity.Promo, error) {
	var d []*entity.Promo
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Update a Promo
func (r *Inmem) Update(e *entity.Promo) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Delete a Promo
func (r *Inmem) Delete(id string) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
