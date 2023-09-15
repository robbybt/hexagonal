package newprovider

import (
	"fmt"
	"hexagonal/domain/shop"
)

type InputGetShop struct {
	ListShopID []int
}

type ShopResponse struct {
	ShopID   int64   `json:"shop_id"`
	ShopURL  string  `json:"shop_url"`
	UserID   int64   `json:"user_id"`
	AdminID  []int64 `json:"admin_id"`
	ShopName string  `json:"shop_name"`
	Status   int32   `json:"status"`
}

func (repo *Repositories) GetShop(i InputGetShop) shop.Shop {
	// restapicalls
	fmt.Println("rest API with", i)
	var resp ShopResponse

	return resp.BuildToShop()
}

func (r *ShopResponse) BuildToShop() shop.Shop {
	return shop.Shop{}
}
