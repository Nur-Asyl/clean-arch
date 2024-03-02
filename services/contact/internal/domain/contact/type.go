package contact

import (
	"errors"
	"fmt"
	"regexp"
)

type Contact struct {
	ID          int
	FullName    string
	PhoneNumber string
}

func NewContact(firstName string, lastName string, phoneNumber string) (*Contact, error) {

	match, _ := regexp.MatchString("^[0-9]+$", phoneNumber)
	if !match {
		return nil, errors.New("phone number must contain only digits")
	}

	fullName := fmt.Sprintf("%s %s", firstName, lastName)

	contact := &Contact{
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}

	return contact, nil
}
