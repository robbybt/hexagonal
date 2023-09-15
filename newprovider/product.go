package newprovider

import (
	"fmt"
	"hexagonal/domain/product"
)

// TODO: ASK

type InputGetProductATC struct {
	ListProductID []int
}

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

func (r *TomeResponse) BuildToProductDataATC() product.ProductDataATC {
	return product.ProductDataATC{}
}

func (repo *Repositories) GetProductATC(i InputGetProductATC) product.ProductDataATC {
	// restapicalls
	fmt.Println("rest API with", i)
	var resp TomeResponse

	return resp.BuildToProductDataATC()
}
