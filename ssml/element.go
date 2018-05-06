package ssml

import "fmt"

const TextElementName = "~~TEXT~~"

type Element interface {
	// Name returns the name of the element.
	Name() string
	// Attributes returns the attributes of the element.
	Attributes() []Attribute
	// SubElements returns the contained sub elements.
	SubElements() []Element
	Empty() bool

	Content() string

	addAttribute(elem Attribute)
	addSubElement(elem Element)

	fmt.Stringer
}

type element struct {
	name       string
	attributes []Attribute
	content    []Element
}

func newElement(name string) *element {
	return &element{
		name:       name,
		attributes: make([]Attribute, 0),
		content:    make([]Element, 0),
	}
}

func (e *element) Name() string {
	return e.name
}

func (e *element) Attributes() []Attribute {
	return e.attributes
}

func (e *element) Empty() bool {
	return len(e.content) == 0
}

func (e *element) Content() string {
	content := ""
	for _, elem := range e.content {
		content += elem.String()
	}
	return content
}

func (e *element) IsPureText() bool {
	return false
}

func (e *element) SubElements() []Element {
	return e.content
}

func (e *element) addAttribute(attr Attribute) {
	e.attributes = append(e.attributes, attr)
}

func (e *element) addSubElement(elem Element) {
	e.content = append(e.content, elem)
}

func (e *element) String() string {
	text, err := RenderElement(e)
	panicOnErr(err)
	return text
}

type textElement struct {
	content string
}

func (e *textElement) Name() string {
	return ""
}

func (e *textElement) Attributes() []Attribute {
	return []Attribute{}
}

func (e *textElement) IsPureText() bool {
	return true
}

func (e *textElement) Empty() bool {
	return false
}

func (e *textElement) Content() string {
	return e.content
}

func (e *textElement) SubElements() []Element {
	return []Element{}
}

func (e *textElement) addAttribute(attr Attribute) {}

func (e *textElement) addSubElement(elem Element) {}

func (e *textElement) String() string {
	str, err := RenderText(e)
	panicOnErr(err)
	return str
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
