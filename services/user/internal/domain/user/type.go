package user

import (
	"errors"
	"regexp"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

func NewUser(firstName, lastName, email, password string) (*User, error) {
	if !validatePassword(password) {
		return nil, errors.New("password length must be greater or equal to 5")
	}
	if !validateEmail(email) {
		return nil, errors.New("email format incorrect")
	}
	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}
	return &user, nil
}

func validatePassword(password string) bool {
	if len(password) < 5 {
		return false
	}
	return true
}

func validateEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)

	return re.MatchString(email)
}
