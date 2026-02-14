package parse

//go:generate ragel -Z -G2 -o lex.go lex.rl
//go:generate go tool goyacc -o cpp.go cpp.y

func Parse(input []byte, debug bool) int {
	if debug {
		yyDebug = 5
		yyErrorVerbose = true
	}
	lex := newLexer(input)
	return yyParse(lex)
}
