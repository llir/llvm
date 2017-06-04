%foo = type { %bar* }

%bar = type { %foo* }

define i32 @main() {
; <label>:0
	%1 = alloca %foo
	%2 = alloca %bar
	%3 = load %foo, %foo* %1
	%4 = load %bar, %bar* %2
	ret i32 42
}
