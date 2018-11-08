define void @f() {
	ret void
}

define i32 @g() {
	ret i32 42
}

define void @h() {
	br label %foo
foo:
	ret void
}

define i32 @i(i1 %cond) {
	br i1 %cond, label %t, label %f
t:
	ret i32 10
f:
	ret i32 20
}

define i32 @j(i32 %x) {
	switch i32 %x, label %d [
		i32 10, label %c1
		i32 20, label %c2
		i32 30, label %c3
		i32 40, label %c4
	]
c1:
	ret i32 1
c2:
	ret i32 2
c3:
	ret i32 3
c4:
	indirectbr i8* blockaddress(@j, %c2), [ label %c2 ]
d:
	ret i32 5
}

define void @k(i32 %x) {
	ret void
}

define void @l() personality i32 10 {
	%1 = invoke i32 @g() to label %continue unwind label %cleanup
	invoke void @k(i32 30) to label %continue unwind label %cleanup
continue:
	ret void
cleanup:
	%exn = landingpad { i8*, i32 } catch i8** null
	resume { i8*, i32 } %exn
}

define void @m() personality i32 10 {
	ret void
handler0:
	%catchpad = catchpad within %cs [i8** null]
	ret void
handler1:
	%cleanuppad = cleanuppad within %cs [i8** null]
	ret void
dispatch1:
	catchret from %catchpad to label %exit
dispatch2:
	cleanupret from %cleanuppad unwind label %exit
dispatch3:
	%cs = catchswitch within none [label %handler0, label %handler1] unwind to caller
exit:
	ret void
}

define void @n() {
	unreachable
}

- [ ] ret
- [ ] br
- [ ] condbr
- [ ] switch
- [x] indirectbr
- [ ] invoke
- [ ] resume
- [ ] catchswitch
- [ ] catchret
- [ ] cleanupret
- [ ] unreachable
