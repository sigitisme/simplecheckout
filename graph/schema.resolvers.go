package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"simplecheckout/graph/generated"
	"simplecheckout/graph/model"
	"simplecheckout/models"
	"simplecheckout/usecase/checkout"
	"simplecheckout/usecase/product"
	"simplecheckout/usecase/promo"
)

func (r *mutationResolver) CreateProduct(ctx context.Context, input *model.NewProduct) (*model.Product, error) {
	p := product.NewService(rp)

	product := &model.Product{}

	createdProduct, err := p.CreateProduct(input.Sku, input.Name, input.Price, input.Quantity)

	if err != nil {
		return product, err
	}

	product = &model.Product{
		ID:       createdProduct.ID,
		Name:     createdProduct.Name,
		Price:    createdProduct.Price,
		Quantity: createdProduct.Quantity,
		Sku:      createdProduct.SKU,
	}

	return product, nil
}

func (r *mutationResolver) CreatePromo(ctx context.Context, input *model.NewPromo) (*model.Promo, error) {
	pr := promo.NewService(rpr)

	createdPromo, err := pr.CreatePromo(input.Sku, input.Scheme, *input.Freebiesku, *input.Minqty, *input.Payonly, *input.Percentageoff)

	if err != nil {
		return &model.Promo{}, err
	}

	promo := &model.Promo{
		ID:            createdPromo.ID,
		Sku:           createdPromo.SKU,
		Scheme:        createdPromo.Scheme,
		Freebiesku:    createdPromo.FreebieSKU,
		Minqty:        createdPromo.MinQty,
		Payonly:       createdPromo.PayOnly,
		Percentageoff: createdPromo.PercentageOff,
	}

	return promo, nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	pr := product.NewService(rp)

	product, err := pr.GetProduct(id)

	if err != nil {
		return &model.Product{}, err
	}

	return &model.Product{
		ID:       product.ID,
		Sku:      product.SKU,
		Name:     product.Name,
		Price:    product.Price,
		Quantity: product.Quantity,
	}, nil
}

func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	pr := product.NewService(rp)

	product, err := pr.ListProducts()

	results := make([]*model.Product, len(product))

	if err != nil {
		return results, err
	}

	for _, p := range product {
		temp := &model.Product{
			ID:       p.ID,
			Sku:      p.SKU,
			Name:     p.Name,
			Price:    p.Price,
			Quantity: p.Quantity,
		}

		results = append(results, temp)
	}

	return results, nil
}

func (r *queryResolver) Promo(ctx context.Context, id string) (*model.Promo, error) {
	pr := promo.NewService(rpr)

	promo, err := pr.GetPromo(id)

	if err != nil {
		return &model.Promo{}, err
	}

	return &model.Promo{
		ID:            promo.ID,
		Sku:           promo.SKU,
		Scheme:        promo.Scheme,
		Freebiesku:    promo.FreebieSKU,
		Minqty:        promo.MinQty,
		Payonly:       promo.PayOnly,
		Percentageoff: promo.PercentageOff,
	}, nil
}

func (r *queryResolver) Promos(ctx context.Context) ([]*model.Promo, error) {
	pr := promo.NewService(rpr)

	promo, err := pr.ListPromos()

	results := make([]*model.Promo, len(promo))

	if err != nil {
		return results, err
	}

	for _, p := range promo {
		temp := &model.Promo{
			ID:            p.ID,
			Sku:           p.SKU,
			Scheme:        p.Scheme,
			Freebiesku:    p.FreebieSKU,
			Minqty:        p.MinQty,
			Payonly:       p.PayOnly,
			Percentageoff: p.PercentageOff,
		}

		results = append(results, temp)
	}

	return results, nil
}

func (r *queryResolver) Checkout(ctx context.Context, input *model.Cart) (*model.Response, error) {
	pu := product.NewService(rp)
	pru := promo.NewService(rpr)

	response := &model.Response{}

	uc := checkout.NewService(pu, pru)

	dbyte, err := json.Marshal(input)

	if err != nil {
		return response, err
	}

	cart := &models.Cart{}

	err = json.Unmarshal(dbyte, cart)

	err = uc.Checkout(cart)

	if err != nil {
		return response, err
	}

	response.Total = uc.Total

	return response, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var rp = product.NewInmem()
var rpr = promo.NewInmem()
