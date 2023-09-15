package shop

type ListShop struct {
	List []Shop
}

type Shop struct {
	ShopID    int64
	ShopURL   string
	UserID    int64
	AdminID   []int64
	ShopName  string
	Status    int32
	IsTokonow bool
}

func (p *ListShop) Map() map[int64]Shop {
	return map[int64]Shop{}
}
