%t = type {i32, i8}

; function return types.
declare void @a()      ; valid
declare i32 () @b()    ; error: invalid function return type; expected void or first class type except label and metadata, got *types.FuncType
declare i32 @c()       ; valid
declare double @d()    ; valid
declare i32* @e()      ; valid
declare <5 x i32> @f() ; valid
declare label @g()     ; error: invalid function return type; expected void or first class type except label and metadata, got *types.LabelType
declare metadata @h()  ; error: invalid function return type; expected void or first class type except label and metadata, got *types.MetadataType
declare [5 x i32] @i() ; valid
declare {i32, i8} @j() ; valid
declare %t @k()        ; valid

; function argument types.
;TODO: declare void @l(void %x)      ; error: invalid function argument; expected first class type, got *types.VoidType
;TODO: declare void @m(i32 () %x)    ; error: invalid function argument; expected first class type, got *types.FuncType
declare void @n(i32 %x)       ; valid
declare void @o(double %x)    ; valid
declare void @p(i32* %x)      ; valid
declare void @q(<5 x i32> %x) ; valid
declare void @r(label %x)     ; valid
declare void @s(metadata %x)  ; valid
declare void @t([5 x i32] %x) ; valid
declare void @u({i32, i8} %x) ; valid
declare void @v(%t %x)        ; valid
