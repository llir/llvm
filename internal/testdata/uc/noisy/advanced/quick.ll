; ModuleID = 'quick.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@eol = common global [2 x i8] zeroinitializer, align 1
@n = common global i32 0, align 4

; Function Attrs: nounwind uwtable
define void @sort(i8* %a, i32 %l, i32 %r) #0 {
  %1 = alloca i8*, align 8
  %2 = alloca i32, align 4
  %3 = alloca i32, align 4
  %i = alloca i32, align 4
  %j = alloca i32, align 4
  %x = alloca i8, align 1
  %w = alloca i8, align 1
  store i8* %a, i8** %1, align 8
  store i32 %l, i32* %2, align 4
  store i32 %r, i32* %3, align 4
  %4 = load i32, i32* %2, align 4
  store i32 %4, i32* %i, align 4
  %5 = load i32, i32* %3, align 4
  store i32 %5, i32* %j, align 4
  %6 = load i32, i32* %2, align 4
  %7 = load i32, i32* %3, align 4
  %8 = add nsw i32 %6, %7
  %9 = sdiv i32 %8, 2
  %10 = sext i32 %9 to i64
  %11 = load i8*, i8** %1, align 8
  %12 = getelementptr inbounds i8, i8* %11, i64 %10
  %13 = load i8, i8* %12, align 1
  store i8 %13, i8* %x, align 1
  br label %14

; <label>:14                                      ; preds = %74, %0
  %15 = load i32, i32* %i, align 4
  %16 = load i32, i32* %j, align 4
  %17 = icmp sle i32 %15, %16
  br i1 %17, label %18, label %75

; <label>:18                                      ; preds = %14
  br label %19

; <label>:19                                      ; preds = %29, %18
  %20 = load i32, i32* %i, align 4
  %21 = sext i32 %20 to i64
  %22 = load i8*, i8** %1, align 8
  %23 = getelementptr inbounds i8, i8* %22, i64 %21
  %24 = load i8, i8* %23, align 1
  %25 = sext i8 %24 to i32
  %26 = load i8, i8* %x, align 1
  %27 = sext i8 %26 to i32
  %28 = icmp slt i32 %25, %27
  br i1 %28, label %29, label %32

; <label>:29                                      ; preds = %19
  %30 = load i32, i32* %i, align 4
  %31 = add nsw i32 %30, 1
  store i32 %31, i32* %i, align 4
  br label %19

; <label>:32                                      ; preds = %19
  br label %33

; <label>:33                                      ; preds = %43, %32
  %34 = load i8, i8* %x, align 1
  %35 = sext i8 %34 to i32
  %36 = load i32, i32* %j, align 4
  %37 = sext i32 %36 to i64
  %38 = load i8*, i8** %1, align 8
  %39 = getelementptr inbounds i8, i8* %38, i64 %37
  %40 = load i8, i8* %39, align 1
  %41 = sext i8 %40 to i32
  %42 = icmp slt i32 %35, %41
  br i1 %42, label %43, label %46

; <label>:43                                      ; preds = %33
  %44 = load i32, i32* %j, align 4
  %45 = sub nsw i32 %44, 1
  store i32 %45, i32* %j, align 4
  br label %33

; <label>:46                                      ; preds = %33
  %47 = load i32, i32* %i, align 4
  %48 = load i32, i32* %j, align 4
  %49 = icmp sle i32 %47, %48
  br i1 %49, label %50, label %74

; <label>:50                                      ; preds = %46
  %51 = load i32, i32* %i, align 4
  %52 = sext i32 %51 to i64
  %53 = load i8*, i8** %1, align 8
  %54 = getelementptr inbounds i8, i8* %53, i64 %52
  %55 = load i8, i8* %54, align 1
  store i8 %55, i8* %w, align 1
  %56 = load i32, i32* %j, align 4
  %57 = sext i32 %56 to i64
  %58 = load i8*, i8** %1, align 8
  %59 = getelementptr inbounds i8, i8* %58, i64 %57
  %60 = load i8, i8* %59, align 1
  %61 = load i32, i32* %i, align 4
  %62 = sext i32 %61 to i64
  %63 = load i8*, i8** %1, align 8
  %64 = getelementptr inbounds i8, i8* %63, i64 %62
  store i8 %60, i8* %64, align 1
  %65 = load i8, i8* %w, align 1
  %66 = load i32, i32* %j, align 4
  %67 = sext i32 %66 to i64
  %68 = load i8*, i8** %1, align 8
  %69 = getelementptr inbounds i8, i8* %68, i64 %67
  store i8 %65, i8* %69, align 1
  %70 = load i32, i32* %i, align 4
  %71 = add nsw i32 %70, 1
  store i32 %71, i32* %i, align 4
  %72 = load i32, i32* %j, align 4
  %73 = sub nsw i32 %72, 1
  store i32 %73, i32* %j, align 4
  br label %74

; <label>:74                                      ; preds = %50, %46
  br label %14

; <label>:75                                      ; preds = %14
  %76 = load i8*, i8** %1, align 8
  call void @putstring(i8* %76)
  call void @putstring(i8* getelementptr inbounds ([2 x i8], [2 x i8]* @eol, i32 0, i32 0))
  %77 = load i32, i32* %2, align 4
  %78 = load i32, i32* %j, align 4
  %79 = icmp slt i32 %77, %78
  br i1 %79, label %80, label %84

; <label>:80                                      ; preds = %75
  %81 = load i8*, i8** %1, align 8
  %82 = load i32, i32* %2, align 4
  %83 = load i32, i32* %j, align 4
  call void @sort(i8* %81, i32 %82, i32 %83)
  br label %84

; <label>:84                                      ; preds = %80, %75
  %85 = load i32, i32* %i, align 4
  %86 = load i32, i32* %3, align 4
  %87 = icmp slt i32 %85, %86
  br i1 %87, label %88, label %92

; <label>:88                                      ; preds = %84
  %89 = load i8*, i8** %1, align 8
  %90 = load i32, i32* %i, align 4
  %91 = load i32, i32* %3, align 4
  call void @sort(i8* %89, i32 %90, i32 %91)
  br label %92

; <label>:92                                      ; preds = %88, %84
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
  store i8 10, i8* getelementptr inbounds ([2 x i8], [2 x i8]* @eol, i32 0, i64 0), align 1
  store i8 0, i8* getelementptr inbounds ([2 x i8], [2 x i8]* @eol, i32 0, i64 1), align 1
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
  call void @putstring(i8* %28)
  call void @putstring(i8* getelementptr inbounds ([2 x i8], [2 x i8]* @eol, i32 0, i32 0))
  %29 = getelementptr inbounds [27 x i8], [27 x i8]* %s, i32 0, i32 0
  %30 = load i32, i32* @n, align 4
  %31 = sub nsw i32 %30, 1
  call void @sort(i8* %29, i32 0, i32 %31)
  %32 = getelementptr inbounds [27 x i8], [27 x i8]* %s, i32 0, i32 0
  call void @putstring(i8* %32)
  call void @putstring(i8* getelementptr inbounds ([2 x i8], [2 x i8]* @eol, i32 0, i32 0))
  %33 = load i32, i32* %1
  ret i32 %33
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
