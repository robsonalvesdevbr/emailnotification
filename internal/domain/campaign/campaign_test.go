package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "Robson Candido dos Santos Alves"
	content = "body"
	emails  = []string{"robson@testuser.com", "admin@testuser.com"}
)

func Test_Campaign_Create(t *testing.T) {
	asserts := assert.New(t)

	var contacts []Contact
	for _, email := range emails {
		contacts = append(contacts, Contact{Email: email})
	}

	campaign := NewCampaign(name, content, emails)
	asserts.NotNil(campaign)
	asserts.NotNil(campaign.ID)
	asserts.Equal(name, campaign.Name)
	asserts.Equal(content, campaign.Content)
	asserts.Equal(contacts, campaign.Contacts)
}
