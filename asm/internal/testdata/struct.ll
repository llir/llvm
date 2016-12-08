%struct.foo = type { i32, i32 }
define i32 @main() {
; <label>:0
	%1 = alloca %struct.foo
	%2 = getelementptr %struct.foo, %struct.foo* %1, i32 0, i32 0
	store i32 1, i32* %2
	%3 = getelementptr %struct.foo, %struct.foo* %1, i32 0, i32 1
	store i32 2, i32* %3
	ret i32 42
}
