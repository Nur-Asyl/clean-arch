package group

type Group struct {
	ID   int
	Name string
}

func NewGroup(id int, name string) *Group {
	return &Group{
		ID:   id,
		Name: validateName(name),
	}
}

func validateName(name string) string {
	if len(name) > 250 {
		name = name[:250]
	}
	return name
}
