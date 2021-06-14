package entity

import "time"

// Product is a representation of product
type Product struct {
	ID        string
	SKU       string
	Name      string
	Price     float64
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

//NewProduct creates a new product
func NewProduct(sku string, name string, price float64, qty int) (*Product, error) {
	p := &Product{
		ID:        NewIDString(),
		SKU:       sku,
		Name:      name,
		Price:     price,
		Quantity:  qty,
		CreatedAt: time.Now(),
	}
	err := p.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return p, nil
}

//Validate is to validate a product params
func (p *Product) Validate() error {
	if p.SKU == "" || p.Name == "" || p.Price == 0 || p.Quantity == 0 {
		return ErrInvalidEntity
	}
	return nil
}
