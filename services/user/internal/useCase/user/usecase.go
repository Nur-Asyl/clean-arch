package user

import (
	"architecture_go/services/user/internal/domain/user"
	"architecture_go/services/user/internal/useCase/adapters/storage"
	"context"
	"errors"
)

type UserUseCase struct {
	userRepo storage.User
}

func NewUserUseCase(userRepo storage.User) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, firstName, lastName, email, password string) (*user.User, error) {
	newUser, err := user.NewUser(firstName, lastName, email, password)
	if err != nil {
		return nil, err
	}

	err = uc.userRepo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (uc *UserUseCase) ReadUser(ctx context.Context, email string) (*user.User, error) {
	return uc.userRepo.ReadUser(ctx, email)
}

func (uc *UserUseCase) UpdateUser(ctx context.Context, firstName, lastName, email, password string) error {
	existingUser, err := uc.userRepo.ReadUser(ctx, email)
	if err != nil {
		return err
	}

	if existingUser == nil {
		return errors.New("user not found")
	}

	existingUser.FirstName = firstName
	existingUser.LastName = lastName
	existingUser.Email = email
	existingUser.Password = password

	return uc.userRepo.UpdateUser(ctx, existingUser)
}

func (uc *UserUseCase) DeleteUser(ctx context.Context, email string) error {
	return uc.userRepo.DeleteUser(ctx, email)
}
