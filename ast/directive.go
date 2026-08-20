package ast

import "fmt"

type Directive interface {
	Kind() DirectiveKind
	String() string
}

////////////////////////////////////////////////////////////////

//go:generate go tool stringer -type=DirectiveKind
type DirectiveKind uint8

const (
	Define DirectiveKind = iota
)

////////////////////////////////////////////////////////////////

type DefineDirective struct {
	Identifier string
	Value      Literal
}

func (d DefineDirective) Kind() DirectiveKind {
	return Define
}

func (d DefineDirective) String() string {
	return fmt.Sprintf("#define %s %s", d.Identifier, d.Value.String())
}
