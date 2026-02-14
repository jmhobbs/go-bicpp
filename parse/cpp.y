%{
package parse

import "fmt"
%}

%token IDENTIFIER TOK_CLASS TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON TOK_ASSIGN INTEGER FLOAT TOK_QUOTE

%union{
  identifier string
  integer int
  float float64
}

%token <identifier> IDENTIFIER
%token <integer> INTEGER
%token <float> FLOAT

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
  ;

class_declaration:
  TOK_CLASS IDENTIFIER TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON {
    fmt.Printf("Class declaration: %s\n", $2)
  }
  ;

integer_declaration:
  IDENTIFIER TOK_ASSIGN INTEGER TOK_SEMICOLON {
    fmt.Printf("Variable declaration: %s = %d\n", $1, $3)
  }
  ;

float_declaration:
  IDENTIFIER TOK_ASSIGN FLOAT TOK_SEMICOLON {
    fmt.Printf("Variable declaration: %s = %f\n", $1, $3)
  }
  ;

string_declaration:
  IDENTIFIER TOK_ASSIGN TOK_QUOTE IDENTIFIER TOK_QUOTE TOK_SEMICOLON {
    fmt.Printf("Variable declaration: %s = %s\n", $1, $4)
  }
  ;

%%
