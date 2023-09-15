package business

import (
	"context"
	"hexagonal/domain/product"
	"hexagonal/domain/shop"
	"hexagonal/newprovider"
)

type OWOCUsecaseInterface interface {
	DoOWOC(ctx context.Context) error
}

type OWOCUseCases struct {
	ProductRepos product.ProductRepository
	ShopRepos    shop.ShopRepository
}

func NewOWOCUseCases(repos newprovider.DomainRepository) *OWOCUseCases {
	return &OWOCUseCases{
		ProductRepos: repos,
		ShopRepos:    repos,
	}
}

func (uc *OWOCUseCases) DoOWOC(ctx context.Context) error {
	_, err := uc.ProductRepos.GetProductATC(product.InputGetProductATC{})
	_, err = uc.ShopRepos.GetShop(shop.InputGetShop{})
	return err
}
