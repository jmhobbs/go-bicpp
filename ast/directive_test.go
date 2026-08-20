package ast_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/stretchr/testify/assert"
)

func Test_Define_String(t *testing.T) {
	d := ast.Define{
		Identifier: "myVar",
		Value:      ast.Integer(42),
	}
	assert.Equal(t, "#define myVar 42", d.String())
}
