package ast

type Value interface {
	Kind() ValueKind
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

////////////////////////////////////////////////////////////////

type StringValue string

func (l StringValue) Kind() ValueKind {
	return StringValueKind
}

func (l StringValue) Value() string {
	return string(l)
}

////////////////////////////////////////////////////////////////

type IntegerValue int

func (l IntegerValue) Kind() ValueKind {
	return IntegerValueKind
}

func (l IntegerValue) Value() int {
	return int(l)
}

////////////////////////////////////////////////////////////////

type FloatValue float64

func (l FloatValue) Kind() ValueKind {
	return FloatValueKind
}

func (l FloatValue) Value() float64 {
	return float64(l)
}

////////////////////////////////////////////////////////////////

type ArrayValue []Value

func (l ArrayValue) Kind() ValueKind {
	return ArrayValueKind
}
