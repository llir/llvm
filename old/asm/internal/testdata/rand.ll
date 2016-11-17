; Package rand implements the Borland C/C++ pseudo-random number generator
; algorithm.
;
; Random numbers are generated using a linear congruential generator with a
; multiplier of 0x15A4E35 and an increment of 1.
;
; References:
;  * https://en.wikipedia.org/wiki/Linear_congruential_generator#Parameters_in_common_use

; External declarations from the crt module.
declare i32 @abs(i32)

; seed represents the global seed.
@seed = global i32 0

; set_seed sets the global seed to x.
define void @set_seed(i32 %x) {
	store i32 %x, i32* @seed
	ret void
}

; rand returns a non-negative pseudo-random integer in [0, 2^31).
define i32 @rand() {
	%t1 = load i32, i32* @seed
	%t2 = mul i32 %t1, u0x15A4E35
	%t3 = add i32 %t2, 1
	store i32 %t3, i32* @seed
	%t4 = call i32(i32) @abs(i32 %t3)
	ret i32 %t4
}
