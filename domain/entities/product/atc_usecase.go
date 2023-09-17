package product

type ListProductATC struct {
	List []ProductDataATC
}

type ProductDataATC struct {
	Product

	WarehouseData string
	Volumetric    string
}

func (p *ListProductATC) MapofProductID() map[int64]Product {
	return map[int64]Product{}
}

func (p *ListProductATC) ListofProduct() []Product {
	return []Product{}
}

func (p *ListProductATC) ListWarehouse() []int {
	return []int{}
}

// GetProductVariant will return product data from variant ID
func (p *ProductDataATC) GetProductVariant() ProductDataATC {
	return ProductDataATC{}
}
