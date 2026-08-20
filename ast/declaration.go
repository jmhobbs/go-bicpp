package ast

import (
	"strings"
)

type Declaration interface {
	Kind() DeclarationKind
	String() string
}

////////////////////////////////////////////////////////////////

//go:generate go tool stringer -type=DeclarationKind
type DeclarationKind uint8

const (
	Assignment DeclarationKind = iota
	Class
)

////////////////////////////////////////////////////////////////

type AssignmentDeclaration struct {
	Identifier string
	Value      Literal
}

func (d AssignmentDeclaration) Kind() DeclarationKind {
	return Assignment
}

func (d AssignmentDeclaration) String() string {
	return d.Identifier + " = " + d.Value.String() + ";"
}

////////////////////////////////////////////////////////////////

type ClassDeclaration struct {
	Identifier string
	Parent     string
	Body       BlockExpression
}

func (d ClassDeclaration) Kind() DeclarationKind {
	return Class
}

func (d ClassDeclaration) String() string {
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
