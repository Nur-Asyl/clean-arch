package contact

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/domain/group"
	"architecture_go/services/contact/internal/useCase"
	"architecture_go/services/contact/internal/useCase/adapters/storage"
)

type ContactUseCase struct {
	repo storage.ContactRepository
}

func NewUseCase(repo storage.ContactRepository) useCase.ContactUseCase {
	return &ContactUseCase{repo: repo}
}

func (c *ContactUseCase) CreateContact(firstName, lastName, phoneNumber string) (*contact.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContactUseCase) GetContactByID(contactID int) (*contact.Contact, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContactUseCase) UpdateContact(contactID int, firstName, lastName, phoneNumber string) error {
	//TODO implement me
	panic("implement me")
}

func (c *ContactUseCase) DeleteContact(contactID int) error {
	//TODO implement me
	panic("implement me")
}

func (c ContactUseCase) CreateGroup(name string) (*group.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContactUseCase) GetGroupByID(groupID int) (*group.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ContactUseCase) AddContactToGroup(groupID, contactID int) error {
	//TODO implement me
	panic("implement me")
}
