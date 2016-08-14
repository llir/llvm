; Module crt implements the C run-time library functions.

; abs returns the absolute value of x.
define i32 @abs(i32 %x) {
	%t1 = icmp slt i32 %x, 0
	br i1 %t1, label %true_branch, label %false_branch
true_branch:
	%t2 = sub i32 0, %x
	ret i32 %t2
false_branch:
	ret i32 %x
}
