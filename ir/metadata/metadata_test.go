package metadata_test

import "github.com/llir/llvm/ir/metadata"

// Validate that the relevant types satisfy the metadata.Node interface.
var (
	_ metadata.Node = &metadata.Metadata{}
	_ metadata.Node = &metadata.String{}
)
