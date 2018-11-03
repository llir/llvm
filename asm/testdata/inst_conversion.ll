define void @f() {
	trunc i32 321 to i8
	zext i8 123 to i32
	sext i8 -123 to i32
	fptrunc double 1.0 to float
	fpext float 2.0 to double
	fptoui double 3.0 to i32
	fptosi double -4.0 to i32
	uitofp i32 5 to double
	sitofp i32 -6 to double
	ptrtoint i8* null to i32
	inttoptr i32 1234 to i8*
	bitcast { i32, i32 } { i32 10, i32 20 } to i64
	addrspacecast i8* null to i8 addrspace(1)*
	ret void
}
