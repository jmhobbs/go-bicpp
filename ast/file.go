package ast

import "strings"

// Node is an interface that represents a node in the AST.
type Node interface {
	String() string
}

// File is the root node of the AST, representing a complete configuration file.
type File []Node

func (f File) String() string {
	var builder strings.Builder

	for _, n := range f {
		builder.WriteString(n.String())
		builder.WriteString("\n")
	}

	return builder.String()
}
