package warehouse

type InputGetWarehouseData struct {
	ListWarehouseID []int
}

type WarehouseRepository interface {
	GetWarehouseData(i InputGetWarehouseData) ([]WarehousePartnerData, error)
}
