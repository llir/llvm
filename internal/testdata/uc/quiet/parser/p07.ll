; ModuleID = 'p07.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %1 = alloca i32, align 4
  %x = alloca i32, align 4
  %y = alloca i32, align 4
  store i32 0, i32* %1
  %2 = load i32, i32* %x, align 4
  %3 = icmp ne i32 %2, 0
  br i1 %3, label %4, label %10

; <label>:4                                       ; preds = %0
  br label %5

; <label>:5                                       ; preds = %8, %4
  %6 = load i32, i32* %y, align 4
  %7 = icmp ne i32 %6, 0
  br i1 %7, label %8, label %9

; <label>:8                                       ; preds = %5
  store i32 42, i32* %x, align 4
  br label %5

; <label>:9                                       ; preds = %5
  br label %10

; <label>:10                                      ; preds = %9, %0
  br label %11

; <label>:11                                      ; preds = %18, %10
  %12 = load i32, i32* %x, align 4
  %13 = icmp ne i32 %12, 0
  br i1 %13, label %14, label %19

; <label>:14                                      ; preds = %11
  %15 = load i32, i32* %y, align 4
  %16 = icmp ne i32 %15, 0
  br i1 %16, label %17, label %18

; <label>:17                                      ; preds = %14
  store i32 42, i32* %x, align 4
  br label %18

; <label>:18                                      ; preds = %17, %14
  br label %11

; <label>:19                                      ; preds = %11
  %20 = load i32, i32* %1
  ret i32 %20
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
