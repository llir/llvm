; --- [ Bitwise instructions ] -------------------------------------------------

; ~~~ [ shl ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @shl_1() {
	; Plain instruction.
	%result = shl i32 21, 1
	ret i32 %result
}

define <2 x i32> @shl_2() {
	; Vector operands.
	%result = shl <2 x i32> <i32 21, i32 42>, <i32 1, i32 0>
	ret <2 x i32> %result
}

define i32 @shl_3() {
	; Overflow flags.
	%result = shl nsw nuw i32 21, 1
	ret i32 %result
}

define i32 @shl_4() {
	; Metadata.
	%result = shl i32 21, 1, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

define i32 @shl_5() {
	; Full instruction.
	%result = shl nsw nuw i32 21, 1, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ lshr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @lshr_1() {
	; Plain instruction.
	%result = lshr i32 84, 1
	ret i32 %result
}

define <2 x i32> @lshr_2() {
	; Vector operands.
	%result = lshr <2 x i32> <i32 84, i32 42>, <i32 1, i32 0>
	ret <2 x i32> %result
}

define i32 @lshr_3() {
	; Exact.
	%result = lshr exact i32 84, 1
	ret i32 %result
}

define i32 @lshr_4() {
	; Metadata.
	%result = lshr i32 84, 1, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

define i32 @lshr_5() {
	; Full instruction.
	%result = lshr exact i32 84, 1, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ ashr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @ashr_1() {
	; Plain instruction.
	%result = ashr i32 84, 1
	ret i32 %result
}

define <2 x i32> @ashr_2() {
	; Vector operands.
	%result = ashr <2 x i32> <i32 84, i32 42>, <i32 1, i32 0>
	ret <2 x i32> %result
}

define i32 @ashr_3() {
	; Exact.
	%result = ashr exact i32 84, 1
	ret i32 %result
}

define i32 @ashr_4() {
	; Metadata.
	%result = ashr i32 84, 1, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

define i32 @ashr_5() {
	; Negative operand.
	%result = ashr i32 -84, 1
	ret i32 %result
}

define i32 @ashr_6() {
	; Full instruction.
	%result = ashr exact i32 84, 1, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ and ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @and_1() {
	; Plain instruction.
	%result = and i32 58, 239 ; 0b00111010 & 0b11101111
	ret i32 %result
}

define <2 x i32> @and_2() {
	; Vector operands.
	%result = and <2 x i32> <i32 58, i32 170>, <i32 239, i32 127> ; 0b00111010 & 0b11101111, 0b10101010 & 0b01111111
	ret <2 x i32> %result
}

define i32 @and_3() {
	; Metadata.
	%result = and i32 58, 239, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ or ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @or_1() {
	; Plain instruction.
	%result = or i32 10, 32 ; 0b00001010 | 0b0100000
	ret i32 %result
}

define <2 x i32> @or_2() {
	; Vector operands.
	%result = or <2 x i32> <i32 10, i32 40>, <i32 32, i32 2> ; 0b00001010 | 0b0100000, 0b00101000 & 0b00000010
	ret <2 x i32> %result
}

define i32 @or_3() {
	; Metadata.
	%result = or i32 10, 32, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ xor ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @xor_1() {
	; Plain instruction.
	%result = xor i32 255, 213 ; 0b11111111 ^ 0b11010101
	ret i32 %result
}

define <2 x i32> @xor_2() {
	; Vector operands.
	%result = xor <2 x i32> <i32 255, i32 123>, <i32 213, i32 81> ; 0b11111111 ^ 0b11010101, 0b01111011 ^ 0b01010001
	ret <2 x i32> %result
}

define i32 @xor_3() {
	; Metadata.
	%result = xor i32 255, 213, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}
