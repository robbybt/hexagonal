package dproduct

type ListProduct struct {
	Data []ProductDataATC
}

type ProductDataATC struct {
	Product

	WarehouseData string
	Volumetric    string
}

type Product struct {
	ProductName     string
	ProductAlias    string
	ProductID       int64
	ParentID        int64
	ParentName      string
	Variant         productVariant
	SKU             string
	ShopID          int64
	SellerID        int64
	ProductPriceFmt string
}

type productVariant struct {
	ParentID   int64
	IsParent   bool
	IsVariant  bool
	ChildrenID []int64
}
