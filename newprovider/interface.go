package newprovider

import (
	"hexagonal/domain/campaign"
	"hexagonal/domain/cart"
	"hexagonal/domain/ims"
	"hexagonal/domain/product"
	"hexagonal/domain/restriction"
	"hexagonal/domain/shop"
	"hexagonal/domain/tokonow"
	"hexagonal/domain/warehouse"
)

type DomainRepository interface {
	product.ProductRepository
	shop.ShopRepository
	ims.ImsRepository
	campaign.CampaignRepository
	cart.CartRepository
	warehouse.WarehouseRepository
	tokonow.TokonowRepository
	restriction.RestrictionRepository
}
