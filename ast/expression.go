package ast

import "strings"

type BlockExpression []Declaration

func (b BlockExpression) String() string {
	lines := make([]string, len(b)+2)
	lines[0] = "{"
	lines[len(lines)-1] = "}"
	for i, d := range b {
		lines[i+1] = d.String()
	}
	return strings.Join(lines, "\n")
}

type ArrayExpression []Literal

func (a ArrayExpression) String() string {
	values := make([]string, len(a))
	for i, value := range a {
		values[i] = value.String()
	}
	return "{" + strings.Join(values, ", ") + "}"
}
