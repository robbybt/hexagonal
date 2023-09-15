package newprovider

import (
	"context"
	"fmt"
	"hexagonal/domain/cart"
)

type CartResponse struct {
	CartID            int64  `json:"cart_id"`
	ProductID         int64  `json:"product_id"`
	Quantity          int32  `json:"quantity"`
	Notes             string `json:"notes"`
	ShopID            int64  `json:"shop_id"`
	ShopIsGold        int32  `json:"shop_is_gold,omitempty"`
	IsPreorder        int32  `json:"product_is_preorder,omitempty"`
	WarehouseID       int64  `json:"warehouse_id"`
	CustomerID        int64  `json:"customer_id"`
	ShippingAddressID int64  `json:"shipping_address_id"`
	BundleID          int64  `json:"bundle_id"`
	BundleGroupID     string `json:"bundle_group_id"`
	BundleQuantity    int64  `json:"bundle_quantity"`
}

func (repo *Repositories) GetCartDetailByUserCartStatuses(ctx context.Context, userIDs []int64, statuses []int, isAllowBundle bool) ([]cart.Cart, error) {
	// restapicalls
	fmt.Println("rest API with", ctx)
	var resp CartResponse

	return resp.BuildToCart(), nil
}

func (r *CartResponse) BuildToCart() []cart.Cart {
	return []cart.Cart{}
}
