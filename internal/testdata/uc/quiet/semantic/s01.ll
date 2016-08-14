; ModuleID = 's01.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@x = common global i32 0, align 4
@y = common global i8 0, align 1

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %1 = alloca i32, align 4
  %z = alloca i32, align 4
  %w = alloca i8, align 1
  store i32 0, i32* %1
  %2 = load i32, i32* @x, align 4
  %3 = load i8, i8* @y, align 1
  %4 = sext i8 %3 to i32
  %5 = add nsw i32 %2, %4
  %6 = load i32, i32* %z, align 4
  %7 = add nsw i32 %5, %6
  %8 = load i8, i8* %w, align 1
  %9 = sext i8 %8 to i32
  %10 = add nsw i32 %7, %9
  store i32 %10, i32* @x, align 4
  store i32 42, i32* %z, align 4
  store i32 42, i32* @x, align 4
  %11 = load i32, i32* @x, align 4
  %12 = load i32, i32* %z, align 4
  %13 = icmp eq i32 %11, %12
  %14 = zext i1 %13 to i32
  %15 = icmp eq i32 %14, 42
  %16 = zext i1 %15 to i32
  %17 = load i32, i32* @x, align 4
  store i32 99, i32* %z, align 4
  %18 = icmp eq i32 %17, 99
  %19 = zext i1 %18 to i32
  br label %20

; <label>:20                                      ; preds = %23, %0
  %21 = load i32, i32* @x, align 4
  %22 = icmp ne i32 %21, 0
  br i1 %22, label %23, label %24

; <label>:23                                      ; preds = %20
  store i32 0, i32* @x, align 4
  br label %20

; <label>:24                                      ; preds = %20
  store i8 4, i8* @y, align 1
  %25 = load i32, i32* @x, align 4
  %26 = load i8, i8* @y, align 1
  %27 = sext i8 %26 to i32
  %28 = icmp sgt i32 %25, %27
  %29 = zext i1 %28 to i32
  %30 = trunc i32 %29 to i8
  store i8 %30, i8* %w, align 1
  %31 = load i32, i32* @x, align 4
  %32 = icmp slt i32 0, %31
  %33 = zext i1 %32 to i32
  %34 = icmp slt i32 %33, 10
  %35 = zext i1 %34 to i32
  %36 = trunc i32 %35 to i8
  store i8 %36, i8* @y, align 1
  %37 = load i32, i32* %1
  ret i32 %37
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
