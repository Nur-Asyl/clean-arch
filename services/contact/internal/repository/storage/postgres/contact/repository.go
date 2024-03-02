package contact

import (
	"architecture_go/services/contact/internal/domain/contact"
	"context"
	"database/sql"
	"errors"
)

type ContactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) CreateContact(ctx context.Context, contact *contact.Contact) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO contacts (id, full_name, phone_number) VALUES ($1, $2, $3)", contact.ID, contact.FullName, contact.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (r *ContactRepository) ReadContact(ctx context.Context, contactID int) (*contact.Contact, error) {
	var contact contact.Contact
	err := r.db.QueryRowContext(ctx, "SELECT id, full_name, phone_number FROM contacts WHERE id = $1", contactID).Scan(&contact.ID, &contact.FullName, &contact.PhoneNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("contact not found")
		}
		return nil, err
	}
	return &contact, nil
}

func (r *ContactRepository) UpdateContact(ctx context.Context, newContact *contact.Contact) error {
	_, err := r.db.ExecContext(ctx, "UPDATE contacts SET fullname = $1, phone_number = $2 WHERE id = $3", newContact.FullName, newContact.PhoneNumber, newContact.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ContactRepository) DeleteContact(ctx context.Context, contactID int) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM contacts WHERE id = $1", contactID)
	if err != nil {
		return err
	}
	return nil
}
