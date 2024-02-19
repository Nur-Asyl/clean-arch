package group

import (
	"architecture_go/services/contact/internal/domain/group"
	"architecture_go/services/contact/internal/useCase"
	"architecture_go/services/contact/internal/useCase/adapters/storage"
)

type GroupUseCase struct {
	contactRepo storage.Contact
	groupRepo   storage.Group
}

func NewGroupUseCase(contactRepo storage.Contact, groupRepo storage.Group) useCase.GroupUseCase {
	return &GroupUseCase{contactRepo: contactRepo, groupRepo: groupRepo}
}

func (g *GroupUseCase) CreateGroup(name string) (*group.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GroupUseCase) GetGroupByID(groupID int) (*group.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GroupUseCase) AddContactToGroup(groupID, contactID int) error {
	//TODO implement me
	panic("implement me")
}
