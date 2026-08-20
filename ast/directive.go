package ast

import "fmt"

// Define is a preprocessor directive that defines a constant value.
//
//	// example
//	#define true 1
type Define struct {
	Identifier string
	Value      Node
}

func (d Define) String() string {
	return fmt.Sprintf("#define %s %s", d.Identifier, d.Value.String())
}
