package ast_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/stretchr/testify/assert"
)

func Test_Stringers(t *testing.T) {
	tests := []struct {
		Name     string
		Value    ast.Value
		Expected string
	}{
		{
			Name:     "Identifier",
			Value:    ast.IdentifierValue("myVar"),
			Expected: "myVar",
		},
		{
			Name:     "String",
			Value:    ast.StringValue("myString"),
			Expected: `"myString"`,
		},
		{
			Name:     "Integer",
			Value:    ast.IntegerValue(42),
			Expected: "42",
		},
		{
			Name:     "Float",
			Value:    ast.FloatValue(3.9),
			Expected: "3.9",
		},
		{
			Name:     "Float must have one zero",
			Value:    ast.FloatValue(3.0),
			Expected: "3.0",
		},
		{
			Name: "Array",
			Value: ast.ArrayValue{
				ast.IdentifierValue("myVar"),
				ast.IntegerValue(1),
				ast.FloatValue(2.5),
				ast.StringValue("string"),
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
