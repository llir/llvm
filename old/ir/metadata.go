package ir

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/llir/llvm/internal/enc"
	"github.com/llir/llvm/ir/metadata"
)

// ### [ Helper functions ] ####################################################

// metadataString returns the string representation of the given metadata, when
// referred from a value using the given separator.
func metadataString(m map[string]*metadata.Metadata, sep string) string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	buf := &bytes.Buffer{}
	for _, key := range keys {
		md := m[key]
		fmt.Fprintf(buf, "%s %s %s", sep, enc.Metadata(key), md.Ident())
	}
	return buf.String()
}
