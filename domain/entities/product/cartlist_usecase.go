package product

type ListProductCartList struct {
	List []ProductDataCartList
}

type MapProductCartList struct {
	Map map[int64]ProductDataCartList
}

type ProductDataCartList struct {
	Product

	CartlistData string
}

func (list ListProductCartList) MapProductID() MapProductCartList {
	return MapProductCartList{}
}

func (p *ListProductCartList) ListofProduct() []Product {
	return []Product{}
}

func (p *MapProductCartList) ListofProduct() MapProduct {
	return MapProduct{}
}
