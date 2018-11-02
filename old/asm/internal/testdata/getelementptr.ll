define i32 @main() {
; <label>:0
	%1 = alloca [4 x i32]
	%2 = getelementptr [4 x i32], [4 x i32]* %1, i64 0, i64 0
	store i32 0, i32* %2
	%3 = getelementptr [4 x i32], [4 x i32]* %1, i64 0, i64 1
	store i32 1, i32* %3
	%4 = getelementptr [4 x i32], [4 x i32]* %1, i64 0, i64 2
	store i32 2, i32* %4
	%5 = getelementptr [4 x i32], [4 x i32]* %1, i64 0, i64 3
	store i32 3, i32* %5
	ret i32 0
}
