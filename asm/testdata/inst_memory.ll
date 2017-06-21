; --- [ Memory instructions ] --------------------------------------------------

; ~~~ [ alloca ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32* @alloca_1() {
	; Plain instruction.
	%result = alloca i32
	ret i32* %result
}

define i32* @alloca_2() {
	; Number of elements operand.
	%result = alloca i32, i32 10
	ret i32* %result
}

define i32* @alloca_3() {
	; Alignment operand.
	%result = alloca i32, align 8
	ret i32* %result
}

define i32* @alloca_4() {
	; Metadata.
	%result = alloca i32, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32* %result
}

define i32* @alloca_5() {
	; Full instruction.
	%result = alloca i32, i32 10, align 8, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32* %result
}

; ~~~ [ load ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @load_1(i32* %x) {
	; Plain instruction.
	%result = load i32, i32* %x
	ret i32 %result
}

define i32 @load_2(i32* %x) {
	; Volatile.
	%result = load volatile i32, i32* %x
	ret i32 %result
}

define i32 @load_3(i32* %x) {
	; Alignment operand.
	%result = load i32, i32* %x, align 8
	ret i32 %result
}

define i32 @load_4(i32* %x) {
	; Metadata.
	%result = load i32, i32* %x, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

define i32 @load_5(i32* %x) {
	; Full instruction.
	%result = load volatile i32, i32* %x, align 8, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ store ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define void @store_1(i32* %x) {
	; Plain instruction.
	store i32 42, i32* %x
	ret void
}

define void @store_2(i32* %x) {
	; Volatile.
	store volatile i32 42, i32* %x
	ret void
}

define void @store_3(i32* %x) {
	; Alignment operand.
	store i32 42, i32* %x, align 8
	ret void
}

define void @store_4(i32* %x) {
	; Metadata.
	store i32 42, i32* %x, !foo !{!"bar"}, !baz !{!"qux"}
	ret void
}

define void @store_5(i32* %x) {
	; Full instruction.
	store volatile i32 42, i32* %x, align 8, !foo !{!"bar"}, !baz !{!"qux"}
	ret void
}

; ~~~ [ fence ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

; TODO: add test cases for fence instruction.

; ~~~ [ cmpxchg ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

; TODO: add test cases for cmpxchg instruction.

; ~~~ [ atomicrmw ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

; TODO: add test cases for atomicrmw instruction.

; ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32* @getelementptr_1(i32* %x) {
	; Plain instruction.
	%result = getelementptr i32, i32* %x
	ret i32* %result
}

define i32* @getelementptr_2({ i32, { [2 x i32], i8 } }* %x) {
	; Indices.
	%result = getelementptr { i32, { [2 x i32], i8 } }, { i32, { [2 x i32], i8 } }* %x, i32 0, i32 1, i32 0, i32 1
	ret i32* %result
}

define i32* @getelementptr_3(i32* %x) {
	; Inbounds.
	%result = getelementptr inbounds i32, i32* %x
	ret i32* %result
}

define i32* @getelementptr_4(i32* %x) {
	; Metadata.
	%result = getelementptr i32, i32* %x, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32* %result
}

define i32* @getelementptr_5({ i32, { [2 x i32], i8 } }* %x) {
	; Full instruction.
	%result = getelementptr inbounds { i32, { [2 x i32], i8 } }, { i32, { [2 x i32], i8 } }* %x, i32 0, i32 1, i32 0, i32 1, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32* %result
}
