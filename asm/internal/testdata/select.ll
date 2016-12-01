@x = global i32 3
define i32 @main() {
; <label>:0
	%1 = load i32, i32* @x
	%2 = icmp ne i32 %1, 0
	%. = select i1 %2, i32 1, i32 2
	ret i32 %.
}
