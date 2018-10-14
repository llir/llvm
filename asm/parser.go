// Package parser implements a parser for Textmapper source files.
package parser

import (
	"github.com/inspirer/textmapper/tm-go/status"
	"github.com/inspirer/textmapper/tm-parsers/tm"
	"github.com/inspirer/textmapper/tm-parsers/tm/ast"
)

func Parse(filename, content string) (*ast.File, error) {
	var l tm.Lexer
	l.Init(content)
	var p tm.Parser
	b := newBuilder(filename, content)
	p.Init(b.addError, b.addNode)
	err := p.Parse(&l)
	if err != nil {
		return nil, err
	}
	if err := b.status.Err(); err != nil {
		return nil, err
	}

	b.file.parsed = b.chunks
	x := ast.ToTmNode(b.file.root())
	return x.(*ast.File), nil
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

func (b *builder) addError(se tm.SyntaxError) bool {
	r := b.file.sourceRange(se.Offset, se.Endoffset)
	b.status.Add(r, "syntax error")
	return true
}

func (b *builder) addNode(t tm.NodeType, offset, endoffset int) {
	if t == tm.File {
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
