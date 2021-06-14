package product

import (
	"simplecheckout/entity"
)

//Inmem in memory repo
type Inmem struct {
	m map[string]*entity.Product
}

//NewInmem create new repository
func NewInmem() *Inmem {
	var m = map[string]*entity.Product{}
	return &Inmem{
		m: m,
	}
}

//Create a Product
func (r *Inmem) Create(e *entity.Product) (*entity.Product, error) {
	r.m[e.ID] = e
	return e, nil
}

//Get a Product
func (r *Inmem) Get(id string) (*entity.Product, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//GetByName Get a Product By Name
func (r *Inmem) GetByName(name string) (*entity.Product, error) {
	for _, v := range r.m {
		if v.Name == name {
			return v, nil
		}
	}
	return nil, entity.ErrNotFound
}

//GetBySKU Get a Product By Name
func (r *Inmem) GetBySKU(sku string) (*entity.Product, error) {
	for _, v := range r.m {
		if v.SKU == sku {
			return v, nil
		}
	}
	return nil, entity.ErrNotFound
}

//List products
func (r *Inmem) List() ([]*entity.Product, error) {
	var d []*entity.Product
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

//Update a Product
func (r *Inmem) Update(e *entity.Product) error {
	_, err := r.Get(e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}

//Delete a Product
func (r *Inmem) Delete(id string) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}
