// === [ Aggregate instructions ] ==============================================
//
// References:
//    http://llvm.org/docs/LangRef.html#aggregate-operations

package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"github.com/pkg/errors"
)

// --- [ extractvalue ] --------------------------------------------------------

// InstExtractValue represents an extractvalue instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#extractvalue-instruction
type InstExtractValue struct {
	// Parent basic block.
	Parent *BasicBlock
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Typ types.Type
	// Vector.
	X value.Value
	// Indices.
	Indices []int64
	// Map from metadata identifier (e.g. !dbg) to metadata associated with the
	// instruction.
	Metadata map[string]*metadata.Metadata
}

// NewExtractValue returns a new extractvalue instruction based on the given
// vector and indices.
func NewExtractValue(x value.Value, indices []int64) *InstExtractValue {
	typ, err := aggregateElemType(x.Type(), indices)
	if err != nil {
		panic(err)
	}
	return &InstExtractValue{
		Typ:      typ,
		X:        x,
		Indices:  indices,
		Metadata: make(map[string]*metadata.Metadata),
	}
}

// Type returns the type of the instruction.
func (inst *InstExtractValue) Type() types.Type {
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstExtractValue) Ident() string {
	return enc.Local(inst.Name)
}

// GetName returns the name of the local variable associated with the
// instruction.
func (inst *InstExtractValue) GetName() string {
	return inst.Name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstExtractValue) SetName(name string) {
	inst.Name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstExtractValue) String() string {
	indices := &bytes.Buffer{}
	for _, index := range inst.Indices {
		fmt.Fprintf(indices, ", %d", index)
	}
	md := metadataString(inst.Metadata, ",")
	return fmt.Sprintf("%s = extractvalue %s %s%s%s",
		inst.Ident(),
		inst.X.Type(),
		inst.X.Ident(),
		indices,
		md)
}

// GetParent returns the parent basic block of the instruction.
func (inst *InstExtractValue) GetParent() *BasicBlock {
	return inst.Parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstExtractValue) SetParent(parent *BasicBlock) {
	inst.Parent = parent
}

// --- [ insertvalue ] ---------------------------------------------------------

// InstInsertValue represents an insertvalue instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#insertvalue-instruction
type InstInsertValue struct {
	// Parent basic block.
	Parent *BasicBlock
	// Name of the local variable associated with the instruction.
	Name string
	// Vector.
	X value.Value
	// Element to insert.
	Elem value.Value
	// Indices.
	Indices []int64
	// Map from metadata identifier (e.g. !dbg) to metadata associated with the
	// instruction.
	Metadata map[string]*metadata.Metadata
}

// NewInsertValue returns a new insertvalue instruction based on the given
// vector, element and indices.
func NewInsertValue(x, elem value.Value, indices []int64) *InstInsertValue {
	return &InstInsertValue{
		X:        x,
		Elem:     elem,
		Indices:  indices,
		Metadata: make(map[string]*metadata.Metadata),
	}
}

// Type returns the type of the instruction.
func (inst *InstInsertValue) Type() types.Type {
	return inst.X.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstInsertValue) Ident() string {
	return enc.Local(inst.Name)
}

// GetName returns the name of the local variable associated with the
// instruction.
func (inst *InstInsertValue) GetName() string {
	return inst.Name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstInsertValue) SetName(name string) {
	inst.Name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstInsertValue) String() string {
	indices := &bytes.Buffer{}
	for _, index := range inst.Indices {
		fmt.Fprintf(indices, ", %d", index)
	}
	md := metadataString(inst.Metadata, ",")
	return fmt.Sprintf("%s = insertvalue %s %s, %s %s%s%s",
		inst.Ident(),
		inst.X.Type(),
		inst.X.Ident(),
		inst.Elem.Type(),
		inst.Elem.Ident(),
		indices,
		md)
}

// GetParent returns the parent basic block of the instruction.
func (inst *InstInsertValue) GetParent() *BasicBlock {
	return inst.Parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstInsertValue) SetParent(parent *BasicBlock) {
	inst.Parent = parent
}

// ### [ Helper functions ] ####################################################

// aggregateElemType returns the element type of the given aggregate type, based
// on the specified indices.
func aggregateElemType(t types.Type, indices []int64) (types.Type, error) {
	if len(indices) == 0 {
		return t, nil
	}
	index := indices[0]
	switch t := t.(type) {
	case *types.ArrayType:
		if index >= t.Len {
			return nil, errors.Errorf("invalid index (%d); exceeds array length (%d)", index, t.Len)
		}
		return aggregateElemType(t.Elem, indices[1:])
	case *types.StructType:
		if index >= int64(len(t.Fields)) {
			return nil, errors.Errorf("invalid index (%d); exceeds struct field count (%d)", index, len(t.Fields))
		}
		return aggregateElemType(t.Fields[index], indices[1:])
	default:
		return nil, errors.Errorf("invalid aggregate value type; expected *types.ArrayType or *types.StructType, got %T", t)
	}
}
