package product

const (
	ProductStateDeleted   = 0
	ProductStateActive    = 1
	ProductStateBest      = 2
	ProductStateWareHouse = 3
	ProductStateHidden    = 4
	ProductStateActiveB2B = 5
	ProductStateModerated = 6
	ProductStatePending   = -1
	ProductStateBanned    = -2
)

type ListProduct struct {
	List []Product
}

type MapProduct struct {
	Map map[int64]Product
}

type Product struct {
	ProductName      string
	ProductAlias     string
	ProductID        int64
	ParentID         int64
	CategoryID       int
	ParentName       string
	Variant          ProductVariant
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

type ProductVariant struct {
	ParentID   int64
	IsParent   bool
	IsVariant  bool
	ChildrenID []int64
}

func (list *ListProduct) MapProductID() MapProduct {
	return MapProduct{}
}

func (list *ListProduct) ListWarehouseID() (listWarehouseID []int) {
	return listWarehouseID
}

// GetProductVariant will return product data from variant id
func (list *Product) GetProductVariant() Product {
	return Product{}
}
