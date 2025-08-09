package campaign

import (
	"github.com/google/uuid"
	"github.com/robsonalvesdevbr/emailnotification/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (uuid.UUID, error) {
	var emails []Contact
	for contacts := range newCampaign.Emails {
		emails = append(emails, Contact{Email: newCampaign.Emails[contacts]})
	}
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, emails)
	if err != nil {
		return uuid.UUID{}, err
	}
	return campaign.ID, nil
}
