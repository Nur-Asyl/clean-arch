package group

type Group struct {
	ID   int
	Name string
}

func NewGroup(name string) *Group {
	return &Group{
		Name: validateName(name),
	}
}

func validateName(name string) string {
	if len(name) > 250 {
		name = name[:250]
	}
	return name
}
