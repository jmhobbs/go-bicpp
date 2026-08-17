package parse

import "github.com/jmhobbs/go-bicpp/ast"

var program ast.Program

//go:generate go tool goyacc -o cpp.go cpp.y
func Parse(input []byte, debug bool) (*ast.Program, error) {
	program = ast.Program{
		Definitions:  make(map[string]ast.Value),
		Declarations: []ast.Declaration{},
	}

	if debug {
		yyDebug = 5
		yyErrorVerbose = true
	}
	lex := newLexer(input)
	if yyParse(lex) != 0 {
		return nil, lex.err
	}
	return &program, nil
}
