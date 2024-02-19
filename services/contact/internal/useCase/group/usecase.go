package group

import (
	"architecture_go/services/contact/internal/domain/group"
	"architecture_go/services/contact/internal/useCase"
	"architecture_go/services/contact/internal/useCase/adapters/storage"
)

type GroupUseCase struct {
	repo storage.GroupRepository
}

func NewUseCase(repo storage.GroupRepository) useCase.GroupUseCase {
	return &GroupUseCase{repo: repo}
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
