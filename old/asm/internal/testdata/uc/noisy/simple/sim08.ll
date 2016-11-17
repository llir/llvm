; ModuleID = 'sim08.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define void @foo(i32 %v1, i32 %v2, i32 %v3, i32 %v4, i32 %v5) #0 {
  %1 = alloca i32, align 4
  %2 = alloca i32, align 4
  %3 = alloca i32, align 4
  %4 = alloca i32, align 4
  %5 = alloca i32, align 4
  store i32 %v1, i32* %1, align 4
  store i32 %v2, i32* %2, align 4
  store i32 %v3, i32* %3, align 4
  store i32 %v4, i32* %4, align 4
  store i32 %v5, i32* %5, align 4
  %6 = load i32, i32* %1, align 4
  call void @putint(i32 %6)
  %7 = load i32, i32* %2, align 4
  call void @putint(i32 %7)
  %8 = load i32, i32* %3, align 4
  call void @putint(i32 %8)
  %9 = load i32, i32* %4, align 4
  call void @putint(i32 %9)
  %10 = load i32, i32* %5, align 4
  call void @putint(i32 %10)
  ret void
}

declare void @putint(i32) #1

; Function Attrs: nounwind uwtable
define i32 @f(i32 %x) #0 {
  %1 = alloca i32, align 4
  store i32 %x, i32* %1, align 4
  %2 = load i32, i32* %1, align 4
  %3 = add nsw i32 %2, 1
  ret i32 %3
}

; Function Attrs: nounwind uwtable
define signext i8 @g(i8 signext %x) #0 {
  %1 = alloca i8, align 1
  %2 = alloca i8, align 1
  store i8 %x, i8* %2, align 1
  %3 = load i8, i8* %2, align 1
  %4 = icmp ne i8 %3, 0
  br i1 %4, label %6, label %5

; <label>:5                                       ; preds = %0
  store i8 1, i8* %1
  br label %15

; <label>:6                                       ; preds = %0
  %7 = load i8, i8* %2, align 1
  %8 = sext i8 %7 to i32
  %9 = sub nsw i32 %8, 1
  %10 = trunc i32 %9 to i8
  %11 = call signext i8 @g(i8 signext %10)
  %12 = sext i8 %11 to i32
  %13 = mul nsw i32 2, %12
  %14 = trunc i32 %13 to i8
  store i8 %14, i8* %1
  br label %15

; <label>:15                                      ; preds = %6, %5
  %16 = load i8, i8* %1
  ret i8 %16
}

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %x = alloca i32, align 4
  call void @foo(i32 0, i32 1, i32 2, i32 3, i32 4)
  store i32 5, i32* %x, align 4
  %1 = load i32, i32* %x, align 4
  %2 = add nsw i32 %1, 0
  %3 = load i32, i32* %x, align 4
  %4 = add nsw i32 %3, 1
  %5 = load i32, i32* %x, align 4
  %6 = add nsw i32 %5, 2
  %7 = load i32, i32* %x, align 4
  %8 = load i32, i32* %x, align 4
  %9 = add nsw i32 %7, %8
  %10 = sub nsw i32 %9, 2
  %11 = load i32, i32* %x, align 4
  %12 = mul nsw i32 %11, 2
  %13 = sub nsw i32 %12, 1
  call void @foo(i32 %2, i32 %4, i32 %6, i32 %10, i32 %13)
  %14 = call i32 @f(i32 0)
  %15 = call i32 @f(i32 0)
  %16 = call i32 @f(i32 %15)
  %17 = call i32 @f(i32 0)
  %18 = call i32 @f(i32 %17)
  %19 = call i32 @f(i32 %18)
  %20 = call signext i8 @g(i8 signext 2)
  %21 = sext i8 %20 to i32
  call void @foo(i32 0, i32 %14, i32 %16, i32 %19, i32 %21)
  %22 = call signext i8 @g(i8 signext 2)
  %23 = sext i8 %22 to i32
  %24 = call signext i8 @g(i8 signext 0)
  %25 = sext i8 %24 to i32
  %26 = add nsw i32 %23, %25
  %27 = call signext i8 @g(i8 signext 2)
  %28 = sext i8 %27 to i32
  %29 = call signext i8 @g(i8 signext 1)
  %30 = sext i8 %29 to i32
  %31 = add nsw i32 %28, %30
  %32 = call signext i8 @g(i8 signext 0)
  %33 = sext i8 %32 to i32
  %34 = call signext i8 @g(i8 signext 1)
  %35 = sext i8 %34 to i32
  %36 = add nsw i32 %33, %35
  %37 = call signext i8 @g(i8 signext 2)
  %38 = sext i8 %37 to i32
  %39 = add nsw i32 %36, %38
  %40 = call signext i8 @g(i8 signext 3)
  %41 = sext i8 %40 to i32
  %42 = call signext i8 @g(i8 signext 4)
  %43 = sext i8 %42 to i32
  %44 = sub nsw i32 %43, 7
  call void @foo(i32 %26, i32 %31, i32 %39, i32 %41, i32 %44)
  ret i32 0
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
