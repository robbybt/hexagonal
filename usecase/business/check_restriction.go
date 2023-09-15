package check_restriction

import (
	"context"
	"hexagonal/domain/product"
	"hexagonal/domain/restriction"
	"hexagonal/domain/tokonow"
	"hexagonal/newprovider"
)

type CheckRestrictionUsecaseInterface interface {
	CheckRestrictedCategoryProduct(ctx context.Context, listProd []product.Product) (validationMessage string)
	CheckRestrictedProductForShopFollower()
	CheckRestrictedProductForLoyaltyDays()
	MultiValidateRestriction()
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

func (uc *CheckRestrictionUseCases) CheckRestrictedProductForShopFollower() {
	// TODO implement me
	panic("implement me")
}

func (uc *CheckRestrictionUseCases) CheckRestrictedProductForLoyaltyDays() {
	// TODO implement me
	panic("implement me")
}

func (uc *CheckRestrictionUseCases) MultiValidateRestriction() {
	// TODO implement me
	panic("implement me")
}

func CheckStatusFSLocationRestriction() string {
	return ""
}
