package business

import (
	"context"
	"hexagonal/domain/entities/product"
	"hexagonal/domain/entities/restriction"
	"hexagonal/domain/entities/tokonow"
)

type CheckRestrictionUsecaseInterface interface {
	CheckRestrictedCategoryProduct(ctx context.Context, listProd []product.Product) (validationMessage string)
	CheckRestrictedProductForShopFollower()
	CheckRestrictedProductForLoyaltyDays()
	MultiValidateRestriction()
	CheckStatusFSLocationRestriction() string
}

type CheckRestrictionRepository struct {
	TokonowRepo     tokonow.TokonowRepository
	RestrictionRepo restriction.RestrictionRepository
}

type checkRestrictionUseCases struct {
	CheckRestrictionRepository
}

func NewCheckRestrictionUseCases(repos CheckRestrictionRepository) CheckRestrictionUsecaseInterface {
	return &checkRestrictionUseCases{repos}
}

// CheckRestrictedCategoryProduct will return validation message according to category
func (uc *checkRestrictionUseCases) CheckRestrictedCategoryProduct(ctx context.Context, listProd []product.Product) (validationMessage string) {
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
func (uc *checkRestrictionUseCases) CheckRestrictedProductForShopFollower() {
	// TODO implement me
	panic("implement me")
}

// CheckRestrictedProductForLoyaltyDays will return validation message according to category
func (uc *checkRestrictionUseCases) CheckRestrictedProductForLoyaltyDays() {
	// TODO implement me
	panic("implement me")
}

// MultiValidateRestriction will return validation message according to category
func (uc *checkRestrictionUseCases) MultiValidateRestriction() {
	// TODO implement me
	panic("implement me")
}

// CheckStatusFSLocationRestriction will check all status location
func (uc *checkRestrictionUseCases) CheckStatusFSLocationRestriction() string {
	return ""
}
