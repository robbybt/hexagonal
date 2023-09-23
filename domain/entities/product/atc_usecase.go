package product

type ListProductATC struct {
	List []ProductDataATC
}

type MapProductATC struct {
	Map map[int64]ProductDataATC
}

type ProductDataATC struct {
	Product

	WarehouseData string
	Volumetric    string
}

func (list ListProductATC) MapProductID() MapProductATC {
	return MapProductATC{}
}

func (p *ListProductATC) ListofProduct() ListProduct {
	return ListProduct{}
}

func (p *MapProductATC) ListofProduct() MapProduct {
	return MapProduct{}
}
