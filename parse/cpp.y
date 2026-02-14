%{
package parse

import "fmt"
%}

%token IDENTIFIER INTEGER FLOAT STRING
%token TOK_CLASS TOK_ARRAY TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON TOK_ASSIGN TOK_QUOTE TOK_COMMA

%union{
  identifier string

  integerValue int
  floatValue float64
  stringValue string

  anything any
  anythings []any
}

%token <identifier> IDENTIFIER

%token <stringValue> STRING
%token <integerValue> INTEGER
%token <floatValue> FLOAT

%type <anything>  value
%type <anythings> array_values

%%

program
  : statements
  ;

statements
  : statements statement
  | statement
  ;

statement:
  class_declaration
  | integer_declaration
  | float_declaration
  | string_declaration
  | array_declaration
  ;

class_declaration
  : TOK_CLASS IDENTIFIER TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON {
    fmt.Printf("!!! Class declaration: %s\n", $2)
  }
  ;

integer_declaration
  : IDENTIFIER TOK_ASSIGN INTEGER TOK_SEMICOLON {
    fmt.Printf("!!! Variable declaration: %s = %d\n", $1, $3)
  }
  ;

float_declaration
  : IDENTIFIER TOK_ASSIGN FLOAT TOK_SEMICOLON {
    fmt.Printf("!!! Variable declaration: %s = %f\n", $1, $3)
  }
  ;

string_declaration
  : IDENTIFIER TOK_ASSIGN STRING TOK_SEMICOLON {
    fmt.Printf("!!! Variable declaration: %s = %s\n", $1, $3)
  }
  ;

value
  : INTEGER {
    $$ = $1
  }
  | FLOAT {
    $$ = $1
  }
  | STRING {
  fmt.Printf("!!! String value: %s\n", $1)
    $$ = $1
  }
  ;

array_values
  : value {
    $$ = []any{$1}
  }
  | array_values TOK_COMMA value {
    $$ = append($$, $3)
  }
  ;

array_declaration:
  IDENTIFIER TOK_ARRAY TOK_ASSIGN TOK_BLOCK_OPEN array_values TOK_BLOCK_CLOSE TOK_SEMICOLON {
    fmt.Printf("!!! Array declaration: %v[] = %v\n", $1, $5)
  }
  ;

%%
