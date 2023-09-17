package business

import (
	"context"
	"hexagonal/domain/entities/product"
	"hexagonal/domain/entities/restriction"
	"hexagonal/domain/entities/tokonow"
	"hexagonal/newprovider"
)

type CheckRestrictionUsecaseInterface interface {
	CheckRestrictedCategoryProduct(ctx context.Context, listProd []product.Product) (validationMessage string)
	CheckRestrictedProductForShopFollower()
	CheckRestrictedProductForLoyaltyDays()
	MultiValidateRestriction()
	CheckStatusFSLocationRestriction() string
}

type CheckRestrictionUseCases struct {
	TokonowRepo     tokonow.TokonowRepository
	RestrictionRepo restriction.RestrictionRepository
}

func NewCheckRestrictionUseCases(repos newprovider.DomainRepository) *CheckRestrictionUseCases {
	return &CheckRestrictionUseCases{
		TokonowRepo:     repos,
		RestrictionRepo: repos,
	}
}

// CheckRestrictedCategoryProduct will return validation message according to category
func (uc *CheckRestrictionUseCases) CheckRestrictedCategoryProduct(ctx context.Context, listProd []product.Product) (validationMessage string) {
	var reqRestriction restriction.InputValidateRestrictionCategory
	for _, p := range listProd {
		tokonow, err := uc.TokonowRepo.GetTokonowData(p.CategoryID)
		if err != nil {
			return
		}
		reqRestriction.TokonowData[restriction.ProdCategoryID(p.CategoryID)] = tokonow
	}
	validateRestriction, err := uc.RestrictionRepo.ValidateRestrictionCategory(reqRestriction)
	if err != nil {
		return
	}
	validationMessage = validateRestriction.GetRestrictedCategoryMessage("", 10)
	return
}

// CheckRestrictedProductForShopFollower will return validation message according to category
func (uc *CheckRestrictionUseCases) CheckRestrictedProductForShopFollower() {
	// TODO implement me
	panic("implement me")
}

// CheckRestrictedProductForLoyaltyDays will return validation message according to category
func (uc *CheckRestrictionUseCases) CheckRestrictedProductForLoyaltyDays() {
	// TODO implement me
	panic("implement me")
}

// MultiValidateRestriction will return validation message according to category
func (uc *CheckRestrictionUseCases) MultiValidateRestriction() {
	// TODO implement me
	panic("implement me")
}

// CheckStatusFSLocationRestriction will check all status location
func (uc *CheckRestrictionUseCases) CheckStatusFSLocationRestriction() string {
	return ""
}
