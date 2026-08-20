package ast

import (
	"fmt"
	"strings"
)

// Identifier is a string representing an identifier.
type Identifier string

func (l Identifier) Value() string {
	return string(l)
}

func (l Identifier) String() string {
	return string(l)
}

////////////////////////////////////////////////////////////////

// String represents a string literal value.
type String string

func (l String) Value() string {
	return string(l)
}

func (l String) String() string {
	return fmt.Sprintf("%q", string(l))
}

////////////////////////////////////////////////////////////////

// Integer represents an integer literal value.
type Integer int

func (l Integer) Value() int {
	return int(l)
}

func (l Integer) String() string {
	return fmt.Sprintf("%d", int(l))
}

////////////////////////////////////////////////////////////////

// Float represents a float literal value.
type Float float64

func (l Float) Value() float64 {
	return float64(l)
}

func (l Float) String() string {
	s := strings.TrimRight(fmt.Sprintf("%0.2f", float64(l)), "0")
	if strings.HasSuffix(s, ".") {
		return s + "0"
	}
	return s
}

////////////////////////////////////////////////////////////////

type Array []Node

func (a Array) String() string {
	values := make([]string, len(a))
	for i, value := range a {
		values[i] = value.String()
	}
	return "{" + strings.Join(values, ", ") + "}"
}
