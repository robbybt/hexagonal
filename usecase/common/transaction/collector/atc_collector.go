package collector

import (
	"hexagonal/domain/entities/campaign"
	"hexagonal/domain/entities/ims"
	"hexagonal/domain/entities/product"
	"hexagonal/domain/entities/shop"
)

type CollectorATCUsecaseInterface interface {
	FetchDataListProduct(
		listProductID, listShopID []int,
	) (
		listProdATC product.ListProductATC,
		listShop shop.ListShop,
		imsData ims.ProductsWarehouseData,
		cmp campaign.CampaignData,
		err error,
	)
}

type CollectorATCRepository struct {
	ProductRepository  product.ProductRepository
	ShopRepository     shop.ShopRepository
	ImsRepository      ims.ImsRepository
	CampaignRepository campaign.CampaignRepository
}

type collectorATCUseCases struct {
	CollectorATCRepository
}

func NewCollectorATCUseCases(repos CollectorATCRepository) CollectorATCUsecaseInterface {
	return &collectorATCUseCases{repos}
}

func (c *collectorATCUseCases) FetchDataListProduct(
	listProductID, listShopID []int,
) (
	listProdATC product.ListProductATC,
	listShop shop.ListShop,
	imsData ims.ProductsWarehouseData,
	cmp campaign.CampaignData,
	err error,
) {
	listProdATC, err = c.ProductRepository.GetProductATC(product.InputGetProductATC{
		ListProductID: listProductID,
	})
	if err != nil {
		return
	}

	listShop, err = c.ShopRepository.GetShop(shop.InputGetShop{
		ListShopID: listShopID,
	})
	if err != nil {
		return
	}

	imsData, err = c.ImsRepository.GetNearestWarehouse(ims.InputGetNearestWarehouse{
		ListProductID: listProductID,
	})
	if err != nil {
		return
	}

	// GetCampaign need ims data so need to be called after get ims data
	cmp, err = c.CampaignRepository.GetCampaign(campaign.InputGetCampaign{
		WarehouseData: imsData,
		ListProductID: listProductID,
	})
	if err != nil {
		return
	}
	return
}
