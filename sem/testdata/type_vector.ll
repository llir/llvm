%t = type {i32, i8}

; Vector types.
;TODO: @a = external global <5 x void>      ; error: invalid vector element type; expected integer, floating-point or pointer type, got *types.VoidType
;TODO: @b = external global <5 x i32 ()>    ; error: invalid vector element type; expected integer, floating-point or pointer type, got *types.FuncType
@c = external global <5 x i32>       ; valid
@d = external global <5 x double>    ; valid
@e = external global <5 x i32*>      ; valid
@f = external global <5 x <5 x i32>> ; error: invalid vector element type; expected integer, floating-point or pointer type, got *types.VectorType
@g = external global <5 x label>     ; error: invalid vector element type; expected integer, floating-point or pointer type, got *types.LabelType
@h = external global <5 x metadata>  ; error: invalid vector element type; expected integer, floating-point or pointer type, got *types.MetadataType
@i = external global <5 x [5 x i32]> ; error: invalid vector element type; expected integer, floating-point or pointer type, got *types.ArrayType
@j = external global <5 x {i32, i8}> ; error: invalid vector element type; expected integer, floating-point or pointer type, got *types.StructType
@k = external global <5 x %t>        ; error: invalid vector element type; expected integer, floating-point or pointer type, got *types.StructType
