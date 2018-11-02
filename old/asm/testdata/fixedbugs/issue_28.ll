declare void @g()

define void @f() {
	call void () @g()
	ret void

	call void () @g()
	ret void
}
