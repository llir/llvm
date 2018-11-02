; Vector constants.
;TODO: @a = global <3 x i32> <i32 1, i32 2>        ; error: number of vector elements mismatch for type `[3 x i32]`; expected 3, got 2
@b = global <3 x i32> <i32 1, i8 2, i32 3>  ; error: vector element type `i32` and element type `i8` mismatch
@c = global <3 x i32> <i32 1, i32 2, i32 3> ; valid
