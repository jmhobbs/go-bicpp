package ast

import "strings"

type Program struct {
	Definitions  []Definition
	Declarations []Declaration
}

func (p *Program) String() string {
	var builder strings.Builder

	for _, def := range p.Definitions {
		builder.WriteString("#define " + def.Identifier + " ")
		builder.WriteString(def.Value.String())
		builder.WriteString("\n")
	}

	for _, decl := range p.Declarations {
		builder.WriteString(decl.String())
		builder.WriteString("\n")
	}

	return builder.String()
}
