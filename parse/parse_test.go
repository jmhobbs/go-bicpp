package parse_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/jmhobbs/go-bicpp/parse"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Parse_Exhaustive(t *testing.T) {
	doc := `#define true 1
#define name "CfgModule"
#define value 42.01
#define arr {1, "two", 3.01}

class CfgBase;

class CfgModule : CfgBase {
	scope = 2;
	model = "\dz\path\to\thing.p3d";
	emptyArray[] = {};
	nestedArray[] = {1, {2, 3}, 4};
};

// this is a nested class
class CfgAnother {
	myVar = 420;
	class CfgInner {
		scope = 2;
	};
};
`

	f, err := parse.Parse([]byte(doc), true)
	require.NoError(t, err)

	t.Log(f.String())
	assert.Equal(
		t,
		&ast.File{
			Directives: []ast.Node{
				ast.Define{Identifier: "true", Value: ast.Integer(1)},
				ast.Define{Identifier: "name", Value: ast.String("CfgModule")},
				ast.Define{Identifier: "value", Value: ast.Float(42.01)},
				ast.Define{
					Identifier: "arr",
					Value: ast.Array{
						ast.Integer(1),
						ast.String("two"),
						ast.Float(3.01),
					},
				},
			},
			Declarations: []ast.Node{
				ast.Class{Identifier: "CfgBase"},
				ast.Class{
					Identifier: "CfgModule",
					Parent:     "CfgBase",
					Body: ast.Block{
						ast.Assignment{Identifier: "scope", Value: ast.Integer(2)},
						ast.Assignment{Identifier: "model", Value: ast.String(`\dz\path\to\thing.p3d`)},
						ast.Assignment{Identifier: "emptyArray", Value: ast.Array{}},
						ast.Assignment{
							Identifier: "nestedArray",
							Value: ast.Array{
								ast.Integer(1),
								ast.Array{
									ast.Integer(2),
									ast.Integer(3),
								},
								ast.Integer(4),
							},
						},
					},
				},
				ast.Comment(" this is a nested class"),
				ast.Class{
					Identifier: "CfgAnother",
					Body: ast.Block{
						ast.Assignment{Identifier: "myVar", Value: ast.Integer(420)},
						ast.Class{
							Identifier: "CfgInner",
							Body: ast.Block{
								ast.Assignment{Identifier: "scope", Value: ast.Integer(2)},
							},
						},
					},
				},
			},
		},
		f,
	)
}

func Test_Parse_Define(t *testing.T) {
	p, err := parse.Parse([]byte(`#define true 1`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Directives[0], ast.Define{Identifier: "true", Value: ast.Integer(1)})
}

func Test_Parse_ForwardClass(t *testing.T) {
	p, err := parse.Parse([]byte(`class CfgModule;`), true)
	require.NoError(t, err)
	assert.Equal(t, ast.Class{Identifier: "CfgModule"}, p.Declarations[0])
}

func Test_Parse_Class(t *testing.T) {
	p, err := parse.Parse([]byte(`class CfgModule {};`), true)
	require.NoError(t, err)
	assert.Equal(
		t,
		ast.Class{
			Identifier: "CfgModule",
			Body:       ast.Block{},
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
		ast.Class{
			Identifier: "CfgWhatever",
			Body: ast.Block{
				ast.Assignment{Identifier: "myVar", Value: ast.Integer(420)},
				ast.Class{
					Identifier: "CfgAnother",
					Body: ast.Block{
						ast.Assignment{Identifier: "scope", Value: ast.Integer(2)},
						ast.Assignment{Identifier: "model", Value: ast.String(`\dz\path\to\thing.p3d`)},
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
		ast.Class{
			Identifier: "CfgModule",
			Parent:     "CfgBase",
			Body:       ast.Block{},
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
	assert.Equal(t, p.Declarations[0], ast.Assignment{Identifier: "intValue", Value: ast.Integer(42)})
}

func Test_Parse_FloatValue(t *testing.T) {
	p, err := parse.Parse([]byte(`floatValue = 42.0;`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Declarations[0], ast.Assignment{Identifier: "floatValue", Value: ast.Float(42.0)})
}

func Test_Parse_StringValue(t *testing.T) {
	p, err := parse.Parse([]byte(`stringValue = "text with spaces";`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Declarations[0], ast.Assignment{Identifier: "stringValue", Value: ast.String("text with spaces")})
}

func Test_Parse_ArrayValue(t *testing.T) {
	p, err := parse.Parse([]byte(`arrValue[] = {1, 2.5, "string"};`), true)
	require.NoError(t, err)
	assert.Equal(t, p.Declarations[0], ast.Assignment{
		Identifier: "arrValue",
		Value: ast.Array{
			ast.Integer(1),
			ast.Float(2.5),
			ast.String("string"),
		},
	})
}

func Test_Parse_Comment(t *testing.T) {
	doc := `// this is a comment
notThis = "though";
`

	p, err := parse.Parse([]byte(doc), true)
	require.NoError(t, err)
	assert.Equal(
		t,
		&ast.File{
			Directives: []ast.Node{},
			Declarations: []ast.Node{
				ast.Comment(" this is a comment"),
				ast.Assignment{Identifier: "notThis", Value: ast.String("though")},
			},
		},
		p,
	)
}

func Test_Parse_Comment_Empty(t *testing.T) {
	p, err := parse.Parse([]byte(`//`), true)
	require.NoError(t, err)
	assert.Equal(t, ast.Comment(""), p.Declarations[0])
}
