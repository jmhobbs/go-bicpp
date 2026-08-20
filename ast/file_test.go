package ast_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/stretchr/testify/assert"
)

func Test_File_String(t *testing.T) {
	f := ast.File{
		Directives: []ast.Node{
			ast.Define{Identifier: "true", Value: ast.Integer(1)},
		},
		Declarations: []ast.Node{
			ast.Assignment{Identifier: "myVar", Value: ast.Integer(420)},
		},
	}

	expected := "#define true 1\nmyVar = 420;\n"

	assert.Equal(t, expected, f.String())
}
