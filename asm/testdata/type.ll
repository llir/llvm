; === [ Types ] ================================================================

; Void type
;%t1 = type void

; Function type
%t2 = type i32 (i32)

; Integer type
%t3 = type i32

; Floating-point type
%t4 = type double

; MMX type
; TODO: add test case for MMX type.
;%t5 = type x86_mmx

; Pointer type
%t6 = type i32*

; Vector type
%t7 = type <2 x i32>

; Label type
%t8 = type label

; Token type
; TODO: add test case for token type.
;%t9 = type token

; Metadata type
%t10 = type metadata

; Array type
%t11 = type [2 x i32]

; Struct type
%t12 = type { i32, double }

; Packed struct type
%t13 = type <{ i32, i8, i32 }>

%t14 = type %t3

; --- [ Void type ] ------------------------------------------------------------

declare void @f1()

; --- [ Function type ] --------------------------------------------------------

; --- [ Integer type ] ---------------------------------------------------------

declare i1 @f2()
declare i2 @f3()
declare i4 @f4()
declare i8 @f5()
declare i16 @f6()
declare i32 @f7()
declare i64 @f8()
declare i128 @f9()

; --- [ Floating-point type ] --------------------------------------------------

declare half @f10()
declare float @f11()
declare double @f12()
declare fp128 @f13()
declare x86_fp80 @f14()
declare ppc_fp128 @f15()

; --- [ MMX type ] -------------------------------------------------------------

; TODO: add test case for MMX type.

; --- [ Pointer type ] ---------------------------------------------------------

declare i8* @f16()

; --- [ Vector type ] ----------------------------------------------------------

declare <2 x i8> @f17()

; --- [ Label type ] -----------------------------------------------------------

define void @f18() {
	br label %foo
foo:
	ret void
}

; --- [ Token type ] -----------------------------------------------------------

; TODO: add test case for token type.

; --- [ Metadata type ] --------------------------------------------------------

declare void @f19(metadata %x)

; --- [ Array type ] -----------------------------------------------------------

declare [2 x i32] @f20()

; --- [ Struct type ] ----------------------------------------------------------

declare { i32 } @f21()
declare { i32, i8, [2 x i32], { i32, <2 x i8> } } @f22()
declare <{ i32, i8, i32 }> @f23()

; --- [ Named type ] -----------------------------------------------------------

declare %t3 @f25()
declare %t4 @f26()
; TODO: add test case for MMX type.
;declare %t5 @f27()
declare %t6 @f28()
declare %t7 @f29()
; TODO: add test case for token type.
;declare %t9 @f31()
declare void @f32(%t10 %x)
declare %t11 @f33()
declare %t12 @f34()
declare %t13 @f35()
