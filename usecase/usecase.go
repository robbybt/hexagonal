package usecase

import (
	"hexagonal/newprovider"
)

type UseCases struct {
	AtcUsecaseInterface
	CartlistUsecaseInterface
}

func NewUseCases(repos newprovider.DomainRepository) *UseCases {
	atc := NewAtcUsecase(atcRepository{
		productRepository:     repos,
		shopRepository:        repos,
		imsRepository:         repos,
		campaignRepository:    repos,
		cartRepository:        repos,
		warehouseRepository:   repos,
		tokonowRepository:     repos,
		restrictionRepository: repos,
	})
	cartlist := NewCartListUsecase(cartlistRepository{
		productRepository: repos,
		shopRepository:    repos,
		imsRepository:     repos,
	})
	return &UseCases{
		atc,
		cartlist,
	}
}
