package ast_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/stretchr/testify/assert"
)

func Test_Block_String(t *testing.T) {
	b := ast.Block{
		ast.Assignment{Identifier: "myVar", Value: ast.Integer(42)},
	}

	expected := "{\nmyVar = 42;\n}"
	assert.Equal(t, expected, b.String())
}

func Test_Block_Empty_String(t *testing.T) {
	b := ast.Block{}

	expected := "{\n}"
	assert.Equal(t, expected, b.String())
}

func Test_ArrayBlock_String(t *testing.T) {
	a := ast.Array{
		ast.Integer(42),
		ast.Float(3.14),
		ast.Array{
			ast.String("nested"),
		},
		ast.String("hello"),
	}

	expected := "{42, 3.14, {\"nested\"}, \"hello\"}"
	assert.Equal(t, expected, a.String())
}
