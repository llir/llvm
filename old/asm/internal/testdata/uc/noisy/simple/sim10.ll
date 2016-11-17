; ModuleID = 'sim10.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %nameq = alloca [12 x i8], align 1
  %ageq = alloca [10 x i8], align 1
  %youare = alloca [10 x i8], align 1
  %cr = alloca [2 x i8], align 1
  %name = alloca [80 x i8], align 16
  %age = alloca i32, align 4
  %1 = getelementptr inbounds [12 x i8], [12 x i8]* %nameq, i32 0, i64 0
  store i8 89, i8* %1, align 1
  %2 = getelementptr inbounds [12 x i8], [12 x i8]* %nameq, i32 0, i64 1
  store i8 111, i8* %2, align 1
  %3 = getelementptr inbounds [12 x i8], [12 x i8]* %nameq, i32 0, i64 2
  store i8 117, i8* %3, align 1
  %4 = getelementptr inbounds [12 x i8], [12 x i8]* %nameq, i32 0, i64 3
  store i8 114, i8* %4, align 1
  %5 = getelementptr inbounds [12 x i8], [12 x i8]* %nameq, i32 0, i64 4
  store i8 32, i8* %5, align 1
  %6 = getelementptr inbounds [12 x i8], [12 x i8]* %nameq, i32 0, i64 5
  store i8 110, i8* %6, align 1
  %7 = getelementptr inbounds [12 x i8], [12 x i8]* %nameq, i32 0, i64 6
  store i8 97, i8* %7, align 1
  %8 = getelementptr inbounds [12 x i8], [12 x i8]* %nameq, i32 0, i64 7
  store i8 109, i8* %8, align 1
  %9 = getelementptr inbounds [12 x i8], [12 x i8]* %nameq, i32 0, i64 8
  store i8 101, i8* %9, align 1
  %10 = getelementptr inbounds [12 x i8], [12 x i8]* %nameq, i32 0, i64 9
  store i8 63, i8* %10, align 1
  %11 = getelementptr inbounds [12 x i8], [12 x i8]* %nameq, i32 0, i64 10
  store i8 32, i8* %11, align 1
  %12 = getelementptr inbounds [12 x i8], [12 x i8]* %nameq, i32 0, i64 11
  store i8 0, i8* %12, align 1
  %13 = getelementptr inbounds [10 x i8], [10 x i8]* %ageq, i32 0, i64 0
  store i8 89, i8* %13, align 1
  %14 = getelementptr inbounds [10 x i8], [10 x i8]* %ageq, i32 0, i64 1
  store i8 111, i8* %14, align 1
  %15 = getelementptr inbounds [10 x i8], [10 x i8]* %ageq, i32 0, i64 2
  store i8 117, i8* %15, align 1
  %16 = getelementptr inbounds [10 x i8], [10 x i8]* %ageq, i32 0, i64 3
  store i8 114, i8* %16, align 1
  %17 = getelementptr inbounds [10 x i8], [10 x i8]* %ageq, i32 0, i64 4
  store i8 32, i8* %17, align 1
  %18 = getelementptr inbounds [10 x i8], [10 x i8]* %ageq, i32 0, i64 5
  store i8 97, i8* %18, align 1
  %19 = getelementptr inbounds [10 x i8], [10 x i8]* %ageq, i32 0, i64 6
  store i8 103, i8* %19, align 1
  %20 = getelementptr inbounds [10 x i8], [10 x i8]* %ageq, i32 0, i64 7
  store i8 101, i8* %20, align 1
  %21 = getelementptr inbounds [10 x i8], [10 x i8]* %ageq, i32 0, i64 8
  store i8 32, i8* %21, align 1
  %22 = getelementptr inbounds [10 x i8], [10 x i8]* %ageq, i32 0, i64 9
  store i8 0, i8* %22, align 1
  %23 = getelementptr inbounds [10 x i8], [10 x i8]* %youare, i32 0, i64 0
  store i8 89, i8* %23, align 1
  %24 = getelementptr inbounds [10 x i8], [10 x i8]* %youare, i32 0, i64 1
  store i8 111, i8* %24, align 1
  %25 = getelementptr inbounds [10 x i8], [10 x i8]* %youare, i32 0, i64 2
  store i8 117, i8* %25, align 1
  %26 = getelementptr inbounds [10 x i8], [10 x i8]* %youare, i32 0, i64 3
  store i8 32, i8* %26, align 1
  %27 = getelementptr inbounds [10 x i8], [10 x i8]* %youare, i32 0, i64 4
  store i8 97, i8* %27, align 1
  %28 = getelementptr inbounds [10 x i8], [10 x i8]* %youare, i32 0, i64 5
  store i8 114, i8* %28, align 1
  %29 = getelementptr inbounds [10 x i8], [10 x i8]* %youare, i32 0, i64 6
  store i8 101, i8* %29, align 1
  %30 = getelementptr inbounds [10 x i8], [10 x i8]* %youare, i32 0, i64 7
  store i8 58, i8* %30, align 1
  %31 = getelementptr inbounds [10 x i8], [10 x i8]* %youare, i32 0, i64 8
  store i8 32, i8* %31, align 1
  %32 = getelementptr inbounds [10 x i8], [10 x i8]* %youare, i32 0, i64 9
  store i8 0, i8* %32, align 1
  %33 = getelementptr inbounds [2 x i8], [2 x i8]* %cr, i32 0, i64 0
  store i8 10, i8* %33, align 1
  %34 = getelementptr inbounds [2 x i8], [2 x i8]* %cr, i32 0, i64 1
  store i8 0, i8* %34, align 1
  %35 = getelementptr inbounds [12 x i8], [12 x i8]* %nameq, i32 0, i32 0
  call void @putstring(i8* %35)
  %36 = getelementptr inbounds [80 x i8], [80 x i8]* %name, i32 0, i32 0
  %37 = call i32 @getstring(i8* %36)
  %38 = getelementptr inbounds [10 x i8], [10 x i8]* %ageq, i32 0, i32 0
  call void @putstring(i8* %38)
  %39 = call i32 @getint()
  store i32 %39, i32* %age, align 4
  %40 = getelementptr inbounds [10 x i8], [10 x i8]* %youare, i32 0, i32 0
  call void @putstring(i8* %40)
  %41 = getelementptr inbounds [80 x i8], [80 x i8]* %name, i32 0, i32 0
  call void @putstring(i8* %41)
  %42 = getelementptr inbounds [2 x i8], [2 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %42)
  %43 = getelementptr inbounds [10 x i8], [10 x i8]* %youare, i32 0, i32 0
  call void @putstring(i8* %43)
  %44 = load i32, i32* %age, align 4
  call void @putint(i32 %44)
  %45 = getelementptr inbounds [2 x i8], [2 x i8]* %cr, i32 0, i32 0
  call void @putstring(i8* %45)
  ret i32 0
}

declare void @putstring(i8*) #1

declare i32 @getstring(i8*) #1

declare i32 @getint() #1

declare void @putint(i32) #1

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
