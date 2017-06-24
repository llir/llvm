$f70 = comdat any
$com1 = comdat exactmatch

; Plain function declaration.
declare void @f1()

; Metadata.
declare !foo !{!"bar"} !baz !{!"qux"} void @f2()

; External linkage.
declare extern_weak void @f3()
declare external void @f4()

; Visibility.
declare default void @f5()
declare hidden void @f6()
declare protected void @f7()

; DLL storage class.
declare dllimport void @f8()
declare dllexport void @f9()

; Calling conventions.
declare amdgpu_cs void @f10()
declare amdgpu_gs void @f11()
declare amdgpu_kernel void @f12()
declare amdgpu_ps void @f13()
declare amdgpu_vs void @f14()
declare anyregcc void @f15()
declare arm_aapcs_vfpcc void @f16()
declare arm_aapcscc void @f17()
declare arm_apcscc void @f18()
declare avr_intrcc void @f19()
declare avr_signalcc void @f20()
declare cc 0 void @f21()
declare ccc void @f22()
declare coldcc void @f23()
declare cxx_fast_tlscc void @f24()
declare fastcc void @f25()
declare ghccc void @f26()
declare hhvm_ccc void @f27()
declare hhvmcc void @f28()
declare intel_ocl_bicc void @f29()
declare msp430_intrcc void @f30()
declare preserve_allcc void @f31()
declare preserve_mostcc void @f32()
declare ptx_device void @f33()
declare ptx_kernel void @f34()
declare spir_func void @f35()
declare spir_kernel void @f36()
declare swiftcc void @f37()
declare webkit_jscc void @f38()
declare x86_64_sysvcc void @f39()
declare x86_64_win64cc void @f40()
declare x86_fastcallcc void @f41()
declare x86_intrcc void @f42()
declare x86_regcallcc void @f43()
declare x86_stdcallcc void @f44()
declare x86_thiscallcc void @f45()
declare x86_vectorcallcc void @f46()

; Return parameter attributes.
declare "foo" "bar"="baz" align 8 dereferenceable(11) dereferenceable_or_null(22) inreg noalias i32 @f47()
declare nonnull i32 ()* @f48()
declare signext i32 @f49()
declare zeroext i32 @f50()

; Parameters.
declare void @f51()
declare void @f52(i32)
declare void @f53(i32 %x)
declare void @f54(i32, i32)
declare void @f55(i32, i32 %y)
declare void @f56(i32 %x, i32)
declare void @f57(i32 %x, i32 %y)

; Variadic function.
declare void @f58(...)
declare void @f59(i32, ...)
declare void @f60(i32 %x, ...)
declare void @f61(i32, i32, ...)
declare void @f62(i32, i32 %y, ...)
declare void @f63(i32 %x, i32, ...)
declare void @f64(i32 %x, i32 %y, ...)

; Parameter attributes.
declare void @f65(i32 "foo" "bar"="baz" align 8 byval dereferenceable(11) dereferenceable_or_null(22) inalloca inreg nest noalias nocapture nonnull readnone readonly returned signext sret swifterror swiftself writeonly zeroext %x)

; Unnamed address.
declare void @f66() local_unnamed_addr
declare void @f67() unnamed_addr

; Function attributes.
declare void @f68() "foo" "bar"="baz" #0 #1 alignstack(8) allocsize(16) allocsize(32, 64) alwaysinline argmemonly cold convergent inaccessiblemem_or_argmemonly inaccessiblememonly inlinehint jumptable minsize naked nobuiltin noduplicate noimplicitfloat noinline nonlazybind norecurse noredzone noreturn nounwind optnone optsize readnone readonly returns_twice safestack sanitize_address sanitize_memory sanitize_thread ssp sspreq sspstrong uwtable writeonly

; Section.
declare void @f69() section "foo"

; Comdat.
declare void @f70() comdat
declare void @f71() comdat($com1)

; Align.
declare void @f72() align 8

; Garbage collection.
declare void @f73() gc "foo"

; Function prefix data.
declare void @f74() prefix i32 42

; Function prologue data.
declare void @f75() prologue i32 42

; Function personality data.
declare void @f76() personality i32 42

; Full function declaration.
declare !foo !{!"bar"} !baz !{!"qux"} external default dllimport ccc "foo" "bar"="baz" align 8 dereferenceable(11) dereferenceable_or_null(22) inreg noalias i32 @f77(i32 %x, i32 "foo" "bar"="baz" align 8 byval dereferenceable(11) dereferenceable_or_null(22) inalloca inreg nest noalias nocapture nonnull readnone readonly returned signext sret swifterror swiftself writeonly zeroext %y, ...) unnamed_addr "foo" "bar"="baz" #0 #1 alignstack(8) allocsize(16) allocsize(32, 64) alwaysinline argmemonly cold convergent inaccessiblemem_or_argmemonly inaccessiblememonly inlinehint jumptable minsize naked nobuiltin noduplicate noimplicitfloat noinline nonlazybind norecurse noredzone noreturn nounwind optnone optsize readnone readonly returns_twice safestack sanitize_address sanitize_memory sanitize_thread ssp sspreq sspstrong uwtable writeonly section "foo" comdat($com1) align 8 gc "foo" prefix i32 42 prologue i32 42 personality i32 42

; Plain function definition.
define void @f78() {
	ret void
}

; Linkage.
define available_externally void @f80() {
	ret void
}

define internal void @f82() {
	ret void
}

define linkonce void @f83() {
	ret void
}

define linkonce_odr void @f84() {
	ret void
}

define private void @f85() {
	ret void
}

define weak void @f86() {
	ret void
}

define weak_odr void @f87() {
	ret void
}

; Metadata.
define void @f88() !foo !{!"bar"} !baz !{!"qux"} {
	ret void
}

; Full function definition.
define available_externally default dllimport ccc "foo" "bar"="baz" align 8 dereferenceable(11) dereferenceable_or_null(22) inreg noalias i32 @f89(i32 %x, i32 "foo" "bar"="baz" align 8 byval dereferenceable(11) dereferenceable_or_null(22) inalloca inreg nest noalias nocapture nonnull readnone readonly returned signext sret swifterror swiftself writeonly zeroext %y, ...) unnamed_addr "foo" "bar"="baz" #0 #1 alignstack(8) allocsize(16) allocsize(32, 64) alwaysinline argmemonly cold convergent inaccessiblemem_or_argmemonly inaccessiblememonly inlinehint jumptable minsize naked nobuiltin noduplicate noimplicitfloat noinline nonlazybind norecurse noredzone noreturn nounwind optnone optsize readnone readonly returns_twice safestack sanitize_address sanitize_memory sanitize_thread ssp sspreq sspstrong uwtable writeonly section "foo" comdat($com1) align 8 gc "foo" prefix i32 42 prologue i32 42 personality i32 42 !foo !{!"bar"} !baz !{!"qux"} {
	ret i32 42
}

attributes #0 = { "foo" }
attributes #1 = { "foo" "bar"="baz" #0 alignstack=8 allocsize(16) allocsize(16, 32) alwaysinline argmemonly builtin cold convergent inaccessiblemem_or_argmemonly inaccessiblememonly inlinehint jumptable minsize naked nobuiltin noduplicate noimplicitfloat noinline nonlazybind norecurse noredzone noreturn nounwind optnone optsize readnone readonly returns_twice safestack sanitize_address sanitize_memory sanitize_thread ssp sspreq sspstrong uwtable writeonly }
