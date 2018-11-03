package asm

import (
	"github.com/llir/ll/ast"
	"github.com/llir/llvm/ir"
)

// irMetadataAttachments returns the IR metadata attachments corresponding to
// the given AST metadata attachments.
func irMetadataAttachments(ns []ast.MetadataAttachment) []ir.MetadataAttachment {
	// TODO: implement.
	return nil
	var mds []ir.MetadataAttachment
	for _, n := range ns {
		md := irMetadataAttachment(n)
		mds = append(mds, md)
	}
	return mds
}

// irMetadataAttachment returns the IR metadata attachment corresponding to
// the given AST metadata attachment.
func irMetadataAttachment(n ast.MetadataAttachment) ir.MetadataAttachment {
	panic("not yet implemented")
}
