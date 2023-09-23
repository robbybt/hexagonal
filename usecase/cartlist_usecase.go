package usecase

import (
	"context"
	"hexagonal/domain/entities/ims"
	"hexagonal/domain/entities/product"
	"hexagonal/domain/entities/shop"
)

type CartlistUsecaseInterface interface {
	DoCartList(ctx context.Context, req RequestFrontEndCartlist) (resp ResponseCartlist, err error)
}

type cartlistRepository struct {
	productRepository product.ProductRepository
	shopRepository    shop.ShopRepository
	imsRepository     ims.ImsRepository
}

type cartlistUsecase struct {
	cartlistRepository
}

func NewCartListUsecase(repository cartlistRepository) CartlistUsecaseInterface {
	return &cartlistUsecase{repository}
}

func (uc *cartlistUsecase) DoCartList(ctx context.Context, req RequestFrontEndCartlist) (resp ResponseCartlist, err error) {
	return ResponseCartlist{}, nil
}
