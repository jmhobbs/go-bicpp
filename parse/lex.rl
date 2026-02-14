package parse

import (
    "fmt"
)

%%{ 
    machine biscript;
    write data;
    access lex.;
    variable p lex.p;
    variable pe lex.pe;
}%%

type lexer struct {
    data []byte
    p, pe, cs int
    ts, te, act int
}

func newLexer(data []byte) *lexer {
    lex := &lexer{ 
        data: data,
        pe: len(data),
    }
    %% write init;
    return lex
}

/*
class           return TOK_CLASS;
[a-zA-Z]+       return IDENTIFIER;
{               return TOK_BLOCK_OPEN;
}               return TOK_BLOCK_CLOSE;
;               return TOK_SEMICOLON;
\n
[ \t]+
*/

func (lex *lexer) Lex(out *yySymType) int {
    eof := lex.pe
    tok := 0
    %%{ 
        main := |*
            'class' => { tok = TOK_CLASS; fbreak;};
            [a-zA-Z]+ => { out.identifier = string(lex.data[lex.ts:lex.te]); tok = IDENTIFIER; fbreak;};
            '{' => { tok = TOK_BLOCK_OPEN; fbreak;};
            '}' => { tok = TOK_BLOCK_CLOSE; fbreak; };
            ';' => { tok = TOK_SEMICOLON; fbreak; };
            space;
        *|;

         write exec;
    }%%

    return tok;
}

func (lex *lexer) Error(e string) {
    fmt.Println("error:", e)
}
