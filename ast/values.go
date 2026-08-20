package ast

import (
	"fmt"
	"strings"
)

////////////////////////////////////////////////////////////////

type Identifier string

func (l Identifier) Value() string {
	return string(l)
}

func (l Identifier) String() string {
	return string(l)
}

////////////////////////////////////////////////////////////////

type String string

func (l String) Value() string {
	return string(l)
}

func (l String) String() string {
	return fmt.Sprintf("%q", string(l))
}

////////////////////////////////////////////////////////////////

type Integer int

func (l Integer) Value() int {
	return int(l)
}

func (l Integer) String() string {
	return fmt.Sprintf("%d", int(l))
}

////////////////////////////////////////////////////////////////

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

type Array struct {
	Body ArrayBlock
}

func (a Array) String() string {
	return a.Body.String()
}
