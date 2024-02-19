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

func NewContact(id int, firstName string, lastName string, phoneNumber string) (*Contact, error) {

	match, _ := regexp.MatchString("^[0-9]+$", phoneNumber)
	if !match {
		return nil, errors.New("phone number must contain only digits")
	}

	fullName := fmt.Sprintf("%s %s", firstName, lastName)

	contact := &Contact{
		ID:          id,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}

	return contact, nil
}

type Group struct {
	ID   int
	Name string
}

func NewGroup(id int, name string) (*Group, error) {

	if len(name) > 250 {
		return nil, errors.New("group name is more than the maximum length of 250 characters")
	}

	group := &Group{
		ID:   id,
		Name: name,
	}

	return group, nil
}
