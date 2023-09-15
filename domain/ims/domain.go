package ims

type ProductsWarehouseData struct {
	ProductsWarehouseData  map[int64]map[int64]ProductWarehouse
	IsEmptyRespByProductID map[int64]bool
}

type ProductWarehouse struct {
	ProductID          int64
	Stock              int64
	OriginalStock      int64
	StockWording       string
	Price              float64
	OriginalPrice      float64
	WarehouseID        int64
	IsOutOfCoverage    bool
	BundleInfo         BundleInfoWarehouse
	DiscountPercentage float64
	EventID            int64
	EventType          int64
	RerouteReason      int32
	GroupKey           string
	GroupType          int32
	GroupWarehouseID   int64
	ShopID             int64
	UnavailableReason  int32
	OriginWarehouseID  int64
}

type BundleInfoWarehouse struct {
	BundleID      int64
	BundleQuota   int32
	OriginalQuota int32
	MaxOrder      int32
	MinOrder      int32
}
