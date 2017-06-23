; --- [ Memory expressions ] ---------------------------------------------------

@x = constant i32 42

@y = constant { i32, { [2 x i32], i8 } } zeroinitializer

; ~~~ [ getelementptr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32* @getelementptr_1() {
	; Plain expression.
	ret i32* getelementptr (i32, i32* @x)
}

define i32* @getelementptr_2() {
	; Indices.
	ret i32* getelementptr ({ i32, { [2 x i32], i8 } }, { i32, { [2 x i32], i8 } }* @y, i32 0, i32 1, i32 0, i32 1)
}

define i32* @getelementptr_3() {
	; Inbounds.
	ret i32* getelementptr inbounds (i32, i32* @x)
}
