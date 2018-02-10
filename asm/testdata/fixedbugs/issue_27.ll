; minimal test case adapted from the @main function of base32.ll, as part of
; coreutils in https://github.com/decomp/testdata

define i32 @main(i32, i8**) {
entry:
	br label %loop_init

loop_init:
	br label %loop_post

loop_cond:
	%cond = icmp ult i32 %i.0, 42
	br i1 %cond, label %loop_post, label %loop_exit

loop_post:
	%i.1 = phi i32 [ %i.0, %loop_cond ], [ 0, %loop_init ]
	%i.0 = add i32 %i.1, 1
	br label %loop_cond

loop_exit:
	ret i32 %i.0
}
