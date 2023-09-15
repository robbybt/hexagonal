package product

const (
	ProductStateDeleted   = 0
	ProductStateActive    = 1
	ProductStateBest      = 2
	ProductStateWareHouse = 3
	ProductStateHidden    = 4
	ProductStateActiveB2B = 5
	ProductStateModerated = 6
	ProductStatePending   = -1
	ProductStateBanned    = -2
)

// IsProductActive to check is product active
func (p *Product) IsProductActive() bool {
	return p.Status == ProductStateActive
}
