package asm

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/mewkiz/pkg/diffutil"
	"github.com/mewkiz/pkg/osutil"
)

func TestParseFile(t *testing.T) {
	golden := []struct {
		path string
	}{
		{path: "testdata/inst_binary.ll"},
		{path: "testdata/inst_bitwise.ll"},
		{path: "testdata/inst_vector.ll"},
		{path: "testdata/inst_aggregate.ll"},
		{path: "testdata/inst_memory.ll"},
		{path: "testdata/inst_conversion.ll"},
		{path: "testdata/inst_other.ll"},
		{path: "testdata/terminator.ll"},
		// LLVM Features.
		//{path: "testdata/Feature/alias2.ll"}, // TODO: fix grammar. syntax error at line 12
		//{path: "testdata/Feature/aliases.ll"}, // TODO: fix grammar. syntax error at line 29
		//{path: "testdata/Feature/alignment.ll"}, // TODO: fix grammar. syntax error at line 7
		{path: "testdata/Feature/attributes.ll"},
		{path: "testdata/Feature/basictest.ll"},
		{path: "testdata/Feature/callingconventions.ll"},
		{path: "testdata/Feature/calltest.ll"},
		{path: "testdata/Feature/casttest.ll"},
		{path: "testdata/Feature/cfgstructures.ll"},
		{path: "testdata/Feature/cold.ll"},
		{path: "testdata/Feature/comdat.ll"},
		//{path: "testdata/Feature/constexpr.ll"}, // TODO: re-enable when signed hex integer literals are supported.
		{path: "testdata/Feature/constpointer.ll"},
		{path: "testdata/Feature/const_pv.ll"},
		{path: "testdata/Feature/elf-linker-options.ll"},
		{path: "testdata/Feature/escaped_label.ll"},
		{path: "testdata/Feature/exception.ll"},
		//{path: "testdata/Feature/float.ll"}, // TODO: re-enable when hex float literals are supported.
		//{path: "testdata/Feature/fold-fpcast.ll"}, // TODO: re-enable when hex float literals are supported.
		{path: "testdata/Feature/forwardreftest.ll"},
		{path: "testdata/Feature/fp-intrinsics.ll"},
		{path: "testdata/Feature/global_pv.ll"},
		//{path: "testdata/Feature/globalredefinition3.ll"}, // TODO: figure out how to test. should report error "redefinition of global '@B'"
		{path: "testdata/Feature/global_section.ll"},
		{path: "testdata/Feature/globalvars.ll"},
		{path: "testdata/Feature/indirectcall2.ll"},
		{path: "testdata/Feature/indirectcall.ll"},
		{path: "testdata/Feature/inlineasm.ll"},
		{path: "testdata/Feature/instructions.ll"},
		{path: "testdata/Feature/intrinsic-noduplicate.ll"},
		//{path: "testdata/Feature/intrinsics.ll"}, // TODO: re-enable when floats are printed using the same format as Clang.
		{path: "testdata/Feature/load_module.ll"},
		{path: "testdata/Feature/md_on_instruction.ll"},
		{path: "testdata/Feature/memorymarkers.ll"},
		{path: "testdata/Feature/metadata.ll"},
		{path: "testdata/Feature/minsize_attr.ll"},
		{path: "testdata/Feature/NamedMDNode2.ll"},
		{path: "testdata/Feature/NamedMDNode.ll"},
		{path: "testdata/Feature/newcasts.ll"},
		{path: "testdata/Feature/optnone.ll"},
		{path: "testdata/Feature/optnone-llc.ll"},
		{path: "testdata/Feature/optnone-opt.ll"},
		//{path: "testdata/Feature/packed.ll"}, // TODO: re-enable when floats are printed using the same format as Clang.
		{path: "testdata/Feature/packed_struct.ll"},
		{path: "testdata/Feature/paramattrs.ll"},
		//{path: "testdata/Feature/ppcld.ll"}, // TODO: re-enable when floats are printed using the same format as Clang.
		{path: "testdata/Feature/prefixdata.ll"},
		{path: "testdata/Feature/prologuedata.ll"},
		{path: "testdata/Feature/properties.ll"},
		{path: "testdata/Feature/prototype.ll"},
		{path: "testdata/Feature/recursivetype.ll"},
		{path: "testdata/Feature/seh-nounwind.ll"},
		{path: "testdata/Feature/simplecalltest.ll"},
		{path: "testdata/Feature/smallest.ll"},
		{path: "testdata/Feature/small.ll"},
		//{path: "testdata/Feature/sparcld.ll"}, // TODO: re-enable when floats are printed using the same format as Clang.
		{path: "testdata/Feature/strip_names.ll"},
		//{path: "testdata/Feature/terminators.ll"}, // TODO: fix grammar. syntax error at line 35
		//{path: "testdata/Feature/testalloca.ll"}, // TODO: re-enable when floats are printed using the same format as Clang.
		{path: "testdata/Feature/testconstants.ll"},
		{path: "testdata/Feature/testlogical.ll"},
		//{path: "testdata/Feature/testtype.ll"}, // TODO: fix nil pointer dereference
		{path: "testdata/Feature/testvarargs.ll"},
		{path: "testdata/Feature/undefined.ll"},
		{path: "testdata/Feature/unreachable.ll"},
		{path: "testdata/Feature/varargs.ll"},
		{path: "testdata/Feature/varargs_new.ll"},
		//{path: "testdata/Feature/vector-cast-constant-exprs.ll"}, // TODO: re-enable when floats are printed using the same format as Clang.
		{path: "testdata/Feature/weak_constant.ll"},
		//{path: "testdata/Feature/weirdnames.ll"}, // TODO: re-enable when floats are printed using the same format as Clang.
		//{path: "testdata/Feature/x86ld.ll"}, // TODO: re-enable when floats are printed using the same format as Clang.
	}
	for _, g := range golden {
		m, err := ParseFile(g.path)
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
			if err := diffutil.Diff(want, got, false, filepath.Base(path)); err != nil {
				panic(err)
			}
			t.Errorf("module mismatch %q; expected `%s`, got `%s`", path, want, got)
			continue
		}
	}
}
