; --- [ Binary instructions ] --------------------------------------------------

; ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @add_1() {
	; Plain instruction.
	%result = add i32 30, 12
	ret i32 %result
}

define <2 x i32> @add_2() {
	; Vector operands.
	%result = add <2 x i32> <i32 30, i32 12>, <i32 12, i32 30>
	ret <2 x i32> %result
}

define i32 @add_3() {
	; Overflow flags.
	%result = add nsw nuw i32 30, 12
	ret i32 %result
}

define i32 @add_4() {
	; Metadata.
	%result = add i32 30, 12, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

define i32 @add_5() {
	; Full instruction.
	%result = add nsw nuw i32 30, 12, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define double @fadd_1() {
	; Plain instruction.
	%result = fadd double 30.0, 12.0
	ret double %result
}

define <2 x double> @fadd_2() {
	; Vector operands.
	%result = fadd <2 x double> <double 30.0, double 12.0>, <double 12.0, double 30.0>
	ret <2 x double> %result
}

define double @fadd_3() {
	; Fast-math flags.
	%result = fadd arcp fast ninf nnan nsz double 30.0, 12.0
	ret double %result
}

define double @fadd_4() {
	; Metadata.
	%result = fadd double 30.0, 12.0, !foo !{!"bar"}, !baz !{!"qux"}
	ret double %result
}

define double @fadd_5() {
	; Full instruction.
	%result = fadd arcp fast ninf nnan nsz double 30.0, 12.0, !foo !{!"bar"}, !baz !{!"qux"}
	ret double %result
}

; ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @sub_1() {
	; Plain instruction.
	%result = sub i32 50, 8
	ret i32 %result
}

define <2 x i32> @sub_2() {
	; Vector operands.
	%result = sub <2 x i32> <i32 50, i32 -8>, <i32 8, i32 -50>
	ret <2 x i32> %result
}

define i32 @sub_3() {
	; Overflow flags.
	%result = sub nsw nuw i32 50, 8
	ret i32 %result
}

define i32 @sub_4() {
	; Metadata.
	%result = sub i32 50, 8, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

define i32 @sub_5() {
	; Full instruction.
	%result = sub nsw nuw i32 50, 8, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define double @fsub_1() {
	; Plain instruction.
	%result = fsub double 50.0, 8.0
	ret double %result
}

define <2 x double> @fsub_2() {
	; Vector operands.
	%result = fsub <2 x double> <double 50.0, double -8.0>, <double 8.0, double -50.0>
	ret <2 x double> %result
}

define double @fsub_3() {
	; Fast-math flags.
	%result = fsub arcp fast ninf nnan nsz double 50.0, 8.0
	ret double %result
}

define double @fsub_4() {
	; Metadata.
	%result = fsub double 50.0, 8.0, !foo !{!"bar"}, !baz !{!"qux"}
	ret double %result
}

define double @fsub_5() {
	; Full instruction.
	%result = fsub arcp fast ninf nnan nsz double 50.0, 8.0, !foo !{!"bar"}, !baz !{!"qux"}
	ret double %result
}

; ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @mul_1() {
	; Plain instruction.
	%result = mul i32 21, 2
	ret i32 %result
}

define <2 x i32> @mul_2() {
	; Vector operands.
	%result = mul <2 x i32> <i32 21, i32 2>, <i32 2, i32 21>
	ret <2 x i32> %result
}

define i32 @mul_3() {
	; Overflow flags.
	%result = mul nsw nuw i32 21, 2
	ret i32 %result
}

define i32 @mul_4() {
	; Metadata.
	%result = mul i32 21, 2, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

define i32 @mul_5() {
	; Full instruction.
	%result = mul nsw nuw i32 21, 2, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define double @fmul_1() {
	; Plain instruction.
	%result = fmul double 21.0, 2.0
	ret double %result
}

define <2 x double> @fmul_2() {
	; Vector operands.
	%result = fmul <2 x double> <double 21.0, double 2.0>, <double 2.0, double 21.0>
	ret <2 x double> %result
}

define double @fmul_3() {
	; Fast-math flags.
	%result = fmul arcp fast ninf nnan nsz double 21.0, 2.0
	ret double %result
}

define double @fmul_4() {
	; Metadata.
	%result = fmul double 21.0, 2.0, !foo !{!"bar"}, !baz !{!"qux"}
	ret double %result
}

define double @fmul_5() {
	; Full instruction.
	%result = fmul arcp fast ninf nnan nsz double 21.0, 2.0, !foo !{!"bar"}, !baz !{!"qux"}
	ret double %result
}

; ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @udiv_1() {
	; Plain instruction.
	%result = udiv i32 84, 2
	ret i32 %result
}

define <2 x i32> @udiv_2() {
	; Vector operands.
	%result = udiv <2 x i32> <i32 84, i32 42>, <i32 2, i32 1>
	ret <2 x i32> %result
}

define i32 @udiv_3() {
	; Exact.
	%result = udiv exact i32 84, 2
	ret i32 %result
}

define i32 @udiv_4() {
	; Metadata.
	%result = udiv i32 84, 2, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

define i32 @udiv_5() {
	; Full instruction.
	%result = udiv exact i32 84, 2, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @sdiv_1() {
	; Plain instruction.
	%result = sdiv i32 -84, -2
	ret i32 %result
}

define <2 x i32> @sdiv_2() {
	; Vector operands.
	%result = sdiv <2 x i32> <i32 -84, i32 42>, <i32 -2, i32 1>
	ret <2 x i32> %result
}

define i32 @sdiv_3() {
	; Exact.
	%result = sdiv exact i32 -84, -2
	ret i32 %result
}

define i32 @sdiv_4() {
	; Metadata.
	%result = sdiv i32 -84, -2, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

define i32 @sdiv_5() {
	; Full instruction.
	%result = sdiv exact i32 -84, -2, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define double @fdiv_1() {
	; Plain instruction.
	%result = fdiv double 84.0, 2.0
	ret double %result
}

define <2 x double> @fdiv_2() {
	; Vector operands.
	%result = fdiv <2 x double> <double 84.0, double 42.0>, <double 2.0, double 1.0>
	ret <2 x double> %result
}

define double @fdiv_3() {
	; Fast-math flags.
	%result = fdiv arcp fast ninf nnan nsz double 84.0, 2.0
	ret double %result
}

define double @fdiv_4() {
	; Metadata.
	%result = fdiv double 84.0, 2.0, !foo !{!"bar"}, !baz !{!"qux"}
	ret double %result
}

define double @fdiv_5() {
	; Full instruction.
	%result = fdiv arcp fast ninf nnan nsz double 84.0, 2.0, !foo !{!"bar"}, !baz !{!"qux"}
	ret double %result
}

; ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @urem_1() {
	; Plain instruction.
	%result = urem i32 85, 43
	ret i32 %result
}

define <2 x i32> @urem_2() {
	; Vector operands.
	%result = urem <2 x i32> <i32 85, i32 162>, <i32 43, i32 60>
	ret <2 x i32> %result
}

define i32 @urem_3() {
	; Metadata.
	%result = urem i32 85, 43, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @srem_1() {
	; Plain instruction.
	%result = srem i32 85, -43
	ret i32 %result
}

define <2 x i32> @srem_2() {
	; Vector operands.
	%result = srem <2 x i32> <i32 85, i32 162>, <i32 -43, i32 -60>
	ret <2 x i32> %result
}

define i32 @srem_3() {
	; Metadata.
	%result = srem i32 85, -43, !foo !{!"bar"}, !baz !{!"qux"}
	ret i32 %result
}

; ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define double @frem_1() {
	; Plain instruction.
	%result = frem double 85.0, 43.0
	ret double %result
}

define <2 x double> @frem_2() {
	; Vector operands.
	%result = frem <2 x double> <double 85.0, double 162.0>, <double 43.0, double 60.0>
	ret <2 x double> %result
}

define double @frem_3() {
	; Fast-math flags.
	%result = frem arcp fast ninf nnan nsz double 85.0, 43.0
	ret double %result
}

define double @frem_4() {
	; Metadata.
	%result = frem double 85.0, 43.0, !foo !{!"bar"}, !baz !{!"qux"}
	ret double %result
}

define double @frem_5() {
	; Full instruction.
	%result = frem arcp fast ninf nnan nsz double 85.0, 43.0, !foo !{!"bar"}, !baz !{!"qux"}
	ret double %result
}
