package checkout

import (
	"simplecheckout/entity"
	"simplecheckout/models"
	"simplecheckout/usecase/product"
	"simplecheckout/usecase/promo"
)

// Service of checkout
type Service struct {
	productservice product.UseCase
	promoservice   promo.UseCase
	Total          float64
}

//NewService create new service
func NewService(p product.UseCase, pr promo.UseCase) *Service {
	return &Service{
		productservice: p,
		promoservice:   pr,
	}
}

// Checkout process flow..
func (s *Service) Checkout(cart *models.Cart) error {

	added := make(map[string]float64)
	freebies := make(map[string]int)

	for _, c := range cart.Contents {
		product, err := s.productservice.GetProductBySKU(c.SKU)

		if err == nil {
			//check product if any promo
			var totalPrices float64
			var qty = c.Quantity
			var price = product.Price
			var sku = product.SKU

			promo, err := s.promoservice.GetPromoBySKU(sku)

			if err == nil {
				//calculate promo
				switch promo.Scheme {
				case entity.Buy1Get1:
					freebies[promo.FreebieSKU] = qty

					if m, ok := added[promo.FreebieSKU]; ok {
						totalPrices -= float64(qty) * m
					}
				case entity.BuyXPayY:
					//buy 3 pay 2,
					//buy 4 => buy 3 pay 2, buy 1 pay 1, so buy 4 pay 3
					//buy 6 pay 4
					if qty == promo.MinQty {
						qty = promo.PayOnly
					}

					div := qty / promo.MinQty //ex: qty=6, div=2
					rem := qty % promo.MinQty //ex: qty=6, rem=0

					if div >= 1 {
						qty = div*promo.PayOnly + rem //2*2 + 0 = 4. buy 6 pay 4
					}

				case entity.BuyMinXOffY:
					if qty == promo.MinQty {
						percentageAfterDiscount := float64(1 - float64(promo.PercentageOff)*0.01)
						price = percentageAfterDiscount * price
					}
				}
			}

			//check freebies
			if m, ok := freebies[sku]; ok {
				//if freebies found and match sku, substract qty with quantity freebies
				qty -= m
			}

			added[sku] = price

			//calculate prices per product order
			totalPrices += float64(qty) * price

			//sum of all orders
			s.Total += totalPrices
		}
	}

	return nil
}
