%struct.t = type { i32 }
@x = constant %struct.t { i32 42 }
define void @f(%struct.t) {
; <label>:1
	ret void
}
define void @g() {
; <label>:0
	call void @f(%struct.t { i32 42 })
	ret void
}
