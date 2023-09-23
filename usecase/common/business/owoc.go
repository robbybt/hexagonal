package business

import (
	"context"
	"hexagonal/domain/entities/product"
	"hexagonal/domain/entities/shop"
	"hexagonal/domain/entities/tokonow"
)

type OWOCUsecaseInterface interface {
	DoOWOC(ctx context.Context, listProduct []product.Product, listShgp []shop.Shop) error
}

type OwocRepository struct {
	Tokonow tokonow.TokonowRepository
}

type owocUseCases struct {
	OwocRepository
}

func NewOWOCUseCases(repos OwocRepository) OWOCUsecaseInterface {
	return &owocUseCases{repos}
}

// DoOWOC will get tokonow first with list product and shop
func (uc *owocUseCases) DoOWOC(ctx context.Context, listProduct []product.Product, listShgp []shop.Shop) error {
	_, err := uc.Tokonow.GetTokonowData(1)
	return err
}
