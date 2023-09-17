package shopgroup

import (
	"hexagonal/domain/entities/product"
	"hexagonal/domain/entities/shop"
)

type CartGroup struct {
	UserID    int64 // user.User
	ShopGroup ShopGroup
}

type ShopGroup struct {
	Shop        shop.Shop
	GroupDetail GroupDetail
}

type GroupDetail struct {
	Product product.Product
}
