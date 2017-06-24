; === [ Constants ] ============================================================

; --- [ Integer constant ] -----------------------------------------------------

@g1 = global i32 42
@g2 = global i32 -42

; --- [ Floating-point constant ] ----------------------------------------------

; Fraction floating-point literal.
@g3 = global double 42.0
@g4 = global double +42.0
@g5 = global double -42.0
@g6 = global double 42.
@g7 = global double +42.
@g8 = global double -42.

; Scientific floating-point literal.
@g9 = global double 4.2e1
@g10 = global double 4.2E1
@g11 = global double 4.2e+1
@g12 = global double 4.2E+1
@g13 = global double 420.0e-1
@g14 = global double 420.0E-1
@g15 = global double +4.2e1
@g16 = global double +4.2E1
@g17 = global double +4.2e+1
@g18 = global double +4.2E+1
@g19 = global double +420.0e-1
@g20 = global double +420.0E-1
@g21 = global double -4.2e1
@g22 = global double -4.2E1
@g23 = global double -4.2e+1
@g24 = global double -4.2E+1
@g25 = global double -420.0e-1
@g26 = global double -420.0E-1

; Hexadecimal floating-point literal.
@g27 = global double 0x0000000000000000
@g28 = global x86_fp80 0xK00000000000000000000
; TODO: add test case for 0xL floating-point constant.
;@g29 = global fp128 0xL00000000000000000000000000000000
; TODO: add test case for 0xM floating-point constant.
;@g30 = global ppc_fp128 0xM00000000000000000000000000000000
@g31 = global half 0xH0000

; --- [ Pointer constant ] -----------------------------------------------------

; Null pointer.
@g32 = global i32* null

; Global variable.
@g33 = global i32** @g32

; Function.
@g34 = global void ()* @f1

; --- [ Vector constant ] ------------------------------------------------------

@g35 = global <1 x i32> <i32 42>
@g36 = global <2 x i32> <i32 42, i32 11>

; --- [ Array constant ] -------------------------------------------------------

; Array constant.
@g37 = global [0 x i32] []
@g38 = global [1 x i32] [i32 42]
@g39 = global [2 x i32] [i32 42, i32 11]

; Character array.
@g40 = global [0 x i8] c""
@g41 = global [1 x i8] c"f"
@g42 = global [2 x i8] c"fo"
@g43 = global [3 x i8] c"foo"

; --- [ Struct constant ] ------------------------------------------------------

; Struct constant.
@g44 = global {} {}
@g45 = global { i32 } { i32 42 }
@g46 = global { i32, i32 } { i32 42, i32 42 }

; Nested struct.
@g47 = global { i32, { i8 } } { i32 42, { i8 } { i8 42 } }

; Packed.
@g48 = global <{}> <{}>
@g49 = global <{ i32, i8, i32 }> <{ i32 42, i8 5, i32 11 }>

; --- [ Zero initializer constant ] --------------------------------------------

@g50 = global { i32, i8, { i32, i32 }, i8 } zeroinitializer

; --- [ Undefined value constant ] ---------------------------------------------

@g51 = global i8* undef

define void @f1() {
	ret void
}
