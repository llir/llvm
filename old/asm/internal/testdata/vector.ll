define void @f() {
; <label>:0
	%tmp1 = alloca <5 x i32>
	%x = load <5 x i32>, <5 x i32>* %tmp1
	%tmp2 = alloca <5 x i32>
	%y = load <5 x i32>, <5 x i32>* %tmp2
	%1 = extractelement <5 x i32> %x, i32 1
	%2 = insertelement <5 x i32> %x, i32 7, i32 1
	%3 = shufflevector <5 x i32> %x, <5 x i32> %y, <5 x i32> <i32 4, i32 3, i32 2, i32 1, i32 0>
	ret void
}
