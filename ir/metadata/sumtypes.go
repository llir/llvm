package metadata

import "fmt"

// TODO: implement.

type MetadataNode interface {
	fmt.Stringer
}

// MDTuple or SpecializedMDNode
type MDNode interface {
	fmt.Stringer
}

type MDField interface {
	fmt.Stringer
}

type SpecializedMDNode interface {
	fmt.Stringer
}

type MDFieldOrInt interface {
	fmt.Stringer
}

type DIExpressionField interface {
	fmt.Stringer
	// IsDIExpressionField ensures that only DIExpression fields can be assigned
	// to the metadata.DIExpressionField interface.
	IsDIExpressionField()
}

type Metadata interface {
	fmt.Stringer
}
