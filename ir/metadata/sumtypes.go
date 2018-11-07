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
}

type IntOrMDField interface {
}

type DIExpressionField interface {
	fmt.Stringer
}
