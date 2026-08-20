package ast_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/stretchr/testify/assert"
)

func Test_BlockExpression_String(t *testing.T) {
	b := ast.BlockExpression{
		ast.AssignmentDeclaration{Identifier: "myVar", Value: ast.IntegerLiteral(42)},
	}

	expected := "{\nmyVar = 42;\n}"
	assert.Equal(t, expected, b.String())
}

func Test_BlockExpression_Empty_String(t *testing.T) {
	b := ast.BlockExpression{}

	expected := "{\n}"
	assert.Equal(t, expected, b.String())
}

func Test_ArrayExpression_String(t *testing.T) {
	a := ast.ArrayExpression{
		ast.IntegerLiteral(42),
		ast.FloatLiteral(3.14),
		ast.ArrayLiteral{
			Body: ast.ArrayExpression{
				ast.StringLiteral("nested"),
			},
		},
		ast.StringLiteral("hello"),
	}

	expected := "{42, 3.14, {\"nested\"}, \"hello\"}"
	assert.Equal(t, expected, a.String())
}
