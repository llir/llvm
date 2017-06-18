// === [ Vector instructions ] =================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#vector-operations

package ir

import (
	"fmt"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ extractelement ] ------------------------------------------------------

// InstExtractElement represents an extractelement instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#extractelement-instruction
type InstExtractElement struct {
	// Parent basic block.
	Parent *BasicBlock
	// Name of the local variable associated with the instruction.
	Name string
	// Type of the instruction.
	Typ types.Type
	// Vector.
	X value.Value
	// Index.
	Index value.Value
	// Map from metadata identifier (e.g. !dbg) to metadata associated with the
	// instruction.
	Metadata map[string]*metadata.Metadata
}

// NewExtractElement returns a new extractelement instruction based on the given
// vector and index.
func NewExtractElement(x, index value.Value) *InstExtractElement {
	t, ok := x.Type().(*types.VectorType)
	if !ok {
		panic(fmt.Errorf("invalid vector type; expected *types.VectorType, got %T", x.Type()))
	}
	return &InstExtractElement{
		Typ:      t.Elem,
		X:        x,
		Index:    index,
		Metadata: make(map[string]*metadata.Metadata),
	}
}

// Type returns the type of the instruction.
func (inst *InstExtractElement) Type() types.Type {
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstExtractElement) Ident() string {
	return enc.Local(inst.Name)
}

// GetName returns the name of the local variable associated with the
// instruction.
func (inst *InstExtractElement) GetName() string {
	return inst.Name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstExtractElement) SetName(name string) {
	inst.Name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstExtractElement) String() string {
	md := metadataString(inst.Metadata, ",")
	return fmt.Sprintf("%s = extractelement %s %s, %s %s%s",
		inst.Ident(),
		inst.X.Type(),
		inst.X.Ident(),
		inst.Index.Type(),
		inst.Index.Ident(),
		md)
}

// GetParent returns the parent basic block of the instruction.
func (inst *InstExtractElement) GetParent() *BasicBlock {
	return inst.Parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstExtractElement) SetParent(parent *BasicBlock) {
	inst.Parent = parent
}

// --- [ insertelement ] -------------------------------------------------------

// InstInsertElement represents an insertelement instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#insertelement-instruction
type InstInsertElement struct {
	// Parent basic block.
	Parent *BasicBlock
	// Name of the local variable associated with the instruction.
	Name string
	// Vector.
	X value.Value
	// Element to insert.
	Elem value.Value
	// Index.
	Index value.Value
	// Map from metadata identifier (e.g. !dbg) to metadata associated with the
	// instruction.
	Metadata map[string]*metadata.Metadata
}

// NewInsertElement returns a new insertelement instruction based on the given
// vector, element and index.
func NewInsertElement(x, elem, index value.Value) *InstInsertElement {
	return &InstInsertElement{
		X:        x,
		Elem:     elem,
		Index:    index,
		Metadata: make(map[string]*metadata.Metadata),
	}
}

// Type returns the type of the instruction.
func (inst *InstInsertElement) Type() types.Type {
	return inst.X.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstInsertElement) Ident() string {
	return enc.Local(inst.Name)
}

// GetName returns the name of the local variable associated with the
// instruction.
func (inst *InstInsertElement) GetName() string {
	return inst.Name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstInsertElement) SetName(name string) {
	inst.Name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstInsertElement) String() string {
	md := metadataString(inst.Metadata, ",")
	return fmt.Sprintf("%s = insertelement %s %s, %s %s, %s %s%s",
		inst.Ident(),
		inst.X.Type(),
		inst.X.Ident(),
		inst.Elem.Type(),
		inst.Elem.Ident(),
		inst.Index.Type(),
		inst.Index.Ident(),
		md)
}

// GetParent returns the parent basic block of the instruction.
func (inst *InstInsertElement) GetParent() *BasicBlock {
	return inst.Parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstInsertElement) SetParent(parent *BasicBlock) {
	inst.Parent = parent
}

// --- [ shufflevector ] -------------------------------------------------------

// InstShuffleVector represents an shufflevector instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#shufflevector-instruction
type InstShuffleVector struct {
	// Parent basic block.
	Parent *BasicBlock
	// Name of the local variable associated with the instruction.
	Name string
	// Vector 1.
	X value.Value
	// Vector 2.
	Y value.Value
	// Shuffle mask.
	Mask value.Value
	// Map from metadata identifier (e.g. !dbg) to metadata associated with the
	// instruction.
	Metadata map[string]*metadata.Metadata
}

// NewShuffleVector returns a new shufflevector instruction based on the given
// vectors and shuffle mask.
func NewShuffleVector(x, y, mask value.Value) *InstShuffleVector {
	return &InstShuffleVector{
		X:        x,
		Y:        y,
		Mask:     mask,
		Metadata: make(map[string]*metadata.Metadata),
	}
}

// Type returns the type of the instruction.
func (inst *InstShuffleVector) Type() types.Type {
	return inst.Mask.Type()
}

// Ident returns the identifier associated with the instruction.
func (inst *InstShuffleVector) Ident() string {
	return enc.Local(inst.Name)
}

// GetName returns the name of the local variable associated with the
// instruction.
func (inst *InstShuffleVector) GetName() string {
	return inst.Name
}

// SetName sets the name of the local variable associated with the instruction.
func (inst *InstShuffleVector) SetName(name string) {
	inst.Name = name
}

// String returns the LLVM syntax representation of the instruction.
func (inst *InstShuffleVector) String() string {
	md := metadataString(inst.Metadata, ",")
	return fmt.Sprintf("%s = shufflevector %s %s, %s %s, %s %s%s",
		inst.Ident(),
		inst.X.Type(),
		inst.X.Ident(),
		inst.Y.Type(),
		inst.Y.Ident(),
		inst.Mask.Type(),
		inst.Mask.Ident(),
		md)
}

// GetParent returns the parent basic block of the instruction.
func (inst *InstShuffleVector) GetParent() *BasicBlock {
	return inst.Parent
}

// SetParent sets the parent basic block of the instruction.
func (inst *InstShuffleVector) SetParent(parent *BasicBlock) {
	inst.Parent = parent
}
