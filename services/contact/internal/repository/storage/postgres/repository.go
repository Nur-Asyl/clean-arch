package postgres

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
)

type ContactRepository struct {
	db *postgres.Storage
}

func NewContactRepository(db *postgres.Storage) ContactRepository {
	return ContactRepository{db: db}
}

func (r *ContactRepository) CreateContact(contact *contact.Contact) (*contact.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ContactRepository) GetContactById(contactID int) (*contact.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ContactRepository) UpdateContact(contact *contact.Contact) error {
	//TODO implement me
	panic("implement me")
}

func (r *ContactRepository) DeleteContact(contactID int) error {
	//TODO implement me
	panic("implement me")
}

func (r *ContactRepository) CreateGroup(group *group.Group) (*group.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ContactRepository) GetGroupById(groupID int) (*group.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ContactRepository) AddContactToGroup(groupID, contactID int) error {
	//TODO implement me
	panic("implement me")
}
