%t = type {i32, i8}

; Pointer types.
@a = external global void*      ; error: invalid pointer element type; expected function type, single value type, or aggregate type, got *types.VoidType
@b = external global i32 ()*    ; valid
@c = external global i32*       ; valid
@d = external global double*    ; valid
@e = external global i32**      ; valid
@f = external global <5 x i32>* ; valid
@g = external global label*     ; error: invalid pointer element type; expected function type, single value type, or aggregate type, got *types.LabelType
@h = external global metadata*  ; error: invalid pointer element type; expected function type, single value type, or aggregate type, got *types.MetadataType
@i = external global [5 x i32]* ; valid
@j = external global {i32, i8}* ; valid
@k = external global %t*        ; valid
