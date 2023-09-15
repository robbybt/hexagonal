package shop

type InputGetShop struct {
	ListShopID []int
}

type ShopRepository interface {
	GetShop(i InputGetShop) (ListShop, error)
}
