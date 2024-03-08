package user

import (
	"architecture_go/services/user/internal/domain/user"
	"context"
	"database/sql"
	"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *user.User) error {
	existingUser, _ := r.ReadUser(ctx, user.Email)
	if existingUser != nil {
		return errors.New("user already exist")
	}
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) ReadUser(ctx context.Context, email string) (*user.User, error) {
	var user user.User
	err := r.db.QueryRowContext(ctx, "SELECT id, first_name, last_name, email, password FROM users WHERE email = $1", email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, newUser *user.User) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET first_name = $1, last_name = $2, email = $3, password = $4 WHERE email = $5", newUser.FirstName, newUser.LastName, newUser.Email, newUser.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, email string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE email = $1", email)
	if err != nil {
		return err
	}
	return nil
}
