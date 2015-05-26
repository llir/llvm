@bar = global i32 20
@s = constant [9 x i8] c"hello %d\0a"
;@bar = external global i32

declare i32 @printf(i8*, i32)

define i32 @main() {
	%s = getelementptr [9 x i8]* @s, i32 0, i32 0
	%foo = load i32* @bar
	call i32 @printf(i8* %s, i32 %foo)
	ret i32 42
}
