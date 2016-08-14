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
	// TODO: Consider using a JSON file to define the relevant data of the source
	// files. Currently, the data is duplicated and present in both the irx and
	// the instruction packages.
	files := []File{
		File{
			Path: "binary.go",
			Desc: "Binary instructions",
			URL:  "http://llvm.org/docs/LangRef.html#binary-operations",
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
			Path: "bitwise.go",
			Desc: "Bitwise binary instructions",
			URL:  "http://llvm.org/docs/LangRef.html#bitwise-binary-operations",
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
	// Instruction description; e.g. `a shift left`
	Desc string
}

// gen generates a source file containing the instructions of the given
// category.
func (f File) Gen() error {
	t := template.New("inst")
	funcs := map[string]interface{}{
		"lower": strings.ToLower,
		"h1":    h1,
		"h2":    h2,
	}
	t.Funcs(funcs)
	if _, err := t.Parse(dataContent[1:]); err != nil {
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

const dataContent = `
// generated by gen.go using 'go generate'; DO NOT EDIT.

// {{ h1 .Desc }}
//
// References:
//    {{ .URL }}

package irx

import (
	"github.com/llir/llvm/ir/instruction"
	"github.com/mewkiz/pkg/errutil"
)

{{- range .Insts }}
// {{ lower .Name | h2 }}

// New{{ .Name }} returns a new {{ lower .Name }} instruction based on the given operand type and
// operands.
func New{{ .Name }}Inst(typ, xVal, yVal interface{}) (*instruction.{{ .Name }}, error) {
	x, err := NewValue(typ, xVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	y, err := NewValue(typ, yVal)
	if err != nil {
		return nil, errutil.Err(err)
	}
	return instruction.New{{ .Name }}(x, y)
}

{{- end }}
`
