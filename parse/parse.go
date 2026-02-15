package parse

//go:generate go tool goyacc -o cpp.go cpp.y
func Parse(input []byte, debug bool) error {
	if debug {
		yyDebug = 5
		yyErrorVerbose = true
	}
	lex := newLexer(input)
	if yyParse(lex) != 0 {
		return lex.err
	}
	return nil
}
