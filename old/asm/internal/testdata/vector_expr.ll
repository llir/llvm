define <6 x i32> @f() {
; <label>:0
	ret <6 x i32> shufflevector (<3 x i32> insertelement (<3 x i32> <i32 11, i32 22, i32 33>, i32 extractelement (<3 x i32> <i32 77, i32 88, i32 99>, i32 1), i32 1), <3 x i32> <i32 44, i32 55, i32 66>, <6 x i32> <i32 2, i32 1, i32 0, i32 5, i32 4, i32 3>)
}
