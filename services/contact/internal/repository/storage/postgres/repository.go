package postgres

import (
	"architecture_go/pkg/store/postgres"
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"architecture_go/services/contact/internal/useCase/adapters/storage"
)

type Repository struct {
	db *postgres.Storage
}

func NewContactRepository(db *postgres.Storage) storage.ContactRepository {
	return &Repository{db: db}
}

func (r *Repository) CreateContact(contact *contact.Contact) (*contact.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetContactById(contactID int) (*contact.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) UpdateContact(contact *contact.Contact) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) DeleteContact(contactID int) error {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) CreateGroup(group *group.Group) (*group.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) GetGroupById(groupID int) (*group.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (r *Repository) AddContactToGroup(groupID, contactID int) error {
	//TODO implement me
	panic("implement me")
}
