source_filename = "foo.c"
target datalayout = "e"
target triple = "x86_64-unknown-linux"

module asm "foo"

%t1 = type i32

%t2 = type opaque

$com1 = comdat any
$com2 = comdat exactmatch
$com3 = comdat largest
$com4 = comdat noduplicates
$com5 = comdat samesize

@g1 = external global i32

@g2 = global i32 0

declare void @exit(i32 %staus) #0

declare i32 @printf(i8*, ...)

define void @f1() {
	ret void
}

define i32 @f2(i32 %x, i32 %y) {
	ret i32 42
}

attributes #0 = { noreturn }

!foo = !{!0}

!0 = !{!"foo"}
