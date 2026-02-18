package ast

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IDK(t *testing.T) {
	var literal Value = StringValue("some string")

	assert.Equal(t, StringValueKind, literal.Kind())
	sl, ok := literal.(StringValue)
	assert.True(t, ok)
	assert.Equal(t, "some string", sl.Value())
}
