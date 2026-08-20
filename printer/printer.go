package printer

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/jmhobbs/go-bicpp/ast"
)

type PrinterOption func(*Printer)

type Printer struct {
	indent                   string
	condenseEmptyClassBodies bool
}

func WithIndent(indent string) PrinterOption {
	return func(p *Printer) {
		p.indent = indent
	}
}

func WithCondenseEmptyClassBodies(condense bool) PrinterOption {
	return func(p *Printer) {
		p.condenseEmptyClassBodies = condense
	}
}

func New(opts ...PrinterOption) *Printer {
	p := &Printer{"  ", true}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

func (p *Printer) Write(out io.Writer, file *ast.File) error {
	var err error
	for _, directive := range file.Directives {
		if _, err = out.Write([]byte(directive.String() + "\n")); err != nil {
			return err
		}
	}
	// add a newline between directives and the body of the file
	if len(file.Directives) > 0 {
		if _, err = out.Write([]byte{'\n'}); err != nil {
			return err
		}
	}

	for _, declaration := range file.Declarations {
		if err = p.writeDeclaration(0, out, declaration); err != nil {
			return err
		}
	}
	return nil
}

func (p *Printer) writeDeclaration(depth int, out io.Writer, declaration ast.Declaration) error {
	_, err := out.Write([]byte(strings.Repeat(p.indent, depth)))
	if err != nil {
		return err
	}

	switch declaration.Kind() {
	case ast.Assignment:
		if _, err = out.Write([]byte(declaration.String() + "\n")); err != nil {
			return err
		}
	case ast.Class:
		classDecl := declaration.(ast.ClassDeclaration)
		if _, err = fmt.Fprintf(out, "class %s", classDecl.Identifier); err != nil {
			return err
		}
		if classDecl.Parent != "" {
			if _, err = fmt.Fprintf(out, " : %s", classDecl.Parent); err != nil {
				return err
			}
		}
		if classDecl.Body != nil {
			// special case, keep it on one line if it's empty
			if p.condenseEmptyClassBodies && len(classDecl.Body) == 0 {
				if _, err = out.Write([]byte{' ', '{', '}'}); err != nil {
					return err
				}
			} else {
				if _, err = out.Write([]byte{'\n'}); err != nil {
					return err
				}
				if _, err = out.Write([]byte(strings.Repeat(p.indent, depth))); err != nil {
					return err
				}
				if _, err = out.Write([]byte{'{', '\n'}); err != nil {
					return err
				}
				for _, nested := range classDecl.Body {
					if err = p.writeDeclaration(depth+1, out, nested); err != nil {
						return err
					}
				}
				if _, err = out.Write([]byte(strings.Repeat(p.indent, depth))); err != nil {
					return err
				}
				if _, err = out.Write([]byte{'}'}); err != nil {
					return err
				}
			}
		}
		if _, err = out.Write([]byte{';', '\n'}); err != nil {
			return err
		}
	default:
		return errors.New("unknown declaration kind")
	}

	return nil
}
