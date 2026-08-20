package ast_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/stretchr/testify/assert"
)

func Test_AssignmentDeclaration_String(t *testing.T) {
	d := ast.AssignmentDeclaration{
		Identifier: "myVar",
		Value:      ast.IntegerLiteral(42),
	}

	assert.Equal(t, "myVar = 42;", d.String())
}

func Test_ClassDeclaration_String(t *testing.T) {
	tests := []struct {
		Name        string
		Declaration ast.ClassDeclaration
		Expected    string
	}{
		{
			Name: "Forward Declaration",
			Declaration: ast.ClassDeclaration{
				Identifier: "CfgModule",
			},
			Expected: "class CfgModule;",
		},
		{
			Name: "Empty Class",
			Declaration: ast.ClassDeclaration{
				Identifier: "CfgModule",
				Body:       ast.BlockExpression{},
			},
			Expected: "class CfgModule\n{\n};",
		},
		{
			Name: "Empty Inherited Class",
			Declaration: ast.ClassDeclaration{
				Identifier: "CfgModule",
				Parent:     "CfgBase",
				Body:       ast.BlockExpression{},
			},
			Expected: "class CfgModule : CfgBase\n{\n};",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(t, tt.Expected, tt.Declaration.String())
		})
	}
}
