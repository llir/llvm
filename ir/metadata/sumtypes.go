package metadata

import "fmt"

// TODO: implement.

type Node interface {
	fmt.Stringer
}

// Tuple or SpecializedNode
type MDNode interface {
	fmt.Stringer
}

type Field interface {
	fmt.Stringer
}

type SpecializedNode interface {
	fmt.Stringer
}

type FieldOrInt interface {
	fmt.Stringer
}

type DIExpressionField interface {
	fmt.Stringer
	// IsDIExpressionField ensures that only DIExpression fields can be assigned
	// to the metadata.DIExpressionField interface.
	IsDIExpressionField()
}

// IsDIExpressionField ensures that only DIExpression fields can be assigned to
// the metadata.DIExpressionField interface.
func (UintLit) IsDIExpressionField() {}

type Metadata interface {
	fmt.Stringer
}
