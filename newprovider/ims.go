package newprovider

import (
	"fmt"
	"hexagonal/domain/ims"
)

type ImsGRPC struct {
	ProductID          int64   `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
	Stock              int64   `protobuf:"varint,2,opt,name=stock,proto3" json:"stock,omitempty"`
	StockWording       string  `protobuf:"bytes,3,opt,name=stockWording,proto3" json:"stockWording,omitempty"`
	Price              float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	OriginalPrice      float64 `protobuf:"fixed64,7,opt,name=originalPrice,proto3" json:"originalPrice,omitempty"`
	DiscountPercentage float64 `protobuf:"fixed64,8,opt,name=discountPercentage,proto3" json:"discountPercentage,omitempty"`
	OriginalStock      int64   `protobuf:"varint,10,opt,name=originalStock,proto3" json:"originalStock,omitempty"`
}

type ImsResponse struct {
	ImsGRPC
}

func (repo *Repositories) GetNearestWarehouse(i ims.InputGetNearestWarehouse) (ims.ProductsWarehouseData, error) {
	// restapicalls
	fmt.Println("rest API with", i)
	var resp ImsResponse

	return resp.BuildToProductsWarehouseData(), nil
}

func (r *ImsResponse) BuildToProductsWarehouseData() ims.ProductsWarehouseData {
	return ims.ProductsWarehouseData{}
}
