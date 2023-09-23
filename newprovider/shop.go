package newprovider

import (
	"fmt"
	"hexagonal/domain/entities/shop"
	shop2 "hexagonal/domain/entities/shop"
)

type shopRequest struct {
	ShopID int `json:"shop_id"`
}

type shopResponse struct {
	ShopID   int64   `json:"shop_id"`
	ShopURL  string  `json:"shop_url"`
	UserID   int64   `json:"user_id"`
	AdminID  []int64 `json:"admin_id"`
	ShopName string  `json:"shop_name"`
	Status   int32   `json:"status"`
}

func (r *shopRequest) setupValue(i shop.InputGetShop) {
	r.ShopID = i.ListShopID[0]
}

func (repo *repositories) GetShop(i shop.InputGetShop) (shop.ListShop, error) {
	// restapicalls
	var req shopRequest
	req.setupValue(i)

	fmt.Println("rest API with", i)
	var resp shopResponse

	return resp.BuildToShop(), nil
}

func (r *shopResponse) BuildToShop() shop.ListShop {
	return shop2.ListShop{}
}
