package contact

import (
	"architecture_go/services/contact/internal/domain/group"
	"context"
	"database/sql"
	"errors"
)

type GroupRepository struct {
	db *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{db: db}
}

func (r *GroupRepository) CreateGroup(ctx context.Context, group *group.Group) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO groups (name) VALUES ($1)", group.Name)
	if err != nil {
		return err
	}
	return nil
}

func (r *GroupRepository) ReadGroup(ctx context.Context, groupID int) (*group.Group, error) {
	var id int
	var name string

	err := r.db.QueryRowContext(ctx, "SELECT id, name FROM groups WHERE id = $1", groupID).Scan(&id, &name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("group not found")
		}
		return nil, err
	}
	return group.NewGroup(name), nil
}

func (r *GroupRepository) AddContactToGroup(ctx context.Context, contactID, groupID int) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO contact_groups (contact_id, group_id) VALUES ($1, $2)", contactID, groupID)
	if err != nil {
		return err
	}
	return nil
}
