; ModuleID = 'sim07.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %1 = alloca i32, align 4
  %i = alloca i32, align 4
  %space = alloca [2 x i8], align 1
  %X = alloca [2 x i8], align 1
  %Y = alloca [2 x i8], align 1
  %W = alloca [2 x i8], align 1
  %nl = alloca [2 x i8], align 1
  store i32 0, i32* %1
  %2 = getelementptr inbounds [2 x i8], [2 x i8]* %space, i32 0, i64 0
  store i8 32, i8* %2, align 1
  %3 = getelementptr inbounds [2 x i8], [2 x i8]* %space, i32 0, i64 1
  store i8 0, i8* %3, align 1
  %4 = getelementptr inbounds [2 x i8], [2 x i8]* %X, i32 0, i64 0
  store i8 88, i8* %4, align 1
  %5 = getelementptr inbounds [2 x i8], [2 x i8]* %X, i32 0, i64 1
  store i8 0, i8* %5, align 1
  %6 = getelementptr inbounds [2 x i8], [2 x i8]* %Y, i32 0, i64 0
  store i8 89, i8* %6, align 1
  %7 = getelementptr inbounds [2 x i8], [2 x i8]* %Y, i32 0, i64 1
  store i8 0, i8* %7, align 1
  %8 = getelementptr inbounds [2 x i8], [2 x i8]* %W, i32 0, i64 0
  store i8 87, i8* %8, align 1
  %9 = getelementptr inbounds [2 x i8], [2 x i8]* %W, i32 0, i64 1
  store i8 0, i8* %9, align 1
  %10 = getelementptr inbounds [2 x i8], [2 x i8]* %nl, i32 0, i64 0
  store i8 10, i8* %10, align 1
  %11 = getelementptr inbounds [2 x i8], [2 x i8]* %nl, i32 0, i64 1
  store i8 0, i8* %11, align 1
  store i32 0, i32* %i, align 4
  br label %12

; <label>:12                                      ; preds = %39, %0
  %13 = load i32, i32* %i, align 4
  %14 = icmp ne i32 %13, 21
  br i1 %14, label %15, label %43

; <label>:15                                      ; preds = %12
  %16 = load i32, i32* %i, align 4
  call void @putint(i32 %16)
  %17 = getelementptr inbounds [2 x i8], [2 x i8]* %space, i32 0, i32 0
  call void @putstring(i8* %17)
  %18 = load i32, i32* %i, align 4
  %19 = sdiv i32 %18, 2
  %20 = mul nsw i32 %19, 2
  %21 = load i32, i32* %i, align 4
  %22 = icmp eq i32 %20, %21
  br i1 %22, label %23, label %25

; <label>:23                                      ; preds = %15
  %24 = getelementptr inbounds [2 x i8], [2 x i8]* %X, i32 0, i32 0
  call void @putstring(i8* %24)
  br label %25

; <label>:25                                      ; preds = %23, %15
  %26 = load i32, i32* %i, align 4
  %27 = sdiv i32 %26, 3
  %28 = mul nsw i32 %27, 3
  %29 = load i32, i32* %i, align 4
  %30 = icmp eq i32 %28, %29
  br i1 %30, label %31, label %33

; <label>:31                                      ; preds = %25
  %32 = getelementptr inbounds [2 x i8], [2 x i8]* %Y, i32 0, i32 0
  call void @putstring(i8* %32)
  br label %39

; <label>:33                                      ; preds = %25
  %34 = load i32, i32* %i, align 4
  %35 = icmp sgt i32 %34, 10
  br i1 %35, label %36, label %38

; <label>:36                                      ; preds = %33
  %37 = getelementptr inbounds [2 x i8], [2 x i8]* %W, i32 0, i32 0
  call void @putstring(i8* %37)
  br label %38

; <label>:38                                      ; preds = %36, %33
  br label %39

; <label>:39                                      ; preds = %38, %31
  %40 = getelementptr inbounds [2 x i8], [2 x i8]* %nl, i32 0, i32 0
  call void @putstring(i8* %40)
  %41 = load i32, i32* %i, align 4
  %42 = add nsw i32 %41, 1
  store i32 %42, i32* %i, align 4
  br label %12

; <label>:43                                      ; preds = %12
  %44 = load i32, i32* %1
  ret i32 %44
}

declare void @putint(i32) #1

declare void @putstring(i8*) #1

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
