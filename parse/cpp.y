%{
package parse

import "fmt"
%}

%token IDENTIFIER TOK_CLASS TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON

%union{
  identifier string
}

%token <identifier> IDENTIFIER

%%

class_declaration:
  TOK_CLASS IDENTIFIER TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON {
    fmt.Printf("Class declaration: %s\n", $2)
  }
  ;

%%
