package cart

// Cart type for representing each item of cart product, feel free to add fields
type Cart struct {
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
