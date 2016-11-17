; ModuleID = 'sim05.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %cr = alloca [1 x i8], align 1
  %x = alloca i32, align 4
  %y = alloca i32, align 4
  %1 = getelementptr inbounds [1 x i8], [1 x i8]* %cr, i32 0, i64 0
  store i8 10, i8* %1, align 1
  %2 = getelementptr inbounds [1 x i8], [1 x i8]* %cr, i32 0, i64 1
  store i8 0, i8* %2, align 1
  call void @putint(i32 42)
  %3 = getelementptr inbounds [1 x i8], [1 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %3)
  call void @putint(i32 42)
  %4 = getelementptr inbounds [1 x i8], [1 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %4)
  call void @putint(i32 42)
  %5 = getelementptr inbounds [1 x i8], [1 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %5)
  call void @putint(i32 42)
  %6 = getelementptr inbounds [1 x i8], [1 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %6)
  call void @putint(i32 42)
  %7 = getelementptr inbounds [1 x i8], [1 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %7)
  call void @putint(i32 42)
  %8 = getelementptr inbounds [1 x i8], [1 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %8)
  call void @putint(i32 42)
  %9 = getelementptr inbounds [1 x i8], [1 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %9)
  store i32 6, i32* %y, align 4
  store i32 6, i32* %x, align 4
  %10 = load i32, i32* %x, align 4
  %11 = load i32, i32* %y, align 4
  %12 = add nsw i32 %11, 1
  %13 = mul nsw i32 %10, %12
  call void @putint(i32 %13)
  %14 = getelementptr inbounds [1 x i8], [1 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %14)
  store i32 5, i32* %y, align 4
  store i32 8, i32* %x, align 4
  %15 = load i32, i32* %x, align 4
  %16 = load i32, i32* %y, align 4
  %17 = mul nsw i32 %15, %16
  %18 = add nsw i32 %17, 2
  call void @putint(i32 %18)
  %19 = getelementptr inbounds [1 x i8], [1 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %19)
  call void @putint(i32 3141592)
  %20 = getelementptr inbounds [1 x i8], [1 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %20)
  call void @putint(i32 1)
  call void @putint(i32 1)
  call void @putint(i32 0)
  call void @putint(i32 0)
  call void @putint(i32 0)
  call void @putint(i32 1)
  call void @putint(i32 0)
  call void @putint(i32 0)
  call void @putint(i32 0)
  %21 = getelementptr inbounds [1 x i8], [1 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %21)
  ret i32 0
}

declare void @putint(i32) #1

declare void @putstring(i8*) #1

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
