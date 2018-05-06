package ssml

import (
	"text/template"
)

const (
	attributeTemplate = "attribute"
	elementTemplate   = "element"
	nestedTemplate    = "nestedElement"
	textTemplate      = "text"
)

func build() *template.Template {
	var err error

	tmpl := template.New("root")
	_, err = tmpl.New(elementTemplate).Parse(
		`<{{.Name|html}}{{range .Attributes}}{{template "attribute" .}}{{end}}{{if .Empty}} />{{else}}>{{.Content}}</{{.Name}}>{{end}}`,
	)
	panicOnErr(err)
	_, err = tmpl.New(attributeTemplate).Parse(
		` {{.Name|html}}="{{.Value|html}}"`,
	)
	panicOnErr(err)
	_, err = tmpl.New(textTemplate).Parse(
		`{{.Content|html}}`,
	)
	panicOnErr(err)

	return tmpl
}
