; ModuleID = 'sim06.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@t = common global [2 x i8] zeroinitializer, align 1

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %1 = alloca i32, align 4
  %b = alloca i32, align 4
  store i32 0, i32* %1
  store i32 10, i32* %b, align 4
  store i8 0, i8* getelementptr inbounds ([2 x i8], [2 x i8]* @t, i32 0, i64 1), align 1
  br label %2

; <label>:2                                       ; preds = %5, %0
  %3 = load i32, i32* %b, align 4
  %4 = icmp ne i32 %3, 0
  br i1 %4, label %5, label %12

; <label>:5                                       ; preds = %2
  %6 = load i32, i32* %b, align 4
  %7 = add nsw i32 48, %6
  %8 = sub nsw i32 %7, 1
  %9 = trunc i32 %8 to i8
  store i8 %9, i8* getelementptr inbounds ([2 x i8], [2 x i8]* @t, i32 0, i64 0), align 1
  call void @putstring(i8* getelementptr inbounds ([2 x i8], [2 x i8]* @t, i32 0, i32 0))
  %10 = load i32, i32* %b, align 4
  %11 = sub nsw i32 %10, 1
  store i32 %11, i32* %b, align 4
  br label %2

; <label>:12                                      ; preds = %2
  store i8 10, i8* getelementptr inbounds ([2 x i8], [2 x i8]* @t, i32 0, i64 0), align 1
  call void @putstring(i8* getelementptr inbounds ([2 x i8], [2 x i8]* @t, i32 0, i32 0))
  %13 = load i32, i32* %1
  ret i32 %13
}

declare void @putstring(i8*) #1

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
