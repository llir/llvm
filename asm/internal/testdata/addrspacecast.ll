define i32 @main() {
; <label>:0
	%1 = alloca i32
	%2 = addrspacecast i32* %1 to i32 addrspace(1)*
	%3 = addrspacecast i32 addrspace(1)* %2 to i64 addrspace(2)*
	ret i32 0
}
