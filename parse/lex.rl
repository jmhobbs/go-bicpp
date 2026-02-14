package parse

import (
    "fmt"
    "strconv"
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

func (lex *lexer) Lex(out *yySymType) int {
    eof := lex.pe
    tok := 0
    %%{ 
        main := |*
            '"' => { tok = TOK_QUOTE; fbreak; };
            'class' => { tok = TOK_CLASS; fbreak;};
            '[]' => { tok = TOK_ARRAY; fbreak;};
            [a-zA-Z]+ => { out.identifier = string(lex.data[lex.ts:lex.te]); tok = IDENTIFIER; fbreak;};
            [0-9]+'.'[0-9]+ => {
              n, err := strconv.ParseFloat(string(lex.data[lex.ts:lex.te]), 64);
              if err != nil {
                panic(err)
              }
              out.float = n;
              tok = FLOAT;
              fbreak;
            };
            [0-9]+ => {
              n, err := strconv.Atoi(string(lex.data[lex.ts:lex.te]));
              if err != nil {
                panic(err)
              }
              out.integer = n;
              tok = INTEGER;
              fbreak;
            };
            ',' => { tok = TOK_COMMA; fbreak; };
            '{' => { tok = TOK_BLOCK_OPEN; fbreak;};
            '}' => { tok = TOK_BLOCK_CLOSE; fbreak; };
            ';' => { tok = TOK_SEMICOLON; fbreak; };
            '=' => { tok = TOK_ASSIGN; fbreak; };
            space;
        *|;

         write exec;
    }%%

    return tok;
}

func (lex *lexer) Error(e string) {
    fmt.Println("error:", e)
}
