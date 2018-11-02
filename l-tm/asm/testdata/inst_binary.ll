define void @f() {
	add i32 1, 2
	fadd double 3.0, 4.0
	sub i32 5, 6
	fsub double 7.0, 8.0
	mul i32 9, 10
	fmul double 11.0, 12.0
	udiv i32 13, 14
	sdiv i32 15, 16
	fdiv double 17.0, 18.0
	urem i32 19, 20
	srem i32 21, 22
	frem double 23.0, 24.0
	ret void
}
