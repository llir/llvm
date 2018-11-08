package asm

import (
	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
)

// irMetadataAttachments returns the IR metadata attachments corresponding to
// the given AST metadata attachments.
func (gen *generator) irMetadataAttachments(ns []ast.MetadataAttachment) []*ir.MetadataAttachment {
	var mds []*ir.MetadataAttachment
	for _, n := range ns {
		md := gen.irMetadataAttachment(n)
		mds = append(mds, md)
	}
	return mds
}

// irMetadataAttachment returns the IR metadata attachment corresponding to
// the given AST metadata attachment.
func (gen *generator) irMetadataAttachment(n ast.MetadataAttachment) *ir.MetadataAttachment {
	panic("not yet implemented")
}
