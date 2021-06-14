package entity

import "time"

// Buy1Get1 ...
var Buy1Get1 = "Buy1Get1"

// BuyXPayY ...
var BuyXPayY = "BuyXPayY"

// BuyMinXOffY ...
var BuyMinXOffY = "BuyMinXOffY"

// Promo is a representation of Promo
type Promo struct {
	ID            string
	SKU           string
	Scheme        string
	FreebieSKU    string
	MinQty        int
	PayOnly       int
	PercentageOff int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

//NewPromo creates a new Promo
func NewPromo(sku string, scheme string, freebieSKU string, minQty int, payOnly int, percentageOff int) (*Promo, error) {
	p := &Promo{
		ID:            NewIDString(),
		SKU:           sku,
		Scheme:        scheme,
		FreebieSKU:    freebieSKU,
		MinQty:        minQty,
		PayOnly:       payOnly,
		PercentageOff: percentageOff,
		CreatedAt:     time.Now(),
	}
	err := p.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return p, nil
}

//NewPromoBuy1Get1 creates a new Promo
func NewPromoBuy1Get1(sku string, freebieSKU string) (*Promo, error) {
	p := &Promo{
		ID:         NewIDString(),
		SKU:        sku,
		Scheme:     Buy1Get1,
		FreebieSKU: freebieSKU,
		CreatedAt:  time.Now(),
	}
	err := p.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return p, nil
}

//NewPromoBuyMinXOffY creates a new Promo
func NewPromoBuyMinXOffY(sku string, minQty int, percentageOff int) (*Promo, error) {
	p := &Promo{
		ID:            NewIDString(),
		SKU:           sku,
		Scheme:        BuyMinXOffY,
		MinQty:        minQty,
		PercentageOff: percentageOff,
		CreatedAt:     time.Now(),
	}
	err := p.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return p, nil
}

//NewPromoBuyXPayY creates a new Promo
func NewPromoBuyXPayY(sku string, minQty int, payOnly int) (*Promo, error) {
	p := &Promo{
		ID:        NewIDString(),
		SKU:       sku,
		Scheme:    BuyXPayY,
		MinQty:    minQty,
		PayOnly:   payOnly,
		CreatedAt: time.Now(),
	}
	err := p.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return p, nil
}

//Validate is to validate a Promo params
func (p *Promo) Validate() error {
	if p.SKU == "" || p.Scheme == "" {
		return ErrInvalidEntity
	}

	switch p.Scheme {
	case Buy1Get1:
		if p.FreebieSKU == "" {
			return ErrInvalidEntity
		}
	case BuyXPayY:
		if p.MinQty == 0 || p.PayOnly == 0 {
			return ErrInvalidEntity
		}
	case BuyMinXOffY:
		if p.MinQty == 0 || p.PercentageOff == 0 {
			return ErrInvalidEntity
		}
	default:
		return ErrInvalidEntity
	}

	return nil
}
