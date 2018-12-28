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
func (gen *generator) irMetadataAttachments(olds []ast.MetadataAttachment) ([]*metadata.Attachment, error) {
	if len(olds) == 0 {
		return nil, nil
	}
	mds := make([]*metadata.Attachment, len(olds))
	for i, old := range olds {
		md, err := gen.irMetadataAttachment(old)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		mds[i] = md
	}
	return mds, nil
}

// irMetadataAttachment returns the IR metadata attachment corresponding to
// the given AST metadata attachment.
func (gen *generator) irMetadataAttachment(old ast.MetadataAttachment) (*metadata.Attachment, error) {
	// Name.
	name := metadataName(old.Name())
	// Node.
	node, err := gen.irMDNode(old.MDNode())
	if err != nil {
		return nil, errors.WithStack(err)
	}
	md := &metadata.Attachment{
		Name: name,
		Node: node,
	}
	return md, nil
}

// irMDNode returns the IR metadata node corresponding to the given AST metadata
// node.
func (gen *generator) irMDNode(old ast.MDNode) (metadata.MDNode, error) {
	switch old := old.(type) {
	case *ast.MDTuple:
		return gen.irMDTuple(old)
	case *ast.MetadataID:
		return gen.metadataDefFromID(*old)
	case ast.SpecializedMDNode:
		return gen.irSpecializedMDNode(old)
	default:
		panic(fmt.Errorf("support for metadata node %T not yet implemented", old))
	}
}

// irMDTuple returns the IR metadata tuple corresponding to the given AST
// metadata tuple.
func (gen *generator) irMDTuple(old *ast.MDTuple) (*metadata.Tuple, error) {
	tuple := &metadata.Tuple{}
	if oldFields := old.MDFields().MDFields(); len(oldFields) > 0 {
		tuple.Fields = make([]metadata.Field, len(oldFields))
		for i, oldField := range oldFields {
			field, err := gen.irMDField(oldField)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			tuple.Fields[i] = field
		}
	}
	return tuple, nil
}

// irMDField returns the IR metadata field corresponding to the given AST
// metadata field.
func (gen *generator) irMDField(old ast.MDField) (metadata.Field, error) {
	switch old := old.(type) {
	case *ast.NullLit:
		return metadata.Null, nil
	case ast.Metadata:
		return gen.irMetadata(old)
	default:
		panic(fmt.Errorf("support for metadata field %T not yet implemented", old))
	}
}

// irMetadata returns the IR metadata corresponding to the given AST metadata.
func (fgen *funcGen) irMetadata(old ast.Metadata) (metadata.Metadata, error) {
	switch old := old.(type) {
	case *ast.TypeValue:
		return fgen.irTypeValue(*old)
	default:
		return fgen.gen.irMetadata(old)
	}
}

// irMetadata returns the IR metadata corresponding to the given AST metadata.
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
		s := stringLit(old.Val())
		return &metadata.String{Value: s}, nil
	case *ast.MDTuple:
		return gen.irMDTuple(old)
	case *ast.MetadataID:
		return gen.metadataDefFromID(*old)
	case ast.SpecializedMDNode:
		return gen.irSpecializedMDNode(old)
	default:
		panic(fmt.Errorf("support for metadata %T not yet implemented", old))
	}
}

// irMetadataNode returns the IR metadata node corresponding to the given AST
// metadata node.
func (gen *generator) irMetadataNode(old ast.MetadataNode) (metadata.Node, error) {
	switch old := old.(type) {
	case *ast.MetadataID:
		return gen.metadataDefFromID(*old)
	case *ast.DIExpression:
		return gen.irDIExpression(old)
	default:
		panic(fmt.Errorf("support for metadata node %T not yet implemented", old))
	}
}

// ### [ Helper functions ] ####################################################

// metadataDefFromID returns the IR metadata definition associated with the
// given AST metadata ID.
func (gen *generator) metadataDefFromID(old ast.MetadataID) (*metadata.Def, error) {
	id := metadataID(old)
	node, ok := gen.new.metadataDefs[id]
	if !ok {
		return nil, errors.Errorf("unable to locate metadata ID %q", enc.MetadataID(id))
	}
	return node, nil
}
