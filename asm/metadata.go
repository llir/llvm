package asm

import (
	"fmt"

	"github.com/llir/ll/ast"
	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/metadata"
	"github.com/pkg/errors"
)

// irMetadataAttachments returns the IR metadata attachments corresponding to
// the given AST metadata attachments.
func (gen *generator) irMetadataAttachments(ns []ast.MetadataAttachment) ([]*metadata.MetadataAttachment, error) {
	var mds []*metadata.MetadataAttachment
	for _, n := range ns {
		md, err := gen.irMetadataAttachment(n)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		mds = append(mds, md)
	}
	return mds, nil
}

// irMetadataAttachment returns the IR metadata attachment corresponding to
// the given AST metadata attachment.
func (gen *generator) irMetadataAttachment(old ast.MetadataAttachment) (*metadata.MetadataAttachment, error) {
	// Name.
	name := metadataName(old.Name())
	// Node.
	node, err := gen.irMDNode(old.MDNode())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	md := &metadata.MetadataAttachment{
		Name: name,
		Node: node,
	}
	return md, nil
}

func (gen *generator) irMDNode(old ast.MDNode) (metadata.MDNode, error) {
	switch old := old.(type) {
	case *ast.MDTuple:
		return gen.irMDTuple(old)
	case *ast.MetadataID:
		id := metadataID(*old)
		node, ok := gen.new.metadataDefs[id]
		if !ok {
			panic(fmt.Errorf("unable to locate metadata ID %q", enc.MetadataID(id)))
		}
		return node, nil
	case ast.SpecializedMDNode:
		return gen.irSpecializedMDNode(old)
	default:
		panic(fmt.Errorf("support for metadata node %T not yet implemented", old))
	}
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
	case *ast.NullLit:
		return metadata.Null, nil
	case ast.Metadata:
		return gen.irMetadata(old)
	default:
		panic(fmt.Errorf("support for metadata field %T not yet implemented", old))
	}
}

func (fgen *funcGen) irMetadata(old ast.Metadata) (metadata.Metadata, error) {
	switch old := old.(type) {
	case *ast.TypeValue:
		return fgen.astToIRTypeValue(*old)
	default:
		return fgen.gen.irMetadata(old)
	}
}

func (gen *generator) irMetadata(old ast.Metadata) (metadata.Metadata, error) {
	switch old := old.(type) {
	case *ast.TypeValue:
		typ, err := gen.irType(old.Typ())
		if err != nil {
			return nil, errors.WithStack(err)
		}
		switch oldVal := old.Val().(type) {
		case ast.Constant:
			return gen.irConstant(typ, oldVal)
		default:
			panic(fmt.Errorf("support for metadata value %T not yet implemented", oldVal))
		}
	case *ast.MDString:
		return &metadata.MDString{Value: stringLit(old.Val())}, nil
	case *ast.MDTuple:
		return gen.irMDTuple(old)
	case *ast.MetadataID:
		id := metadataID(*old)
		node, ok := gen.new.metadataDefs[id]
		if !ok {
			panic(fmt.Errorf("unable to locate metadata ID %q", enc.MetadataID(id)))
		}
		return node, nil
	case ast.SpecializedMDNode:
		return gen.irSpecializedMDNode(old)
	default:
		panic(fmt.Errorf("support for metadata %T not yet implemented", old))
	}
}

func (gen *generator) irMetadataNode(old ast.MetadataNode) (metadata.MetadataNode, error) {
	switch old := old.(type) {
	case *ast.MetadataID:
		id := metadataID(*old)
		node, ok := gen.new.metadataDefs[id]
		if !ok {
			return nil, errors.Errorf("unable to locate metadata ID %q", enc.MetadataID(id))
		}
		return node, nil
	case *ast.DIExpression:
		return gen.irDIExpression(old)
	default:
		panic(fmt.Errorf("support for metadata node %T not yet implemented", old))
	}
}
