; ModuleID = 'm02.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define i32 @f(i32 %x) #0 {
  %1 = alloca i32, align 4
  %2 = alloca i32, align 4
  store i32 %x, i32* %2, align 4
  %3 = load i32, i32* %2, align 4
  %4 = icmp sgt i32 %3, 0
  br i1 %4, label %5, label %10

; <label>:5                                       ; preds = %0
  %6 = load i32, i32* %2, align 4
  %7 = sub nsw i32 %6, 1
  %8 = call i32 @f(i32 %7)
  %9 = add nsw i32 2, %8
  store i32 %9, i32* %1
  br label %11

; <label>:10                                      ; preds = %0
  store i32 112, i32* %1
  br label %11

; <label>:11                                      ; preds = %10, %5
  %12 = load i32, i32* %1
  ret i32 %12
}

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %1 = call i32 @f(i32 8)
  ret i32 0
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
