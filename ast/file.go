package ast

import "strings"

type File struct {
	Directives   []Directive
	Declarations []Declaration
}

func (f *File) String() string {
	var builder strings.Builder

	for _, d := range f.Directives {
		builder.WriteString(d.String())
		builder.WriteString("\n")
	}

	for _, d := range f.Declarations {
		builder.WriteString(d.String())
		builder.WriteString("\n")
	}

	return builder.String()
}
