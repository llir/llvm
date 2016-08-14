; ModuleID = 'p05.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@c = common global [10 x i32] zeroinitializer, align 16
@d = common global [10 x i8] zeroinitializer, align 1

; Function Attrs: nounwind uwtable
define void @f(i32* %h, i8* %i) #0 {
  %1 = alloca i32*, align 8
  %2 = alloca i8*, align 8
  store i32* %h, i32** %1, align 8
  store i8* %i, i8** %2, align 8
  ret void
}

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  ret i32 0
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
