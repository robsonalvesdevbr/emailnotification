package campaign

import (
	"testing"

	"github.com/robsonalvesdevbr/emailnotification/internal/contract"
	"github.com/stretchr/testify/assert"
)

func Test_Campaign_Service_Create(t *testing.T) {
	asserts := assert.New(t)
	service := Service{}
	newCampaign := contract.NewCampaign{
		Name:    "Test Y",
		Content: "body",
		Emails:  []string{"test@testuser.com"},
	}
	id, err := service.Create(newCampaign)
	asserts.Nil(err)
	asserts.NotNil(id)
}
