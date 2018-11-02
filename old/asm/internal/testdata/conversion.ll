define i32 @main() {
; <label>:0
	%1 = trunc i64 5 to i32
	%2 = zext i32 3 to i64
	%3 = sext i32 3 to i64
	%4 = fptrunc double 5.0 to float
	%5 = fpext float 3.0 to double
	%6 = fptoui double 5.0 to i64
	%7 = fptosi double 3.0 to i64
	%8 = uitofp i64 5 to double
	%9 = sitofp i64 3 to double
	%10 = ptrtoint i8* null to i64
	%11 = inttoptr i64 0 to i8*
	ret i32 0
}
