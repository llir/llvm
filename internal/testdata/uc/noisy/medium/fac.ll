; ModuleID = 'fac.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@a = common global i32 0, align 4

; Function Attrs: nounwind uwtable
define i32 @fac(i32 %n) #0 {
  %1 = alloca i32, align 4
  %2 = alloca i32, align 4
  store i32 %n, i32* %2, align 4
  %3 = load i32, i32* %2, align 4
  %4 = icmp eq i32 %3, 0
  br i1 %4, label %5, label %6

; <label>:5                                       ; preds = %0
  store i32 1, i32* %1
  br label %12

; <label>:6                                       ; preds = %0
  %7 = load i32, i32* %2, align 4
  %8 = load i32, i32* %2, align 4
  %9 = sub nsw i32 %8, 1
  %10 = call i32 @fac(i32 %9)
  %11 = mul nsw i32 %7, %10
  store i32 %11, i32* %1
  br label %12

; <label>:12                                      ; preds = %6, %5
  %13 = load i32, i32* %1
  ret i32 %13
}

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %1 = call i32 @getint()
  store i32 %1, i32* @a, align 4
  %2 = load i32, i32* @a, align 4
  %3 = call i32 @fac(i32 %2)
  call void @putint(i32 %3)
  ret i32 0
}

declare i32 @getint() #1

declare void @putint(i32) #1

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
