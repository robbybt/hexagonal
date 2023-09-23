package newprovider

import (
	"fmt"
	"hexagonal/domain/entities/product"
)

type tomeRequest struct {
	ProductName     string `json:"json1"`
	ProductAlias    string `json:"json2"`
	ProductID       int    `json:"json3"`
	ParentID        int64  `json:"json4"`
	ParentName      string `json:"json5"`
	SKU             string `json:"json7"`
	ShopID          int64  `json:"json8"`
	SellerID        int64  `json:"json9"`
	ProductPriceFmt string `json:"json"`
}

type tomeResponse struct {
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

func (r *tomeRequest) setupValue(i product.InputGetProductATC) {
	r.ProductID = i.ListProductID[0]
}

func (repo *repositories) GetProductATC(i product.InputGetProductATC) (product.ListProductATC, error) {
	// restapicalls
	var req tomeRequest
	req.setupValue(i)

	fmt.Println("rest API with", i)
	var resp tomeResponse

	return resp.BuildToProductDataATC(), nil
}

func (r *tomeResponse) BuildToProductDataATC() product.ListProductATC {
	return product.ListProductATC{}
}

func (r *tomeResponse) BuildToProduct() product.Product {
	return product.Product{}
}
