package newprovider

import (
	"fmt"
	product2 "hexagonal/domain/entities/product"
)

type TomeResponse struct {
	ProductName     string `json:"json1"`
	ProductAlias    string `json:"json2"`
	ProductID       int64  `json:"json3"`
	ParentID        int64  `json:"json4"`
	ParentName      string `json:"json5"`
	SKU             string `json:"json7"`
	ShopID          int64  `json:"json8"`
	SellerID        int64  `json:"json9"`
	ProductPriceFmt string `json:"json"`
}

func (r *TomeResponse) BuildToProductDataATC() product2.ListProductATC {
	return product2.ListProductATC{}
}

func (repo *Repositories) GetProductATC(i product2.InputGetProductATC) (product2.ListProductATC, error) {
	// restapicalls
	fmt.Println("rest API with", i)
	var resp TomeResponse

	return resp.BuildToProductDataATC(), nil
}
