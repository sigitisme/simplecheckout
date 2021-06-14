package product_test

import (
	"simplecheckout/entity"
	"simplecheckout/usecase/product"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	r := product.NewInmem()

	macbook := &entity.Product{
		ID:        entity.NewIDString(),
		SKU:       "43N23P",
		Name:      "Macbook Pro",
		Price:     5399.99,
		Quantity:  5,
		CreatedAt: time.Now(),
	}

	p := product.NewService(r)

	_, err := p.CreateProduct(macbook.SKU, macbook.Name, macbook.Price, macbook.Quantity)

	assert.Nil(t, err)
	assert.False(t, macbook.CreatedAt.IsZero())
	assert.True(t, macbook.UpdatedAt.IsZero())
}

func TestGet(t *testing.T) {
	r := product.NewInmem()

	macbook := &entity.Product{
		ID:       entity.NewIDString(),
		SKU:      "43N23P",
		Name:     "Macbook Pro",
		Price:    5399.99,
		Quantity: 5,
	}

	raspberry := &entity.Product{
		ID:       entity.NewIDString(),
		SKU:      "234234",
		Name:     "Raspberry Pi B",
		Price:    30.00,
		Quantity: 2,
	}

	googleHome := &entity.Product{
		ID:       entity.NewIDString(),
		SKU:      "120P90",
		Name:     "Google Home",
		Price:    49.99,
		Quantity: 10,
	}

	alexa := &entity.Product{
		ID:       entity.NewIDString(),
		SKU:      "A304SD",
		Name:     "Alexa Speaker",
		Price:    109.50,
		Quantity: 10,
	}

	p := product.NewService(r)

	product, _ := p.CreateProduct(macbook.SKU, macbook.Name, macbook.Price, macbook.Quantity)
	p.CreateProduct(raspberry.SKU, raspberry.Name, raspberry.Price, raspberry.Quantity)
	p.CreateProduct(googleHome.SKU, googleHome.Name, googleHome.Price, googleHome.Quantity)
	p.CreateProduct(alexa.SKU, alexa.Name, alexa.Price, alexa.Quantity)

	t.Run("get", func(t *testing.T) {
		saved, err := p.GetProduct(product.ID)
		assert.Nil(t, err)
		assert.Equal(t, macbook.Name, saved.Name)
	})
}
