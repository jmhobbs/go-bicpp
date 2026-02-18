package ast

type Declaration interface {
	Kind() DeclarationKind
}

////////////////////////////////////////////////////////////////

type DeclarationKind uint8

const (
	VariableDeclarationKind DeclarationKind = iota
	ClassDeclarationKind
)

////////////////////////////////////////////////////////////////

type VariableDeclaration struct {
	Identifier string
	Value      Value
}

func (d VariableDeclaration) Kind() DeclarationKind {
	return VariableDeclarationKind
}

////////////////////////////////////////////////////////////////

type ClassDeclaration struct {
	Identifier string
	Fields     []VariableDeclaration
}

func (d ClassDeclaration) Kind() DeclarationKind {
	return ClassDeclarationKind
}
