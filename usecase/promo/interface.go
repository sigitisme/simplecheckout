package promo

import "simplecheckout/entity"

//Repository interface
type Repository interface {
	Create(e *entity.Promo) (*entity.Promo, error)
	Get(id string) (*entity.Promo, error)
	GetBySKU(sku string) (*entity.Promo, error)
	List() ([]*entity.Promo, error)
	Update(e *entity.Promo) error
	Delete(id string) error
}

//UseCase interface
type UseCase interface {
	CreatePromo(sku string, scheme string, freebieSKU string, minQty int, payOnly int, percentageOff int) (*entity.Promo, error)
	CreatePromoBuy1Get1(sku string, freebies string) (*entity.Promo, error)
	CreatePromoBuyXPayY(sku string, minqty int, payonly int) (*entity.Promo, error)
	CreatePromoBuyMinXOffY(sku string, minqty int, percentageoff int) (*entity.Promo, error)
	GetPromo(id string) (*entity.Promo, error)
	GetPromoBySKU(sku string) (*entity.Promo, error)
	ListPromos() ([]*entity.Promo, error)
	UpdatePromo(e *entity.Promo) error
	DeletePromo(id string) error
}
