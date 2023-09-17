package restriction

import (
	"hexagonal/domain/entities/tokonow"
)

type ProdCategoryID int

type InputValidateRestrictionCategory struct {
	CategoryId    int64
	IsAdult       int
	IsKyc         bool
	IsBlacklisted bool
	MinAge        int
	TokonowData   map[ProdCategoryID]tokonow.TokonowCategory
}

type RestrictionRepository interface {
	ValidateRestrictionCategory(i InputValidateRestrictionCategory) (ValidateRestrictionResponse, error)
}
