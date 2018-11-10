package ir

import "github.com/llir/llvm/ir/constant"

// NewAlias appends a new alias to the module based on the given alias name and
// aliasee.
func (m *Module) NewAlias(name string, aliasee constant.Constant) *Alias {
	alias := NewAlias(name, aliasee)
	m.Aliases = append(m.Aliases, alias)
	return alias
}
