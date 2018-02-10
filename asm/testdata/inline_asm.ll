define i32 @main() {
	%result = call i32 asm sideeffect "mov $1, $0; add $2, $0", "=r,i,i"(i32 32, i32 10)
	ret i32 %result
}
