package ast

import (
	"strings"
)

type Declaration interface {
	Kind() DeclarationKind
	String() string
}

////////////////////////////////////////////////////////////////

type DeclarationKind uint8

const (
	VariableDeclarationKind DeclarationKind = iota
	ClassDeclarationKind
	ForwardClassDeclarationKind
)

////////////////////////////////////////////////////////////////

type VariableDeclaration struct {
	Identifier string
	Value      Value
}

func (d VariableDeclaration) Kind() DeclarationKind {
	return VariableDeclarationKind
}

func (d VariableDeclaration) String() string {
	// TODO: Arrays
	return d.Identifier + " = " + d.Value.String() + ";"
}

////////////////////////////////////////////////////////////////

type ForwardClassDeclaration struct {
	Identifier string
}

func (d ForwardClassDeclaration) Kind() DeclarationKind {
	return ForwardClassDeclarationKind
}

func (d ForwardClassDeclaration) String() string {
	return "class " + d.Identifier + ";"
}

type ClassDeclaration struct {
	Identifier string
	Parent     string
	Fields     []Declaration
}

func (d ClassDeclaration) Kind() DeclarationKind {
	return ClassDeclarationKind
}

func (d ClassDeclaration) String() string {
	lines := make([]string, len(d.Fields)+3)
	lines[0] = "class " + d.Identifier
	if d.Parent != "" {
		lines[0] = lines[0] + " : " + d.Parent
	}
	lines[1] = "{"
	lines[len(lines)-1] = "}"
	for i, f := range d.Fields {
		lines[i+2] = f.String()
	}
	return strings.Join(lines, "\n")
}
