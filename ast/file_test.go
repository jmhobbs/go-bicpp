package ast_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/stretchr/testify/assert"
)

func Test_File_String(t *testing.T) {
	f := ast.File{
		Directives: []ast.Directive{
			ast.DefineDirective{Identifier: "true", Value: ast.IntegerLiteral(1)},
		},
		Declarations: []ast.Declaration{
			ast.AssignmentDeclaration{Identifier: "myVar", Value: ast.IntegerLiteral(420)},
		},
	}

	expected := "#define true 1\nmyVar = 420;\n"

	assert.Equal(t, expected, f.String())
}
