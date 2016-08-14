// lparse is a parser and pretty-printer for LLVM IR assembler.
//
// Usage: lparse [OPTION]... FILE...
//
// If FILE is -, read standard input.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/llir/spec/gocc/errors"
	"github.com/llir/spec/gocc/lexer"
	"github.com/llir/spec/gocc/parser"
	"github.com/mewkiz/pkg/errutil"
	"github.com/mewkiz/pkg/ioutilx"
)

func usage() {
	const use = `
Usage: lparse [OPTION]... FILE...

If FILE is -, read standard input.
`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Parse input.
	for _, path := range flag.Args() {
		err := parseFile(path)
		if err != nil {
			log.Print(err)
		}
	}
}

// parseFile parses the given file and pretty-prints its LLVM IR assembler to
// standard output.
func parseFile(path string) error {
	// Create lexer for the input.
	buf, err := ioutilx.ReadFile(path)
	if err != nil {
		return errutil.Err(err)
	}
	if path == "-" {
		elog.Print("Parsing from standard input")
	} else {
		elog.Printf("Parsing %q", path)
	}
	s := lexer.NewLexer(buf)

	// Parse input.
	p := parser.NewParser()
	module, err := p.Parse(s)
	if err != nil {
		if err, ok := err.(*errors.Error); ok {
			return parser.NewError(err)
		}
		return errutil.Err(err)
	}
	fmt.Println(module)

	return nil
}

// elog represents a logger with no prefix or flags, which logs to standard
// error.
var elog = log.New(os.Stderr, "", 0)
