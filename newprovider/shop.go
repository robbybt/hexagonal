package newprovider

import (
	"fmt"
	shop2 "hexagonal/domain/entities/shop"
)

type ShopResponse struct {
	ShopID   int64   `json:"shop_id"`
	ShopURL  string  `json:"shop_url"`
	UserID   int64   `json:"user_id"`
	AdminID  []int64 `json:"admin_id"`
	ShopName string  `json:"shop_name"`
	Status   int32   `json:"status"`
}

func (repo *Repositories) GetShop(i shop2.InputGetShop) (shop2.ListShop, error) {
	// restapicalls
	fmt.Println("rest API with", i)
	var resp ShopResponse

	return resp.BuildToShop(), nil
}

func (r *ShopResponse) BuildToShop() shop2.ListShop {
	return shop2.ListShop{}
}
