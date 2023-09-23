package usecase

import (
	"context"
	"fmt"
	"hexagonal/domain/entities/campaign"
	"hexagonal/domain/entities/cart"
	"hexagonal/domain/entities/ims"
	"hexagonal/domain/entities/product"
	"hexagonal/domain/entities/restriction"
	"hexagonal/domain/entities/shop"
	"hexagonal/domain/entities/tokonow"
	"hexagonal/domain/entities/warehouse"
	"hexagonal/lib/ctxlib"
	"hexagonal/lib/ratelimiter"
	"hexagonal/usecase/common/business"
	"hexagonal/usecase/common/transaction/collector"
	"hexagonal/usecase/common/transaction/validation"
)

type AtcUsecaseInterface interface {
	DoAtc(ctx context.Context, req RequestFrontEndATC) (resp ResponseATC, err error)
}

type atcRepository struct {
	productRepository     product.ProductRepository
	shopRepository        shop.ShopRepository
	imsRepository         ims.ImsRepository
	campaignRepository    campaign.CampaignRepository
	cartRepository        cart.CartRepository
	warehouseRepository   warehouse.WarehouseRepository
	tokonowRepository     tokonow.TokonowRepository
	restrictionRepository restriction.RestrictionRepository
}

type atcUsecase struct {
	atcRepository
	bmgm             business.BMGMUsecaseInterface
	checkrestriction business.CheckRestrictionUsecaseInterface
	owoc             business.OWOCUsecaseInterface
	collectorATC     collector.CollectorATCUsecaseInterface
}

func NewAtcUsecase(repository atcRepository) AtcUsecaseInterface {
	// init common business
	bmgm := business.NewBMGMUseCases(business.BmgmRepository{
		CampaignRepos: repository.campaignRepository,
	})
	checkrestriction := business.NewCheckRestrictionUseCases(business.CheckRestrictionRepository{
		TokonowRepo:     repository.tokonowRepository,
		RestrictionRepo: repository.restrictionRepository,
	})
	owoc := business.NewOWOCUseCases(business.OwocRepository{
		Tokonow: repository.tokonowRepository,
	})
	collectorATC := collector.NewCollectorATCUseCases(collector.CollectorATCRepository{
		ProductRepository:  repository.productRepository,
		ShopRepository:     repository.shopRepository,
		ImsRepository:      repository.imsRepository,
		CampaignRepository: repository.campaignRepository,
	})
	return &atcUsecase{
		atcRepository:    repository,
		bmgm:             bmgm,
		checkrestriction: checkrestriction,
		owoc:             owoc,
		collectorATC:     collectorATC,
	}
}

func (uc *atcUsecase) DoAtc(ctx context.Context, req RequestFrontEndATC) (resp ResponseATC, err error) {
	ctx, tSpan := ctxlib.SetupContextTracerUsecase(ctx)
	defer tSpan.Finish()

	resp.setInitialResponse(false)

	if r := req.validateInput(ctx); len(r) > 0 {
		resp.setResponseInputValidationError()
		return resp, fmt.Errorf("error")
	}

	listProdATC, listShop, imsData, cmp, err := uc.collectorATC.FetchDataListProduct([]int{req.productID}, []int{req.shopID})
	if err != nil {
		fmt.Println(err.Error())
		return resp, err
	}

	if ratelimiter.CheckRequest(ctx, ratelimiter.ATCSource, func(ctx context.Context, param ratelimiter.LimiterCheckParam) (ratelimiter.LimiterCheckResult, error) {
		return ratelimiter.LimiterCheckResult{}, nil
	}) {
		resp.setResponseTooManyRequest()
		fmt.Println("kena ratelimiter")
		return resp, fmt.Errorf("error")
	}

	// validate whether users do ATC from wishlist, last seen, recommendation list, or etc..
	listProd := listProdATC.ListofProduct()
	prod := listProd.List[0]
	if validation.ValidateIsOneOfATCFromExternalSource("cd.ATCFromExternalSource") {
		if prod.HaveVariant() {
			// Change product data to product variant data
			prod = prod.GetProductVariant()
		}
		req.qty = prod.GetMinimumOrder()
	}

	ListWarehouseData, err := uc.warehouseRepository.GetWarehouseData(warehouse.InputGetWarehouseData{
		ListWarehouseID: listProd.ListWarehouseID(),
	})

	if !validation.IsEnableMultiValidateRestriction() {
		uc.checkrestriction.CheckRestrictedCategoryProduct(ctx, listProd.List)
	} else {
		uc.checkrestriction.CheckRestrictedCategoryProduct(ctx, listProd.List)
	}

	fmt.Println(listProd, listShop, cmp, imsData, ListWarehouseData)
	return ResponseATC{}, nil
}
