package ast

type Program struct {
	Definitions  map[string]Value
	Declarations map[string]Declaration
}

func (p *Program) Define(identifier string, value Value) {
	p.Definitions[identifier] = value
}

func (p *Program) Declare(identifier string, value Value) {
	p.Declarations[identifier] = VariableDeclaration{Identifier: identifier, Value: value}
}
