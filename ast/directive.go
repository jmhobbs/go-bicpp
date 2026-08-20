package ast

import "fmt"

type Define struct {
	Identifier string
	Value      Node
}

func (d Define) String() string {
	return fmt.Sprintf("#define %s %s", d.Identifier, d.Value.String())
}
