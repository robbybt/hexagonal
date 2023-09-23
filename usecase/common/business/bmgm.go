package business

import (
	"context"
	"hexagonal/domain/entities/campaign"
)

type BMGMUsecaseInterface interface {
	DoBMGM(ctx context.Context) error
}

type BmgmRepository struct {
	CampaignRepos campaign.CampaignRepository
}
type bmgmUseCases struct {
	BmgmRepository
}

func NewBMGMUseCases(repos BmgmRepository) BMGMUsecaseInterface {
	return &bmgmUseCases{repos}
}

// DoBMGM will get campaign and ...
func (uc *bmgmUseCases) DoBMGM(ctx context.Context) error {
	_, err := uc.CampaignRepos.GetCampaign(campaign.InputGetCampaign{})
	return err
}
