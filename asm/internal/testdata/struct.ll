%struct.foo = type { %struct.bar, i32 }
%struct.bar = type { i32, i32 }
%struct.qux = type { %struct.anon, i32 }
%struct.anon = type { i32, i32 }
define i32 @main() {
; <label>:0
	%1 = alloca %struct.foo
	%2 = alloca %struct.qux
	%3 = getelementptr %struct.foo, %struct.foo* %1, i32 0, i32 0
	%4 = getelementptr %struct.bar, %struct.bar* %3, i32 0, i32 0
	store i32 1, i32* %4
	%5 = getelementptr %struct.foo, %struct.foo* %1, i32 0, i32 0
	%6 = getelementptr %struct.bar, %struct.bar* %5, i32 0, i32 1
	store i32 2, i32* %6
	%7 = getelementptr %struct.foo, %struct.foo* %1, i32 0, i32 1
	store i32 3, i32* %7
	%8 = getelementptr %struct.qux, %struct.qux* %2, i32 0, i32 0
	%9 = getelementptr %struct.anon, %struct.anon* %8, i32 0, i32 0
	store i32 4, i32* %9
	%10 = getelementptr %struct.qux, %struct.qux* %2, i32 0, i32 0
	%11 = getelementptr %struct.anon, %struct.anon* %10, i32 0, i32 1
	store i32 5, i32* %11
	%12 = getelementptr %struct.qux, %struct.qux* %2, i32 0, i32 1
	store i32 6, i32* %12
	ret i32 42
}
