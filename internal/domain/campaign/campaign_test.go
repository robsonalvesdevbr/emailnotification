package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "Robson Candido dos Santos Alves"
	content = "body"
	emails  = []Contact{{Email: "robson@testuser.com"}, {Email: "admin@testuser.com"}}
)

func Test_Campaign_Create(t *testing.T) {
	asserts := assert.New(t)

	campaign := NewCampaign(name, content, emails)

	asserts.NotNil(campaign)
	asserts.NotNil(campaign.ID)
	asserts.Equal(name, campaign.Name)
	asserts.Equal(content, campaign.Content)
	asserts.Equal(emails, campaign.Contacts)
}
