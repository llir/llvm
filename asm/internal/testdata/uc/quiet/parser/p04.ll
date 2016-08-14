; ModuleID = 'p04.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %x = alloca i32, align 4
  %y = alloca i32, align 4
  %z = alloca i32, align 4
  %1 = load i32, i32* %x, align 4
  %2 = load i32, i32* %y, align 4
  %3 = sub nsw i32 %1, %2
  %4 = load i32, i32* %z, align 4
  %5 = sub nsw i32 %3, %4
  %6 = sub nsw i32 %5, 42
  %7 = load i32, i32* %x, align 4
  %8 = icmp ne i32 %7, 0
  %9 = xor i1 %8, true
  %10 = zext i1 %9 to i32
  %11 = load i32, i32* %y, align 4
  %12 = mul nsw i32 %10, %11
  %13 = load i32, i32* %z, align 4
  %14 = add nsw i32 %12, %13
  %15 = load i32, i32* %x, align 4
  %16 = icmp slt i32 %14, %15
  %17 = zext i1 %16 to i32
  %18 = load i32, i32* %x, align 4
  %19 = load i32, i32* %y, align 4
  %20 = load i32, i32* %x, align 4
  %21 = icmp ne i32 %20, 0
  %22 = xor i1 %21, true
  %23 = zext i1 %22 to i32
  %24 = mul nsw i32 %19, %23
  %25 = add nsw i32 %18, %24
  %26 = icmp slt i32 42, %25
  %27 = zext i1 %26 to i32
  %28 = icmp ne i32 %17, %27
  %29 = zext i1 %28 to i32
  ret i32 0
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
