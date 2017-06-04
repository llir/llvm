@x = external global i32

define i32 @f() {
; <label>:0
	%1 = load i32, i32* @x
	ret i32 %1
}
