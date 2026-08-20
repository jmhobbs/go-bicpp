package ast

import "strings"

// Node is an interface that represents a node in the AST.
type Node interface {
	String() string
}

// File is the root node of the AST, representing a complete configuration file.
type File struct {
	Directives   []Node
	Declarations []Node
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
