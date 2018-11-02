; --- [ Other expressions ] ----------------------------------------------------

; ~~~ [ icmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i1 @icmp_1() {
	; Plain expression.
	ret i1 icmp ne (i32 42, i32 5)
}

define <2 x i1> @icmp_2() {
	; Vector operands.
	ret <2 x i1> icmp eq (<2 x i32> <i32 42, i32 11>, <2 x i32> <i32 42, i32 22>)
}

define i1 @icmp_3() {
	; Predicates.
	ret i1 icmp eq (i32 10, i32 15)
}

define i1 @icmp_4() {
	; Predicates.
	ret i1 icmp ne (i32 10, i32 15)
}

define i1 @icmp_5() {
	; Predicates.
	ret i1 icmp ugt (i32 10, i32 15)
}

define i1 @icmp_6() {
	; Predicates.
	ret i1 icmp uge (i32 10, i32 15)
}

define i1 @icmp_7() {
	; Predicates.
	ret i1 icmp ult (i32 10, i32 15)
}

define i1 @icmp_8() {
	; Predicates.
	ret i1 icmp ule (i32 10, i32 15)
}

define i1 @icmp_9() {
	; Predicates.
	ret i1 icmp sgt (i32 10, i32 15)
}

define i1 @icmp_10() {
	; Predicates.
	ret i1 icmp sge (i32 10, i32 15)
}

define i1 @icmp_11() {
	; Predicates.
	ret i1 icmp slt (i32 10, i32 15)
}

define i1 @icmp_12() {
	; Predicates.
	ret i1 icmp sle (i32 10, i32 15)
}

; ~~~ [ fcmp ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i1 @fcmp_1() {
	; Plain expression.
	ret i1 fcmp one (double 42.0, double 5.0)
}

define <2 x i1> @fcmp_2() {
	; Vector operands.
	ret <2 x i1> fcmp oeq (<2 x double> <double 42.0, double 11.0>, <2 x double> <double 42.0, double 22.0>)
}

define i1 @fcmp_3() {
	; Predicates.
	ret i1 fcmp false (double 10.0, double 15.0)
}

define i1 @fcmp_4() {
	; Predicates.
	ret i1 fcmp oeq (double 10.0, double 15.0)
}

define i1 @fcmp_5() {
	; Predicates.
	ret i1 fcmp ogt (double 10.0, double 15.0)
}

define i1 @fcmp_6() {
	; Predicates.
	ret i1 fcmp oge (double 10.0, double 15.0)
}

define i1 @fcmp_7() {
	; Predicates.
	ret i1 fcmp olt (double 10.0, double 15.0)
}

define i1 @fcmp_8() {
	; Predicates.
	ret i1 fcmp ole (double 10.0, double 15.0)
}

define i1 @fcmp_9() {
	; Predicates.
	ret i1 fcmp one (double 10.0, double 15.0)
}

define i1 @fcmp_10() {
	; Predicates.
	ret i1 fcmp ord (double 10.0, double 15.0)
}

define i1 @fcmp_11() {
	; Predicates.
	ret i1 fcmp ueq (double 10.0, double 15.0)
}

define i1 @fcmp_12() {
	; Predicates.
	ret i1 fcmp ugt (double 10.0, double 15.0)
}

define i1 @fcmp_13() {
	; Predicates.
	ret i1 fcmp uge (double 10.0, double 15.0)
}

define i1 @fcmp_14() {
	; Predicates.
	ret i1 fcmp ult (double 10.0, double 15.0)
}

define i1 @fcmp_15() {
	; Predicates.
	ret i1 fcmp ule (double 10.0, double 15.0)
}

define i1 @fcmp_16() {
	; Predicates.
	ret i1 fcmp une (double 10.0, double 15.0)
}

define i1 @fcmp_17() {
	; Predicates.
	ret i1 fcmp uno (double 10.0, double 15.0)
}

define i1 @fcmp_18() {
	; Predicates.
	ret i1 fcmp true (double 10.0, double 15.0)
}

; ~~~ [ select ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @select_1() {
	; Plain expression.
	ret i32 select (i1 true, i32 42, i32 37)
}

define <2 x i32> @select_2() {
	; Vector operands.
	ret <2 x i32> select (i1 false, <2 x i32> <i32 42, i32 37>, <2 x i32> <i32 11, i32 22>)
}
