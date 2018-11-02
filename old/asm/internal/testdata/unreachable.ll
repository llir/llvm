define i32 @main() {
; <label>:0
	br i1 true, label %always, label %never
never:
	unreachable
always:
	ret i32 42
}
