package ast

import (
	"fmt"
	"strings"
)

type Literal interface {
	Kind() LiteralKind
	String() string
}

////////////////////////////////////////////////////////////////

type LiteralKind uint8

const (
	IntegerLiteralKind LiteralKind = iota
	FloatLiteralKind
	StringLiteralKind
	ArrayLiteralKind
	IdentifierLiteralKind
)

////////////////////////////////////////////////////////////////

type IdentifierLiteral string

func (l IdentifierLiteral) Kind() LiteralKind {
	return IdentifierLiteralKind
}

func (l IdentifierLiteral) Value() string {
	return string(l)
}

func (l IdentifierLiteral) String() string {
	return string(l)
}

////////////////////////////////////////////////////////////////

type StringLiteral string

func (l StringLiteral) Kind() LiteralKind {
	return StringLiteralKind
}

func (l StringLiteral) Value() string {
	return string(l)
}

func (l StringLiteral) String() string {
	return fmt.Sprintf("%q", string(l))
}

////////////////////////////////////////////////////////////////

type IntegerLiteral int

func (l IntegerLiteral) Kind() LiteralKind {
	return IntegerLiteralKind
}

func (l IntegerLiteral) Value() int {
	return int(l)
}

func (l IntegerLiteral) String() string {
	return fmt.Sprintf("%d", int(l))
}

////////////////////////////////////////////////////////////////

type FloatLiteral float64

func (l FloatLiteral) Kind() LiteralKind {
	return FloatLiteralKind
}

func (l FloatLiteral) Value() float64 {
	return float64(l)
}

func (l FloatLiteral) String() string {
	s := strings.TrimRight(fmt.Sprintf("%0.2f", float64(l)), "0")
	if strings.HasSuffix(s, ".") {
		return s + "0"
	}
	return s
}

////////////////////////////////////////////////////////////////

type ArrayLiteral struct {
	Body ArrayExpression
}

func (a ArrayLiteral) Kind() LiteralKind {
	return ArrayLiteralKind
}

func (a ArrayLiteral) String() string {
	return a.Body.String()
}
