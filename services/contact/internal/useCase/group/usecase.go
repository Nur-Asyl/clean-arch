package group

import (
	"architecture_go/services/contact/internal/domain/group"
	"architecture_go/services/contact/internal/useCase"
	"architecture_go/services/contact/internal/useCase/adapters/storage"
	"context"
)

type GroupUseCase struct {
	contactRepo storage.Contact
	groupRepo   storage.Group
}

func NewGroupUseCase(contactRepo storage.Contact, groupRepo storage.Group) useCase.GroupUseCase {
	return &GroupUseCase{contactRepo: contactRepo, groupRepo: groupRepo}
}

func (uc *GroupUseCase) CreateGroup(ctx context.Context, name string) (*group.Group, error) {
	newGroup := &group.Group{
		Name: name,
	}

	err := uc.groupRepo.CreateGroup(ctx, newGroup)
	if err != nil {
		return nil, err
	}

	return newGroup, nil
}

func (uc *GroupUseCase) ReadGroup(ctx context.Context, groupID int) (*group.Group, error) {
	return uc.groupRepo.ReadGroup(ctx, groupID)
}

func (uc *GroupUseCase) AddContactToGroup(ctx context.Context, groupID, contactID int) error {
	return uc.groupRepo.AddContactToGroup(ctx, groupID, contactID)
}
