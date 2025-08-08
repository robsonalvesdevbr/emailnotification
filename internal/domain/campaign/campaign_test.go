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

func Test_Campaign_CreateValid(t *testing.T) {
	asserts := assert.New(t)

	campaign, err := NewCampaign(name, content, emails)

	asserts.Nil(err)

	asserts.NotNil(campaign)
	asserts.NotNil(campaign.ID)
	asserts.Equal(name, campaign.Name)
	asserts.Equal(content, campaign.Content)
	asserts.Equal(emails, campaign.Contacts)
}

func Test_Campaign_Create_EmptyName(t *testing.T) {
	asserts := assert.New(t)
	_, err := NewCampaign("", content, emails)
	asserts.NotNil(err)
	asserts.Equal("name is required", err.Error())
}

func Test_Campaign_Create_EmptyContent(t *testing.T) {
	asserts := assert.New(t)
	_, err := NewCampaign(name, "", emails)
	asserts.NotNil(err)
	asserts.Equal("content is required", err.Error())
}

func Test_Campaign_Create_EmptyContacts(t *testing.T) {
	asserts := assert.New(t)
	_, err := NewCampaign(name, content, []Contact{})
	asserts.NotNil(err)
	asserts.Equal("contacts is required", err.Error())
}
