package entity_test

import (
	"simplecheckout/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	b, err := entity.NewProduct("120P90", "Google Home", 49.99, 10)
	assert.Nil(t, err)
	assert.Equal(t, b.SKU, "120P90")
	assert.Equal(t, b.Name, "Google Home")
	assert.Equal(t, b.Price, 49.99)
	assert.Equal(t, b.Quantity, 10)
	assert.NotNil(t, b.ID)
}

func TestProductValidate(t *testing.T) {
	type test struct {
		sku   string
		name  string
		price float64
		qty   int
		want  error
	}

	tests := []test{
		{
			sku:   "120P90",
			name:  "Google Home",
			price: 49.99,
			qty:   10,
			want:  nil,
		},
		{
			sku:   "",
			name:  "Google Home",
			price: 49.99,
			qty:   10,
			want:  entity.ErrInvalidEntity,
		},
		{
			sku:   "120P90",
			name:  "",
			price: 49.99,
			qty:   10,
			want:  entity.ErrInvalidEntity,
		},
		{
			sku:   "120P90",
			name:  "Google Home",
			price: 0,
			qty:   10,
			want:  entity.ErrInvalidEntity,
		},
		{
			sku:   "120P90",
			name:  "Google Home",
			price: 49.99,
			qty:   0,
			want:  entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {
		_, err := entity.NewProduct(tc.sku, tc.name, tc.price, tc.qty)
		assert.Equal(t, err, tc.want)
	}
}
