package storage

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"context"
)

type Contact interface {
	CreateContact(ctx context.Context, contact *contact.Contact) error
	ReadContact(ctx context.Context, contactID int) (*contact.Contact, error)
	UpdateContact(ctx context.Context, contact *contact.Contact) error
	DeleteContact(ctx context.Context, contactID int) error
}

type Group interface {
	CreateGroup(ctx context.Context, group *group.Group) error
	ReadGroup(ctx context.Context, groupID int) (*group.Group, error)
	AddContactToGroup(ctx context.Context, groupID, contactID int) error
}
