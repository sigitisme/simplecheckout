package product

import "simplecheckout/entity"

//Repository interface
type Repository interface {
	Create(e *entity.Product) (*entity.Product, error)
	Get(id string) (*entity.Product, error)
	GetByName(name string) (*entity.Product, error)
	GetBySKU(sku string) (*entity.Product, error)
	List() ([]*entity.Product, error)
	Update(e *entity.Product) error
	Delete(id string) error
}

//UseCase interface
type UseCase interface {
	CreateProduct(sku string, name string, price float64, quantity int) (*entity.Product, error)
	GetProduct(id string) (*entity.Product, error)
	GetProductByName(name string) (*entity.Product, error)
	GetProductBySKU(sku string) (*entity.Product, error)
	ListProducts() ([]*entity.Product, error)
	UpdateProduct(e *entity.Product) error
	DeleteProduct(id string) error
}
