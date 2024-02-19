package storage

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
)

type Contact interface {
	CreateContact(contact *contact.Contact) (*contact.Contact, error)
	GetContactById(contactID int) (*contact.Contact, error)
	UpdateContact(contact *contact.Contact) error
	DeleteContact(contactID int) error
}

type Group interface {
	CreateGroup(group *group.Group) (*group.Group, error)
	GetGroupById(groupID int) (*group.Group, error)
	AddContactToGroup(groupID, contactID int) error
}
