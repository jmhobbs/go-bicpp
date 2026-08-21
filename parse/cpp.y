%{
package parse

import "github.com/jmhobbs/go-bicpp/ast"
%}

%token CLASS IDENTIFIER INTEGER FLOAT STRING
%token TOK_ARRAY TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON TOK_ASSIGN TOK_QUOTE TOK_COMMA TOK_COLON TOK_DEFINE COMMENT INLINE_COMMENT

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

%token <stringValue> STRING COMMENT INLINE_COMMENT
%token <integerValue> INTEGER
%token <floatValue> FLOAT

%type <node> literal inline_comment
%type <nodes> array_values array

%type <node> define_macro
%type <nodes> directives

%type <node> declaration
%type <nodes> declarations

%type <node> class_declaration, variable_declaration, array_declaration, comment_declaration

%%

file
  : directives {
    file = $1
  }
  | directives declarations {
    file = append($1, $2...)
  }
  | declarations {
    file = $1
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
  | comment_declaration
  | inline_comment
  ;

define_macro
  : TOK_DEFINE IDENTIFIER literal {
    $$ = ast.Define{Identifier: $2, Value: $3}
  }
  | TOK_DEFINE IDENTIFIER array {
    $$ = ast.Define{Identifier: $2, Value: ast.Array($3)}
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

inline_comment
  : variable_declaration INLINE_COMMENT {
    $$ = ast.CommentedNode{
      Node: $1,
      Comment: ast.Comment($2),
    }
  }
  ;

variable_declaration
  : IDENTIFIER TOK_ASSIGN literal TOK_SEMICOLON  {
    $$ =  ast.Assignment{
      Identifier: $1,
      Value: $3,
    }
  }
  ;

array_declaration:
  IDENTIFIER TOK_ARRAY TOK_ASSIGN array TOK_SEMICOLON {
    $$ =  ast.Assignment{
      Identifier: $1,
      Value: ast.Array($4),
    }
  }
  ;

comment_declaration
  : COMMENT {
    $$ = ast.Comment($1)
  }
  ;

literal
  : IDENTIFIER {
    $$ = ast.Identifier($1)
  }
  | INTEGER {
    $$ = ast.Integer($1)
  }
  | FLOAT {
    $$ = ast.Float($1)
  }
  | STRING {
    $$ = ast.String($1)
  }
  ;

array
  : TOK_BLOCK_OPEN array_values TOK_BLOCK_CLOSE {
    $$ = ast.Array($2)
  }
  | TOK_BLOCK_OPEN TOK_BLOCK_CLOSE {
    $$ = ast.Array{}
  }
  ;

array_values
  : literal {
    $$ = []ast.Node{$1}
  }
  | array {
    $$ = $1
  }
  | array_values TOK_COMMA literal {
    $$ = append($$, $3)
  }
  | array_values TOK_COMMA array {
    $$ = append($$, ast.Array($3))
  }
  ;

%%
