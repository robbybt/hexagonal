package usecase

import (
	"context"
	"fmt"
	"hexagonal/domain/entities/campaign"
	"hexagonal/domain/entities/ims"
	"hexagonal/domain/entities/product"
	"hexagonal/domain/entities/shop"
	"hexagonal/domain/entities/warehouse"
	"hexagonal/lib/ctxlib"
	"hexagonal/lib/ratelimiter"
	"hexagonal/newprovider"
	"hexagonal/usecase/common/business"
	"hexagonal/usecase/common/transaction/validation"
)

type UseCases struct {
	Repos            newprovider.DomainRepository
	Bmgm             business.BMGMUsecaseInterface
	CheckRestriction business.CheckRestrictionUsecaseInterface
	Owoc             business.OWOCUsecaseInterface
}

func NewUseCases(repos newprovider.DomainRepository) *UseCases {
	return &UseCases{
		Repos:            repos,
		Bmgm:             business.NewBMGMUseCases(repos),
		CheckRestriction: business.NewCheckRestrictionUseCases(repos),
		Owoc:             business.NewOWOCUseCases(repos),
	}
}

type RequestFrontEndATC struct {
	productID int
	qty       int
	shopID    int
}

type responseATC struct {
	ListProductID []int
	ListShopID    []int
}

func (uc *UseCases) DoAtc(ctx context.Context, req RequestFrontEndATC) (resp responseATC, err error) {
	ctx, tSpan := ctxlib.SetupContextTracerUsecase(ctx)
	defer tSpan.Finish()

	resp.setInitialAtcResponse(false)

	if r := req.validateInputAddProductToCart(ctx); len(r) > 0 {
		resp.setAtcResponseInputValidationError()
		return resp, nil
	}

	listProdATC, err := uc.Repos.GetProductATC(product.InputGetProductATC{
		ListProductID: []int{req.productID},
	})
	if err != nil {
		return resp, err
	}
	prod := listProdATC.List[0]
	listProd := listProdATC.ListofProduct()

	listShop, err := uc.Repos.GetShop(shop.InputGetShop{
		ListShopID: []int{req.shopID},
	})
	if err != nil {
		return resp, err
	}
	mapShop := listShop.Map()
	fmt.Println(mapShop)

	imsData, err := uc.Repos.GetNearestWarehouse(ims.InputGetNearestWarehouse{
		ListProductID: []int{req.productID},
	})
	if err != nil {
		return resp, err
	}

	// GetCampaign need ims data so need to be called after get ims data
	cmp, err := uc.Repos.GetCampaign(campaign.InputGetCampaign{
		WarehouseData: imsData,
		ListProductID: []int{req.productID},
	})
	if err != nil {
		return resp, err
	}

	// set campign data to product
	// set ims data to product
	if ratelimiter.CheckRequest(ctx, ratelimiter.ATCSource, func(ctx context.Context, param ratelimiter.LimiterCheckParam) (ratelimiter.LimiterCheckResult, error) {
		return ratelimiter.LimiterCheckResult{}, nil
	}) {
		resp.setAtcResponseTooManyRequest()
		fmt.Println("kena ratelimiter")
		return resp, nil
	}

	// validate whether users do ATC from wishlist, last seen, recommendation list, or etc..
	if validation.ValidateIsOneOfATCFromExternalSource("cd.ATCFromExternalSource") {
		// if product has variants; set a default variant for it
		if prod.HaveVariant() {
			// Change product data to product variant data
			prod = prod.GetProductVariant()
		}
		req.qty = prod.GetMinimumOrder()
		/*
			if IsFromATCExternalEndpoint(cd.ATCFromExternalSource) {
				cd.ShopID = productData.ShopID
			}
		*/
	}

	ListWarehouseData, err := uc.Repos.GetWarehouseData(warehouse.InputGetWarehouseData{
		ListWarehouseID: listProdATC.ListWarehouse(),
	})

	if !validation.IsEnableMultiValidateRestriction() {
		uc.CheckRestriction.CheckRestrictedCategoryProduct(ctx, listProd)
	} else {
		uc.CheckRestriction.CheckRestrictedCategoryProduct(ctx, listProd)
	}

	fmt.Println(listProd, listShop, cmp, ListWarehouseData)
	return responseATC{}, nil
}

func (req *RequestFrontEndATC) validateInputAddProductToCart(ctx context.Context) []string {
	return []string{}
}

func (resp *responseATC) setInitialAtcResponse(isFromExternalEndpoint bool) {
	resp.ListShopID = nil
}

func (resp *responseATC) setAtcResponseInputValidationError() {
	// set response value
}

func (resp *responseATC) setAtcResponseTooManyRequest() {
	// set response value
}
