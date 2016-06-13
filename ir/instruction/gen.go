//+build ignore

// gen.go generates the data representations of binary instructions and binary
// bitwise instructions.
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"strings"
	"text/template"

	"github.com/mewkiz/pkg/errutil"
)

func main() {
	files := []File{
		File{
			Template: "binary.tmpl",
			Path:     "binary.go",
			Desc:     "Binary instructions",
			URL:      "http://llvm.org/docs/LangRef.html#binary-operations",
			Insts: []Instruction{
				{
					Name: "Add",
					Desc: "an addition",
				},
				{
					Name: "FAdd",
					Desc: "a floating-point addition",
				},
				{
					Name: "Sub",
					Desc: "a subtraction",
				},
				{
					Name: "FSub",
					Desc: "a floating-point subtraction",
				},
				{
					Name: "Mul",
					Desc: "a multiplication",
				},
				{
					Name: "FMul",
					Desc: "a floating-point multiplication",
				},
				{
					Name: "UDiv",
					Desc: "an unsigned division",
				},
				{
					Name: "SDiv",
					Desc: "a signed division",
				},
				{
					Name: "FDiv",
					Desc: "a floating-point division",
				},
				{
					Name: "URem",
					Desc: "an unsigned remainder",
				},
				{
					Name: "SRem",
					Desc: "a signed remainder",
				},
				{
					Name: "FRem",
					Desc: "a floating-point remainder",
				},
			},
		},
		File{
			Template: "binary.tmpl",
			Path:     "bitwise.go",
			Desc:     "Bitwise binary instructions",
			URL:      "http://llvm.org/docs/LangRef.html#bitwise-binary-operations",
			Insts: []Instruction{
				{
					Name: "ShL",
					Desc: "a shift left",
				},
				{
					Name: "LShR",
					Desc: "a logical shift right",
				},
				{
					Name: "AShR",
					Desc: "an arithmetic shift right",
				},
				{
					Name: "And",
					Desc: "an AND",
				},
				{
					Name: "Or",
					Desc: "an OR",
				},
				{
					Name: "Xor",
					Desc: "an exclusive-OR",
				},
			},
		},
		File{
			Template: "conversion.tmpl",
			Path:     "conversion.go",
			Desc:     "Conversion instructions",
			URL:      "http://llvm.org/docs/LangRef.html#conversion-operations",
			Insts: []Instruction{
				{
					Name: "Trunc",
					Desc: "a truncation",
				},
				{
					Name: "ZExt",
					Desc: "a zero extension",
				},
				{
					Name: "SExt",
					Desc: "a sign extension",
				},
				{
					Name: "FPTrunc",
					Desc: "a floating-point truncation",
				},
				{
					Name: "FPExt",
					Desc: "a floating-point extension",
				},
				{
					Name: "FPToUI",
					Desc: "a floating-point to unsigned integer conversion",
				},
				{
					Name: "FPToSI",
					Desc: "a floating-point to signed integer conversion",
				},
				{
					Name: "UIToFP",
					Desc: "an unsigned integer to floating-point conversion",
				},
				{
					Name: "SIToFP",
					Desc: "a signed integer to floating-point conversion",
				},
				{
					Name: "PtrToInt",
					Desc: "a pointer to integer conversion",
				},
				{
					Name: "IntToPtr",
					Desc: "an integer to pointer conversion",
				},
				{
					Name: "BitCast",
					Desc: "a bitcast",
				},
				{
					Name: "AddrSpaceCast",
					Desc: "an address space cast",
				},
			},
		},
	}
	for _, file := range files {
		if err := file.Gen(); err != nil {
			log.Fatal(err)
		}
	}
}

// A File represents a source file containing the instructions of a given
// category.
type File struct {
	// Template path.
	Template string
	// File path.
	Path string
	// Instruction category description; e.g. `Binary instructions`
	Desc string
	// URL to the corresponding section of the LLVM IR Language Reference Manual.
	URL string
	// Instructions.
	Insts []Instruction
}

// An Instruction represents an LLVM IR instruction.
type Instruction struct {
	// CamelCased instruction name; e.g. `ShL`
	Name string
	// Instructio	n description; e.g. `a shift left`
	Desc string
}

// gen generates a source file containing the instructions of the given
// category.
func (f File) Gen() error {
	t := template.New(f.Template)
	funcs := map[string]interface{}{
		"lower": strings.ToLower,
		"h1":    h1,
		"h2":    h2,
	}
	t.Funcs(funcs)
	if _, err := t.ParseFiles(f.Template); err != nil {
		return errutil.Err(err)
	}
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, f); err != nil {
		return errutil.Err(err)
	}
	src, err := format.Source(buf.Bytes())
	if err != nil {
		return errutil.Err(err)
	}
	if err := ioutil.WriteFile(f.Path, src, 0644); err != nil {
		return errutil.Err(err)
	}
	return nil
}

// header returns a header based on the given title and underline character.
func header(title string, c rune) string {
	pad := 80 - len("// ") - len("=== [  ] ") - len(title)
	return fmt.Sprintf("%s [ %s ] %s", strings.Repeat(string(c), 3), title, strings.Repeat(string(c), pad))
}

// h1 returns a level 1 header based on the given title.
func h1(title string) string {
	return header(title, '=')
}

// h2 returns a level 2 header based on the given title.
func h2(title string) string {
	return header(title, '-')
}
