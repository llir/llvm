define i32 @main() {
; <label>:0
	%1 = shl i32 5, 3
	%2 = lshr i32 5, 3
	%3 = ashr i32 5, 3
	%4 = and i32 5, 3
	%5 = or i32 5, 3
	%6 = xor i32 5, 3
	ret i32 0
}
