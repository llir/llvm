; ModuleID = 's05.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %s = alloca [27 x i8], align 16
  %t = alloca i8, align 1
  %1 = load i8, i8* %t, align 1
  %2 = sext i8 %1 to i32
  %3 = add nsw i32 97, %2
  %4 = trunc i32 %3 to i8
  %5 = getelementptr inbounds [27 x i8], [27 x i8]* %s, i32 0, i64 0
  store i8 %4, i8* %5, align 1
  ret i32 0
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
