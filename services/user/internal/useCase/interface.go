package useCase

import (
	"architecture_go/services/user/internal/domain/user"
	"context"
)

type UserUseCase interface {
	CreateUser(ctx context.Context, firstName, lastName, email, password string) (*user.User, error)
	ReadUser(ctx context.Context, email string) (*user.User, error)
	UpdateUser(ctx context.Context, firstName, lastName, email, password string) error
	DeleteUser(ctx context.Context, email string) error
}
