package ast

import (
	"fmt"
	"strings"
)

type Value interface {
	Kind() ValueKind
	String() string
}

////////////////////////////////////////////////////////////////

type ValueKind uint8

const (
	IntegerValueKind ValueKind = iota
	FloatValueKind
	StringValueKind
	ArrayValueKind
	IdentifierValueKind
)

////////////////////////////////////////////////////////////////

type IdentifierValue string

func (l IdentifierValue) Kind() ValueKind {
	return IdentifierValueKind
}

func (l IdentifierValue) Value() string {
	return string(l)
}

func (l IdentifierValue) String() string {
	return string(l)
}

////////////////////////////////////////////////////////////////

type StringValue string

func (l StringValue) Kind() ValueKind {
	return StringValueKind
}

func (l StringValue) Value() string {
	return string(l)
}

func (l StringValue) String() string {
	return fmt.Sprintf("%q", string(l))
}

////////////////////////////////////////////////////////////////

type IntegerValue int

func (l IntegerValue) Kind() ValueKind {
	return IntegerValueKind
}

func (l IntegerValue) Value() int {
	return int(l)
}

func (l IntegerValue) String() string {
	return fmt.Sprintf("%d", int(l))
}

////////////////////////////////////////////////////////////////

type FloatValue float64

func (l FloatValue) Kind() ValueKind {
	return FloatValueKind
}

func (l FloatValue) Value() float64 {
	return float64(l)
}

func (l FloatValue) String() string {
	s := strings.TrimRight(fmt.Sprintf("%0.2f", float64(l)), "0")
	if strings.HasSuffix(s, ".") {
		return s + "0"
	}
	return s
}

////////////////////////////////////////////////////////////////

type ArrayValue []Value

func (l ArrayValue) Kind() ValueKind {
	return ArrayValueKind
}

func (l ArrayValue) String() string {
	values := make([]string, len(l))
	for i, value := range l {
		values[i] = value.String()
	}
	return "{" + strings.Join(values, ", ") + "}"
}
