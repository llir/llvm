define i32 @main() {
; <label>:0
	%1 = add i32 5, 3
	%2 = sub i32 5, 3
	%3 = mul i32 5, 3
	%4 = sdiv i32 5, 3
	%5 = srem i32 5, 3
	%6 = udiv i32 5, 3
	%7 = urem i32 5, 3
	%8 = fadd float 5.0, 3.0
	%9 = fsub float 5.0, 3.0
	%10 = fmul float 5.0, 3.0
	%11 = fdiv float 5.0, 3.0
	ret i32 0
}
