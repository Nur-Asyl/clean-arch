package useCase

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"context"
)

type ContactUseCase interface {
	CreateContact(ctx context.Context, fullname, phoneNumber string) (*contact.Contact, error)
	ReadContact(ctx context.Context, contactID int) (*contact.Contact, error)
	UpdateContact(ctx context.Context, contactID int, fullname, phoneNumber string) error
	DeleteContact(ctx context.Context, contactID int) error
}

type GroupUseCase interface {
	CreateGroup(ctx context.Context, name string) (*group.Group, error)
	ReadGroup(ctx context.Context, groupID int) (*group.Group, error)
	AddContactToGroup(ctx context.Context, groupID, contactID int) error
}
