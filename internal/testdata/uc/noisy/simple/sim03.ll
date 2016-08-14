; ModuleID = 'sim03.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@a = common global [10 x i32] zeroinitializer, align 16

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %b = alloca [10 x i32], align 16
  store i32 123, i32* getelementptr inbounds ([10 x i32], [10 x i32]* @a, i32 0, i64 7), align 4
  %1 = getelementptr inbounds [10 x i32], [10 x i32]* %b, i32 0, i64 5
  store i32 456, i32* %1, align 4
  %2 = load i32, i32* getelementptr inbounds ([10 x i32], [10 x i32]* @a, i32 0, i64 7), align 4
  call void @putint(i32 %2)
  %3 = getelementptr inbounds [10 x i32], [10 x i32]* %b, i32 0, i64 5
  %4 = load i32, i32* %3, align 4
  call void @putint(i32 %4)
  ret i32 0
}

declare void @putint(i32) #1

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
