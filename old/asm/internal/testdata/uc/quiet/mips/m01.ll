; ModuleID = 'm01.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@b = common global i32 0, align 4
@addi = common global i8 0, align 1
@la = common global i32 0, align 4

; Function Attrs: nounwind uwtable
define void @jal() #0 {
  %1 = load i32, i32* @b, align 4
  %2 = load i8, i8* @addi, align 1
  %3 = sext i8 %2 to i32
  %4 = mul nsw i32 %1, %3
  %5 = load i32, i32* @la, align 4
  %6 = add nsw i32 %4, %5
  store i32 %6, i32* @b, align 4
  ret void
}

; Function Attrs: nounwind uwtable
define i32 @mov(i32 %lb) #0 {
  %1 = alloca i32, align 4
  %2 = alloca i32, align 4
  store i32 %lb, i32* %2, align 4
  %3 = load i32, i32* %2, align 4
  %4 = trunc i32 %3 to i8
  store i8 %4, i8* @addi, align 1
  %5 = load i32, i32* %1
  ret i32 %5
}

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  store i32 8, i32* @la, align 4
  call void @jal()
  %1 = load i32, i32* @la, align 4
  %2 = call i32 @mov(i32 %1)
  ret i32 0
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
