package ast

import "strings"

type Program struct {
	Definitions  map[string]Value
	Declarations []Declaration
}

func (p *Program) String() string {
	var builder strings.Builder

	for identifier, def := range p.Definitions {
		builder.WriteString("#define " + identifier + " ")
		builder.WriteString(def.String())
		builder.WriteString("\n")
	}

	for _, decl := range p.Declarations {
		builder.WriteString(decl.String())
		builder.WriteString("\n")
	}

	return builder.String()
}

func (p *Program) Define(identifier string, value Value) {
	p.Definitions[identifier] = value
}

func (p *Program) Declare(identifier string, value Value) {
	p.Declarations = append(p.Declarations, VariableDeclaration{Identifier: identifier, Value: value})
}

func (p *Program) ForwardDeclareClass(identifier string) {
	p.Declarations = append(p.Declarations, ForwardClassDeclaration{Identifier: identifier})
}

func (p *Program) DeclareClass(identifier, parent string) {
	p.Declarations = append(p.Declarations, ClassDeclaration{
		Identifier: identifier,
		Parent:     parent,
		Fields:     []VariableDeclaration{},
	})
}
