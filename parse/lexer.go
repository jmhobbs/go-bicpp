package parse

import "strings"

//go:generate ragel -Z -G2 -o lex.go lex.rl
type lexer struct {
	data        []byte
	p, pe, cs   int
	ts, te, act int
	err         error
}

func (lex *lexer) Error(e string) {
	line, col := lex.position(lex.ts)
	before, current, after := lex.context(line, 5)
	lex.err = ParseError{
		Message: e,
		Line:    line,
		Column:  col,
		Current: current,
		Before:  before,
		After:   after,
	}
}

func (lex *lexer) position(offset int) (line, col int) {
	line, col = 1, 1
	for _, b := range lex.data[:offset] {
		if b == '\n' {
			line++
			col = 1
		} else {
			col++
		}
	}
	return line, col
}

func (lex *lexer) context(target, contextLines int) (before []string, line string, after []string) {
	lines := strings.Split(string(lex.data), "\n")

	idx := target - 1
	if idx < 0 || idx >= len(lines) {
		return nil, "", nil
	}

	start := idx - contextLines
	if start < 0 {
		start = 0
	}
	before = lines[start:idx]

	line = lines[idx]

	end := idx + 1 + contextLines
	if end > len(lines) {
		end = len(lines)
	}
	after = lines[idx+1 : end]

	return before, line, after
}
