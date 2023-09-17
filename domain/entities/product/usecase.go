package product

// IsProductActive to check is product active
func (p *Product) IsProductActive() bool {
	return p.Status == ProductStateActive
}

func (p *Product) HaveVariant() bool {
	return false
}

func (p *Product) GetMinimumOrder() int {
	return 1
}
