package ast_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/stretchr/testify/assert"
)

func Test_DefineDirective_String(t *testing.T) {
	d := ast.DefineDirective{
		Identifier: "myVar",
		Value:      ast.IntegerLiteral(42),
	}
	assert.Equal(t, "#define myVar 42", d.String())
}
