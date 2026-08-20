package ast_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/stretchr/testify/assert"
)

func Test_Assignment_String(t *testing.T) {
	d := ast.Assignment{
		Identifier: "myVar",
		Value:      ast.Integer(42),
	}

	assert.Equal(t, "myVar = 42;", d.String())
}

func Test_Class_String(t *testing.T) {
	tests := []struct {
		Name        string
		Declaration ast.Class
		Expected    string
	}{
		{
			Name: "Forward Declaration",
			Declaration: ast.Class{
				Identifier: "CfgModule",
			},
			Expected: "class CfgModule;",
		},
		{
			Name: "Empty Class",
			Declaration: ast.Class{
				Identifier: "CfgModule",
				Body:       ast.Block{},
			},
			Expected: "class CfgModule\n{\n};",
		},
		{
			Name: "Empty Inherited Class",
			Declaration: ast.Class{
				Identifier: "CfgModule",
				Parent:     "CfgBase",
				Body:       ast.Block{},
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

func Test_Comment_String(t *testing.T) {
	c := ast.Comment("this is a comment")
	assert.Equal(t, "// this is a comment", c.String())
}
