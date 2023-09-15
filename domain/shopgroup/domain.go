package shopgroup

import (
	"hexagonal/domain/product"
	"hexagonal/domain/shop"
)

type CartGroup struct {
	UserID    int64
	ShopGroup ShopGroup
}

type ShopGroup struct {
	Shop        shop.Shop
	GroupDetail GroupDetail
}

type GroupDetail struct {
	Product product.Product
}
