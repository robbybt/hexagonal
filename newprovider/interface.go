package newprovider

import (
	"hexagonal/domain/campaign"
	"hexagonal/domain/ims"
	"hexagonal/domain/product"
	"hexagonal/domain/shop"
)

type DomainRepository interface {
	GetProductATC(i InputGetProductATC) (product.ListProductAtc, error)
	GetShop(i InputGetShop) (shop.ListShop, error)
	GetNearestWarehouse(i InputGetNearestWarehouse) (ims.ProductsWarehouseData, error)
	GetCampaign(i InputGetCampaign) (campaign.CampaignData, error)
}
