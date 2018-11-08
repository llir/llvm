package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/metadata"
	"github.com/pkg/errors"
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

func irDIExpression(old *ast.DIExpression) *metadata.DIExpression {
	panic("not yet implemented")
}

func (gen *generator) irMDTuple(old *ast.MDTuple) (*metadata.MDTuple, error) {
	tuple := &metadata.MDTuple{}
	for _, oldField := range old.MDFields().MDFields() {
		field, err := gen.irMDField(oldField)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		tuple.Fields = append(tuple.Fields, field)
	}
	return tuple, nil
}

func (gen *generator) irMDField(old ast.MDField) (metadata.MDField, error) {
	switch old := old.(type) {
	case ast.NullLit:
		// TODO: add support.
		panic("support for metadata field null literal not yet implemented")
	case ast.Metadata:
		return gen.irMetadata(old)
	default:
		panic(fmt.Errorf("support for metadata field %T not yet implemented", old))
	}
}

func (gen *generator) irMetadata(old ast.Metadata) (metadata.Metadata, error) {
	switch old := old.(type) {
	case *ast.TypeValue:
		// TODO: figure out how to handle local values.
		panic("support for metadata type value not yet implemented")
	case *ast.MDString:
		return &metadata.MDString{Value: stringLit(old.Val())}, nil
	case *ast.MDTuple:
		return gen.irMDTuple(old)
	case *ast.MetadataID:
		id := metadataID(*old)
		node, ok := gen.new.metadataDefs[id]
		if !ok {
			panic(fmt.Errorf("unable to locate metadata ID %q", enc.Metadata(id)))
		}
		return node, nil
	case ast.SpecializedMDNode:
		return gen.irSpecializedMDNode(old)
	default:
		panic(fmt.Errorf("support for metadata %T not yet implemented", old))
	}
}

func (gen *generator) irSpecializedMDNode(old ast.SpecializedMDNode) (metadata.SpecializedMDNode, error) {
	panic("not yet implemented")
}

func (gen *generator) irMetadataNode(old ast.MetadataNode) (metadata.MetadataNode, error) {
	switch old := old.(type) {
	case *ast.MetadataID:
		id := metadataID(*old)
		node, ok := gen.new.metadataDefs[id]
		if !ok {
			return nil, errors.Errorf("unable to locate metadata ID %q", enc.Metadata(id))
		}
		return node, nil
	case *ast.DIExpression:
		return irDIExpression(old), nil
	default:
		panic(fmt.Errorf("support for metadata node %T not yet implemented", old))
	}
}
