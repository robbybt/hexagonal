package campaign

import (
	"hexagonal/domain/entities/ims"
)

type InputGetCampaign struct {
	WarehouseData ims.ProductsWarehouseData
	ListProductID []int
}

type CampaignRepository interface {
	GetCampaign(i InputGetCampaign) (CampaignData, error)
}
