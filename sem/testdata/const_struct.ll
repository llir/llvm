; Struct constants.
;TODO: @a = global {i32, i16, i8} {i32 1, i16 2}       ; error: number of struct fields mismatch for type `{i32, i16, i8}`; expected 3, got 2
;TODO: @b = global {i32, i16, i8} {i32 1, i16 2, i4 3} ; error: struct field type `i8` and field type `i4` mismatch
@c = global {i32, i16, i8} {i32 1, i16 2, i8 3} ; valid
