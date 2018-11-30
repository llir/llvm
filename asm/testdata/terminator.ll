define void @f(i8* %target) {
; <label>:0
	indirectbr i8* %target, [label %foo]

foo:
	br label %bar

bar:
	ret void
}
