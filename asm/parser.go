// Package asm implements a parser for LLVM IR assembly files.
package asm

import (
	"io/ioutil"

	"github.com/inspirer/textmapper/tm-go/status"
	"github.com/mewmew/l-tm/asm/ll"
	"github.com/mewmew/l-tm/asm/ll/ast"
	"github.com/pkg/errors"
)

// ParseFile parses the given LLVM IR assembly file into an LLVM IR module.
func ParseFile(path string) (*ast.Module, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	content := string(buf)
	return Parse(path, content)
}

// Parse parses the given LLVM IR assembly file into an LLVM IR module, reading
// from content.
func Parse(filename, content string) (*ast.Module, error) {
	var l ll.Lexer
	l.Init(content)
	var p ll.Parser
	b := newBuilder(filename, content)
	//p.Init(b.addError, b.addNode)
	p.Init(b.addNode)
	err := p.Parse(&l)
	if err != nil {
		return nil, err
	}
	if err := b.status.Err(); err != nil {
		return nil, err
	}

	b.file.parsed = b.chunks
	x := ast.ToLlvmNode(b.file.root())
	return x.(*ast.Module), nil
}

type builder struct {
	file   *file
	chunks []chunk
	stack  []int
	status status.Status
}

func newBuilder(filename, content string) *builder {
	return &builder{
		file:   newFile(filename, content),
		chunks: []chunk{{offset: -1}},
		stack:  make([]int, 1, 512),
	}
}

func (b *builder) addError(se ll.SyntaxError) bool {
	r := b.file.sourceRange(se.Offset, se.Endoffset)
	b.status.Add(r, "syntax error")
	return true
}

func (b *builder) addNode(t ll.NodeType, offset, endoffset int) {
	if t == ll.Module {
		offset, endoffset = 0, len(b.file.content)
	}

	index := len(b.chunks)
	start := len(b.stack)
	end := start
	for o := b.chunks[b.stack[start-1]].offset; o >= offset; o = b.chunks[b.stack[start-1]].offset {
		start--
		if o >= endoffset {
			end--
		}
	}
	firstChild := 0
	if start < end {
		firstChild = b.stack[start]
		for _, i := range b.stack[start:end] {
			b.chunks[i].parent = index
		}
	}
	b.chunks[b.stack[start-1]].next = index
	if end == len(b.stack) {
		b.stack = append(b.stack[:start], index)
	} else if start < end {
		b.stack[start] = index
		l := copy(b.stack[start+1:], b.stack[end:])
		b.stack = b.stack[:start+1+l]
	} else {
		b.stack = append(b.stack, 0)
		copy(b.stack[start+1:], b.stack[start:])
		b.stack[start] = index
	}
	b.chunks = append(b.chunks, chunk{
		t:          t,
		offset:     offset,
		endoffset:  endoffset,
		firstChild: firstChild,
	})
}
