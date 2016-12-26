%t = type {i32, i8}

; function return types.
declare void @a()      ; valid
declare label @b()     ; error: invalid function return type; expected void or first class type except label and metadata, got *types.LabelType
declare metadata @c()  ; error: invalid function return type; expected void or first class type except label and metadata, got *types.MetadataType
declare i32 @d()       ; valid
declare double @e()    ; valid
declare i32 () @f()    ; error: invalid function return type; expected void or first class type except label and metadata, got *types.FuncType
declare i32* @g()      ; valid
declare <5 x i32> @h() ; valid
declare [5 x i32] @i() ; valid
declare {i32, i8} @j() ; valid
declare %t @k()        ; valid

; function argument types.
declare void @m(label %x)     ; valid
declare void @n(metadata %x)  ; valid
declare void @o(i32 %x)       ; valid
declare void @p(double %x)    ; valid
declare void @r(i32* %x)      ; valid
declare void @s(<5 x i32> %x) ; valid
declare void @t([5 x i32] %x) ; valid
declare void @u({i32, i8} %x) ; valid
declare void @v(%t %x)        ; valid
