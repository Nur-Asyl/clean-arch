package contact

import (
	"architecture_go/services/contact/internal/domain/contact"
	"architecture_go/services/contact/internal/useCase/adapters/storage"
	"context"
	"errors"
)

type ContactUseCase struct {
	contactRepo storage.Contact
}

func NewContactUseCase(contactRepo storage.Contact) *ContactUseCase {
	return &ContactUseCase{
		contactRepo: contactRepo,
	}
}

func (uc *ContactUseCase) CreateContact(ctx context.Context, firstName, lastName, phoneNumber string) (*contact.Contact, error) {
	newContact, err := contact.NewContact(firstName, lastName, phoneNumber)

	err = uc.contactRepo.CreateContact(ctx, newContact)
	if err != nil {
		return nil, err
	}

	return newContact, nil
}

func (uc *ContactUseCase) ReadContact(ctx context.Context, contactID int) (*contact.Contact, error) {
	return uc.contactRepo.ReadContact(ctx, contactID)
}

func (uc *ContactUseCase) UpdateContact(ctx context.Context, contactID int, fullname, phoneNumber string) error {
	existingContact, err := uc.contactRepo.ReadContact(ctx, contactID)
	if err != nil {
		return err
	}

	if existingContact == nil {
		return errors.New("contact not found")
	}

	existingContact.FullName = fullname
	existingContact.PhoneNumber = phoneNumber

	return uc.contactRepo.UpdateContact(ctx, existingContact)
}

func (uc *ContactUseCase) DeleteContact(ctx context.Context, contactID int) error {
	return uc.contactRepo.DeleteContact(ctx, contactID)
}
