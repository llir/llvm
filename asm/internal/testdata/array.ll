@x = global [4 x i32] [i32 0, i32 1, i32 2, i32 3]

define i32 @f() {
; <label>:0
	%1 = load i32, i32* getelementptr ([4 x i32], [4 x i32]* @x, i64 0, i64 0)
	ret i32 %1
}
