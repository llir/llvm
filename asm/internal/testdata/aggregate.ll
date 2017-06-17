define void @f() {
; <label>:0
	%tmp1 = alloca [5 x i32]
	%x = load [5 x i32], [5 x i32]* %tmp1
	%tmp2 = alloca [5 x i32]
	%y = load [5 x i32], [5 x i32]* %tmp2
	%1 = extractvalue [5 x i32] %x, 1
	%2 = insertvalue [5 x i32] %x, i32 7, 1
	ret void
}

define void @g() {
; <label>:0
	%tmp1 = alloca { i8, i32 }
	%x = load { i8, i32 }, { i8, i32 }* %tmp1
	%tmp2 = alloca { i8, i32 }
	%y = load { i8, i32 }, { i8, i32 }* %tmp2
	%1 = extractvalue { i8, i32 } %x, 1
	%2 = insertvalue { i8, i32 } %x, i32 7, 1
	ret void
}
