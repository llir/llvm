; ModuleID = 'r05.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@a = common global [10 x i32] zeroinitializer, align 16

; Function Attrs: nounwind uwtable
define void @f(i32* %x) #0 {
  %1 = alloca i32*, align 8
  store i32* %x, i32** %1, align 8
  %2 = load i32*, i32** %1, align 8
  %3 = getelementptr inbounds i32, i32* %2, i64 5
  %4 = load i32, i32* %3, align 4
  %5 = add nsw i32 %4, 7
  %6 = load i32*, i32** %1, align 8
  %7 = getelementptr inbounds i32, i32* %6, i64 3
  store i32 %5, i32* %7, align 4
  ret void
}

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  call void @f(i32* getelementptr inbounds ([10 x i32], [10 x i32]* @a, i32 0, i32 0))
  ret i32 0
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
