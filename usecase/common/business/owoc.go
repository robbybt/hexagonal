package business

import (
	"context"
	"hexagonal/domain/entities/product"
	"hexagonal/domain/entities/shop"
	"hexagonal/domain/entities/tokonow"
	"hexagonal/newprovider"
)

type OWOCUsecaseInterface interface {
	DoOWOC(ctx context.Context) error
}

type OWOCUseCases struct {
	Tokonow tokonow.TokonowRepository
}

func NewOWOCUseCases(repos newprovider.DomainRepository) *OWOCUseCases {
	return &OWOCUseCases{
		Tokonow: repos,
	}
}

// DoOWOC will get tokonow first with list product and shop
func (uc *OWOCUseCases) DoOWOC(ctx context.Context, listProduct []product.Product, listShgp []shop.Shop) error {
	_, err := uc.Tokonow.GetTokonowData(1)
	return err
}
