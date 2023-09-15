package usecase

import (
	"context"
	"fmt"
	"hexagonal/domain/campaign"
	"hexagonal/domain/ims"
	"hexagonal/domain/product"
	"hexagonal/domain/shop"
	"hexagonal/domain/warehouse"
	"hexagonal/lib/ctxlib"
	"hexagonal/lib/ratelimiter"
	"hexagonal/newprovider"
	"hexagonal/usecase/business"
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

func (uc *UseCases) DoAtc(ctx context.Context, req RequestFrontEndATC) (responseATC, error) {
	ctx, tSpan := ctxlib.SetupContextTracerUsecase(ctx)
	defer tSpan.Finish()

	resp := buildInitialAtcResponse(false)

	if r := validateInputAddProductToCart(ctx, req); len(r) > 0 {
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
	if ValidateIsOneOfATCFromExternalSource("cd.ATCFromExternalSource") { // if product has variants; set a default variant for it
		if prod.HaveVariant() {
			// Change product data to product variant data
			prod = getNewProductDataVariant(req, prod.Variant.ChildrenID[0])
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

	// Set IsAllowManageWarehouse if admin
	// will skip it

	if !isEnableMultiValidateRestriction() {
		uc.CheckRestriction.CheckRestrictedCategoryProduct(ctx, listProd)
	} else {
		uc.CheckRestriction.CheckRestrictedCategoryProduct(ctx, listProd)
	}

	// ValidateAddProductToCart

	fmt.Println(listProd, listShop, cmp, ListWarehouseData)

	return responseATC{}, nil
}

func validateInputAddProductToCart(ctx context.Context, req RequestFrontEndATC) []string {
	return []string{}
}

func buildInitialAtcResponse(isFromExternalEndpoint bool) responseATC {
	return responseATC{}
}

func (r *responseATC) setAtcResponseInputValidationError() {
	// set response value
}

func (r *responseATC) setAtcResponseTooManyRequest() {
	// set response value
}

func ValidateIsOneOfATCFromExternalSource(s string) bool {
	return false
}

// get product data from variant ID
func getNewProductDataVariant(req RequestFrontEndATC, i int64) product.ProductDataATC {
	return product.ProductDataATC{}
}

func isEnableMultiValidateRestriction() bool { return false }
