package storage

import "architecture_go/services/contact/internal/domain/contact"

type ContactRepository interface {
	Create(contact *contact.Contact) (*contact.Contact, error)
	GetById(contactID int) (*contact.Contact, error)
	Update(contact *contact.Contact) error
	Delete(contactID int) error
}

type GroupRepository interface {
	Create(group *contact.Group) (*contact.Group, error)
	GetById(groupID int) (*contact.Group, error)
	AddContactToGroup(groupID, contactID int) error
}
