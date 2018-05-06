package ssml

import (
	"fmt"
	"testing"
)

func TestRender(t *testing.T) {
	elem := &element{
		name: "speak",
		content: []Element{
			&textElement{
				content: "helo < asdf",
			},
			&element{
				name: "break",
			},
			&element{
				name: "emph",
				attributes: []Attribute{
					&attribute{
						name:  "strength",
						value: "str\"ong",
					},
				},

				content: []Element{
					&textElement{
						content: "helo < asdf",
					},
				},
			},
		},
	}

	str, err := RenderElement(elem)
	fmt.Println(err)
	fmt.Println(str)
}
