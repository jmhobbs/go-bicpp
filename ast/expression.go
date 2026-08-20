package ast

import "strings"

type Block []Node

func (b Block) String() string {
	lines := make([]string, len(b)+2)
	lines[0] = "{"
	lines[len(lines)-1] = "}"
	for i, d := range b {
		lines[i+1] = d.String()
	}
	return strings.Join(lines, "\n")
}
