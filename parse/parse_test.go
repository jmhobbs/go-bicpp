package parse_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/parse"
	"github.com/stretchr/testify/require"
)

func Test_Parse_Define(t *testing.T) {
	require.NoError(t, parse.Parse([]byte(`#define true 0`), true))
}

func Test_Parse_Class(t *testing.T) {
	require.NoError(t, parse.Parse([]byte(`class CfgModule;`), true))

	require.NoError(t, parse.Parse([]byte(`class CfgModule {};`), true))

	require.NoError(t, parse.Parse([]byte(`class CfgModule: CfgBase {};`), true))
}

func Test_Parse_IntValue(t *testing.T) {
	require.NoError(t, parse.Parse([]byte(`intValue = 42;`), true))
}

func Test_Parse_FloatValue(t *testing.T) {
	require.NoError(t, parse.Parse([]byte(`floatValue = 42.0;`), true))
}

func Test_Parse_StringValue(t *testing.T) {
	require.NoError(t, parse.Parse([]byte(`stringValue = "text with spaces";`), true))
}

func Test_Parse_ArrayValue(t *testing.T) {
	require.NoError(t, parse.Parse([]byte(`arrValue[] = {1, 2.5, "string"};`), true))
}
