package checkout

import "simplecheckout/models"

//UseCase use case interface
type UseCase interface {
	Checkout(cart models.Cart) error
}
