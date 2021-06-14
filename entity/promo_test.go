package entity_test

import (
	"simplecheckout/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPromo(t *testing.T) {
	b, err := entity.NewPromo("43N23P", entity.Buy1Get1, "234234", 0, 0, 0)
	assert.Nil(t, err)
	assert.Equal(t, b.SKU, "43N23P")
	assert.Equal(t, b.Scheme, entity.Buy1Get1)
	assert.Equal(t, b.FreebieSKU, "234234")
	assert.NotNil(t, b.ID)
}

func TestNewPromoBuy1Get1(t *testing.T) {
	b, err := entity.NewPromoBuy1Get1("43N23P", "234234")
	assert.Nil(t, err)
	assert.Equal(t, b.SKU, "43N23P")
	assert.Equal(t, b.Scheme, entity.Buy1Get1)
	assert.Equal(t, b.FreebieSKU, "234234")
	assert.NotNil(t, b.ID)
}

func TestNewPromoBuyXPayY(t *testing.T) {
	b, err := entity.NewPromoBuyXPayY("120P90", 3, 2)
	assert.Nil(t, err)
	assert.Equal(t, b.SKU, "120P90")
	assert.Equal(t, b.Scheme, entity.BuyXPayY)
	assert.Equal(t, b.MinQty, 3)
	assert.Equal(t, b.PayOnly, 2)
	assert.NotNil(t, b.ID)
}

func TestNewPromoBuyMinXOffY(t *testing.T) {
	b, err := entity.NewPromoBuyMinXOffY("A304SD", 3, 10)
	assert.Nil(t, err)
	assert.Equal(t, b.SKU, "A304SD")
	assert.Equal(t, b.Scheme, entity.BuyMinXOffY)
	assert.Equal(t, b.MinQty, 3)
	assert.Equal(t, b.PercentageOff, 10)
	assert.NotNil(t, b.ID)
}

func TestPromoValidate(t *testing.T) {
	type test struct {
		sku           string
		scheme        string
		freebiesku    string
		minqty        int
		payonly       int
		percentageoff int
		want          error
	}

	tests := []test{
		{
			sku:        "43N23P",
			scheme:     entity.Buy1Get1,
			freebiesku: "234234",
			want:       nil,
		},
		{
			sku:        "43N23P",
			scheme:     entity.Buy1Get1,
			freebiesku: "",
			want:       entity.ErrInvalidEntity,
		},
		{
			sku:        "",
			scheme:     entity.Buy1Get1,
			freebiesku: "234234",
			want:       entity.ErrInvalidEntity,
		},
		{
			sku:     "43N23P",
			scheme:  entity.BuyXPayY,
			minqty:  3,
			payonly: 2,
			want:    nil,
		},
		{
			sku:     "120P90",
			scheme:  entity.BuyXPayY,
			minqty:  3,
			payonly: 2,
			want:    nil,
		},
		{
			sku:     "120P90",
			scheme:  entity.BuyXPayY,
			minqty:  0,
			payonly: 2,
			want:    entity.ErrInvalidEntity,
		},
		{
			sku:     "120P90",
			scheme:  entity.BuyXPayY,
			minqty:  3,
			payonly: 0,
			want:    entity.ErrInvalidEntity,
		},
		{
			sku:           "A304SD",
			scheme:        entity.BuyMinXOffY,
			minqty:        3,
			percentageoff: 10,
			want:          nil,
		},
		{
			sku:           "A304SD",
			scheme:        entity.BuyMinXOffY,
			minqty:        0,
			percentageoff: 10,
			want:          entity.ErrInvalidEntity,
		},
		{
			sku:           "A304SD",
			scheme:        entity.BuyMinXOffY,
			minqty:        3,
			percentageoff: 0,
			want:          entity.ErrInvalidEntity,
		},
		{
			sku:           "A304SD",
			scheme:        "InvalidPromo",
			minqty:        3,
			percentageoff: 10,
			want:          entity.ErrInvalidEntity,
		},
	}
	for _, tc := range tests {
		_, err := entity.NewPromo(tc.sku, tc.scheme, tc.freebiesku, tc.minqty, tc.payonly, tc.percentageoff)
		assert.Equal(t, err, tc.want)
	}
}
