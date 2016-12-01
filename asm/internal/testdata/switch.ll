@a = global i32 0
define i32 @main() {
; <label>:0
	%1 = load i32, i32* @a
	switch i32 %1, label %6 [
		i32 0, label %2
		i32 1, label %3
		i32 2, label %4
		i32 3, label %5
	]
; <label>:2
	br label %7
; <label>:3
	br label %7
; <label>:4
	br label %7
; <label>:5
	br label %7
; <label>:6
	br label %7
; <label>:7
	%.0 = phi i32 [ 50, %6 ], [ 40, %5 ], [ 30, %4 ], [ 20, %3 ], [ 10, %2 ]
	ret i32 %.0
}
