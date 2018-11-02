; --- [ Vector instructions ] --------------------------------------------------

; ~~~ [ extractelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @extractelement_1() {
	; Plain instruction.
	%result = extractelement <2 x i32> <i32 21, i32 42>, i32 1
	ret i32 %result
}

define i32 @extractelement_2() {
	; Metadata.
	%result = extractelement <2 x i32> <i32 21, i32 42>, i32 1, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ insertelement ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define <2 x i32> @insertelement_1() {
	; Plain instruction.
	%result = insertelement <2 x i32> <i32 21, i32 42>, i32 42, i32 0
	ret <2 x i32> %result
}

define <2 x i32> @insertelement_2() {
	; Metadata.
	%result = insertelement <2 x i32> <i32 21, i32 42>, i32 42, i32 0, !foo !{!"bar"}, !baz !{!"qux"}
	ret <2 x i32> %result
}

; ~~~ [ shufflevector ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define <2 x i32> @shufflevector_1() {
	; Plain instruction.
	%result = shufflevector <2 x i32> <i32 21, i32 42>, <2 x i32> <i32 42, i32 84>, <2 x i32> <i32 1, i32 2>
	ret <2 x i32> %result
}

define <2 x i32> @shufflevector_2() {
	; Metadata.
	%result = shufflevector <2 x i32> <i32 21, i32 42>, <2 x i32> <i32 42, i32 84>, <2 x i32> <i32 1, i32 2>, !foo !{!"bar"}, !baz !{!"qux"}
	ret <2 x i32> %result
}
