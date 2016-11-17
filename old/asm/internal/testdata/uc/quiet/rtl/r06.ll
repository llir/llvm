; ModuleID = 'r06.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define i32 @fac(i32 %n) #0 {
  %1 = alloca i32, align 4
  %2 = alloca i32, align 4
  %i = alloca i32, align 4
  %p = alloca i32, align 4
  store i32 %n, i32* %2, align 4
  %3 = load i32, i32* %2, align 4
  %4 = icmp slt i32 %3, 0
  br i1 %4, label %5, label %6

; <label>:5                                       ; preds = %0
  store i32 0, i32* %1
  br label %19

; <label>:6                                       ; preds = %0
  store i32 0, i32* %i, align 4
  store i32 1, i32* %p, align 4
  br label %7

; <label>:7                                       ; preds = %11, %6
  %8 = load i32, i32* %i, align 4
  %9 = load i32, i32* %2, align 4
  %10 = icmp slt i32 %8, %9
  br i1 %10, label %11, label %17

; <label>:11                                      ; preds = %7
  %12 = load i32, i32* %i, align 4
  %13 = add nsw i32 %12, 1
  store i32 %13, i32* %i, align 4
  %14 = load i32, i32* %p, align 4
  %15 = load i32, i32* %i, align 4
  %16 = mul nsw i32 %14, %15
  store i32 %16, i32* %p, align 4
  br label %7

; <label>:17                                      ; preds = %7
  %18 = load i32, i32* %p, align 4
  store i32 %18, i32* %1
  br label %19

; <label>:19                                      ; preds = %17, %5
  %20 = load i32, i32* %1
  ret i32 %20
}

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %1 = call i32 @fac(i32 5)
  ret i32 0
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
