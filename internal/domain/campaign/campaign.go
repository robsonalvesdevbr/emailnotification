package campaign

import (
	"time"

	"github.com/google/uuid"
)

type Campaign struct {
	ID       uuid.UUID
	Name     string
	CreateAt time.Time
	Content  string
	Contacts []Contact
}

func NewCampaign(name, content string, contacts []string) *Campaign {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	var contactsList []Contact
	for _, contact := range contacts {
		contactsList = append(contactsList, Contact{Email: contact})
	}

	return &Campaign{
		ID:       id,
		Name:     name,
		Content:  content,
		CreateAt: time.Now(),
		Contacts: contactsList,
	}
}
