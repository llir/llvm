; main1 is a simple tool for printing random numbers.

; External declarations from the rand module.
declare void @set_seed(i32)
declare i32 @rand()

; External declarations from libc.
declare i32 @printf(i8*, ...)

@format1 = constant [18 x i8] c"initial seed: %d\0A\00"
@format2 = constant [14 x i8] c"seed: 0x%08X\0A\00"

define i32 @main() {
	; printf("initial seed: %d\n", 666);
	%t1 = getelementptr [18 x i8], [18 x i8]* @format1, i32 0, i32 0
	%initial_seed = add i32 666, 0
	call i32(i8*, ...) @printf(i8* %t1, i32 %initial_seed)

	; set_seed(666);
	call void(i32) @set_seed(i32 %initial_seed)
	br label %loop_init

	; for (int i = 0; i < 10; i++) {
loop_init:
	br label %loop_cond

loop_cond:
	%i.0 = phi i32 [0, %loop_init], [%i.1, %loop_post]
	%cond = icmp slt i32 %i.0, 10
	br i1 %cond, label %loop_body, label %loop_exit

loop_body:
	;    printf("seed: %0x%08X\n", rand());
	%t2 = getelementptr [14 x i8], [14 x i8]* @format2, i32 0, i32 0
	%t3 = call i32() @rand()
	call i32(i8*, ...) @printf(i8* %t2, i32 %t3)
	br label %loop_post

loop_post:
	%i.1 = add i32 %i.0, 1
	br label %loop_cond

loop_exit:
	; }

	ret i32 0
}
