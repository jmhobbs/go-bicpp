package ast_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/stretchr/testify/assert"
)

func Test_Stringers(t *testing.T) {
	tests := []struct {
		Name     string
		Value    ast.Literal
		Expected string
	}{
		{
			Name:     "Identifier",
			Value:    ast.IdentifierLiteral("myVar"),
			Expected: "myVar",
		},
		{
			Name:     "String",
			Value:    ast.StringLiteral("myString"),
			Expected: `"myString"`,
		},
		{
			Name:     "Integer",
			Value:    ast.IntegerLiteral(42),
			Expected: "42",
		},
		{
			Name:     "Float",
			Value:    ast.FloatLiteral(3.9),
			Expected: "3.9",
		},
		{
			Name:     "Float must have one zero",
			Value:    ast.FloatLiteral(3.0),
			Expected: "3.0",
		},
		{
			Name: "Array",
			Value: ast.ArrayLiteral{
				Body: ast.ArrayExpression{
					ast.IdentifierLiteral("myVar"),
					ast.IntegerLiteral(1),
					ast.FloatLiteral(2.5),
					ast.StringLiteral("string"),
				},
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
