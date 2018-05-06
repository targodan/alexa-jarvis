package ssml

import (
    "bytes"
	"text/template"
)

var rootTmpl *template.Template

func RenderElement(elem Element) (string, error) {
    return render(elementTemplate, elem)
}

func RenderAttribute(elem Attribute) (string, error) {
    return render(attributeTemplate, elem)
}

func RenderText(elem Element) (string, error) {
    return render(textTemplate, elem)
}

func render(name string, data interface{}) (string, error) {
	if rootTmpl == nil {
		rootTmpl = build()
	}
	buf := &bytes.Buffer{}
	err := rootTmpl.ExecuteTemplate(buf, name, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
