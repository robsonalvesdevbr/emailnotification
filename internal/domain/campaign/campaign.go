package campaign

import (
	"errors"
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

func NewCampaign(name, content string, contacts []Contact) (*Campaign, error) {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}

	if name == "" {
		return nil, errors.New("name is required")
	}
	if content == "" {
		return nil, errors.New("content is required")
	}
	if len(contacts) == 0 {
		return nil, errors.New("contacts is required")
	}

	return &Campaign{
		ID:       id,
		Name:     name,
		Content:  content,
		CreateAt: time.Now(),
		Contacts: contacts,
	}, nil
}
