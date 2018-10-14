package parser

import (
	"sort"
	"strings"

	"github.com/inspirer/textmapper/tm-go/status"
	"github.com/inspirer/textmapper/tm-parsers/tm"
	"github.com/inspirer/textmapper/tm-parsers/tm/ast"
	"github.com/inspirer/textmapper/tm-parsers/tm/selector"
)

// file holds an AST of a single file and can covert offset ranges into status.SourceRange.
type file struct {
	filename string
	content  string
	lines    []int
	parsed   []chunk
}

func (f *file) root() ast.Node {
	return node{f, len(f.parsed) - 1}
}

type chunk struct {
	t          tm.NodeType
	offset     int
	endoffset  int
	next       int
	firstChild int
	parent     int
}

func newFile(filename, content string) *file {
	return &file{
		filename: filename,
		content:  content,
		lines:    lineOffsets(content),
	}
}

func (f *file) sourceRange(offset, endoffset int) status.SourceRange {
	line := sort.Search(len(f.lines), func(i int) bool { return f.lines[i] > offset }) - 1
	return status.SourceRange{
		Filename:  f.filename,
		Offset:    offset,
		EndOffset: endoffset,
		Line:      line + 1,
		Column:    offset - f.lines[line] + 1,
	}
}

func lineOffsets(str string) []int {
	var lines = make([]int, 1, 128)

	var off int
	for {
		i := strings.IndexByte(str[off:], '\n')
		if i == -1 {
			break
		}
		off += i + 1
		lines = append(lines, off)
	}
	return lines
}

type node struct {
	file  *file
	index int
}

// Type implements ast.Node
func (n node) Type() tm.NodeType {
	if n.file == nil {
		// TODO: introduce InvalidType
		return tm.SyntaxProblem
	}
	return n.file.parsed[n.index].t
}

// Offset implements ast.Node
func (n node) Offset() int {
	if n.file == nil {
		return 0
	}
	return n.file.parsed[n.index].offset
}

// Endoffset implements ast.Node
func (n node) Endoffset() int {
	if n.file == nil {
		return 0
	}
	return n.file.parsed[n.index].endoffset
}

// Child implements ast.Node
func (n node) Child(sel selector.Selector) ast.Node {
	if n.file == nil {
		return node{}
	}
	for i := n.file.parsed[n.index].firstChild; i > 0; i = n.file.parsed[i].next {
		if sel(n.file.parsed[i].t) {
			return node{n.file, i}
		}
	}
	return node{}
}

// Children implements ast.Node
func (n node) Children(sel selector.Selector) []ast.Node {
	if n.file == nil {
		return nil
	}
	var ret []ast.Node
	for i := n.file.parsed[n.index].firstChild; i > 0; i = n.file.parsed[i].next {
		if sel(n.file.parsed[i].t) {
			ret = append(ret, node{n.file, i})
		}
	}
	return ret
}

// Next implements ast.Node
func (n node) Next(sel selector.Selector) ast.Node {
	if n.file == nil {
		return node{}
	}
	for i := n.file.parsed[n.index].next; i > 0; i = n.file.parsed[i].next {
		if sel(n.file.parsed[i].t) {
			return node{n.file, i}
		}
	}
	return node{}
}

// NextAll implements ast.Node
func (n node) NextAll(sel selector.Selector) []ast.Node {
	if n.file == nil {
		return nil
	}
	var ret []ast.Node
	for i := n.file.parsed[n.index].next; i > 0; i = n.file.parsed[i].next {
		if sel(n.file.parsed[i].t) {
			ret = append(ret, node{n.file, i})
		}
	}
	return nil
}

// Text implements ast.Node
func (n node) Text() string {
	if n.file == nil {
		return ""
	}
	start, end := n.file.parsed[n.index].offset, n.file.parsed[n.index].endoffset
	return n.file.content[start:end]
}

// SourceRange implements status.SourceNode
func (n node) SourceRange() status.SourceRange {
	return n.file.sourceRange(n.file.parsed[n.index].offset, n.file.parsed[n.index].endoffset)
}
