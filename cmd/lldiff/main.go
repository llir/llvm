// The lldiff tool displays the difference between LLVM IR input and llir/llvm
// output.
//
// The input of lldiff is LLVM IR assembly and the output is the difference
// between the input and the llir/llvm string representation of the same LLVM IR
// module.
//
// Usage:
//
//    lldiff FILE.ll...
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/llir/llvm/asm"
	"github.com/mewkiz/pkg/term"
	"github.com/sergi/go-diff/diffmatchpatch"
)

func usage() {
	const use = `
Display the difference between LLVM IR input and llir/llvm output.

Usage:

	lldiff [OPTION]... FILE.ll...

Flags:
`
	fmt.Fprintln(os.Stderr, use[1:])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	dmp := diffmatchpatch.New()
	for _, path := range flag.Args() {
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("%q: unable to read file; %v", path, err)
		}
		want := string(buf)
		module, err := asm.ParseString(path, want)
		if err != nil {
			log.Fatalf("%q: unable to parse module; %v", path, err)
		}
		got := module.String()
		if got != want {
			fmt.Printf(term.Red("--- input: %q\n"), path)
			fmt.Println(term.Green("+++ output: *llir/llvm/ir.Module"))
			fmt.Println()
			diffs := dmp.DiffMain(want, got, false)
			fmt.Println(dmp.DiffPrettyText(diffs))
		}
	}
}
