; ModuleID = 'bubble.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@eol = common global [1 x i8] zeroinitializer, align 1
@n = common global i32 0, align 4

; Function Attrs: nounwind uwtable
define void @bubble(i8* %a) #0 {
  %1 = alloca i8*, align 8
  %i = alloca i32, align 4
  %j = alloca i32, align 4
  %t = alloca i8, align 1
  store i8* %a, i8** %1, align 8
  %2 = load i8*, i8** %1, align 8
  call void @putstring(i8* %2)
  call void @putstring(i8* getelementptr inbounds ([1 x i8], [1 x i8]* @eol, i32 0, i32 0))
  %3 = load i32, i32* @n, align 4
  %4 = sub nsw i32 %3, 1
  store i32 %4, i32* %i, align 4
  br label %5

; <label>:5                                       ; preds = %53, %0
  %6 = load i32, i32* %i, align 4
  %7 = icmp sgt i32 %6, 0
  br i1 %7, label %8, label %57

; <label>:8                                       ; preds = %5
  store i32 0, i32* %j, align 4
  br label %9

; <label>:9                                       ; preds = %50, %8
  %10 = load i32, i32* %j, align 4
  %11 = load i32, i32* %i, align 4
  %12 = icmp slt i32 %10, %11
  br i1 %12, label %13, label %53

; <label>:13                                      ; preds = %9
  %14 = load i32, i32* %j, align 4
  %15 = sext i32 %14 to i64
  %16 = load i8*, i8** %1, align 8
  %17 = getelementptr inbounds i8, i8* %16, i64 %15
  %18 = load i8, i8* %17, align 1
  %19 = sext i8 %18 to i32
  %20 = load i32, i32* %j, align 4
  %21 = add nsw i32 %20, 1
  %22 = sext i32 %21 to i64
  %23 = load i8*, i8** %1, align 8
  %24 = getelementptr inbounds i8, i8* %23, i64 %22
  %25 = load i8, i8* %24, align 1
  %26 = sext i8 %25 to i32
  %27 = icmp sgt i32 %19, %26
  br i1 %27, label %28, label %50

; <label>:28                                      ; preds = %13
  %29 = load i32, i32* %j, align 4
  %30 = sext i32 %29 to i64
  %31 = load i8*, i8** %1, align 8
  %32 = getelementptr inbounds i8, i8* %31, i64 %30
  %33 = load i8, i8* %32, align 1
  store i8 %33, i8* %t, align 1
  %34 = load i32, i32* %j, align 4
  %35 = add nsw i32 %34, 1
  %36 = sext i32 %35 to i64
  %37 = load i8*, i8** %1, align 8
  %38 = getelementptr inbounds i8, i8* %37, i64 %36
  %39 = load i8, i8* %38, align 1
  %40 = load i32, i32* %j, align 4
  %41 = sext i32 %40 to i64
  %42 = load i8*, i8** %1, align 8
  %43 = getelementptr inbounds i8, i8* %42, i64 %41
  store i8 %39, i8* %43, align 1
  %44 = load i8, i8* %t, align 1
  %45 = load i32, i32* %j, align 4
  %46 = add nsw i32 %45, 1
  %47 = sext i32 %46 to i64
  %48 = load i8*, i8** %1, align 8
  %49 = getelementptr inbounds i8, i8* %48, i64 %47
  store i8 %44, i8* %49, align 1
  br label %50

; <label>:50                                      ; preds = %28, %13
  %51 = load i32, i32* %j, align 4
  %52 = add nsw i32 %51, 1
  store i32 %52, i32* %j, align 4
  br label %9

; <label>:53                                      ; preds = %9
  %54 = load i8*, i8** %1, align 8
  call void @putstring(i8* %54)
  call void @putstring(i8* getelementptr inbounds ([1 x i8], [1 x i8]* @eol, i32 0, i32 0))
  %55 = load i32, i32* %i, align 4
  %56 = sub nsw i32 %55, 1
  store i32 %56, i32* %i, align 4
  br label %5

; <label>:57                                      ; preds = %5
  ret void
}

declare void @putstring(i8*) #1

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %1 = alloca i32, align 4
  %s = alloca [27 x i8], align 16
  %i = alloca i32, align 4
  %t = alloca i8, align 1
  %q = alloca i32, align 4
  store i32 0, i32* %1
  store i8 10, i8* getelementptr inbounds ([1 x i8], [1 x i8]* @eol, i32 0, i64 0), align 1
  store i8 0, i8* getelementptr inbounds ([1 x i8], [1 x i8]* @eol, i64 1, i64 0), align 1
  store i32 26, i32* @n, align 4
  %2 = load i32, i32* @n, align 4
  %3 = sext i32 %2 to i64
  %4 = getelementptr inbounds [27 x i8], [27 x i8]* %s, i32 0, i64 %3
  store i8 0, i8* %4, align 1
  store i32 0, i32* %i, align 4
  store i32 11, i32* %q, align 4
  br label %5

; <label>:5                                       ; preds = %9, %0
  %6 = load i32, i32* %i, align 4
  %7 = load i32, i32* @n, align 4
  %8 = icmp slt i32 %6, %7
  br i1 %8, label %9, label %27

; <label>:9                                       ; preds = %5
  %10 = load i32, i32* %q, align 4
  %11 = load i32, i32* %q, align 4
  %12 = sdiv i32 %11, 26
  %13 = mul nsw i32 %12, 26
  %14 = sub nsw i32 %10, %13
  %15 = trunc i32 %14 to i8
  store i8 %15, i8* %t, align 1
  %16 = load i8, i8* %t, align 1
  %17 = sext i8 %16 to i32
  %18 = add nsw i32 97, %17
  %19 = trunc i32 %18 to i8
  %20 = load i32, i32* %i, align 4
  %21 = sext i32 %20 to i64
  %22 = getelementptr inbounds [27 x i8], [27 x i8]* %s, i32 0, i64 %21
  store i8 %19, i8* %22, align 1
  %23 = load i32, i32* %i, align 4
  %24 = add nsw i32 %23, 1
  store i32 %24, i32* %i, align 4
  %25 = load i32, i32* %q, align 4
  %26 = add nsw i32 %25, 17
  store i32 %26, i32* %q, align 4
  br label %5

; <label>:27                                      ; preds = %5
  %28 = getelementptr inbounds [27 x i8], [27 x i8]* %s, i32 0, i32 0
  call void @bubble(i8* %28)
  %29 = load i32, i32* %1
  ret i32 %29
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
