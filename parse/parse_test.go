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
	assert.Equal(t, p.Definitions[0], ast.Definition{Identifier: "true", Value: ast.IntegerValue(1)})
}

func Test_Parse_ForwardClass(t *testing.T) {
	p, err := parse.Parse([]byte(`class CfgModule;`), true)
	require.NoError(t, err)
	assert.Equal(t, ast.ForwardClassDeclaration{Identifier: "CfgModule"}, p.Declarations[0])
}

func Test_Parse_Class(t *testing.T) {
	p, err := parse.Parse([]byte(`class CfgModule {};`), true)
	require.NoError(t, err)
	assert.Equal(t, ast.ClassDeclaration{Identifier: "CfgModule", Fields: []ast.Declaration{}}, p.Declarations[0])
}

func Test_Parse_Nested_Class(t *testing.T) {
	p, err := parse.Parse([]byte(`class CfgWhatever {
	myVar = 420;
	class CfgAnother {
		scope = 2;
		model = "\dz\path\to\thing.p3d";
	};
};
`), true)
	require.NoError(t, err)
	assert.Equal(
		t,
		ast.ClassDeclaration{
			Identifier: "CfgWhatever",
			Fields: []ast.Declaration{
				ast.VariableDeclaration{Identifier: "myVar", Value: ast.IntegerValue(420)},
				ast.ClassDeclaration{
					Identifier: "CfgAnother",
					Fields: []ast.Declaration{
						ast.VariableDeclaration{Identifier: "scope", Value: ast.IntegerValue(2)},
						ast.VariableDeclaration{Identifier: "model", Value: ast.StringValue(`\dz\path\to\thing.p3d`)},
					},
				},
			},
		}, p.Declarations[0])

}

func Test_Parse_InheritedClass(t *testing.T) {
	p, err := parse.Parse([]byte(`class CfgModule: CfgBase {};`), true)
	require.NoError(t, err)
	assert.Equal(t, ast.ClassDeclaration{Identifier: "CfgModule", Parent: "CfgBase", Fields: []ast.Declaration{}}, p.Declarations[0])
}

func Test_Parse_IntValue(t *testing.T) {
	p, err := parse.Parse([]byte(`intValue = 42;`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Declarations[0], ast.VariableDeclaration{Identifier: "intValue", Value: ast.IntegerValue(42)})
}

func Test_Parse_FloatValue(t *testing.T) {
	p, err := parse.Parse([]byte(`floatValue = 42.0;`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Declarations[0], ast.VariableDeclaration{Identifier: "floatValue", Value: ast.FloatValue(42.0)})
}

func Test_Parse_StringValue(t *testing.T) {
	p, err := parse.Parse([]byte(`stringValue = "text with spaces";`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Declarations[0], ast.VariableDeclaration{Identifier: "stringValue", Value: ast.StringValue("text with spaces")})
}

func Test_Parse_ArrayValue(t *testing.T) {
	p, err := parse.Parse([]byte(`arrValue[] = {1, 2.5, "string"};`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Declarations[0], ast.VariableDeclaration{
		Identifier: "arrValue",
		Value: ast.ArrayValue{
			ast.IntegerValue(1),
			ast.FloatValue(2.5),
			ast.StringValue("string"),
		},
	})
}
