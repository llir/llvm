; --- [ Aggregate expressions ] ------------------------------------------------

; ~~~ [ extractvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @extractvalue_1() {
	; Plain expression.
	ret i32 extractvalue ({ i8, i32 } { i8 21, i32 42 }, 1)
}

define i32 @extractvalue_2() {
	; Nested struct and array operand.
	ret i32 extractvalue ({ i32, { [2 x i32], i8 } } { i32 0, { [2 x i32], i8 } { [2 x i32] [i32 100, i32 42], i8 11 } }, 1, 0, 1)
}

; ~~~ [ insertvalue ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define { i32, i32 } @insertvalue_1() {
	; Plain expression.
	ret { i32, i32 } insertvalue ({ i32, i32 } { i32 21, i32 42 }, i32 42, 0)
}

define { i32, { [2 x i32], i32 } } @insertvalue_2() {
	; Nested struct and array operand.
	ret { i32, { [2 x i32], i32 } } insertvalue ({ i32, { [2 x i32], i32 } } { i32 42, { [2 x i32], i32 } { [2 x i32] [i32 100, i32 42], i32 42 } }, i32 42, 1, 0, 0)
}
