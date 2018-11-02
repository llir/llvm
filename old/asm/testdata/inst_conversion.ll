; --- [ Conversion instructions ] ----------------------------------------------

; ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i8 @trunc_1() {
	; Plain instruction.
	%result = trunc i32 298 to i8
	ret i8 %result
}

define i8 @trunc_2() {
	; Metadata.
	%result = trunc i32 298 to i8, !foo !{!"bar"}, !baz !{!"qux"}
	ret i8 %result
}

; ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @zext_1() {
	; Plain instruction.
	%result = zext i8 42 to i32
	ret i32 %result
}

define i32 @zext_2() {
	; Metadata.
	%result = zext i8 42 to i32, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @sext_1() {
	; Plain instruction.
	%result = sext i8 -42 to i32
	ret i32 %result
}

define i32 @sext_2() {
	; Metadata.
	%result = sext i8 -42 to i32, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define float @fptrunc_1() {
	; Plain instruction.
	%result = fptrunc double 42.0 to float
	ret float %result
}

define float @fptrunc_2() {
	; Metadata.
	%result = fptrunc double 42.0 to float, !foo !{!"bar"}, !baz !{!"qux"}
	ret float %result
}

; ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define double @fpext_1() {
	; Plain instruction.
	%result = fpext float 42.0 to double
	ret double %result
}

define double @fpext_2() {
	; Metadata.
	%result = fpext float 42.0 to double, !foo !{!"bar"}, !baz !{!"qux"}
	ret double %result
}

; ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @fptoui_1() {
	; Plain instruction.
	%result = fptoui double 42.0 to i32
	ret i32 %result
}

define i32 @fptoui_2() {
	; Metadata.
	%result = fptoui double 42.0 to i32, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @fptosi_1() {
	; Plain instruction.
	%result = fptosi double -42.0 to i32
	ret i32 %result
}

define i32 @fptosi_2() {
	; Metadata.
	%result = fptosi double -42.0 to i32, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define double @uitofp_1() {
	; Plain instruction.
	%result = uitofp i32 42 to double
	ret double %result
}

define double @uitofp_2() {
	; Metadata.
	%result = uitofp i32 42 to double, !foo !{!"bar"}, !baz !{!"qux"}
	ret double %result
}

; ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define double @sitofp_1() {
	; Plain instruction.
	%result = sitofp i32 -42 to double
	ret double %result
}

define double @sitofp_2() {
	; Metadata.
	%result = sitofp i32 -42 to double, !foo !{!"bar"}, !baz !{!"qux"}
	ret double %result
}

; ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i64 @ptrtoint_1() {
	; Plain instruction.
	%result = ptrtoint i8* null to i64
	ret i64 %result
}

define i64 @ptrtoint_2(i8* %x) {
	; Local operand.
	%result = ptrtoint i8* %x to i64
	ret i64 %result
}

define void @f() {
	ret void
}

define i64 @ptrtoint_3() {
	; Global operand.
	%result = ptrtoint void ()* @f to i64
	ret i64 %result
}

define i64 @ptrtoint_4() {
	; Metadata.
	%result = ptrtoint i8* null to i64, !foo !{!"bar"}, !baz !{!"qux"}
	ret i64 %result
}

; ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i8* @inttoptr_1() {
	; Plain instruction.
	%result = inttoptr i32 0 to i8*
	ret i8* %result
}

define i8* @inttoptr_2() {
	; Metadata.
	%result = inttoptr i32 0 to i8*, !foo !{!"bar"}, !baz !{!"qux"}
	ret i8* %result
}

; ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define <4 x i8> @bitcast_1() {
	; Plain instruction.
	%result = bitcast i32 707406378 to <4 x i8>
	ret <4 x i8> %result
}

define <4 x i8> @bitcast_2() {
	; Metadata.
	%result = bitcast i32 707406378 to <4 x i8>, !foo !{!"bar"}, !baz !{!"qux"}
	ret <4 x i8> %result
}

; ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i8 addrspace(2)* @addrspacecast_1() {
	; Plain instruction.
	%result = addrspacecast i8 addrspace(1)* null to i8 addrspace(2)*
	ret i8 addrspace(2)* %result
}

define i8 addrspace(2)* @addrspacecast_2() {
	; Metadata.
	%result = addrspacecast i8 addrspace(1)* null to i8 addrspace(2)*, !foo !{!"bar"}, !baz !{!"qux"}
	ret i8 addrspace(2)* %result
}
