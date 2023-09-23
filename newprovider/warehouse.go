package newprovider

import (
	"fmt"
	"hexagonal/domain/entities/warehouse"
	warehouse2 "hexagonal/domain/entities/warehouse"
)

type warehouseResponse struct {
	warehouseDataResponse
}

type warehouseDataResponse struct {
	Warehouses []*warehouseData `protobuf:"bytes,1,rep,name=Warehouses,proto3" json:"Warehouses,omitempty"`
}

type warehouseData struct {
	WarehouseID               int64   `protobuf:"varint,1,opt,name=WarehouseID,proto3" json:"WarehouseID,omitempty"`
	PartnerID                 int     `protobuf:"bytes,2,opt,name=PartnerID,proto3" json:"PartnerID,omitempty"`
	ShopID                    int     `protobuf:"bytes,3,opt,name=ShopID,proto3" json:"ShopID,omitempty"`
	WarehouseName             string  `protobuf:"bytes,4,opt,name=WarehouseName,proto3" json:"WarehouseName,omitempty"`
	DistrictID                int64   `protobuf:"varint,5,opt,name=DistrictID,proto3" json:"DistrictID,omitempty"`
	DistrictName              string  `protobuf:"bytes,6,opt,name=DistrictName,proto3" json:"DistrictName,omitempty"`
	CityID                    int64   `protobuf:"varint,7,opt,name=CityID,proto3" json:"CityID,omitempty"`
	CityName                  string  `protobuf:"bytes,8,opt,name=CityName,proto3" json:"CityName,omitempty"`
	ProvinceID                int64   `protobuf:"varint,9,opt,name=ProvinceID,proto3" json:"ProvinceID,omitempty"`
	ProvinceName              string  `protobuf:"bytes,10,opt,name=ProvinceName,proto3" json:"ProvinceName,omitempty"`
	Status                    int32   `protobuf:"varint,11,opt,name=Status,proto3" json:"Status,omitempty"`
	Postal                    string  `protobuf:"bytes,12,opt,name=Postal,proto3" json:"Postal,omitempty"`
	IsDefault                 int32   `protobuf:"varint,13,opt,name=IsDefault,proto3" json:"IsDefault,omitempty"`
	Latlon                    string  `protobuf:"bytes,14,opt,name=Latlon,proto3" json:"Latlon,omitempty"`
	Latitude                  string  `protobuf:"bytes,15,opt,name=Latitude,proto3" json:"Latitude,omitempty"`
	Longitude                 string  `protobuf:"bytes,16,opt,name=Longitude,proto3" json:"Longitude,omitempty"`
	Email                     string  `protobuf:"bytes,17,opt,name=Email,proto3" json:"Email,omitempty"`
	AddressDetail             string  `protobuf:"bytes,18,opt,name=AddressDetail,proto3" json:"AddressDetail,omitempty"`
	Country                   string  `protobuf:"bytes,19,opt,name=Country,proto3" json:"Country,omitempty"`
	IsFulfillment             bool    `protobuf:"varint,20,opt,name=IsFulfillment,proto3" json:"IsFulfillment,omitempty"`
	Phone                     string  `protobuf:"bytes,21,opt,name=Phone,proto3" json:"Phone,omitempty"`
	TkpdPreferredLogisticSpID []int64 `protobuf:"varint,22,rep,packed,name=TkpdPreferredLogisticSpID,proto3" json:"TkpdPreferredLogisticSpID,omitempty"`
	WarehouseType             int32   `protobuf:"varint,23,opt,name=WarehouseType,proto3" json:"WarehouseType,omitempty"`
	GreaterCityID             int32   `protobuf:"varint,24,opt,name=GreaterCityID,proto3" json:"GreaterCityID,omitempty"`
	GreaterCityName           string  `protobuf:"bytes,25,opt,name=GreaterCityName,proto3" json:"GreaterCityName,omitempty"`
}

func (repo *repositories) GetWarehouseData(i warehouse.InputGetWarehouseData) ([]warehouse.WarehousePartnerData, error) {
	// restapicalls
	fmt.Println("rest API with", i)
	var resp warehouseResponse

	return resp.BuildToWarehousePartnerData(), nil
}

func (r *warehouseResponse) BuildToWarehousePartnerData() []warehouse.WarehousePartnerData {
	return []warehouse2.WarehousePartnerData{}
}
