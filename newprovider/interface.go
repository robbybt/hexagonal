package newprovider

import (
	"hexagonal/domain/entities/campaign"
	"hexagonal/domain/entities/cart"
	"hexagonal/domain/entities/ims"
	"hexagonal/domain/entities/product"
	"hexagonal/domain/entities/restriction"
	"hexagonal/domain/entities/shop"
	"hexagonal/domain/entities/tokonow"
	"hexagonal/domain/entities/warehouse"
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
