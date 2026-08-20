package ast

import (
	"strings"
)

type Assignment struct {
	Identifier string
	Value      Node
}

func (d Assignment) String() string {
	return d.Identifier + " = " + d.Value.String() + ";"
}

////////////////////////////////////////////////////////////////

type Class struct {
	Identifier string
	Parent     string
	Body       Block
}

func (d Class) String() string {
	var builder strings.Builder
	builder.WriteString("class " + d.Identifier)
	if d.Parent != "" {
		builder.WriteString(" : " + d.Parent)
	}
	if d.Body != nil {
		builder.WriteString("\n")
		builder.WriteString(d.Body.String())
	}
	builder.WriteString(";")
	return builder.String()
}
