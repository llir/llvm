%t = type {i32, i8}

; Array types.
;TODO: @a = external global [5 x void]      ; error: invalid array element type; expected single value or aggregate type, got *types.VoidType
;TODO: @b = external global [5 x i32 ()]    ; error: invalid array element type; expected single value or aggregate type, got *types.FuncType
@c = external global [5 x i32]       ; valid
@d = external global [5 x double]    ; valid
@e = external global [5 x i32*]      ; valid
@f = external global [5 x <5 x i32>] ; valid
@g = external global [5 x label]     ; error: invalid array element type; expected single value or aggregate type, got *types.LabelType
@h = external global [5 x metadata]  ; error: invalid array element type; expected single value or aggregate type, got *types.MetadataType
@i = external global [5 x [5 x i32]] ; valid
@j = external global [5 x {i32, i8}] ; valid
@k = external global [5 x %t]        ; valid
