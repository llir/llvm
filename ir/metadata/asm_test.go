package metadata_test

import (
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/llir/llvm/asm"
	"github.com/mewkiz/pkg/diffutil"
	"github.com/mewkiz/pkg/osutil"
)

// words specifies whether to colour words in diff output.
var words bool

func init() {
	flag.BoolVar(&words, "words", false, "colour words in diff output")
	flag.Parse()
}

func TestModule(t *testing.T) {
	golden := []struct {
		path string
	}{
		{path: "../../testdata/coreutils/test/cat.ll"},
	}
	for _, g := range golden {
		if filepath.HasPrefix(g.path, "../../testdata/") {
			if !osutil.Exists("../../testdata/") {
				// Skip test cases from the llir/testdata submodule if not present.
				// Users may add this submodule using git clone --recursive.
				continue
			}
		}
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
