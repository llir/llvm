; ModuleID = 'sim04.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@x = common global [7 x i8] zeroinitializer, align 1

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %y = alloca [10 x i8], align 1
  store i8 72, i8* getelementptr inbounds ([7 x i8], [7 x i8]* @x, i32 0, i64 0), align 1
  store i8 101, i8* getelementptr inbounds ([7 x i8], [7 x i8]* @x, i32 0, i64 1), align 1
  store i8 108, i8* getelementptr inbounds ([7 x i8], [7 x i8]* @x, i32 0, i64 2), align 1
  store i8 108, i8* getelementptr inbounds ([7 x i8], [7 x i8]* @x, i32 0, i64 3), align 1
  store i8 111, i8* getelementptr inbounds ([7 x i8], [7 x i8]* @x, i32 0, i64 4), align 1
  store i8 10, i8* getelementptr inbounds ([7 x i8], [7 x i8]* @x, i32 0, i64 5), align 1
  store i8 0, i8* getelementptr inbounds ([7 x i8], [7 x i8]* @x, i32 0, i64 6), align 1
  %1 = getelementptr inbounds [10 x i8], [10 x i8]* %y, i32 0, i64 0
  store i8 71, i8* %1, align 1
  %2 = getelementptr inbounds [10 x i8], [10 x i8]* %y, i32 0, i64 1
  store i8 111, i8* %2, align 1
  %3 = getelementptr inbounds [10 x i8], [10 x i8]* %y, i32 0, i64 2
  store i8 111, i8* %3, align 1
  %4 = getelementptr inbounds [10 x i8], [10 x i8]* %y, i32 0, i64 3
  store i8 100, i8* %4, align 1
  %5 = getelementptr inbounds [10 x i8], [10 x i8]* %y, i32 0, i64 4
  store i8 32, i8* %5, align 1
  %6 = getelementptr inbounds [10 x i8], [10 x i8]* %y, i32 0, i64 5
  store i8 98, i8* %6, align 1
  %7 = getelementptr inbounds [10 x i8], [10 x i8]* %y, i32 0, i64 6
  store i8 121, i8* %7, align 1
  %8 = getelementptr inbounds [10 x i8], [10 x i8]* %y, i32 0, i64 7
  store i8 101, i8* %8, align 1
  %9 = getelementptr inbounds [10 x i8], [10 x i8]* %y, i32 0, i64 8
  store i8 10, i8* %9, align 1
  %10 = getelementptr inbounds [10 x i8], [10 x i8]* %y, i32 0, i64 9
  store i8 0, i8* %10, align 1
  call void @putstring(i8* getelementptr inbounds ([7 x i8], [7 x i8]* @x, i32 0, i32 0))
  %11 = getelementptr inbounds [10 x i8], [10 x i8]* %y, i32 0, i32 0
  call void @putstring(i8* %11)
  ret i32 0
}

declare void @putstring(i8*) #1

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
