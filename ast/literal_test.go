package ast_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/stretchr/testify/assert"
)

func Test_Stringers(t *testing.T) {
	tests := []struct {
		Name     string
		Value    ast.Node
		Expected string
	}{
		{
			Name:     "Identifier",
			Value:    ast.Identifier("myVar"),
			Expected: "myVar",
		},
		{
			Name:     "String",
			Value:    ast.String("myString"),
			Expected: `"myString"`,
		},
		{
			Name:     "Integer",
			Value:    ast.Integer(42),
			Expected: "42",
		},
		{
			Name:     "Float",
			Value:    ast.Float(3.9),
			Expected: "3.9",
		},
		{
			Name:     "Float must have one zero",
			Value:    ast.Float(3.0),
			Expected: "3.0",
		},
		{
			Name: "Array",
			Value: ast.Array{
				ast.Identifier("myVar"),
				ast.Integer(1),
				ast.Float(2.5),
				ast.String("string"),
			},
			Expected: `{myVar, 1, 2.5, "string"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			assert.Equal(t, tt.Expected, tt.Value.String())
		})
	}
}
