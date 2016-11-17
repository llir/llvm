; ModuleID = 'p02.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %1 = alloca i32, align 4
  %x = alloca i32, align 4
  store i32 0, i32* %1
  br label %2

; <label>:2                                       ; preds = %5, %0
  %3 = load i32, i32* %x, align 4
  %4 = icmp slt i32 %3, 10
  br i1 %4, label %5, label %8

; <label>:5                                       ; preds = %2
  %6 = load i32, i32* %x, align 4
  %7 = add nsw i32 %6, 3
  store i32 %7, i32* %x, align 4
  br label %2

; <label>:8                                       ; preds = %2
  %9 = load i32, i32* %x, align 4
  %10 = add nsw i32 %9, 3
  store i32 %10, i32* %x, align 4
  %11 = load i32, i32* %1
  ret i32 %11
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
