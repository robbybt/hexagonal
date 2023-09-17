package product

type InputGetProductATC struct {
	ListProductID []int
}

type ProductRepository interface {
	GetProductATC(i InputGetProductATC) (ListProductATC, error)
}
