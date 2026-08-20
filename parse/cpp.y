%{
package parse

import "github.com/jmhobbs/go-bicpp/ast"
%}

%token CLASS IDENTIFIER INTEGER FLOAT STRING
%token TOK_ARRAY TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON TOK_ASSIGN TOK_QUOTE TOK_COMMA TOK_COLON TOK_DEFINE

%union{
  identifier string

  integerValue int
  floatValue float64
  stringValue string

  literal ast.Literal
  literals []ast.Literal

  directive ast.Directive
  directives []ast.Directive

  decl ast.Declaration
  decls []ast.Declaration
}

%token <identifier> CLASS
%token <identifier> IDENTIFIER

%token <stringValue> STRING
%token <integerValue> INTEGER
%token <floatValue> FLOAT

%type <literal> literal
%type <literal> value
%type <literals> array_values

%type <directive> define_macro
%type <directives> directives

%type <decl> declaration
%type <decls> declarations

%type <decl> class_declaration, variable_declaration, array_declaration

%%

file
  : directives {
    file.Directives = $1
  }
  | directives declarations {
    file.Directives = $1
    file.Declarations = $2
  }
  | declarations {
    file.Declarations = $1
  }
  ;

directives
  : directives define_macro {
    $$ = append($1, $2)
  }
  | define_macro {
    $$ = []ast.Directive{$1}
  }
  ;

declarations
  : declarations declaration {
    $$ = append($1, $2)
  }
  | declaration {
    $$ = []ast.Declaration{$1}
  }
  ;

declaration
  : class_declaration
  | variable_declaration
  | array_declaration
  ;

define_macro
  : TOK_DEFINE IDENTIFIER value {
    $$ = ast.DefineDirective{Identifier: $2, Value: $3}
  }
  ;

class_declaration
  : CLASS TOK_SEMICOLON {
    $$ = ast.ClassDeclaration{Identifier: $1}
  }
  | CLASS TOK_COLON IDENTIFIER TOK_BLOCK_OPEN declarations TOK_BLOCK_CLOSE TOK_SEMICOLON {
    $$ = ast.ClassDeclaration{
      Identifier: $1,
      Parent: $3,
      Body: ast.BlockExpression($5),
    }
  }
  | CLASS TOK_COLON IDENTIFIER TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON {
    $$ = ast.ClassDeclaration{
      Identifier: $1,
      Parent: $3,
      Body: ast.BlockExpression{},
    }
  }
  | CLASS TOK_BLOCK_OPEN declarations TOK_BLOCK_CLOSE TOK_SEMICOLON {
    $$ = ast.ClassDeclaration{
      Identifier: $1,
      Body: ast.BlockExpression($3),
    }
  }
  | CLASS TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON {
    $$ = ast.ClassDeclaration{
      Identifier: $1,
      Body: ast.BlockExpression{},
    }
  }
  ;

variable_declaration
  : IDENTIFIER TOK_ASSIGN value TOK_SEMICOLON  {
    $$ =  ast.AssignmentDeclaration{
      Identifier: $1,
      Value: $3,
    }
  }
  ;

array_declaration:
  IDENTIFIER TOK_ARRAY TOK_ASSIGN TOK_BLOCK_OPEN array_values TOK_BLOCK_CLOSE TOK_SEMICOLON {
    $$ =  ast.AssignmentDeclaration{
      Identifier: $1,
      Value: ast.ArrayLiteral{
        Body: ast.ArrayExpression($5),
      },
    }
  }
  ;

literal
  : INTEGER {
    $$ = ast.IntegerLiteral($1)
  }
  | FLOAT {
    $$ = ast.FloatLiteral($1)
  }
  | STRING {
    $$ = ast.StringLiteral($1)
  }
  ;

value
  : IDENTIFIER {
    $$ = ast.IdentifierLiteral($1)
  }
  | literal
  ;

array_values
  : value {
    $$ = []ast.Literal{$1}
  }
  | array_values TOK_COMMA value {
    $$ = append($$, $3)
  }
  ;

%%
