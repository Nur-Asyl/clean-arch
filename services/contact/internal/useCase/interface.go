package useCase

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
)

type ContactUseCase interface {
	CreateContact(firstName, lastName, phoneNumber string) (*contact.Contact, error)
	GetContactByID(contactID int) (*contact.Contact, error)
	UpdateContact(contactID int, firstName, lastName, phoneNumber string) error
	DeleteContact(contactID int) error
}

type GroupUseCase interface {
	CreateGroup(name string) (*group.Group, error)
	GetGroupByID(groupID int) (*group.Group, error)
	AddContactToGroup(groupID, contactID int) error
}
