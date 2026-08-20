package ast

import (
	"strings"
)

// Assignment represents a variable/field assignment using the `=` operator.
//
//	// example
//	myVar = 1234;
type Assignment struct {
	Identifier string
	Value      Node
}

func (d Assignment) String() string {
	return d.Identifier + " = " + d.Value.String() + ";"
}

////////////////////////////////////////////////////////////////

// Class represents a class definition.
type Class struct {
	Identifier string
	Parent     string
	Body       Block
}

// IsForwardClassDeclaration returns true if the class is a forward declaration with no body.
//
//	// example
//	class CfgBase;
func (c Class) IsForwardClassDeclaration() bool {
	return c.Body == nil
}

// HasAncestor returns true if the class has a parent class.
//
//	// example
//	class CfgChild : CfgParent { ... };
func (c Class) HasAncestor() bool {
	return c.Parent != ""
}

func (c Class) String() string {
	var builder strings.Builder
	builder.WriteString("class " + c.Identifier)
	if c.HasAncestor() {
		builder.WriteString(" : " + c.Parent)
	}
	if !c.IsForwardClassDeclaration() {
		builder.WriteString("\n")
		builder.WriteString(c.Body.String())
	}
	builder.WriteString(";")
	return builder.String()
}

////////////////////////////////////////////////////////////////

type Comment string

func (c Comment) String() string {
	return "// " + string(c)
}
