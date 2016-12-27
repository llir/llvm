@.str = constant [17 x i8] c"hello \E4\B8\96 world\0A\00"
define i32 @main() {
; <label>:0
	%1 = call i32 (i8*, ...) @printf(i8* getelementptr ([17 x i8], [17 x i8]* @.str, i64 0, i64 0))
	ret i32 0
}
declare i32 @printf(i8*, ...)
