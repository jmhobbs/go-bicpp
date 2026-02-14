package parse_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/parse"
	"github.com/stretchr/testify/require"
)

func Test_Parse_IntValue(t *testing.T) {
	e := parse.Parse([]byte(`intValue = 42;`), true)
	require.Equal(t, 0, e)
}

func Test_Parse_FloatValue(t *testing.T) {
	e := parse.Parse([]byte(`floatValue = 42.0;`), true)
	require.Equal(t, 0, e)
}

func Test_Parse_StringValue(t *testing.T) {
	e := parse.Parse([]byte(`stringValue = "text";`), true)
	require.Equal(t, 0, e)
}

func Test_Parse_ArrayValue(t *testing.T) {
	e := parse.Parse([]byte(`arrValue[] = {1, 2.5, "string"};`), true)
	require.Equal(t, 0, e)
}
