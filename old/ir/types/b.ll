; ModuleID = 'a.ll'

%regset = type { i32, i16, i32, i16, i16, i16, i16, i64, i64, i64, i64, i64, i64, i64, i64, i64, i16, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i32, i32, i32, i32, i32, i32, i32, i32, i80, i80, i80, i80, i80, i80, i80, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i64, i80, i80, i80, i80, i80, i80, i80, i80, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512, i512 }

define void @foo(%regset*) {
  ret void
}

define i32 @main() {
  call void @foo(%regset* null)
  ret i32 42
}
