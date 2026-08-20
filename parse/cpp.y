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

  node ast.Node
  nodes []ast.Node
}

%token <identifier> CLASS
%token <identifier> IDENTIFIER

%token <stringValue> STRING
%token <integerValue> INTEGER
%token <floatValue> FLOAT

%type <node> literal
%type <node> value
%type <nodes> array_values

%type <node> define_macro
%type <nodes> directives

%type <node> declaration
%type <nodes> declarations

%type <node> class_declaration, variable_declaration, array_declaration

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
    $$ = []ast.Node{$1}
  }
  ;

declarations
  : declarations declaration {
    $$ = append($1, $2)
  }
  | declaration {
    $$ = []ast.Node{$1}
  }
  ;

declaration
  : class_declaration
  | variable_declaration
  | array_declaration
  ;

define_macro
  : TOK_DEFINE IDENTIFIER value {
    $$ = ast.Define{Identifier: $2, Value: $3}
  }
  ;

class_declaration
  : CLASS TOK_SEMICOLON {
    $$ = ast.Class{Identifier: $1}
  }
  | CLASS TOK_COLON IDENTIFIER TOK_BLOCK_OPEN declarations TOK_BLOCK_CLOSE TOK_SEMICOLON {
    $$ = ast.Class{
      Identifier: $1,
      Parent: $3,
      Body: ast.Block($5),
    }
  }
  | CLASS TOK_COLON IDENTIFIER TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON {
    $$ = ast.Class{
      Identifier: $1,
      Parent: $3,
      Body: ast.Block{},
    }
  }
  | CLASS TOK_BLOCK_OPEN declarations TOK_BLOCK_CLOSE TOK_SEMICOLON {
    $$ = ast.Class{
      Identifier: $1,
      Body: ast.Block($3),
    }
  }
  | CLASS TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON {
    $$ = ast.Class{
      Identifier: $1,
      Body: ast.Block{},
    }
  }
  ;

variable_declaration
  : IDENTIFIER TOK_ASSIGN value TOK_SEMICOLON  {
    $$ =  ast.Assignment{
      Identifier: $1,
      Value: $3,
    }
  }
  ;

array_declaration:
  IDENTIFIER TOK_ARRAY TOK_ASSIGN TOK_BLOCK_OPEN array_values TOK_BLOCK_CLOSE TOK_SEMICOLON {
    $$ =  ast.Assignment{
      Identifier: $1,
      Value: ast.Array{
        Body: ast.ArrayBlock($5),
      },
    }
  }
  ;

literal
  : INTEGER {
    $$ = ast.Integer($1)
  }
  | FLOAT {
    $$ = ast.Float($1)
  }
  | STRING {
    $$ = ast.String($1)
  }
  ;

value
  : IDENTIFIER {
    $$ = ast.Identifier($1)
  }
  | literal
  ;

array_values
  : value {
    $$ = []ast.Node{$1}
  }
  | array_values TOK_COMMA value {
    $$ = append($$, $3)
  }
  ;

%%
