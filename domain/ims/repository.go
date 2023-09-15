package ims

type InputGetNearestWarehouse struct {
	ListProductID []int
}

type ImsRepository interface {
	GetNearestWarehouse(i InputGetNearestWarehouse) (ProductsWarehouseData, error)
}
