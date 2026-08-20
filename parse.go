// Package bicpp provides an entry point for parsing Bohemia Interactive .cpp files.
//
// The parser here is based on the docs at https://community.bistudio.com/wiki/CPP_File_Format
// as well as real .cpp files seen "in the wild" of DayZ.
package bicpp

import (
	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/jmhobbs/go-bicpp/parse"
)

func Parse(input []byte) (ast.File, error) {
	return parse.Parse(input, false)
}

func ParseWithDebug(input []byte) (ast.File, error) {
	return parse.Parse(input, false)
}
