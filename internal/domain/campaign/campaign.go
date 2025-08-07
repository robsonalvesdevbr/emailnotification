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

func NewCampaign(name, content string, contacts []Contact) *Campaign {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	return &Campaign{
		ID:       id,
		Name:     name,
		Content:  content,
		CreateAt: time.Now(),
		Contacts: contacts,
	}
}
