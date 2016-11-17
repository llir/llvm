; ModuleID = 'fib.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define i32 @fib(i32 %n) #0 {
  %1 = alloca i32, align 4
  %2 = alloca i32, align 4
  store i32 %n, i32* %2, align 4
  %3 = load i32, i32* %2, align 4
  %4 = icmp eq i32 %3, 0
  br i1 %4, label %5, label %6

; <label>:5                                       ; preds = %0
  store i32 1, i32* %1
  br label %18

; <label>:6                                       ; preds = %0
  %7 = load i32, i32* %2, align 4
  %8 = icmp eq i32 %7, 1
  br i1 %8, label %9, label %10

; <label>:9                                       ; preds = %6
  store i32 1, i32* %1
  br label %18

; <label>:10                                      ; preds = %6
  %11 = load i32, i32* %2, align 4
  %12 = sub nsw i32 %11, 1
  %13 = call i32 @fib(i32 %12)
  %14 = load i32, i32* %2, align 4
  %15 = sub nsw i32 %14, 2
  %16 = call i32 @fib(i32 %15)
  %17 = add nsw i32 %13, %16
  store i32 %17, i32* %1
  br label %18

; <label>:18                                      ; preds = %10, %9, %5
  %19 = load i32, i32* %1
  ret i32 %19
}

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %1 = alloca i32, align 4
  %i = alloca i32, align 4
  %space = alloca [2 x i8], align 1
  %cr = alloca [2 x i8], align 1
  store i32 0, i32* %1
  %2 = getelementptr inbounds [2 x i8], [2 x i8]* %space, i32 0, i64 0
  store i8 32, i8* %2, align 1
  %3 = getelementptr inbounds [2 x i8], [2 x i8]* %space, i32 0, i64 1
  store i8 0, i8* %3, align 1
  %4 = getelementptr inbounds [2 x i8], [2 x i8]* %cr, i32 0, i64 0
  store i8 10, i8* %4, align 1
  %5 = getelementptr inbounds [2 x i8], [2 x i8]* %cr, i32 0, i64 1
  store i8 0, i8* %5, align 1
  store i32 0, i32* %i, align 4
  br label %6

; <label>:6                                       ; preds = %9, %0
  %7 = load i32, i32* %i, align 4
  %8 = icmp sle i32 %7, 12
  br i1 %8, label %9, label %17

; <label>:9                                       ; preds = %6
  %10 = load i32, i32* %i, align 4
  call void @putint(i32 %10)
  %11 = getelementptr inbounds [2 x i8], [2 x i8]* %space, i32 0, i32 0
  call void @putstring(i8* %11)
  %12 = load i32, i32* %i, align 4
  %13 = call i32 @fib(i32 %12)
  call void @putint(i32 %13)
  %14 = getelementptr inbounds [2 x i8], [2 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %14)
  %15 = load i32, i32* %i, align 4
  %16 = add nsw i32 %15, 1
  store i32 %16, i32* %i, align 4
  br label %6

; <label>:17                                      ; preds = %6
  %18 = load i32, i32* %1
  ret i32 %18
}

declare void @putint(i32) #1

declare void @putstring(i8*) #1

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
