; ModuleID = 'r04.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define void @f(i8* %x) #0 {
  %1 = alloca i8*, align 8
  store i8* %x, i8** %1, align 8
  ret void
}

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %a = alloca [7 x i8], align 1
  %1 = getelementptr inbounds [7 x i8], [7 x i8]* %a, i32 0, i32 0
  call void @f(i8* %1)
  ret i32 0
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
