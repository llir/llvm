package asm

import (
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"

	"github.com/mewkiz/pkg/diffutil"
	"github.com/mewkiz/pkg/osutil"
)

// words specifies whether to colour words in diff output.
var words bool

func init() {
	flag.BoolVar(&words, "words", false, "colour words in diff output")
	flag.Parse()
}

func TestParseFile(t *testing.T) {
	golden := []struct {
		path string
	}{
		{path: "testdata/hexfloat.ll"},
		{path: "testdata/inst_aggregate.ll"},
		{path: "testdata/inst_binary.ll"},
		{path: "testdata/inst_bitwise.ll"},
		{path: "testdata/inst_conversion.ll"},
		{path: "testdata/inst_memory.ll"},
		{path: "testdata/inst_other.ll"},
		{path: "testdata/inst_vector.ll"},
		{path: "testdata/terminator.ll"},

		// DIExpression used in named metdata definition.
		{path: "testdata/diexpression.ll"},

		// frem constant expression.
		{path: "testdata/expr_frem.ll"},

		// LLVM IR compatability.
		{path: "../testdata/llvm/test/Bitcode/compatibility.ll"},

		// Floating-point test cases (issue #31).
		{path: "../testdata/llvm/test/Analysis/CostModel/AMDGPU/fdiv.ll"},

		// float infinity.
		//{path: "../testdata/llvm/test/Assembler/2002-04-07-InfConstant.ll"}, // TODO: fix handling of float infinity hex representation.

		// Distinguish named from unnamed locals (issue #39).
		{path: "../testdata/llvm/test/Analysis/DominanceFrontier/new_pm_test.ll"},

		// Empty array constant.
		{path: "../testdata/llvm/test/Assembler/aggregate-constant-values.ll"},

		// gep with vector indices.
		{path: "../testdata/llvm/test/Assembler/ConstantExprFold.ll"},
		{path: "../testdata/llvm/test/Assembler/getelementptr.ll"},

		// Large values in metadata.
		{path: "../testdata/llvm/test/Assembler/ditype-large-values.ll"},

		// fadd, fmul and fcmp constant expressions.
		{path: "../testdata/llvm/test/DebugInfo/ARM/selectiondag-deadcode.ll"},

		// fsub constant expressions.
		{path: "../testdata/llvm/test/Transforms/InstCombine/fma.ll"},

		// Constant expressions.
		{path: "../testdata/llvm/test/Transforms/ConstProp/constant-expr.ll"},

		// Vector constant expressions.
		{path: "../testdata/llvm/test/Transforms/InstCombine/vec_demanded_elts.ll"},
		{path: "../testdata/llvm/test/Transforms/InstCombine/vector_insertelt_shuffle.ll"},

		// Use of address space in function declaration and dereferenable
		// parameter attribute.
		{path: "../testdata/llvm/test/Transforms/InstSimplify/compare.ll"},

		// LLVM test/Features.
		{path: "../testdata/llvm/test/Feature/OperandBundles/adce.ll"},
		{path: "../testdata/llvm/test/Feature/OperandBundles/basic-aa-argmemonly.ll"},
		{path: "../testdata/llvm/test/Feature/OperandBundles/dse.ll"},
		{path: "../testdata/llvm/test/Feature/OperandBundles/early-cse.ll"},
		{path: "../testdata/llvm/test/Feature/OperandBundles/function-attrs.ll"},
		{path: "../testdata/llvm/test/Feature/OperandBundles/inliner-conservative.ll"},
		{path: "../testdata/llvm/test/Feature/OperandBundles/merge-func.ll"},
		{path: "../testdata/llvm/test/Feature/OperandBundles/pr26510.ll"},
		{path: "../testdata/llvm/test/Feature/OperandBundles/special-state.ll"},
		//{path: "../testdata/llvm/test/Feature/alias2.ll"}, // TODO: fix grammar. syntax error at line 12. The issue is that the aliasee (a bitcast expression in this case is missing a type (in this case, the type should be i16*). `@a1 = alias i16, bitcast (i32* @v1 to i16*)`
		//{path: "../testdata/llvm/test/Feature/aliases.ll"}, // TODO: fix grammar. syntax error at line 29. The issue is that the aliasee (a bitcast expression in this case is missing a type (in this case, the type should be i64*). `@A = alias i64, bitcast (i32* @bar to i64*)`
		//{path: "../testdata/llvm/test/Feature/alignment.ll"}, // TODO: fix grammar. syntax error at line 7. The issue is that there is a parsing ambiguity between GlobalAttr and FuncAttr, both of which may be empty and both of which may contain Align.
		{path: "../testdata/llvm/test/Feature/attributes.ll"},
		{path: "../testdata/llvm/test/Feature/basictest.ll"},
		{path: "../testdata/llvm/test/Feature/callingconventions.ll"},
		{path: "../testdata/llvm/test/Feature/calltest.ll"},
		{path: "../testdata/llvm/test/Feature/casttest.ll"},
		{path: "../testdata/llvm/test/Feature/cfgstructures.ll"},
		{path: "../testdata/llvm/test/Feature/cold.ll"},
		{path: "../testdata/llvm/test/Feature/comdat.ll"},
		//{path: "../testdata/llvm/test/Feature/constexpr.ll"}, // TODO: re-enable when signed hex integer literals are supported.
		{path: "../testdata/llvm/test/Feature/constpointer.ll"},
		{path: "../testdata/llvm/test/Feature/const_pv.ll"},
		{path: "../testdata/llvm/test/Feature/elf-linker-options.ll"},
		{path: "../testdata/llvm/test/Feature/escaped_label.ll"},
		{path: "../testdata/llvm/test/Feature/exception.ll"},
		{path: "../testdata/llvm/test/Feature/float.ll"},
		{path: "../testdata/llvm/test/Feature/fold-fpcast.ll"},
		{path: "../testdata/llvm/test/Feature/forwardreftest.ll"},
		{path: "../testdata/llvm/test/Feature/fp-intrinsics.ll"},
		{path: "../testdata/llvm/test/Feature/global_pv.ll"},
		{path: "../testdata/llvm/test/Feature/global_section.ll"},
		{path: "../testdata/llvm/test/Feature/globalvars.ll"},
		{path: "../testdata/llvm/test/Feature/indirectcall2.ll"},
		{path: "../testdata/llvm/test/Feature/indirectcall.ll"},
		{path: "../testdata/llvm/test/Feature/inlineasm.ll"},
		{path: "../testdata/llvm/test/Feature/instructions.ll"},
		{path: "../testdata/llvm/test/Feature/intrinsic-noduplicate.ll"},
		{path: "../testdata/llvm/test/Feature/intrinsics.ll"},
		{path: "../testdata/llvm/test/Feature/load_module.ll"},
		{path: "../testdata/llvm/test/Feature/md_on_instruction.ll"},
		{path: "../testdata/llvm/test/Feature/memorymarkers.ll"},
		{path: "../testdata/llvm/test/Feature/metadata.ll"},
		{path: "../testdata/llvm/test/Feature/minsize_attr.ll"},
		{path: "../testdata/llvm/test/Feature/NamedMDNode2.ll"},
		{path: "../testdata/llvm/test/Feature/NamedMDNode.ll"},
		{path: "../testdata/llvm/test/Feature/newcasts.ll"},
		{path: "../testdata/llvm/test/Feature/optnone.ll"},
		{path: "../testdata/llvm/test/Feature/optnone-llc.ll"},
		{path: "../testdata/llvm/test/Feature/optnone-opt.ll"},
		{path: "../testdata/llvm/test/Feature/packed.ll"},
		{path: "../testdata/llvm/test/Feature/packed_struct.ll"},
		{path: "../testdata/llvm/test/Feature/paramattrs.ll"},
		{path: "../testdata/llvm/test/Feature/ppcld.ll"},
		{path: "../testdata/llvm/test/Feature/prefixdata.ll"},
		{path: "../testdata/llvm/test/Feature/prologuedata.ll"},
		{path: "../testdata/llvm/test/Feature/properties.ll"},
		{path: "../testdata/llvm/test/Feature/prototype.ll"},
		{path: "../testdata/llvm/test/Feature/recursivetype.ll"},
		{path: "../testdata/llvm/test/Feature/seh-nounwind.ll"},
		{path: "../testdata/llvm/test/Feature/simplecalltest.ll"},
		{path: "../testdata/llvm/test/Feature/smallest.ll"},
		{path: "../testdata/llvm/test/Feature/small.ll"},
		{path: "../testdata/llvm/test/Feature/sparcld.ll"},
		{path: "../testdata/llvm/test/Feature/strip_names.ll"},
		{path: "../testdata/llvm/test/Feature/terminators.ll"},
		{path: "../testdata/llvm/test/Feature/testalloca.ll"},
		{path: "../testdata/llvm/test/Feature/testconstants.ll"},
		{path: "../testdata/llvm/test/Feature/testlogical.ll"},
		//{path: "../testdata/llvm/test/Feature/testtype.ll"}, // TODO: fix nil pointer dereference
		{path: "../testdata/llvm/test/Feature/testvarargs.ll"},
		{path: "../testdata/llvm/test/Feature/undefined.ll"},
		{path: "../testdata/llvm/test/Feature/unreachable.ll"},
		{path: "../testdata/llvm/test/Feature/varargs.ll"},
		{path: "../testdata/llvm/test/Feature/varargs_new.ll"},
		{path: "../testdata/llvm/test/Feature/vector-cast-constant-exprs.ll"},
		{path: "../testdata/llvm/test/Feature/weak_constant.ll"},
		{path: "../testdata/llvm/test/Feature/weirdnames.ll"},
		{path: "../testdata/llvm/test/Feature/x86ld.ll"},

		// LLVM test/Assembler.
		{path: "../testdata/llvm/test/Assembler/2002-03-08-NameCollision.ll"},
		{path: "../testdata/llvm/test/Assembler/2002-03-08-NameCollision2.ll"},
		{path: "../testdata/llvm/test/Assembler/2002-04-07-HexFloatConstants.ll"},
		{path: "../testdata/llvm/test/Assembler/2002-04-29-NameBinding.ll"},
		{path: "../testdata/llvm/test/Assembler/2002-05-02-InvalidForwardRef.ll"},
		{path: "../testdata/llvm/test/Assembler/2002-07-14-OpaqueType.ll"},
		{path: "../testdata/llvm/test/Assembler/2002-07-25-QuoteInString.ll"},
		//{path: "../testdata/llvm/test/Assembler/2002-07-25-ReturnPtrFunction.ll"}, // TODO: figure out how function return type should be printed in call instructions.
		{path: "../testdata/llvm/test/Assembler/2002-07-31-SlashInString.ll"},
		{path: "../testdata/llvm/test/Assembler/2002-08-15-CastAmbiguity.ll"},
		{path: "../testdata/llvm/test/Assembler/2002-08-15-ConstantExprProblem.ll"},
		{path: "../testdata/llvm/test/Assembler/2002-08-15-UnresolvedGlobalReference.ll"},
		{path: "../testdata/llvm/test/Assembler/2002-08-16-ConstExprInlined.ll"},
		{path: "../testdata/llvm/test/Assembler/2002-08-19-BytecodeReader.ll"},
		{path: "../testdata/llvm/test/Assembler/2002-08-22-DominanceProblem.ll"},
		{path: "../testdata/llvm/test/Assembler/2002-10-08-LargeArrayPerformance.ll"},
		{path: "../testdata/llvm/test/Assembler/2002-10-13-ConstantEncodingProblem.ll"},
		{path: "../testdata/llvm/test/Assembler/2002-12-15-GlobalResolve.ll"},
		{path: "../testdata/llvm/test/Assembler/2003-01-30-UnsignedString.ll"},
		{path: "../testdata/llvm/test/Assembler/2003-04-15-ConstantInitAssertion.ll"},
		{path: "../testdata/llvm/test/Assembler/2003-04-25-UnresolvedGlobalReference.ll"},
		{path: "../testdata/llvm/test/Assembler/2003-05-03-BytecodeReaderProblem.ll"},
		{path: "../testdata/llvm/test/Assembler/2003-05-12-MinIntProblem.ll"},
		{path: "../testdata/llvm/test/Assembler/2003-05-15-AssemblerProblem.ll"},
		{path: "../testdata/llvm/test/Assembler/2003-05-15-SwitchBug.ll"},
		{path: "../testdata/llvm/test/Assembler/2003-05-21-ConstantShiftExpr.ll"},
		{path: "../testdata/llvm/test/Assembler/2003-05-21-EmptyStructTest.ll"},
		{path: "../testdata/llvm/test/Assembler/2003-05-21-MalformedShiftCrash.ll"},
		{path: "../testdata/llvm/test/Assembler/2003-05-21-MalformedStructCrash.ll"},
		{path: "../testdata/llvm/test/Assembler/2003-08-20-ConstantExprGEP-Fold.ll"},
		{path: "../testdata/llvm/test/Assembler/2003-08-21-ConstantExprCast-Fold.ll"},
		{path: "../testdata/llvm/test/Assembler/2003-11-05-ConstantExprShift.ll"},
		{path: "../testdata/llvm/test/Assembler/2003-11-12-ConstantExprCast.ll"},
		{path: "../testdata/llvm/test/Assembler/2004-01-11-getelementptrfolding.ll"},
		{path: "../testdata/llvm/test/Assembler/2004-01-20-MaxLongLong.ll"},
		{path: "../testdata/llvm/test/Assembler/2004-02-01-NegativeZero.ll"},
		{path: "../testdata/llvm/test/Assembler/2004-02-27-SelfUseAssertError.ll"},
		{path: "../testdata/llvm/test/Assembler/2004-03-07-FunctionAddressAlignment.ll"},
		{path: "../testdata/llvm/test/Assembler/2004-04-04-GetElementPtrIndexTypes.ll"},
		{path: "../testdata/llvm/test/Assembler/2004-06-07-VerifierBug.ll"},
		{path: "../testdata/llvm/test/Assembler/2004-10-22-BCWriterUndefBug.ll"},
		{path: "../testdata/llvm/test/Assembler/2005-01-03-FPConstantDisassembly.ll"},
		{path: "../testdata/llvm/test/Assembler/2005-01-31-CallingAggregateFunction.ll"},
		{path: "../testdata/llvm/test/Assembler/2005-05-05-OpaqueUndefValues.ll"},
		{path: "../testdata/llvm/test/Assembler/2005-12-21-ZeroInitVector.ll"},
		{path: "../testdata/llvm/test/Assembler/2006-12-09-Cast-To-Bool.ll"},
		{path: "../testdata/llvm/test/Assembler/2007-01-05-Cmp-ConstExpr.ll"},
		{path: "../testdata/llvm/test/Assembler/2007-01-16-CrashOnBadCast.ll"},
		{path: "../testdata/llvm/test/Assembler/2007-01-16-CrashOnBadCast2.ll"},
		//{path: "../testdata/llvm/test/Assembler/2007-03-18-InvalidNumberedVar.ll"}, // TODO: fix type checking during translation of local variables.
		{path: "../testdata/llvm/test/Assembler/2007-03-19-NegValue.ll"},
		{path: "../testdata/llvm/test/Assembler/2007-04-20-AlignedLoad.ll"},
		{path: "../testdata/llvm/test/Assembler/2007-04-20-AlignedStore.ll"},
		{path: "../testdata/llvm/test/Assembler/2007-04-25-AssemblerFoldExternWeak.ll"},
		{path: "../testdata/llvm/test/Assembler/2007-05-21-Escape.ll"},
		{path: "../testdata/llvm/test/Assembler/2007-07-19-ParamAttrAmbiguity.ll"},
		{path: "../testdata/llvm/test/Assembler/2007-09-10-AliasFwdRef.ll"},
		{path: "../testdata/llvm/test/Assembler/2007-09-29-GC.ll"},
		//{path: "../testdata/llvm/test/Assembler/2007-12-11-AddressSpaces.ll"}, // TODO: fix translation of pointer with address space to pointer with address space; e.g. i32 addrspace(33)* addrspace(44)*
		{path: "../testdata/llvm/test/Assembler/2008-01-11-VarargAttrs.ll"},
		{path: "../testdata/llvm/test/Assembler/2008-07-10-APInt.ll"},
		{path: "../testdata/llvm/test/Assembler/2008-09-02-FunctionNotes.ll"},
		{path: "../testdata/llvm/test/Assembler/2008-09-02-FunctionNotes2.ll"},
		{path: "../testdata/llvm/test/Assembler/2008-09-29-RetAttr.ll"},
		{path: "../testdata/llvm/test/Assembler/2008-10-14-QuoteInName.ll"},
		{path: "../testdata/llvm/test/Assembler/2009-02-01-UnnamedForwardRef.ll"},
		{path: "../testdata/llvm/test/Assembler/2009-02-28-CastOpc.ll"},
		{path: "../testdata/llvm/test/Assembler/2009-02-28-StripOpaqueName.ll"},
		{path: "../testdata/llvm/test/Assembler/2009-03-24-ZextConstantExpr.ll"},
		{path: "../testdata/llvm/test/Assembler/2009-07-24-ZeroArgGEP.ll"},
		{path: "../testdata/llvm/test/Assembler/2010-02-05-FunctionLocalMetadataBecomesNull.ll"},
		{path: "../testdata/llvm/test/Assembler/addrspacecast-alias.ll"},
		{path: "../testdata/llvm/test/Assembler/aggregate-return-single-value.ll"},
		{path: "../testdata/llvm/test/Assembler/alias-use-list-order.ll"},
		{path: "../testdata/llvm/test/Assembler/align-inst-alloca.ll"},
		{path: "../testdata/llvm/test/Assembler/align-inst.ll"},
		{path: "../testdata/llvm/test/Assembler/alignstack.ll"},
		{path: "../testdata/llvm/test/Assembler/alloca-addrspace-elems.ll"},
		{path: "../testdata/llvm/test/Assembler/alloca-addrspace0.ll"},
		{path: "../testdata/llvm/test/Assembler/alloca-invalid-type-2.ll"},
		{path: "../testdata/llvm/test/Assembler/alloca-size-one.ll"},
		{path: "../testdata/llvm/test/Assembler/anon-functions.ll"},
		{path: "../testdata/llvm/test/Assembler/atomic.ll"},
		{path: "../testdata/llvm/test/Assembler/attribute-builtin.ll"},
		{path: "../testdata/llvm/test/Assembler/auto_upgrade_intrinsics.ll"},
		{path: "../testdata/llvm/test/Assembler/auto_upgrade_nvvm_intrinsics.ll"},
		{path: "../testdata/llvm/test/Assembler/autoupgrade-thread-pointer.ll"},
		{path: "../testdata/llvm/test/Assembler/bcwrap.ll"},
		{path: "../testdata/llvm/test/Assembler/call-nonzero-program-addrspace-2.ll"},
		{path: "../testdata/llvm/test/Assembler/call-nonzero-program-addrspace.ll"},
		{path: "../testdata/llvm/test/Assembler/comment.ll"},
		{path: "../testdata/llvm/test/Assembler/ConstantExprFoldCast.ll"},
		{path: "../testdata/llvm/test/Assembler/ConstantExprFoldSelect.ll"},
		{path: "../testdata/llvm/test/Assembler/ConstantExprNoFold.ll"},
		{path: "../testdata/llvm/test/Assembler/datalayout-alloca-addrspace-mismatch-0.ll"},
		{path: "../testdata/llvm/test/Assembler/datalayout-alloca-addrspace-mismatch-1.ll"},
		{path: "../testdata/llvm/test/Assembler/datalayout-alloca-addrspace-mismatch-2.ll"},
		{path: "../testdata/llvm/test/Assembler/datalayout-alloca-addrspace.ll"},
		{path: "../testdata/llvm/test/Assembler/datalayout-program-addrspace.ll"},
		{path: "../testdata/llvm/test/Assembler/debug-info.ll"},
		{path: "../testdata/llvm/test/Assembler/debug-label-bitcode.ll"},
		{path: "../testdata/llvm/test/Assembler/dicompileunit.ll"},
		{path: "../testdata/llvm/test/Assembler/dicompositetype-members.ll"},
		{path: "../testdata/llvm/test/Assembler/DIEnumerator.ll"},
		{path: "../testdata/llvm/test/Assembler/diexpression.ll"},
		{path: "../testdata/llvm/test/Assembler/difile-escaped-chars.ll"},
		{path: "../testdata/llvm/test/Assembler/diglobalvariable.ll"},
		{path: "../testdata/llvm/test/Assembler/DIGlobalVariableExpression.ll"},
		{path: "../testdata/llvm/test/Assembler/diimportedentity.ll"},
		{path: "../testdata/llvm/test/Assembler/dilexicalblock.ll"},
		{path: "../testdata/llvm/test/Assembler/dilocalvariable-arg-large.ll"},
		{path: "../testdata/llvm/test/Assembler/dilocalvariable.ll"},
		{path: "../testdata/llvm/test/Assembler/dilocation.ll"},
		{path: "../testdata/llvm/test/Assembler/DIMacroFile.ll"},
		{path: "../testdata/llvm/test/Assembler/dimodule.ll"},
		{path: "../testdata/llvm/test/Assembler/dinamespace.ll"},
		{path: "../testdata/llvm/test/Assembler/diobjcproperty.ll"},
		{path: "../testdata/llvm/test/Assembler/distinct-mdnode.ll"},
		{path: "../testdata/llvm/test/Assembler/disubprogram.ll"},
		{path: "../testdata/llvm/test/Assembler/disubrange-empty-array.ll"},
		{path: "../testdata/llvm/test/Assembler/disubroutinetype.ll"},
		{path: "../testdata/llvm/test/Assembler/ditemplateparameter.ll"},
		{path: "../testdata/llvm/test/Assembler/dllimport-dsolocal-diag.ll"},
		{path: "../testdata/llvm/test/Assembler/drop-debug-info-nonzero-alloca.ll"},
		{path: "../testdata/llvm/test/Assembler/drop-debug-info.ll"},
		{path: "../testdata/llvm/test/Assembler/externally-initialized.ll"},
		{path: "../testdata/llvm/test/Assembler/extractvalue-invalid-idx.ll"},
		{path: "../testdata/llvm/test/Assembler/fast-math-flags.ll"},
		{path: "../testdata/llvm/test/Assembler/flags.ll"},
		{path: "../testdata/llvm/test/Assembler/generic-debug-node.ll"},
		//{path: "../testdata/llvm/test/Assembler/getelementptr_invalid_ptr.ll"}, // TODO: investigate why we are able to parse `return i32* %gep`, should not be possible as there is no `return` token in the grammar.
		{path: "../testdata/llvm/test/Assembler/getelementptr_struct.ll"},
		//{path: "../testdata/llvm/test/Assembler/getelementptr_vec_ce.ll"}, // TODO: fix computation of gep type. We currently do not update the return type to vector if the first gep index is a scalar, but the second (or later) indices is a vector. As such, we currently compute the type `i32*` where `<4 x i32*>` should have been computed.
		//{path: "../testdata/llvm/test/Assembler/getelementptr_vec_ce2.ll"}, // TODO: fix computation of gep type. We currently do not update the return type to vector if the first gep index is a scalar, but the second (or later) indices is a vector. As such, we currently compute the type `i32*` where `<4 x i32*>` should have been computed.
		//{path: "../testdata/llvm/test/Assembler/getelementptr_vec_idx1.ll"}, // TODO: fix gep type computation.
		//{path: "../testdata/llvm/test/Assembler/getelementptr_vec_idx2.ll"}, // TODO: fix gep type computation.
		//{path: "../testdata/llvm/test/Assembler/getelementptr_vec_idx3.ll"}, // TODO: fix gep type computation.
		{path: "../testdata/llvm/test/Assembler/getelementptr_vec_idx4.ll"},
		{path: "../testdata/llvm/test/Assembler/getInt.ll"},
		{path: "../testdata/llvm/test/Assembler/global-addrspace-forwardref.ll"},
		{path: "../testdata/llvm/test/Assembler/globalvariable-attributes.ll"},
		{path: "../testdata/llvm/test/Assembler/gv-invalid-type.ll"},
		{path: "../testdata/llvm/test/Assembler/half-constprop.ll"},
		{path: "../testdata/llvm/test/Assembler/half-conv.ll"},
		{path: "../testdata/llvm/test/Assembler/half.ll"},
		{path: "../testdata/llvm/test/Assembler/huge-array.ll"},
		{path: "../testdata/llvm/test/Assembler/ifunc-asm.ll"},
		{path: "../testdata/llvm/test/Assembler/ifunc-dsolocal.ll"},
		{path: "../testdata/llvm/test/Assembler/ifunc-use-list-order.ll"},
		{path: "../testdata/llvm/test/Assembler/inalloca.ll"},
		{path: "../testdata/llvm/test/Assembler/incorrect-tdep-attrs-parsing.ll"},
		{path: "../testdata/llvm/test/Assembler/inline-asm-clobber.ll"},
		{path: "../testdata/llvm/test/Assembler/insertextractvalue.ll"},
		{path: "../testdata/llvm/test/Assembler/insertvalue-invalid-idx.ll"},
		//{path: "../testdata/llvm/test/Assembler/insertvalue-invalid-type-1.ll"}, // TODO: check result type of insertvalue during translation.
		{path: "../testdata/llvm/test/Assembler/insertvalue-invalid-type.ll"},
		{path: "../testdata/llvm/test/Assembler/internal-hidden-alias.ll"},
		{path: "../testdata/llvm/test/Assembler/internal-hidden-function.ll"},
		{path: "../testdata/llvm/test/Assembler/internal-hidden-variable.ll"},
		{path: "../testdata/llvm/test/Assembler/internal-protected-alias.ll"},
		{path: "../testdata/llvm/test/Assembler/internal-protected-function.ll"},
		{path: "../testdata/llvm/test/Assembler/internal-protected-variable.ll"},
		{path: "../testdata/llvm/test/Assembler/invoke-nonzero-program-addrspace.ll"},
		{path: "../testdata/llvm/test/Assembler/large-comdat.ll"},
		{path: "../testdata/llvm/test/Assembler/local-unnamed-addr.ll"},
		{path: "../testdata/llvm/test/Assembler/max-inttype.ll"},
		{path: "../testdata/llvm/test/Assembler/metadata-decl.ll"},
		{path: "../testdata/llvm/test/Assembler/metadata-function-local.ll"},
		{path: "../testdata/llvm/test/Assembler/metadata-null-operands.ll"},
		{path: "../testdata/llvm/test/Assembler/metadata.ll"},
		{path: "../testdata/llvm/test/Assembler/MultipleReturnValueType.ll"},
		{path: "../testdata/llvm/test/Assembler/musttail-invalid-1.ll"},
		{path: "../testdata/llvm/test/Assembler/musttail-invalid-2.ll"},
		{path: "../testdata/llvm/test/Assembler/musttail.ll"},
		{path: "../testdata/llvm/test/Assembler/named-metadata.ll"},
		{path: "../testdata/llvm/test/Assembler/no-mdstring-upgrades.ll"},
		{path: "../testdata/llvm/test/Assembler/numbered-values.ll"},
		{path: "../testdata/llvm/test/Assembler/private-hidden-alias.ll"},
		{path: "../testdata/llvm/test/Assembler/private-hidden-function.ll"},
		{path: "../testdata/llvm/test/Assembler/private-hidden-variable.ll"},
		{path: "../testdata/llvm/test/Assembler/private-protected-alias.ll"},
		{path: "../testdata/llvm/test/Assembler/private-protected-function.ll"},
		{path: "../testdata/llvm/test/Assembler/private-protected-variable.ll"},
		{path: "../testdata/llvm/test/Assembler/select.ll"},
		//{path: "../testdata/llvm/test/Assembler/short-hexpair.ll"}, // TODO: add support for fp128 hex literals.
		{path: "../testdata/llvm/test/Assembler/source-filename-backslash.ll"},
		{path: "../testdata/llvm/test/Assembler/source-filename.ll"},
		//{path: "../testdata/llvm/test/Assembler/thinlto-summary.ll"}, // TODO: add support for ThinLTO module summaries.
		{path: "../testdata/llvm/test/Assembler/tls-models.ll"},
		{path: "../testdata/llvm/test/Assembler/token.ll"},
		{path: "../testdata/llvm/test/Assembler/unnamed-addr.ll"},
		{path: "../testdata/llvm/test/Assembler/unnamed-alias.ll"},
		{path: "../testdata/llvm/test/Assembler/unnamed.ll"},
		{path: "../testdata/llvm/test/Assembler/uselistorder_bb.ll"},
		{path: "../testdata/llvm/test/Assembler/uselistorder.ll"},
		{path: "../testdata/llvm/test/Assembler/vbool-cmp.ll"},
		{path: "../testdata/llvm/test/Assembler/vector-cmp.ll"},
		{path: "../testdata/llvm/test/Assembler/vector-select.ll"},
		{path: "../testdata/llvm/test/Assembler/vector-shift.ll"},
		{path: "../testdata/llvm/test/Assembler/x86mmx.ll"},

		// LLVM test/Bitcode.
		{path: "../testdata/llvm/test/Bitcode/disubrange.ll"},

		// LLVM test/DebugInfo/Generic.
		{path: "../testdata/llvm/test/DebugInfo/Generic/constant-pointers.ll"},
		{path: "../testdata/llvm/test/DebugInfo/Generic/debug-info-enum.ll"},
		{path: "../testdata/llvm/test/DebugInfo/Generic/debug-label-mi.ll"},
		{path: "../testdata/llvm/test/DebugInfo/Generic/debug-names-linkage-name.ll"},
		{path: "../testdata/llvm/test/DebugInfo/Generic/gmlt_profiling.ll"},
		{path: "../testdata/llvm/test/DebugInfo/Generic/invalid.ll"},
		{path: "../testdata/llvm/test/DebugInfo/Generic/template-recursive-void.ll"},

		// LLVM test/DebugInfo.
		{path: "../testdata/llvm/test/DebugInfo/check-debugify-preserves-analyses.ll"},
		{path: "../testdata/llvm/test/DebugInfo/cross-cu-scope.ll"},
		{path: "../testdata/llvm/test/DebugInfo/debugify-bogus-dbg-value.ll"},
		{path: "../testdata/llvm/test/DebugInfo/debugify-each.ll"},
		{path: "../testdata/llvm/test/DebugInfo/debugify-export.ll"},
		{path: "../testdata/llvm/test/DebugInfo/debugify.ll"},
		{path: "../testdata/llvm/test/DebugInfo/debugify-report-missing-locs-only.ll"},
		{path: "../testdata/llvm/test/DebugInfo/dwo.ll"},
		{path: "../testdata/llvm/test/DebugInfo/macro_link.ll"},
		{path: "../testdata/llvm/test/DebugInfo/omit-empty.ll"},
		{path: "../testdata/llvm/test/DebugInfo/pr34186.ll"},
		{path: "../testdata/llvm/test/DebugInfo/pr34672.ll"},
		{path: "../testdata/llvm/test/DebugInfo/skeletoncu.ll"},
		{path: "../testdata/llvm/test/DebugInfo/strip-DIGlobalVariable.ll"},
		{path: "../testdata/llvm/test/DebugInfo/strip-loop-metadata.ll"},
		{path: "../testdata/llvm/test/DebugInfo/strip-module-flags.ll"},
		//{path: "../testdata/llvm/test/DebugInfo/unrolled-loop-remainder.ll"}, // TODO: figure out how to handle duplicate (but distinct) AttrGroupDef

		// LLVM test/DebugInfo/X86.
		{path: "../testdata/llvm/test/DebugInfo/X86/clang-module.ll"},
		{path: "../testdata/llvm/test/DebugInfo/X86/debug-ranges-offset.ll"},
		{path: "../testdata/llvm/test/DebugInfo/X86/DIModuleContext.ll"},
		{path: "../testdata/llvm/test/DebugInfo/X86/DIModule.ll"},
		{path: "../testdata/llvm/test/DebugInfo/X86/dw_op_minus.ll"},
		//{path: "../testdata/llvm/test/DebugInfo/X86/gnu-public-names-empty.ll"}, // TODO: fix grammar. syntax error at line 20. Issue with gnuPubnames in metadata. Should the metadata contain gnuPubnames, or is this now NameTableKind or something else?
		{path: "../testdata/llvm/test/DebugInfo/X86/objc-property-void.ll"},
		{path: "../testdata/llvm/test/DebugInfo/X86/safestack-byval.ll"},
		{path: "../testdata/llvm/test/DebugInfo/X86/sdag-combine.ll"},
		{path: "../testdata/llvm/test/DebugInfo/X86/stack-value-dwarf2.ll"},

		// Coreutils.
		{path: "../testdata/coreutils/test/[.ll"},
		{path: "../testdata/coreutils/test/b2sum.ll"},
		{path: "../testdata/coreutils/test/base32.ll"},
		{path: "../testdata/coreutils/test/base64.ll"},
		{path: "../testdata/coreutils/test/basename.ll"},
		{path: "../testdata/coreutils/test/cat.ll"},
		{path: "../testdata/coreutils/test/chcon.ll"},
		{path: "../testdata/coreutils/test/chgrp.ll"},
		{path: "../testdata/coreutils/test/chmod.ll"},
		{path: "../testdata/coreutils/test/chown.ll"},
		{path: "../testdata/coreutils/test/chroot.ll"},
		{path: "../testdata/coreutils/test/cksum.ll"},
		{path: "../testdata/coreutils/test/comm.ll"},
		{path: "../testdata/coreutils/test/cp.ll"},
		{path: "../testdata/coreutils/test/csplit.ll"},
		{path: "../testdata/coreutils/test/cut.ll"},
		{path: "../testdata/coreutils/test/date.ll"},
		{path: "../testdata/coreutils/test/dd.ll"},
		{path: "../testdata/coreutils/test/df.ll"},
		{path: "../testdata/coreutils/test/dir.ll"},
		{path: "../testdata/coreutils/test/dircolors.ll"},
		{path: "../testdata/coreutils/test/dirname.ll"},
		{path: "../testdata/coreutils/test/du.ll"},
		{path: "../testdata/coreutils/test/echo.ll"},
		{path: "../testdata/coreutils/test/env.ll"},
		{path: "../testdata/coreutils/test/expand.ll"},
		{path: "../testdata/coreutils/test/expr.ll"},
		{path: "../testdata/coreutils/test/factor.ll"},
		{path: "../testdata/coreutils/test/false.ll"},
		{path: "../testdata/coreutils/test/fmt.ll"},
		{path: "../testdata/coreutils/test/fold.ll"},
		{path: "../testdata/coreutils/test/getlimits.ll"},
		{path: "../testdata/coreutils/test/ginstall.ll"},
		{path: "../testdata/coreutils/test/groups.ll"},
		{path: "../testdata/coreutils/test/head.ll"},
		{path: "../testdata/coreutils/test/hostid.ll"},
		{path: "../testdata/coreutils/test/id.ll"},
		{path: "../testdata/coreutils/test/join.ll"},
		{path: "../testdata/coreutils/test/kill.ll"},
		{path: "../testdata/coreutils/test/link.ll"},
		{path: "../testdata/coreutils/test/ln.ll"},
		{path: "../testdata/coreutils/test/logname.ll"},
		{path: "../testdata/coreutils/test/ls.ll"},
		{path: "../testdata/coreutils/test/make-prime-list.ll"},
		{path: "../testdata/coreutils/test/md5sum.ll"},
		{path: "../testdata/coreutils/test/mkdir.ll"},
		{path: "../testdata/coreutils/test/mkfifo.ll"},
		{path: "../testdata/coreutils/test/mknod.ll"},
		{path: "../testdata/coreutils/test/mktemp.ll"},
		{path: "../testdata/coreutils/test/mv.ll"},
		{path: "../testdata/coreutils/test/nice.ll"},
		{path: "../testdata/coreutils/test/nl.ll"},
		{path: "../testdata/coreutils/test/nohup.ll"},
		{path: "../testdata/coreutils/test/nproc.ll"},
		{path: "../testdata/coreutils/test/numfmt.ll"},
		{path: "../testdata/coreutils/test/od.ll"},
		{path: "../testdata/coreutils/test/paste.ll"},
		{path: "../testdata/coreutils/test/pathchk.ll"},
		{path: "../testdata/coreutils/test/pinky.ll"},
		{path: "../testdata/coreutils/test/pr.ll"},
		{path: "../testdata/coreutils/test/printenv.ll"},
		{path: "../testdata/coreutils/test/printf.ll"},
		{path: "../testdata/coreutils/test/ptx.ll"},
		{path: "../testdata/coreutils/test/pwd.ll"},
		{path: "../testdata/coreutils/test/readlink.ll"},
		{path: "../testdata/coreutils/test/realpath.ll"},
		{path: "../testdata/coreutils/test/rm.ll"},
		{path: "../testdata/coreutils/test/rmdir.ll"},
		{path: "../testdata/coreutils/test/runcon.ll"},
		{path: "../testdata/coreutils/test/seq.ll"},
		{path: "../testdata/coreutils/test/sha1sum.ll"},
		{path: "../testdata/coreutils/test/sha224sum.ll"},
		{path: "../testdata/coreutils/test/sha256sum.ll"},
		{path: "../testdata/coreutils/test/sha384sum.ll"},
		{path: "../testdata/coreutils/test/sha512sum.ll"},
		{path: "../testdata/coreutils/test/shred.ll"},
		{path: "../testdata/coreutils/test/shuf.ll"},
		{path: "../testdata/coreutils/test/sleep.ll"},
		{path: "../testdata/coreutils/test/sort.ll"},
		{path: "../testdata/coreutils/test/split.ll"},
		{path: "../testdata/coreutils/test/stat.ll"},
		{path: "../testdata/coreutils/test/stdbuf.ll"},
		{path: "../testdata/coreutils/test/stty.ll"},
		{path: "../testdata/coreutils/test/sum.ll"},
		{path: "../testdata/coreutils/test/sync.ll"},
		{path: "../testdata/coreutils/test/tac.ll"},
		{path: "../testdata/coreutils/test/tail.ll"},
		{path: "../testdata/coreutils/test/tee.ll"},
		{path: "../testdata/coreutils/test/test.ll"},
		{path: "../testdata/coreutils/test/timeout.ll"},
		{path: "../testdata/coreutils/test/touch.ll"},
		{path: "../testdata/coreutils/test/tr.ll"},
		{path: "../testdata/coreutils/test/true.ll"},
		{path: "../testdata/coreutils/test/truncate.ll"},
		{path: "../testdata/coreutils/test/tsort.ll"},
		{path: "../testdata/coreutils/test/tty.ll"},
		{path: "../testdata/coreutils/test/uname.ll"},
		{path: "../testdata/coreutils/test/unexpand.ll"},
		{path: "../testdata/coreutils/test/uniq.ll"},
		{path: "../testdata/coreutils/test/unlink.ll"},
		{path: "../testdata/coreutils/test/uptime.ll"},
		{path: "../testdata/coreutils/test/users.ll"},
		{path: "../testdata/coreutils/test/vdir.ll"},
		{path: "../testdata/coreutils/test/wc.ll"},
		{path: "../testdata/coreutils/test/who.ll"},
		{path: "../testdata/coreutils/test/whoami.ll"},
		{path: "../testdata/coreutils/test/yes.ll"},

		// SQLite.
		{path: "../testdata/sqlite/test/shell.ll"},
	}
	for _, g := range golden {
		if filepath.HasPrefix(g.path, "../testdata/") {
			if !osutil.Exists("../testdata/") {
				// Skip test cases from the llir/testdata submodule if not present.
				// Users may add this submodule using git clone --recursive.
				continue
			}
		}
		log.Printf("=== [ %s ] ===", g.path)
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
			if err := diffutil.Diff(want, got, words, filepath.Base(path)); err != nil {
				panic(err)
			}
			t.Errorf("module mismatch %q; expected `%s`, got `%s`", path, want, got)
			continue
		}
	}
}
