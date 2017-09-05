package irx

import (
	"fmt"

	"github.com/llir/llvm/asm/internal/ast"
	"github.com/llir/llvm/ir/metadata"
)

// irMetadataNode returns the corresponding LLVM IR metadata node of the given
// metadata node.
func (m *Module) irMetadataNode(old ast.MetadataNode) metadata.Node {
	switch old := old.(type) {
	case *ast.Metadata:
		var nodes []metadata.Node
		for _, oldNode := range old.Nodes {
			nodes = append(nodes, m.irMetadataNode(oldNode))
		}
		return &metadata.Metadata{
			ID:    old.ID,
			Nodes: nodes,
		}
	case *ast.MetadataString:
		return &metadata.String{
			Val: old.Val,
		}
	case *ast.MetadataValue:
		return &metadata.Value{
			X: m.irValue(old.X),
		}
	case ast.Constant:
		c := m.irConstant(old)
		md, ok := c.(metadata.Node)
		if !ok {
			panic(fmt.Errorf("invalid constant type; expected metadata.Node, got %T", c))
		}
		return md
	default:
		panic(fmt.Errorf("support for metadata node %T not yet implemented", old))
	}
}
