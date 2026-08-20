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
	assert.Equal(t, p.Directives[0], ast.DefineDirective{Identifier: "true", Value: ast.IntegerLiteral(1)})
}

func Test_Parse_ForwardClass(t *testing.T) {
	p, err := parse.Parse([]byte(`class CfgModule;`), true)
	require.NoError(t, err)
	assert.Equal(t, ast.ClassDeclaration{Identifier: "CfgModule"}, p.Declarations[0])
}

func Test_Parse_Class(t *testing.T) {
	p, err := parse.Parse([]byte(`class CfgModule {};`), true)
	require.NoError(t, err)
	assert.Equal(
		t,
		ast.ClassDeclaration{
			Identifier: "CfgModule",
			Body:       ast.BlockExpression{},
		},
		p.Declarations[0],
	)
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
			Body: ast.BlockExpression{
				ast.AssignmentDeclaration{Identifier: "myVar", Value: ast.IntegerLiteral(420)},
				ast.ClassDeclaration{
					Identifier: "CfgAnother",
					Body: ast.BlockExpression{
						ast.AssignmentDeclaration{Identifier: "scope", Value: ast.IntegerLiteral(2)},
						ast.AssignmentDeclaration{Identifier: "model", Value: ast.StringLiteral(`\dz\path\to\thing.p3d`)},
					},
				},
			},
		}, p.Declarations[0])

}

func Test_Parse_InheritedClass(t *testing.T) {
	p, err := parse.Parse([]byte(`class CfgModule: CfgBase {};`), true)
	require.NoError(t, err)
	assert.Equal(
		t,
		ast.ClassDeclaration{
			Identifier: "CfgModule",
			Parent:     "CfgBase",
			Body:       ast.BlockExpression{},
		},
		p.Declarations[0],
	)
}

func Test_Parse_ErrorLocation(t *testing.T) {
	_, err := parse.Parse([]byte(`
class CfgOuter {
	myVar = 420;
	class CfgInner {
		scope = 2;
	}
};
`), false)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "line 7, column 1")
}

func Test_Parse_IntValue(t *testing.T) {
	p, err := parse.Parse([]byte(`intValue = 42;`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Declarations[0], ast.AssignmentDeclaration{Identifier: "intValue", Value: ast.IntegerLiteral(42)})
}

func Test_Parse_FloatValue(t *testing.T) {
	p, err := parse.Parse([]byte(`floatValue = 42.0;`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Declarations[0], ast.AssignmentDeclaration{Identifier: "floatValue", Value: ast.FloatLiteral(42.0)})
}

func Test_Parse_StringValue(t *testing.T) {
	p, err := parse.Parse([]byte(`stringValue = "text with spaces";`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Declarations[0], ast.AssignmentDeclaration{Identifier: "stringValue", Value: ast.StringLiteral("text with spaces")})
}

func Test_Parse_ArrayValue(t *testing.T) {
	p, err := parse.Parse([]byte(`arrValue[] = {1, 2.5, "string"};`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Declarations[0], ast.AssignmentDeclaration{
		Identifier: "arrValue",
		Value: ast.ArrayLiteral{
			Body: ast.ArrayExpression{
				ast.IntegerLiteral(1),
				ast.FloatLiteral(2.5),
				ast.StringLiteral("string"),
			},
		},
	})
}
