%t = type {i32, i8}

; Global variables.
;TODO: @a = external global void      ; error: invalid global content type; expected single value or aggregate type, got *types.VoidType
;TODO: @b = external global i32 ()    ; error: invalid global content type; expected single value or aggregate type, got *types.FuncType
@c = external global i32       ; valid
@d = external global double    ; valid
@e = external global i32*      ; valid
@f = external global <5 x i32> ; valid
;@g = external global label     ; error: invalid global content type; expected single value or aggregate type, got *types.LabelType
;@h = external global metadata  ; error: invalid global content type; expected single value or aggregate type, got *types.MetadataType
@i = external global [5 x i32] ; valid
@j = external global {i32, i8} ; valid
@k = external global %t        ; valid

;TODO: @l = global {i32, i8} [i8 3, i8 5] ; error: global variable content type `{i32, i8}` and initial value type `[2 x i8]` mismatch
