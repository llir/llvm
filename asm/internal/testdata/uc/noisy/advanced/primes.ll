; ModuleID = 'primes.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@max = common global i32 0, align 4
@notprime = common global [1001 x i8] zeroinitializer, align 16

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %1 = alloca i32, align 4
  %i = alloca i32, align 4
  %j = alloca i32, align 4
  %cr = alloca [2 x i8], align 1
  %space = alloca [2 x i8], align 1
  store i32 0, i32* %1
  %2 = getelementptr inbounds [2 x i8], [2 x i8]* %cr, i32 0, i64 0
  store i8 10, i8* %2, align 1
  %3 = getelementptr inbounds [2 x i8], [2 x i8]* %cr, i32 0, i64 1
  store i8 0, i8* %3, align 1
  %4 = getelementptr inbounds [2 x i8], [2 x i8]* %space, i32 0, i64 0
  store i8 32, i8* %4, align 1
  %5 = getelementptr inbounds [2 x i8], [2 x i8]* %space, i32 0, i64 1
  store i8 0, i8* %5, align 1
  store i32 1001, i32* @max, align 4
  store i32 2, i32* %i, align 4
  br label %6

; <label>:6                                       ; preds = %32, %0
  %7 = load i32, i32* %i, align 4
  %8 = load i32, i32* @max, align 4
  %9 = icmp slt i32 %7, %8
  br i1 %9, label %10, label %35

; <label>:10                                      ; preds = %6
  %11 = load i32, i32* %i, align 4
  %12 = sext i32 %11 to i64
  %13 = getelementptr inbounds [1001 x i8], [1001 x i8]* @notprime, i32 0, i64 %12
  %14 = load i8, i8* %13, align 1
  %15 = icmp ne i8 %14, 0
  br i1 %15, label %32, label %16

; <label>:16                                      ; preds = %10
  %17 = load i32, i32* %i, align 4
  %18 = load i32, i32* %i, align 4
  %19 = add nsw i32 %17, %18
  store i32 %19, i32* %j, align 4
  br label %20

; <label>:20                                      ; preds = %24, %16
  %21 = load i32, i32* %j, align 4
  %22 = load i32, i32* @max, align 4
  %23 = icmp slt i32 %21, %22
  br i1 %23, label %24, label %31

; <label>:24                                      ; preds = %20
  %25 = load i32, i32* %j, align 4
  %26 = sext i32 %25 to i64
  %27 = getelementptr inbounds [1001 x i8], [1001 x i8]* @notprime, i32 0, i64 %26
  store i8 1, i8* %27, align 1
  %28 = load i32, i32* %j, align 4
  %29 = load i32, i32* %i, align 4
  %30 = add nsw i32 %28, %29
  store i32 %30, i32* %j, align 4
  br label %20

; <label>:31                                      ; preds = %20
  br label %32

; <label>:32                                      ; preds = %31, %10
  %33 = load i32, i32* %i, align 4
  %34 = add nsw i32 %33, 1
  store i32 %34, i32* %i, align 4
  br label %6

; <label>:35                                      ; preds = %6
  store i32 2, i32* %i, align 4
  br label %36

; <label>:36                                      ; preds = %61, %35
  %37 = load i32, i32* %i, align 4
  %38 = add nsw i32 %37, 10
  %39 = load i32, i32* @max, align 4
  %40 = icmp slt i32 %38, %39
  br i1 %40, label %41, label %64

; <label>:41                                      ; preds = %36
  %42 = load i32, i32* %i, align 4
  store i32 %42, i32* %j, align 4
  %43 = getelementptr inbounds [2 x i8], [2 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %43)
  br label %44

; <label>:44                                      ; preds = %58, %41
  %45 = load i32, i32* %j, align 4
  %46 = load i32, i32* %i, align 4
  %47 = add nsw i32 %46, 10
  %48 = icmp slt i32 %45, %47
  br i1 %48, label %49, label %61

; <label>:49                                      ; preds = %44
  %50 = load i32, i32* %j, align 4
  %51 = sext i32 %50 to i64
  %52 = getelementptr inbounds [1001 x i8], [1001 x i8]* @notprime, i32 0, i64 %51
  %53 = load i8, i8* %52, align 1
  %54 = icmp ne i8 %53, 0
  br i1 %54, label %58, label %55

; <label>:55                                      ; preds = %49
  %56 = load i32, i32* %j, align 4
  call void @putint(i32 %56)
  %57 = getelementptr inbounds [2 x i8], [2 x i8]* %space, i32 0, i32 0
  call void @putstring(i8* %57)
  br label %58

; <label>:58                                      ; preds = %55, %49
  %59 = load i32, i32* %j, align 4
  %60 = add nsw i32 %59, 1
  store i32 %60, i32* %j, align 4
  br label %44

; <label>:61                                      ; preds = %44
  %62 = load i32, i32* %i, align 4
  %63 = add nsw i32 %62, 10
  store i32 %63, i32* %i, align 4
  br label %36

; <label>:64                                      ; preds = %36
  %65 = load i32, i32* %1
  ret i32 %65
}

declare void @putstring(i8*) #1

declare void @putint(i32) #1

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
