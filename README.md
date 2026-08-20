[![Go Reference](https://pkg.go.dev/badge/github.com/jmhobbs/go-bicpp.svg)](https://pkg.go.dev/github.com/jmhobbs/go-bicpp)
[![Lint & Test](https://github.com/jmhobbs/go-bicpp/actions/workflows/lint-and-test.yml/badge.svg)](https://github.com/jmhobbs/go-bicpp/actions/workflows/lint-and-test.yml)
[![codecov](https://codecov.io/github/jmhobbs/go-bicpp/graph/badge.svg?token=F6SSXORv3k)](https://codecov.io/github/jmhobbs/go-bicpp)

# go-bicpp

A module to parse and generate `.cpp` files used by Bohemia Interactive.

## Usage

```go
package main

import (
  "fmt"

  "github.com/jmhobbs/go-bicpp"
  "github.com/jmhobbs/go-bicpp/printer"
)

func main() {
  file := []byte(`
#define true 0
class CfgVehicles;
class CfgPatches {};
class CfgPatches : CfgVehicles {};
intVal = 42;
strVal = "text with spaces";
arrVal[] = {1,2};
mixedArr[] = {1, 2.0, "text"};
class CfgWhatever {
        myVar = 420;
        class CfgAnother {
                scope = 2;
                model = "\dz\path\to\thing.p3d"
        }
};
`)

  // cpp now has the parsed representation of the .cpp file
  cpp, err := bicpp.Parse(file)
  if err != nil {
    // syntax errors will be of type parse.ParseError, which has a
    // added context. The Error() function prints a "pretty" version
    // of the error
      panic(err)
    }


    // String() will do simple output
    fmt.Println(pgm.String())

    // the printer package does better, and has some options
    p := printer.New(
      printer.WithCondenseEmptyClassBodies(false),
    )
    if err := p.Write(os.Stdout, cpp); err != nil {
      panic(err)
    }
}
```


### Pretty Errors

The parser is strict. When things go awry, it will return an error, which has can print a "pretty" version;

```
panic: parsing error at line 15, column 2

        10 | class CfgWhatever {
        11 |    myVar = 420;
        12 |    class CfgAnother {
        13 |            scope = 2;
        14 |            model = "\dz\path\to\thing.p3d"
        15 |    }
        !! |    `-- syntax error: unexpected "}", expecting ";"
        16 | };
        17 |
```

## Links

 - https://community.bistudio.com/wiki/CPP_File_Format
 - https://www.colm.net/files/ragel/ragel-guide-6.9.pdf
 - https://man7.org/linux/man-pages/man1/yacc.1p.html
 - https://arcb.csc.ncsu.edu/~mueller/codeopt/codeopt00/y_man.pdf
