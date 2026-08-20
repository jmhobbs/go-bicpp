package printer_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/jmhobbs/go-bicpp/printer"
)

var testFile = ast.File{
	ast.Define{Identifier: "true", Value: ast.Integer(1)},
	ast.Class{
		Identifier: "CfgEmpty",
		Body:       ast.Block{},
	},
	ast.Class{
		Identifier: "CfgTest",
		Body: ast.Block{
			ast.Assignment{Identifier: "myVar", Value: ast.Integer(420)},
		},
	},
}

func Test_Printer_Default(t *testing.T) {
	var buf bytes.Buffer

	p := printer.New()
	err := p.Write(&buf, testFile)
	require.NoError(t, err)
	assert.Equal(t, `#define true 1

class CfgEmpty {};
class CfgTest
{
  myVar = 420;
};
`, buf.String())
}

func Test_Printer_MovesDefinesToTop(t *testing.T) {
	var buf bytes.Buffer

	p := printer.New()
	err := p.Write(&buf, ast.File{
		ast.Class{Identifier: "CfgTest"},
		ast.Define{Identifier: "true", Value: ast.Integer(1)},
	})
	require.NoError(t, err)
	assert.Equal(t, `#define true 1

class CfgTest;
`, buf.String())
}

func Test_Printer_WithIndent(t *testing.T) {
	var buf bytes.Buffer

	p := printer.New(printer.WithIndent("••"))
	err := p.Write(&buf, testFile)
	require.NoError(t, err)
	assert.Equal(t, `#define true 1

class CfgEmpty {};
class CfgTest
{
••myVar = 420;
};
`, buf.String())
}

func Test_Printer_WithoutCollapseEmptyClassBodies(t *testing.T) {
	var buf bytes.Buffer

	p := printer.New(printer.WithCondenseEmptyClassBodies(false))
	err := p.Write(&buf, testFile)
	require.NoError(t, err)
	assert.Equal(t, `#define true 1

class CfgEmpty
{
};
class CfgTest
{
  myVar = 420;
};
`, buf.String())
}
