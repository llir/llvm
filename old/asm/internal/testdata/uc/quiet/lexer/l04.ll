; ModuleID = 'l04.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %1 = alloca i32, align 4
  %i = alloca i32, align 4
  %j = alloca i8, align 1
  store i32 0, i32* %1
  store i32 1, i32* %i, align 4
  br label %2

; <label>:2                                       ; preds = %3, %0
  br i1 false, label %3, label %4

; <label>:3                                       ; preds = %2
  store i32 0, i32* %i, align 4
  br label %2

; <label>:4                                       ; preds = %2
  ret i32 42
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
