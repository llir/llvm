; --- [ Other instructions ] ---------------------------------------------------

@g1 = global i32 42

; ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i1 @icmp_1() {
	; Plain instruction.
	%result = icmp ne i32 42, 5
	ret i1 %result
}

define <2 x i1> @icmp_2() {
	; Vector operands.
	%result = icmp eq <2 x i32> <i32 42, i32 11>, <i32 42, i32 22>
	ret <2 x i1> %result
}

define void @icmp_3() {
	; Predicates.
	%1 = icmp eq i32 10, 15
	%2 = icmp ne i32 10, 15
	%3 = icmp ugt i32 10, 15
	%4 = icmp uge i32 10, 15
	%5 = icmp ult i32 10, 15
	%6 = icmp ule i32 10, 15
	%7 = icmp sgt i32 10, 15
	%8 = icmp sge i32 10, 15
	%9 = icmp slt i32 10, 15
	%10 = icmp sle i32 10, 15
	ret void
}

define i1 @icmp_4() {
	; Metadata.
	%result = icmp ne i32 42, 5, !foo !{!"bar"}, !baz !{!"qux"}
	ret i1 %result
}

; ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i1 @fcmp_1() {
	; Plain instruction.
	%result = fcmp one double 42.0, 5.0
	ret i1 %result
}

define <2 x i1> @fcmp_2() {
	; Vector operands.
	%result = fcmp oeq <2 x double> <double 42.0, double 11.0>, <double 42.0, double 22.0>
	ret <2 x i1> %result
}

define void @fcmp_3() {
	; Predicates.
	%1 = fcmp false double 10.0, 15.0
	%2 = fcmp oeq double 10.0, 15.0
	%3 = fcmp ogt double 10.0, 15.0
	%4 = fcmp oge double 10.0, 15.0
	%5 = fcmp olt double 10.0, 15.0
	%6 = fcmp ole double 10.0, 15.0
	%7 = fcmp one double 10.0, 15.0
	%8 = fcmp ord double 10.0, 15.0
	%9 = fcmp ueq double 10.0, 15.0
	%10 = fcmp ugt double 10.0, 15.0
	%11 = fcmp uge double 10.0, 15.0
	%12 = fcmp ult double 10.0, 15.0
	%13 = fcmp ule double 10.0, 15.0
	%14 = fcmp une double 10.0, 15.0
	%15 = fcmp uno double 10.0, 15.0
	%16 = fcmp true double 10.0, 15.0
	ret void
}

define i1 @fcmp_4() {
	; Fast-math flags.
	%result = fcmp arcp fast ninf nnan nsz one double 42.0, 5.0
	ret i1 %result
}

define i1 @fcmp_5() {
	; Metadata.
	%result = fcmp one double 42.0, 5.0, !foo !{!"bar"}, !baz !{!"qux"}
	ret i1 %result
}

define i1 @fcmp_6() {
	; Full instruction.
	%result = fcmp arcp fast ninf nnan nsz one double 42.0, 5.0, !foo !{!"bar"}, !baz !{!"qux"}
	ret i1 %result
}

