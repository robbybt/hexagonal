package bmgm

import (
	"context"
	"hexagonal/domain/campaign"
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

func (uc *BMGMUseCases) DoBMGM(ctx context.Context) error {
	_, err := uc.CampaignRepos.GetCampaign(campaign.InputGetCampaign{})
	return err
}
