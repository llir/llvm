; ModuleID = '8queens.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@n = common global i32 0, align 4
@board = common global [8 x i32] zeroinitializer, align 16

; Function Attrs: nounwind uwtable
define void @printboard(i32* %board) #0 {
  %1 = alloca i32*, align 8
  %i = alloca i32, align 4
  store i32* %board, i32** %1, align 8
  store i32 0, i32* %i, align 4
  br label %2

; <label>:2                                       ; preds = %6, %0
  %3 = load i32, i32* %i, align 4
  %4 = load i32, i32* @n, align 4
  %5 = icmp slt i32 %3, %4
  br i1 %5, label %6, label %14

; <label>:6                                       ; preds = %2
  %7 = load i32, i32* %i, align 4
  %8 = sext i32 %7 to i64
  %9 = load i32*, i32** %1, align 8
  %10 = getelementptr inbounds i32, i32* %9, i64 %8
  %11 = load i32, i32* %10, align 4
  call void @putint(i32 %11)
  %12 = load i32, i32* %i, align 4
  %13 = add nsw i32 %12, 1
  store i32 %13, i32* %i, align 4
  br label %2

; <label>:14                                      ; preds = %2
  ret void
}

declare void @putint(i32) #1

; Function Attrs: nounwind uwtable
define i32 @check(i32 %col, i32 %row) #0 {
  %1 = alloca i32, align 4
  %2 = alloca i32, align 4
  %3 = alloca i32, align 4
  %i = alloca i32, align 4
  %j = alloca i32, align 4
  store i32 %col, i32* %2, align 4
  store i32 %row, i32* %3, align 4
  %4 = load i32, i32* %2, align 4
  %5 = sub nsw i32 %4, 1
  store i32 %5, i32* %i, align 4
  br label %6

; <label>:6                                       ; preds = %40, %0
  %7 = load i32, i32* %i, align 4
  %8 = icmp sge i32 %7, 0
  br i1 %8, label %9, label %43

; <label>:9                                       ; preds = %6
  %10 = load i32, i32* %i, align 4
  %11 = sext i32 %10 to i64
  %12 = getelementptr inbounds [8 x i32], [8 x i32]* @board, i32 0, i64 %11
  %13 = load i32, i32* %12, align 4
  store i32 %13, i32* %j, align 4
  %14 = load i32, i32* %j, align 4
  %15 = load i32, i32* %3, align 4
  %16 = icmp eq i32 %14, %15
  br i1 %16, label %17, label %18

; <label>:17                                      ; preds = %9
  store i32 0, i32* %1
  br label %44

; <label>:18                                      ; preds = %9
  %19 = load i32, i32* %j, align 4
  %20 = load i32, i32* %3, align 4
  %21 = icmp sgt i32 %19, %20
  br i1 %21, label %22, label %31

; <label>:22                                      ; preds = %18
  %23 = load i32, i32* %2, align 4
  %24 = load i32, i32* %i, align 4
  %25 = sub nsw i32 %23, %24
  %26 = load i32, i32* %j, align 4
  %27 = load i32, i32* %3, align 4
  %28 = sub nsw i32 %26, %27
  %29 = icmp eq i32 %25, %28
  br i1 %29, label %30, label %31

; <label>:30                                      ; preds = %22
  store i32 0, i32* %1
  br label %44

; <label>:31                                      ; preds = %22, %18
  %32 = load i32, i32* %2, align 4
  %33 = load i32, i32* %i, align 4
  %34 = sub nsw i32 %32, %33
  %35 = load i32, i32* %3, align 4
  %36 = load i32, i32* %j, align 4
  %37 = sub nsw i32 %35, %36
  %38 = icmp eq i32 %34, %37
  br i1 %38, label %39, label %40

; <label>:39                                      ; preds = %31
  store i32 0, i32* %1
  br label %44

; <label>:40                                      ; preds = %31
  %41 = load i32, i32* %i, align 4
  %42 = sub nsw i32 %41, 1
  store i32 %42, i32* %i, align 4
  br label %6

; <label>:43                                      ; preds = %6
  store i32 1, i32* %1
  br label %44

; <label>:44                                      ; preds = %43, %39, %30, %17
  %45 = load i32, i32* %1
  ret i32 %45
}

; Function Attrs: nounwind uwtable
define i32 @queen(i32 %col, i32 %row) #0 {
  %1 = alloca i32, align 4
  %2 = alloca i32, align 4
  %3 = alloca i32, align 4
  store i32 %col, i32* %2, align 4
  store i32 %row, i32* %3, align 4
  %4 = load i32, i32* %2, align 4
  %5 = load i32, i32* @n, align 4
  %6 = icmp sge i32 %4, %5
  br i1 %6, label %7, label %8

; <label>:7                                       ; preds = %0
  store i32 1, i32* %1
  br label %32

; <label>:8                                       ; preds = %0
  br label %9

; <label>:9                                       ; preds = %28, %8
  %10 = load i32, i32* %3, align 4
  %11 = load i32, i32* @n, align 4
  %12 = icmp slt i32 %10, %11
  br i1 %12, label %13, label %31

; <label>:13                                      ; preds = %9
  %14 = load i32, i32* %3, align 4
  %15 = load i32, i32* %2, align 4
  %16 = sext i32 %15 to i64
  %17 = getelementptr inbounds [8 x i32], [8 x i32]* @board, i32 0, i64 %16
  store i32 %14, i32* %17, align 4
  %18 = load i32, i32* %2, align 4
  %19 = load i32, i32* %3, align 4
  %20 = call i32 @check(i32 %18, i32 %19)
  %21 = icmp ne i32 %20, 0
  br i1 %21, label %22, label %28

; <label>:22                                      ; preds = %13
  %23 = load i32, i32* %2, align 4
  %24 = add nsw i32 %23, 1
  %25 = call i32 @queen(i32 %24, i32 0)
  %26 = icmp ne i32 %25, 0
  br i1 %26, label %27, label %28

; <label>:27                                      ; preds = %22
  store i32 1, i32* %1
  br label %32

; <label>:28                                      ; preds = %22, %13
  %29 = load i32, i32* %3, align 4
  %30 = add nsw i32 %29, 1
  store i32 %30, i32* %3, align 4
  br label %9

; <label>:31                                      ; preds = %9
  store i32 0, i32* %1
  br label %32

; <label>:32                                      ; preds = %31, %27, %7
  %33 = load i32, i32* %1
  ret i32 %33
}

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  store i32 8, i32* @n, align 4
  %1 = call i32 @queen(i32 0, i32 0)
  call void @printboard(i32* getelementptr inbounds ([8 x i32], [8 x i32]* @board, i32 0, i32 0))
  ret i32 0
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
