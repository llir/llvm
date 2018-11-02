; === [ Terminators ] ==========================================================

; ~~~ [ ret ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define i32 @ret_1() {
	; Plain terminator.
	ret i32 42
}

define void @ret_2() {
	; Void return
	ret void
}

define i32 @ret_3() {
	; Metadata.
	ret i32 42, !foo !{!"bar"}, !baz !{!"qux"}
}

; ~~~ [ br ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define void @br_1() {
	; Unconditional branch terminator.
	br label %foo
foo:
	ret void
}

define void @br_2() {
	; Metadata.
	br label %foo, !foo !{!"bar"}, !baz !{!"qux"}
foo:
	ret void
}

define void @br_3(i1 %cond) {
	; Conditional branch terminator.
	br i1 %cond, label %foo, label %bar
foo:
	ret void
bar:
	ret void
}

define void @br_4(i1 %cond) {
	; Metadata.
	br i1 %cond, label %foo, label %bar, !foo !{!"bar"}, !baz !{!"qux"}
foo:
	ret void
bar:
	ret void
}

; ~~~ [ switch ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define void @switch_1() {
	; Plain terminator.
	switch i32 1, label %default [
	]
default:
	ret void
}

define void @switch_2() {
	; Multiple cases.
	switch i32 2, label %default [
		i32 1, label %case1
		i32 2, label %case2
		i32 3, label %case3
	]
default:
	ret void
case1:
	ret void
case2:
	ret void
case3:
	ret void
}

define void @switch_3() {
	; Metadata.
	switch i32 2, label %default [
		i32 1, label %case1
		i32 2, label %case2
		i32 3, label %case3
	], !foo !{!"bar"}, !baz !{!"qux"}
default:
	ret void
case1:
	ret void
case2:
	ret void
case3:
	ret void
}

; ~~~ [ indirectbr ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

; TODO: add test cases for indirectbr terminator.

; ~~~ [ invoke ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

; TODO: add test cases for invoke terminator.

; ~~~ [ resume ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

; TODO: add test cases for resume terminator.

; ~~~ [ catchswitch ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

; TODO: add test cases for catchswitch terminator.

; ~~~ [ catchret ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

; TODO: add test cases for catchret terminator.

; ~~~ [ cleanupret ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

; TODO: add test cases for cleanupret terminator.

; ~~~ [ unreachable ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

define void @unreachable_1() {
	; Plain terminator.
	unreachable
}

define void @unreachable_2() {
	; Metadata.
	unreachable, !foo !{!"bar"}, !baz !{!"qux"}
}
