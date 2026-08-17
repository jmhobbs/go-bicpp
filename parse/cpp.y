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

  value ast.Value
  values ast.ArrayValue

  define ast.Definition
  defines []ast.Definition

  decl ast.Declaration
  decls []ast.Declaration
}

%token <identifier> CLASS
%token <identifier> IDENTIFIER

%token <stringValue> STRING
%token <integerValue> INTEGER
%token <floatValue> FLOAT

%type <value> literal
%type <value> value
%type <values> array_values

%type <define> define_macro
%type <defines> defines

%type <decl> declaration
%type <decls> declarations

%type <decl> class_declaration, variable_declaration, array_declaration

%%

program
  : defines {
    program.Definitions = $1
  }
  | defines declarations {
    program.Definitions = $1
    program.Declarations = $2
  }
  | declarations {
    program.Declarations = $1
  }
  ;

defines
  : defines define_macro {
    $$ = append($1, $2)
  }
  | define_macro {
    $$ = []ast.Definition{$1}
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
    $$ = ast.Definition{Identifier: $2, Value: $3}
  }
  ;

class_declaration
  : CLASS TOK_SEMICOLON {
    $$ = ast.ForwardClassDeclaration{Identifier: $1}
  }
  | CLASS TOK_COLON IDENTIFIER TOK_BLOCK_OPEN declarations TOK_BLOCK_CLOSE TOK_SEMICOLON {
    $$ = ast.ClassDeclaration{Identifier: $1, Parent: $3, Fields: $5}
  }
  | CLASS TOK_COLON IDENTIFIER TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON {
    $$ = ast.ClassDeclaration{Identifier: $1, Parent: $3, Fields: []ast.Declaration{}}
  }
  | CLASS TOK_BLOCK_OPEN declarations TOK_BLOCK_CLOSE TOK_SEMICOLON {
    $$ = ast.ClassDeclaration{Identifier: $1, Fields: $3}
  }
  | CLASS TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON {
    $$ = ast.ClassDeclaration{Identifier: $1, Fields: []ast.Declaration{}}
  }
  ;

variable_declaration
  : IDENTIFIER TOK_ASSIGN value TOK_SEMICOLON  {
    $$ =  ast.VariableDeclaration{Identifier: $1, Value: $3}
  }
  ;

array_declaration:
  IDENTIFIER TOK_ARRAY TOK_ASSIGN TOK_BLOCK_OPEN array_values TOK_BLOCK_CLOSE TOK_SEMICOLON {
    $$ =  ast.VariableDeclaration{Identifier: $1, Value: $5}
  }
  ;

literal
  : INTEGER {
    $$ = ast.IntegerValue($1)
  }
  | FLOAT {
    $$ = ast.FloatValue($1)
  }
  | STRING {
    $$ = ast.StringValue($1)
  }
  ;

value
  : IDENTIFIER {
    $$ = ast.IdentifierValue($1)
  }
  | literal
  ;

array_values
  : value {
    $$ = ast.ArrayValue{$1}
  }
  | array_values TOK_COMMA value {
    $$ = append($$, $3)
  }
  ;

%%
