%{
package parse

import "fmt"
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
}

%token <identifier> CLASS
%token <identifier> IDENTIFIER

%token <stringValue> STRING
%token <integerValue> INTEGER
%token <floatValue> FLOAT

%type <value> literal
%type <value> value
%type <values> array_values

%%

program
  : statements
  ;

statements
  : statements statement
  | statement
  ;

statement
  : define_macro
  | class_declaration
  | variable_declaration
  | array_declaration
  ;

define_macro
  : TOK_DEFINE IDENTIFIER value {
    program.Define($2, $3)
  }
  ;

class_declaration
  : CLASS TOK_SEMICOLON {
    fmt.Printf("!!! Forward class declaration: %s\n", $1)
  }
  | CLASS TOK_COLON IDENTIFIER TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON {
    fmt.Printf("!!! Derived class declaration: %s from %s\n", $1, $3)
  }
  | CLASS TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON {
    fmt.Printf("!!! Class declaration: %s\n", $1)
  }
  ;

variable_declaration
  : IDENTIFIER TOK_ASSIGN value TOK_SEMICOLON {
    program.Declare($1, $3)
  }
  ;

array_declaration:
  IDENTIFIER TOK_ARRAY TOK_ASSIGN TOK_BLOCK_OPEN array_values TOK_BLOCK_CLOSE TOK_SEMICOLON {
    fmt.Printf("!!! Array declaration: %v[] = %v\n", $1, $5)
    program.Declare($1, $5)
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
  | literal {
    $$ = $1
  }

array_values
  : value {
    $$ = ast.ArrayValue{$1}
  }
  | array_values TOK_COMMA value {
    $$ = append($$, $3)
  }
  ;

%%
