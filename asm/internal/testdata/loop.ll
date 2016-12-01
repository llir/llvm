define i32 @main() {
; <label>:0
	br label %1
; <label>:1
	%sum.0 = phi i32 [ 0, %0 ], [ %4, %5 ]
	%i.0 = phi i32 [ 0, %0 ], [ %6, %5 ]
	%2 = icmp slt i32 %i.0, 10
	br i1 %2, label %3, label %7
; <label>:3
	%4 = add i32 %sum.0, %i.0
	br label %5
; <label>:5
	%6 = add i32 %i.0, 1
	br label %1
; <label>:7
	%8 = srem i32 %sum.0, 256
	ret i32 %8
}
