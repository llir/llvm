define [3 x i32] @f() {
; <label>:0
	ret [3 x i32] insertvalue ([3 x i32] [i32 11, i32 22, i32 33], i32 extractvalue ([3 x i32] [i32 77, i32 88, i32 99], 1), 1)
}

define { i8, i32 } @g() {
; <label>:0
	ret { i8, i32 } insertvalue ({ i8, i32 } { i8 11, i32 22 }, i32 extractvalue ({ i8, i32 } { i8 77, i32 88 }, 1), 1)
}
