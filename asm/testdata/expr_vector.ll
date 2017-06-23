; --- [ Vector expressions ] ---------------------------------------------------

; ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @extractelement_1() {
	ret i32 extractelement (<2 x i32> <i32 21, i32 42>, i32 1)
}

; ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define <2 x i32> @insertelement_1() {
	ret <2 x i32> insertelement (<2 x i32> <i32 21, i32 42>, i32 42, i32 0)
}

; ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define <2 x i32> @shufflevector_1() {
	ret <2 x i32> shufflevector (<2 x i32> <i32 21, i32 42>, <2 x i32> <i32 42, i32 84>, <2 x i32> <i32 1, i32 2>)
}
