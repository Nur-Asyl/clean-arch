package useCase

import "architecture_go/services/contact/internal/domain/contact"

type ContactUseCase interface {
	CreateContact(firstName, lastName, phoneNumber string) (*contact.Contact, error)
	GetContactByID(contactID int) (*contact.Contact, error)
	UpdateContact(contactID int, firstName, lastName, phoneNumber string) error
	DeleteContact(contactID int) error
}

type GroupUseCase interface {
	CreateGroup(name string) (*contact.Group, error)
	GetGroupByID(groupID int) (*contact.Group, error)
	AddContactToGroup(groupID, contactID int) error
}
