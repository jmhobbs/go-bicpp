package parse

import "github.com/jmhobbs/go-bicpp/ast"

var file ast.File

//go:generate go tool goyacc -o cpp.go cpp.y
//go:generate gofmt -w cpp.go
func Parse(input []byte, debug bool) (*ast.File, error) {
	file = ast.File{
		Directives:   []ast.Node{},
		Declarations: []ast.Node{},
	}

	yyErrorVerbose = true
	if debug {
		yyDebug = 5
	}
	lex := newLexer(input)
	if yyParse(lex) != 0 {
		return nil, lex.err
	}
	return &file, nil
}
