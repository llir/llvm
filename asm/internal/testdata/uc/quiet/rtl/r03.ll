; ModuleID = 'r03.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@a = common global [10 x i32] zeroinitializer, align 16

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %b = alloca [10 x i8], align 1
  %1 = load i32, i32* getelementptr inbounds ([10 x i32], [10 x i32]* @a, i32 0, i64 5), align 4
  %2 = add nsw i32 %1, 7
  store i32 %2, i32* getelementptr inbounds ([10 x i32], [10 x i32]* @a, i32 0, i64 3), align 4
  %3 = getelementptr inbounds [10 x i8], [10 x i8]* %b, i32 0, i64 5
  %4 = load i8, i8* %3, align 1
  %5 = sext i8 %4 to i32
  %6 = add nsw i32 %5, 7
  %7 = trunc i32 %6 to i8
  %8 = getelementptr inbounds [10 x i8], [10 x i8]* %b, i32 0, i64 3
  store i8 %7, i8* %8, align 1
  ret i32 0
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
