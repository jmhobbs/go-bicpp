package parse_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/jmhobbs/go-bicpp/parse"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Parse_Define(t *testing.T) {
	p, err := parse.Parse([]byte(`#define true 1`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Definitions["true"], ast.IntegerValue(1))
}

func Test_Parse_Class(t *testing.T) {
	_, err := parse.Parse([]byte(`class CfgModule;`), true)
	require.NoError(t, err)

	_, err = parse.Parse([]byte(`class CfgModule {};`), true)
	require.NoError(t, err)

	_, err = parse.Parse([]byte(`class CfgModule: CfgBase {};`), true)
	require.NoError(t, err)
}

func Test_Parse_IntValue(t *testing.T) {
	p, err := parse.Parse([]byte(`intValue = 42;`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Declarations["intValue"], ast.VariableDeclaration{Identifier: "intValue", Value: ast.IntegerValue(42)})
}

func Test_Parse_FloatValue(t *testing.T) {
	p, err := parse.Parse([]byte(`floatValue = 42.0;`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Declarations["floatValue"], ast.VariableDeclaration{Identifier: "floatValue", Value: ast.FloatValue(42.0)})
}

func Test_Parse_StringValue(t *testing.T) {
	p, err := parse.Parse([]byte(`stringValue = "text with spaces";`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Declarations["stringValue"], ast.VariableDeclaration{Identifier: "stringValue", Value: ast.StringValue("text with spaces")})
}

func Test_Parse_ArrayValue(t *testing.T) {
	p, err := parse.Parse([]byte(`arrValue[] = {1, 2.5, "string"};`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Declarations["arrValue"], ast.VariableDeclaration{
		Identifier: "arrValue",
		Value: ast.ArrayValue{
			ast.IntegerValue(1),
			ast.FloatValue(2.5),
			ast.StringValue("string"),
		},
	})
}
