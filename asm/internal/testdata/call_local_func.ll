define i32 @main() {
; <label>:0
	%1 = call i32 @call(i32 (i32, i32)* @add, i32 3, i32 5)
	ret i32 %1
}
define i32 @call(i32 (i32, i32)* %f, i32 %x, i32 %y) {
; <label>:0
	%1 = call i32 %f(i32 %x, i32 %y)
	ret i32 %1
}
define i32 @add(i32 %x, i32 %y) {
; <label>:0
	%1 = add i32 %x, %y
	ret i32 %1
}
