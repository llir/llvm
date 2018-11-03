@s = constant [4 x i8] c"foo\00"

define void @f() {
	%ptr = alloca i32
	load i32, i32* %ptr
	store i32 42, i32* %ptr
	fence acquire
	cmpxchg i32* %ptr, i32 10, i32 20 acquire monotonic
	atomicrmw add i32* %ptr, i32 30 acq_rel
	getelementptr [4 x i8], [4 x i8]* @s, i64 0, i64 0
	ret void
}
