package parse

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Lexer_Context(t *testing.T) {
	thirteenLines := strings.Join([]string{
		"line 1", "line 2", "line 3", "line 4", "line 5",
		"line 6", "line 7", "line 8", "line 9", "line 10",
		"line 11", "line 12", "line 13",
	}, "\n")

	tests := []struct {
		Name           string
		Data           string
		Target         int
		ExpectedBefore []string
		ExpectedLine   string
		ExpectedAfter  []string
	}{
		{
			Name:           "middle of file caps context at 5 lines each side",
			Data:           thirteenLines,
			Target:         7,
			ExpectedBefore: []string{"line 2", "line 3", "line 4", "line 5", "line 6"},
			ExpectedLine:   "line 7",
			ExpectedAfter:  []string{"line 8", "line 9", "line 10", "line 11", "line 12"},
		},
		{
			Name:           "near start of file truncates before",
			Data:           thirteenLines,
			Target:         2,
			ExpectedBefore: []string{"line 1"},
			ExpectedLine:   "line 2",
			ExpectedAfter:  []string{"line 3", "line 4", "line 5", "line 6", "line 7"},
		},
		{
			Name:           "near end of file truncates after",
			Data:           thirteenLines,
			Target:         12,
			ExpectedBefore: []string{"line 7", "line 8", "line 9", "line 10", "line 11"},
			ExpectedLine:   "line 12",
			ExpectedAfter:  []string{"line 13"},
		},
		{
			Name:           "single line file has no context",
			Data:           "only line",
			Target:         1,
			ExpectedBefore: []string{},
			ExpectedLine:   "only line",
			ExpectedAfter:  []string{},
		},
		{
			Name:           "target beyond the end of the file returns no context",
			Data:           thirteenLines,
			Target:         14,
			ExpectedBefore: nil,
			ExpectedLine:   "",
			ExpectedAfter:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			lex := &lexer{data: []byte(test.Data)}
			before, line, after := lex.context(test.Target, 5)
			assert.Equal(t, test.ExpectedBefore, before)
			assert.Equal(t, test.ExpectedLine, line)
			assert.Equal(t, test.ExpectedAfter, after)
		})
	}
}
