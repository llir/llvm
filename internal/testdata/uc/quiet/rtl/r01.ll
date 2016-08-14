; ModuleID = 'r01.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@x = common global i32 0, align 4
@y = common global i8 0, align 1

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %z = alloca i32, align 4
  %w = alloca i8, align 1
  store i32 42, i32* @x, align 4
  store i8 43, i8* @y, align 1
  store i32 65, i32* %z, align 4
  store i8 10, i8* %w, align 1
  ret i32 0
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
