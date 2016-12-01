define i32 @main() {
; <label>:0
	%1 = bitcast i8 255 to i8
	%2 = alloca i32
	%3 = bitcast i32* %2 to i32*
	%4 = alloca <2 x i32>
	%5 = load <2 x i32>, <2 x i32>* %4
	%6 = bitcast <2 x i32> %5 to i64
	%7 = alloca <2 x i32*>
	%8 = load <2 x i32*>, <2 x i32*>* %7
	%9 = bitcast <2 x i32*> %8 to <2 x i64*>
	ret i32 0
}
