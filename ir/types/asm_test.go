package types_test

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/llir/llvm/asm"
	"github.com/mewkiz/pkg/diffutil"
	"github.com/mewkiz/pkg/osutil"
)

// words specifies whether to colour words in diff output.
const words = false

func TestModule(t *testing.T) {
	golden := []struct {
		path string
	}{
		// LLVM IR types.
		{path: "testdata/types.ll"},
	}
	for _, g := range golden {
		log.Printf("=== [ %s ] ===", g.path)
		m, err := asm.ParseFile(g.path)
		if err != nil {
			t.Errorf("unable to parse %q into AST; %+v", g.path, err)
			continue
		}
		path := g.path
		if osutil.Exists(g.path + ".golden") {
			path = g.path + ".golden"
		}
		buf, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("unable to read %q; %+v", path, err)
			continue
		}
		want := string(buf)
		got := m.String()
		if want != got {
			if err := diffutil.Diff(want, got, words, filepath.Base(path)); err != nil {
				panic(err)
			}
			t.Errorf("module mismatch %q; expected `%s`, got `%s`", path, want, got)
			continue
		}
	}
}
