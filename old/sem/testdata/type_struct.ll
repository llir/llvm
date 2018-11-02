%t = type {i32, i8}

; Struct types.
;TODO: @a = external global {void}      ; error: invalid struct field type; expected single value or aggregate type, got *types.VoidType
;TODO: @b = external global {i32 ()}    ; error: invalid struct field type; expected single value or aggregate type, got *types.FuncType
@c = external global {i32}       ; valid
@d = external global {double}    ; valid
@e = external global {i32*}      ; valid
@f = external global {<5 x i32>} ; valid
@g = external global {label}     ; error: invalid struct field type; expected single value or aggregate type, got *types.LabelType
;@h = external global {metadata}  ; error: invalid struct field type; expected single value or aggregate type, got *types.MetadataType
@i = external global {[5 x i32]} ; valid
@j = external global {{i32, i8}} ; valid
@k = external global {%t}        ; valid
