package checkout_test

import (
	"simplecheckout/entity"
	"simplecheckout/models"
	"simplecheckout/usecase/checkout"
	pmock "simplecheckout/usecase/product/mock"
	prmock "simplecheckout/usecase/promo/mock"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestServiceCheckout(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	pMock := pmock.NewMockUseCase(controller)
	prMock := prmock.NewMockUseCase(controller)

	t.Run("Macbook Pro (1 unit) - Raspberry Pi B (1 unit)", func(t *testing.T) {

		uc := checkout.NewService(pMock, prMock)

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

		promoMacbook, _ := entity.NewPromoBuy1Get1(macbook.SKU, raspberry.SKU)

		cart := &models.Cart{
			Contents: []*models.Content{
				&models.Content{Name: macbook.Name, SKU: macbook.SKU, Quantity: 1},
				&models.Content{Name: raspberry.Name, SKU: raspberry.SKU, Quantity: 1},
			},
		}

		pMock.EXPECT().GetProductBySKU(macbook.SKU).Return(macbook, nil)
		pMock.EXPECT().GetProductBySKU(raspberry.SKU).Return(raspberry, nil)

		prMock.EXPECT().GetPromoBySKU(macbook.SKU).Return(promoMacbook, nil)
		prMock.EXPECT().GetPromoBySKU(raspberry.SKU).Return(nil, entity.ErrInvalidEntity)

		err := uc.Checkout(cart)
		assert.Nil(t, err)
		assert.Equal(t, 5399.99, uc.Total)
	})

	t.Run("Macbook Pro (1 unit) - Raspberry Pi B (1 unit)", func(t *testing.T) {

		uc := checkout.NewService(pMock, prMock)

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

		promoMacbook, _ := entity.NewPromoBuy1Get1(macbook.SKU, raspberry.SKU)

		cart := &models.Cart{
			Contents: []*models.Content{
				&models.Content{Name: raspberry.Name, SKU: raspberry.SKU, Quantity: 1},
				&models.Content{Name: macbook.Name, SKU: macbook.SKU, Quantity: 1},
			},
		}

		pMock.EXPECT().GetProductBySKU(macbook.SKU).Return(macbook, nil)
		pMock.EXPECT().GetProductBySKU(raspberry.SKU).Return(raspberry, nil)

		prMock.EXPECT().GetPromoBySKU(macbook.SKU).Return(promoMacbook, nil)
		prMock.EXPECT().GetPromoBySKU(raspberry.SKU).Return(nil, entity.ErrInvalidEntity)

		err := uc.Checkout(cart)
		assert.Nil(t, err)
		assert.Equal(t, 5399.99, uc.Total)
	})

	t.Run("Google Home (3 units) -  Buy 3 Pay 2", func(t *testing.T) {

		uc := checkout.NewService(pMock, prMock)

		googleHome := &entity.Product{
			ID:       entity.NewIDString(),
			SKU:      "120P90",
			Name:     "Google Home",
			Price:    49.99,
			Quantity: 10,
		}

		promoGoogleHome, _ := entity.NewPromoBuyXPayY(googleHome.SKU, 3, 2)

		cart := &models.Cart{
			Contents: []*models.Content{
				&models.Content{Name: googleHome.Name, SKU: googleHome.SKU, Quantity: 3},
			},
		}

		pMock.EXPECT().GetProductBySKU(googleHome.SKU).Return(googleHome, nil)
		prMock.EXPECT().GetPromoBySKU(googleHome.SKU).Return(promoGoogleHome, nil)

		err := uc.Checkout(cart)
		assert.Nil(t, err)
		assert.Equal(t, 99.98, uc.Total)
	})

	t.Run("Google Home (4 units) -  Buy 4 Pay 3", func(t *testing.T) {

		uc := checkout.NewService(pMock, prMock)

		googleHome := &entity.Product{
			ID:       entity.NewIDString(),
			SKU:      "120P90",
			Name:     "Google Home",
			Price:    49.99,
			Quantity: 10,
		}

		promoGoogleHome, _ := entity.NewPromoBuyXPayY(googleHome.SKU, 3, 2)

		cart := &models.Cart{
			Contents: []*models.Content{
				&models.Content{Name: googleHome.Name, SKU: googleHome.SKU, Quantity: 4},
			},
		}

		pMock.EXPECT().GetProductBySKU(googleHome.SKU).Return(googleHome, nil)
		prMock.EXPECT().GetPromoBySKU(googleHome.SKU).Return(promoGoogleHome, nil)

		err := uc.Checkout(cart)
		assert.Nil(t, err)
		assert.Equal(t, 149.97, uc.Total)
	})

	t.Run("Google Home (6 units) -  Buy 6 Pay 4", func(t *testing.T) {

		uc := checkout.NewService(pMock, prMock)

		googleHome := &entity.Product{
			ID:       entity.NewIDString(),
			SKU:      "120P90",
			Name:     "Google Home",
			Price:    49.99,
			Quantity: 10,
		}

		promoGoogleHome, _ := entity.NewPromoBuyXPayY(googleHome.SKU, 3, 2)

		cart := &models.Cart{
			Contents: []*models.Content{
				&models.Content{Name: googleHome.Name, SKU: googleHome.SKU, Quantity: 6},
			},
		}

		pMock.EXPECT().GetProductBySKU(googleHome.SKU).Return(googleHome, nil)
		prMock.EXPECT().GetPromoBySKU(googleHome.SKU).Return(promoGoogleHome, nil)

		err := uc.Checkout(cart)
		assert.Nil(t, err)
		assert.Equal(t, 199.96, uc.Total)
	})

	t.Run("Alexa Speaker (3 units)", func(t *testing.T) {

		uc := checkout.NewService(pMock, prMock)

		alexa := &entity.Product{
			ID:       entity.NewIDString(),
			SKU:      "A304SD",
			Name:     "Alexa Speaker",
			Price:    109.50,
			Quantity: 10,
		}

		promoAlexa, _ := entity.NewPromoBuyMinXOffY(alexa.SKU, 3, 10)

		cart := &models.Cart{
			Contents: []*models.Content{
				&models.Content{Name: alexa.Name, SKU: alexa.SKU, Quantity: 3},
			},
		}

		pMock.EXPECT().GetProductBySKU(alexa.SKU).Return(alexa, nil)
		prMock.EXPECT().GetPromoBySKU(alexa.SKU).Return(promoAlexa, nil)

		err := uc.Checkout(cart)
		assert.Nil(t, err)
		assert.Equal(t, 295.65, uc.Total)
	})
}
