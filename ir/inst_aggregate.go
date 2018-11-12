package ir

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/metadata"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// --- [ Aggregate instructions ] ----------------------------------------------

// ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstExtractValue is an LLVM IR extractvalue instruction.
type InstExtractValue struct {
	// Name of local variable associated with the result.
	LocalName string
	// Aggregate value.
	X value.Value // array or struct
	// Element indices.
	Indices []int64

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewExtractValue returns a new extractvalue instruction based on the given
// aggregate value and indicies.
func NewExtractValue(x value.Value, indices ...int64) *InstExtractValue {
	inst := &InstExtractValue{X: x, Indices: indices}
	// Compute type.
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstExtractValue) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstExtractValue) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = aggregateElemType(inst.X.Type(), inst.Indices)
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstExtractValue) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstExtractValue) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstExtractValue) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstExtractValue) Def() string {
	// 'extractvalue' X=TypeValue Indices=(',' UintLit)+ Metadata=(','
	// MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "extractvalue %s", inst.X)
	for _, index := range inst.Indices {
		fmt.Fprintf(buf, ", %d", index)
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// InstInsertValue is an LLVM IR insertvalue instruction.
type InstInsertValue struct {
	// Name of local variable associated with the result.
	LocalName string
	// Aggregate value.
	X value.Value // array or struct
	// Element to insert.
	Elem value.Value
	// Element indices.
	Indices []int64

	// extra.

	// Type of result produced by the instruction.
	Typ types.Type
	// (optional) Metadata.
	Metadata []*metadata.MetadataAttachment
}

// NewInsertValue returns a new insertvalue instruction based on the given
// aggregate value, element and indicies.
func NewInsertValue(x, elem value.Value, indices ...int64) *InstInsertValue {
	inst := &InstInsertValue{X: x, Elem: elem, Indices: indices}
	// Compute type.
	return inst
}

// String returns the LLVM syntax representation of the instruction as a
// type-value pair.
func (inst *InstInsertValue) String() string {
	return fmt.Sprintf("%s %s", inst.Type(), inst.Ident())
}

// Type returns the type of the instruction.
func (inst *InstInsertValue) Type() types.Type {
	// Cache type if not present.
	if inst.Typ == nil {
		inst.Typ = inst.X.Type()
	}
	return inst.Typ
}

// Ident returns the identifier associated with the instruction.
func (inst *InstInsertValue) Ident() string {
	return enc.Local(inst.LocalName)
}

// Name returns the name of the instruction.
func (inst *InstInsertValue) Name() string {
	return inst.LocalName
}

// SetName sets the name of the instruction.
func (inst *InstInsertValue) SetName(name string) {
	inst.LocalName = name
}

// Def returns the LLVM syntax representation of the instruction.
func (inst *InstInsertValue) Def() string {
	// 'insertvalue' X=TypeValue ',' Elem=TypeValue Indices=(',' UintLit)+
	// Metadata=(',' MetadataAttachment)+?
	buf := &strings.Builder{}
	fmt.Fprintf(buf, "%s = ", inst.Ident())
	fmt.Fprintf(buf, "insertvalue %s, %s", inst.X, inst.Elem)
	for _, index := range inst.Indices {
		fmt.Fprintf(buf, ", %d", index)
	}
	for _, md := range inst.Metadata {
		fmt.Fprintf(buf, ", %s", md)
	}
	return buf.String()
}

// ### [ Helper functions ] ####################################################

// aggregateElemType returns the element type at the position in the aggregate
// type specified by the given indices.
func aggregateElemType(t types.Type, indices []int64) types.Type {
	// Base case.
	if len(indices) == 0 {
		return t
	}
	switch t := t.(type) {
	case *types.ArrayType:
		return aggregateElemType(t.ElemType, indices[1:])
	case *types.StructType:
		return aggregateElemType(t.Fields[indices[0]], indices[1:])
	default:
		panic(fmt.Errorf("support for aggregate type %T not yet implemented", t))
	}
}
