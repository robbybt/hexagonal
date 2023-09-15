package product

type ListProductATC struct {
	List []ProductDataATC
}

type ProductDataATC struct {
	Product

	WarehouseData string
	Volumetric    string
}

type Product struct {
	ProductName      string
	ProductAlias     string
	ProductID        int64
	ParentID         int64
	CategoryID       int
	ParentName       string
	Variant          productVariant
	SKU              string
	ShopID           int64
	SellerID         int64
	ProductPriceFmt  string
	DefaultWarehouse ProductWarehouse
	Status           int
}

type ProductWarehouse struct {
	WarehouseID           int
	ProductPrice          int64
	PriceCurrency         int32
	PriceCurrencyName     string
	ProductPriceIDR       int64
	ProductSwitchInvenage int32
	ProductInvenageValue  int32
}

type productVariant struct {
	ParentID   int64
	IsParent   bool
	IsVariant  bool
	ChildrenID []int64
}

func (p *ListProductATC) MapofProductID() map[int64]Product {
	return map[int64]Product{}
}

func (p *ListProductATC) ListofProduct() []Product {
	return []Product{}
}

func (p *ListProductATC) ListWarehouse() []int {
	return []int{}
}

func (p *Product) HaveVariant() bool {
	return false
}

func (p *Product) GetMinimumOrder() int {
	return 1
}
