package bicpp

import (
	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/jmhobbs/go-bicpp/parse"
)

func Parse(input []byte) (*ast.File, error) {
	return parse.Parse(input, false)
}

func ParseWithDebug(input []byte) (*ast.File, error) {
	return parse.Parse(input, false)
}
