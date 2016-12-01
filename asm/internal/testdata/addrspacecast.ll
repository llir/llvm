define i32 @main() {
; <label>:0
	%1 = alloca i32
	%2 = addrspacecast i32* %1 to i32 addrspace(1)*
	%3 = addrspacecast i32 addrspace(1)* %2 to i64 addrspace(2)*
	%4 = alloca <4 x i32*>
	%5 = load <4 x i32*>, <4 x i32*>* %4
	%6 = addrspacecast <4 x i32*> %5 to <4 x float addrspace(3)*>
	ret i32 0
}
