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
	campaign := NewCampaign(name, content, emails)
	asserts.Equal(name, campaign.Name)
	asserts.Equal(content, campaign.Content)
	//asserts.Equal(emails, campaign.Contacts)

}
