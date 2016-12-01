define i32 @main() {
; <label>:0
	%1 = bitcast i8 255 to i8
	%2 = alloca i32
	%3 = bitcast i32* %2 to i32*
	ret i32 0
}
