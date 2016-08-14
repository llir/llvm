; ModuleID = 'circle.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define void @drawpos(i8 signext %c) #0 {
  %1 = alloca i8, align 1
  %s = alloca [2 x i8], align 1
  store i8 %c, i8* %1, align 1
  %2 = load i8, i8* %1, align 1
  %3 = icmp ne i8 %2, 0
  br i1 %3, label %4, label %6

; <label>:4                                       ; preds = %0
  %5 = getelementptr inbounds [2 x i8], [2 x i8]* %s, i32 0, i64 0
  store i8 35, i8* %5, align 1
  br label %8

; <label>:6                                       ; preds = %0
  %7 = getelementptr inbounds [2 x i8], [2 x i8]* %s, i32 0, i64 0
  store i8 32, i8* %7, align 1
  br label %8

; <label>:8                                       ; preds = %6, %4
  %9 = getelementptr inbounds [2 x i8], [2 x i8]* %s, i32 0, i64 1
  store i8 0, i8* %9, align 1
  %10 = getelementptr inbounds [2 x i8], [2 x i8]* %s, i32 0, i32 0
  call void @putstring(i8* %10)
  ret void
}

declare void @putstring(i8*) #1

; Function Attrs: nounwind uwtable
define void @nl() #0 {
  %cr = alloca [2 x i8], align 1
  %1 = getelementptr inbounds [2 x i8], [2 x i8]* %cr, i32 0, i64 0
  store i8 10, i8* %1, align 1
  %2 = getelementptr inbounds [2 x i8], [2 x i8]* %cr, i32 0, i64 1
  store i8 0, i8* %2, align 1
  %3 = getelementptr inbounds [2 x i8], [2 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %3)
  ret void
}

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %1 = alloca i32, align 4
  %i = alloca i32, align 4
  %j = alloca i32, align 4
  store i32 0, i32* %1
  store i32 -9, i32* %i, align 4
  br label %2

; <label>:2                                       ; preds = %25, %0
  %3 = load i32, i32* %i, align 4
  %4 = icmp sle i32 %3, 9
  br i1 %4, label %5, label %28

; <label>:5                                       ; preds = %2
  store i32 -20, i32* %j, align 4
  call void @nl()
  br label %6

; <label>:6                                       ; preds = %9, %5
  %7 = load i32, i32* %j, align 4
  %8 = icmp sle i32 %7, 20
  br i1 %8, label %9, label %25

; <label>:9                                       ; preds = %6
  %10 = load i32, i32* %i, align 4
  %11 = load i32, i32* %i, align 4
  %12 = mul nsw i32 %10, %11
  %13 = mul nsw i32 %12, 22
  %14 = mul nsw i32 %13, 22
  %15 = sdiv i32 %14, 100
  %16 = load i32, i32* %j, align 4
  %17 = load i32, i32* %j, align 4
  %18 = mul nsw i32 %16, %17
  %19 = add nsw i32 %15, %18
  %20 = icmp sgt i32 %19, 380
  %21 = zext i1 %20 to i32
  %22 = trunc i32 %21 to i8
  call void @drawpos(i8 signext %22)
  %23 = load i32, i32* %j, align 4
  %24 = add nsw i32 %23, 1
  store i32 %24, i32* %j, align 4
  br label %6

; <label>:25                                      ; preds = %6
  %26 = load i32, i32* %i, align 4
  %27 = add nsw i32 %26, 1
  store i32 %27, i32* %i, align 4
  br label %2

; <label>:28                                      ; preds = %2
  call void @nl()
  call void @nl()
  %29 = load i32, i32* %1
  ret i32 %29
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
