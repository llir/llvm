; --- [ Conversion expressions ] -----------------------------------------------

@x = constant i8 0

; ~~~ [ trunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i8 @trunc_1() {
	; Plain expression.
	ret i8 trunc (i32 298 to i8)
}

; ~~~ [ zext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @zext_1() {
	; Plain expression.
	ret i32 zext (i8 42 to i32)
}

; ~~~ [ sext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @sext_1() {
	; Plain expression.
	ret i32 sext (i8 -42 to i32)
}

; ~~~ [ fptrunc ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define float @fptrunc_1() {
	; Plain expression.
	ret float fptrunc (double 42.0 to float)
}

; ~~~ [ fpext ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define double @fpext_1() {
	; Plain expression.
	ret double fpext (float 42.0 to double)
}

; ~~~ [ fptoui ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @fptoui_1() {
	; Plain expression.
	ret i32 fptoui (double 42.0 to i32)
}

; ~~~ [ fptosi ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @fptosi_1() {
	; Plain expression.
	ret i32 fptosi (double -42.0 to i32)
}

; ~~~ [ uitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define double @uitofp_1() {
	; Plain expression.
	ret double uitofp (i32 42 to double)
}

; ~~~ [ sitofp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define double @sitofp_1() {
	; Plain expression.
	ret double sitofp (i32 -42 to double)
}

; ~~~ [ ptrtoint ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i64 @ptrtoint_1() {
	; Plain expression.
	ret i64 ptrtoint (i8* null to i64)
}

define i64 @ptrtoint_2() {
	; Local operand.
	ret i64 ptrtoint (i8* @x to i64)
}

define void @f() {
	ret void
}

define i64 @ptrtoint_3() {
	; Global operand.
	ret i64 ptrtoint (void ()* @f to i64)
}

; ~~~ [ inttoptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i8* @inttoptr_1() {
	; Plain expression.
	ret i8* inttoptr (i32 0 to i8*)
}

; ~~~ [ bitcast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define <4 x i8> @bitcast_1() {
	; Plain expression.
	ret <4 x i8> bitcast (i32 707406378 to <4 x i8>)
}

; ~~~ [ addrspacecast ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i8 addrspace(2)* @addrspacecast_1() {
	; Plain expression.
	ret i8 addrspace(2)* addrspacecast (i8 addrspace(1)* null to i8 addrspace(2)*)
}
