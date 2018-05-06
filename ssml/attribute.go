package ssml

import "fmt"

type Attribute interface {
	Name() string
	Value() string

	fmt.Stringer
}

type attribute struct {
	name  string
	value string
}

func (a *attribute) Name() string {
	return a.name
}

func (a *attribute) Value() string {
	return a.value
}

func (a *attribute) String() string {
	str, err := RenderAttribute(a)
	panicOnErr(err)
	return str
}