; ~~~ [ phi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @phi_1(i1 %cond) {
entry:
	; Plain instruction.
	%result = phi i32 [ 42, %entry ]
	ret i32 %result
}

define i32 @phi_2(i1 %cond) {
	br i1 %cond, label %foo, label %bar
foo:
	br label %baz
bar:
	br label %baz
baz:
	; Multiple incoming branches.
	%result = phi i32 [ 42, %foo ], [ 37, %bar ]
	ret i32 %result
}

define void @phi_3(i1 %cond) {
	br i1 %cond, label %foo, label %bar
foo:
	%x = fadd double 32.0, 10.0
	br label %baz
bar:
	br label %baz
baz:
	; Incoming values of various types.
	%1 = phi i32 [ 42, %foo ], [ 37, %bar ]
	%2 = phi i32* [ null, %foo ], [ @g1, %bar ]
	%3 = phi double [ %x, %foo ], [ 11.0, %bar ]
	%4 = phi { i32 } [ { i32 42 }, %foo ], [ zeroinitializer, %bar ]
	%5 = phi void ()* [ undef, %foo ], [ @j, %bar ]
	ret void
}

define i32 @phi_4(i1 %cond) {
	br i1 %cond, label %foo, label %bar
foo:
	br label %baz
bar:
	br label %baz
baz:
	; Metadata.
	%result = phi i32 [ 42, %foo ], [ 37, %bar ], !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @select_1(i1 %cond) {
	; Plain instruction.
	%result = select i1 %cond, i32 42, i32 37
	ret i32 %result
}

define <2 x i32> @select_2(i1 %cond) {
	; Vector operands.
	%result = select i1 %cond, <2 x i32> <i32 42, i32 37>, <2 x i32> <i32 11, i32 22>
	ret <2 x i32> %result
}

define i32 @select_3(i1 %cond) {
	; Metadata.
	%result = select i1 %cond, i32 42, i32 37, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ call ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @f() {
	ret i32 42
}

define i32 @call_1() {
	; Plain instruction.
	%result = call i32 @f()
	ret i32 %result
}

define i32 @call_2() {
	; Tail.
	%result = tail call i32 @f()
	ret i32 %result
}

define i32 @call_3() {
	; Tail.
	%result = musttail call i32 @f()
	ret i32 %result
}

define i32 @call_4() {
	; Tail.
	%result = notail call i32 @f()
	ret i32 %result
}

define double @g() {
	ret double 42.0
}

define double @call_5() {
	; Fast-math flags.
	%result = call arcp fast ninf nnan nsz double @g()
	ret double %result
}

define void @call_6() {
	; Calling convention.
	%1 = call amdgpu_cs i32 @f()
	%2 = call amdgpu_gs i32 @f()
	%3 = call amdgpu_kernel i32 @f()
	%4 = call amdgpu_ps i32 @f()
	%5 = call amdgpu_vs i32 @f()
	%6 = call anyregcc i32 @f()
	%7 = call arm_aapcs_vfpcc i32 @f()
	%8 = call arm_aapcscc i32 @f()
	%9 = call arm_apcscc i32 @f()
	%10 = call avr_intrcc i32 @f()
	%11 = call avr_signalcc i32 @f()
	%12 = call cc 0 i32 @f()
	%13 = call cc 8 i32 @f()
	%14 = call cc 9 i32 @f()
	%15 = call cc 10 i32 @f()
	%16 = call cc 11 i32 @f()
	%17 = call cc 12 i32 @f()
	%18 = call cc 13 i32 @f()
	%19 = call cc 14 i32 @f()
	%20 = call cc 15 i32 @f()
	%21 = call cc 16 i32 @f()
	%22 = call cc 17 i32 @f()
	%23 = call cc 64 i32 @f()
	%24 = call cc 65 i32 @f()
	%25 = call cc 66 i32 @f()
	%26 = call cc 67 i32 @f()
	%27 = call cc 68 i32 @f()
	%28 = call cc 69 i32 @f()
	%29 = call cc 70 i32 @f()
	%30 = call cc 71 i32 @f()
	%31 = call cc 72 i32 @f()
	%32 = call cc 75 i32 @f()
	%33 = call cc 76 i32 @f()
	%34 = call cc 77 i32 @f()
	%35 = call cc 78 i32 @f()
	%36 = call cc 79 i32 @f()
	%37 = call cc 80 i32 @f()
	%38 = call cc 81 i32 @f()
	%39 = call cc 82 i32 @f()
	%40 = call cc 83 i32 @f()
	%41 = call cc 84 i32 @f()
	%42 = call cc 85 i32 @f()
	%43 = call cc 86 i32 @f()
	%44 = call cc 87 i32 @f()
	%45 = call cc 88 i32 @f()
	%46 = call cc 89 i32 @f()
	%47 = call cc 90 i32 @f()
	%48 = call cc 91 i32 @f()
	%49 = call cc 92 i32 @f()
	%50 = call ccc i32 @f()
	%51 = call coldcc i32 @f()
	%52 = call cxx_fast_tlscc i32 @f()
	%53 = call fastcc i32 @f()
	%54 = call ghccc i32 @f()
	%55 = call hhvm_ccc i32 @f()
	%56 = call hhvmcc i32 @f()
	%57 = call intel_ocl_bicc i32 @f()
	%58 = call msp430_intrcc i32 @f()
	%59 = call preserve_allcc i32 @f()
	%60 = call preserve_mostcc i32 @f()
	%61 = call ptx_device i32 @f()
	%62 = call ptx_kernel i32 @f()
	%63 = call spir_func i32 @f()
	%64 = call spir_kernel i32 @f()
	%65 = call swiftcc i32 @f()
	%66 = call webkit_jscc i32 @f()
	%67 = call x86_64_sysvcc i32 @f()
	%68 = call x86_64_win64cc i32 @f()
	%69 = call x86_fastcallcc i32 @f()
	%70 = call x86_intrcc i32 @f()
	%71 = call x86_regcallcc i32 @f()
	%72 = call x86_stdcallcc i32 @f()
	%73 = call x86_thiscallcc i32 @f()
	%74 = call x86_vectorcallcc i32 @f()
	ret void
}


define i32 ()* @h() {
	ret i32 ()* @f
}

define void @call_7() {
	; Return parameter attributes.
	%1 = call "foo" "bar"="baz" align 8 dereferenceable(11) dereferenceable_or_null(22) inreg noalias i32 @f()
	%2 = call nonnull i32 ()* @h()
	%3 = call signext i32 @f()
	%4 = call zeroext i32 @f()
	ret void
}

define i32 @i(i32 %x, ...) {
	ret i32 42
}

define i32 @call_8() {
	; Callee with variadic arguments.
	%result = call i32 (i32, ...) @i(i32 10, i32 20)
	ret i32 %result
}

define i32 @call_9() {
	; Callee from global identifier.
	%result = call i32 @f()
	ret i32 %result
}

define i32 @call_10(i32 ()* %callee) {
	; Callee from function parameter.
	%result = call i32 %callee()
	ret i32 %result
}

declare void @j()

define i32 @call_11() {
	; Callee from bitcast expression.
	%result = call i32 bitcast (void ()* @j to i32 (i32)*)(i32 42)
	ret i32 %result
}

define i32 @call_12() {
	%callee = bitcast void ()* @j to i32 (i32)*
	; Callee from bitcast instruction.
	%result = call i32 %callee(i32 42)
	ret i32 %result
}

define i32 @call_13(i32 ()** %p) {
	%callee = load i32 ()*, i32 ()** %p
	; Callee from load instruction.
	%result = call i32 %callee()
	ret i32 %result
}

define i32 @k(i32 %x, i32 %y) {
	ret i32 42
}

define i32 @call_14() {
	; Multiple arguments.
	%result = call i32 @k(i32 11, i32 22)
	ret i32 %result
}

define i32 @call_15() {
	; Function attributes.
	%result = call i32 @f() "foo" "bar"="baz" #0 alignstack(8) allocsize(8) allocsize(8, 16) alwaysinline argmemonly builtin cold convergent inaccessiblemem_or_argmemonly inaccessiblememonly inlinehint jumptable minsize naked nobuiltin noduplicate noimplicitfloat noinline nonlazybind norecurse noredzone noreturn nounwind optnone optsize readnone readonly returns_twice safestack sanitize_address sanitize_memory sanitize_thread ssp sspreq sspstrong uwtable writeonly
	ret i32 %result
}

define void @l() {
	ret void
}

define void @call_16() {
	; Callee with void return type.
	call void @l()
	ret void
}

define i32 @call_17() {
	; Metadata.
	%result = call i32 @f(), !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

define double @m(double %x, double %y) {
	ret double 42.0
}

define double @call_18() {
	; Full instruction.
	%result = tail call arcp fast ninf nnan nsz ccc "foo" "bar"="baz" align 8 dereferenceable(11) dereferenceable_or_null(22) inreg noalias double @m(double 11.0, double 22.0) "foo" "bar"="baz" #0 alignstack(8) allocsize(8) allocsize(8, 16) alwaysinline argmemonly builtin cold convergent inaccessiblemem_or_argmemonly inaccessiblememonly inlinehint jumptable minsize naked nobuiltin noduplicate noimplicitfloat noinline nonlazybind norecurse noredzone noreturn nounwind optnone optsize readnone readonly returns_twice safestack sanitize_address sanitize_memory sanitize_thread ssp sspreq sspstrong uwtable writeonly, !foo !{!"bar"}, !baz !{!"qux"}
	ret double %result
}

; ~~~ [ va_arg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

; TODO: add test cases for va_arg instruction.

; ~~~ [ landingpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

; TODO: add test cases for landingpad instruction.

; ~~~ [ catchpad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

; TODO: add test cases for catchpad instruction.

; ~~~ [ cleanuppad ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

; TODO: add test cases for cleanuppad instruction.

attributes #0 = { "qux" }
