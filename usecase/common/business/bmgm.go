package business

import (
	"context"
	"hexagonal/domain/entities/campaign"
	"hexagonal/newprovider"
)

type BMGMUsecaseInterface interface {
	DoBMGM(ctx context.Context) error
}

type BMGMUseCases struct {
	CampaignRepos campaign.CampaignRepository
}

func NewBMGMUseCases(repos newprovider.DomainRepository) *BMGMUseCases {
	return &BMGMUseCases{CampaignRepos: repos}
}

// DoBMGM will get campaign and ...
func (uc *BMGMUseCases) DoBMGM(ctx context.Context) error {
	_, err := uc.CampaignRepos.GetCampaign(campaign.InputGetCampaign{})
	return err
}
