; --- [ Aggregate instructions ] -----------------------------------------------

; ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @extractvalue_1() {
	; Plain instruction.
	%result = extractvalue { i8, i32 } { i8 21, i32 42 }, 1
	ret i32 %result
}

define i32 @extractvalue_2() {
	; Nested struct and array operand.
	%result = extractvalue { i32, { [2 x i32], i8 } } { i32 0, { [2 x i32], i8 } { [2 x i32] [i32 100, i32 42], i8 11 } }, 1, 0, 1
	ret i32 %result
}

define i32 @extractvalue_3() {
	; Metadata.
	%result = extractvalue { i8, i32 } { i8 21, i32 42 }, 1, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define { i32, i32 } @insertvalue_1() {
	; Plain instruction.
	%result = insertvalue { i32, i32 } { i32 21, i32 42 }, i32 42, 0
	ret { i32, i32 } %result
}

define { i32, { [2 x i32], i32 } } @insertvalue_2() {
	; Nested struct and array operand.
	%result = insertvalue { i32, { [2 x i32], i32 } } { i32 42, { [2 x i32], i32 } { [2 x i32] [i32 100, i32 42], i32 42 } }, i32 42, 1, 0, 0
	ret { i32, { [2 x i32], i32 } } %result
}

define { i32, i32 } @insertvalue_3() {
	; Metadata.
	%result = insertvalue { i32, i32 } { i32 21, i32 42 }, i32 42, 0, !foo !{!"bar"}, !baz !{!"qux"}
	ret { i32, i32 } %result
}
