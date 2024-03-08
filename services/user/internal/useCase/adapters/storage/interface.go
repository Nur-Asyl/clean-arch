package storage

import (
	"architecture_go/services/user/internal/domain/user"
	"context"
)

type User interface {
	CreateUser(ctx context.Context, user *user.User) error
	ReadUser(ctx context.Context, email string) (*user.User, error)
	UpdateUser(ctx context.Context, user *user.User) error
	DeleteUser(ctx context.Context, email string) error
}
