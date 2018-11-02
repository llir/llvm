define i32 @main() {
; <label>:0
	%1 = fcmp oeq float 5.0, 3.0
	br i1 %1, label %2, label %3
; <label>:2
	br label %3
; <label>:3
	%4 = fcmp ogt float 5.0, 3.0
	br i1 %4, label %5, label %6
; <label>:5
	br label %6
; <label>:6
	%7 = fcmp oge float 5.0, 3.0
	br i1 %7, label %8, label %9
; <label>:8
	br label %9
; <label>:9
	%10 = fcmp olt float 5.0, 3.0
	br i1 %10, label %11, label %12
; <label>:11
	br label %12
; <label>:12
	%13 = fcmp ole float 5.0, 3.0
	br i1 %13, label %14, label %15
; <label>:14
	br label %15
; <label>:15
	%16 = fcmp une float 5.0, 3.0
	br i1 %16, label %17, label %18
; <label>:17
	br label %18
; <label>:18
	ret i32 0
}
