package group

import "errors"

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
