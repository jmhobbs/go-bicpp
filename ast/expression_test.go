package ast_test

import (
	"testing"

	"github.com/jmhobbs/go-bicpp/ast"
	"github.com/stretchr/testify/assert"
)

func Test_Block_String(t *testing.T) {
	b := ast.Block{
		ast.Assignment{Identifier: "myVar", Value: ast.Integer(42)},
	}

	expected := "{\nmyVar = 42;\n}"
	assert.Equal(t, expected, b.String())
}

func Test_Block_Empty_String(t *testing.T) {
	b := ast.Block{}

	expected := "{\n}"
	assert.Equal(t, expected, b.String())
}

func Test_CommentedNode_String(t *testing.T) {
	c := ast.CommentedNode{
		Node:    ast.Assignment{Identifier: "myVar", Value: ast.Integer(42)},
		Comment: ast.Comment("this is a comment"),
	}
	expected := "myVar = 42; // this is a comment"
	assert.Equal(t, expected, c.String())
}
