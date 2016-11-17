; ModuleID = 'sim09.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@x = common global [10 x i32] zeroinitializer, align 16
@xc = common global [10 x i8] zeroinitializer, align 1

; Function Attrs: nounwind uwtable
define void @f(i32* %a) #0 {
  %1 = alloca i32*, align 8
  store i32* %a, i32** %1, align 8
  %2 = load i32*, i32** %1, align 8
  %3 = getelementptr inbounds i32, i32* %2, i64 3
  %4 = load i32, i32* %3, align 4
  call void @putint(i32 %4)
  ret void
}

declare void @putint(i32) #1

; Function Attrs: nounwind uwtable
define void @g(i32* %b) #0 {
  %1 = alloca i32*, align 8
  store i32* %b, i32** %1, align 8
  %2 = load i32*, i32** %1, align 8
  call void @f(i32* %2)
  ret void
}

; Function Attrs: nounwind uwtable
define void @fc(i8* %a) #0 {
  %1 = alloca i8*, align 8
  store i8* %a, i8** %1, align 8
  %2 = load i8*, i8** %1, align 8
  call void @putstring(i8* %2)
  ret void
}

declare void @putstring(i8*) #1

; Function Attrs: nounwind uwtable
define void @gc(i8* %b) #0 {
  %1 = alloca i8*, align 8
  store i8* %b, i8** %1, align 8
  %2 = load i8*, i8** %1, align 8
  call void @fc(i8* %2)
  ret void
}

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %y = alloca [10 x i32], align 16
  %yc = alloca [10 x i8], align 1
  store i32 12, i32* getelementptr inbounds ([10 x i32], [10 x i32]* @x, i32 0, i64 3), align 4
  %1 = getelementptr inbounds [10 x i32], [10 x i32]* %y, i32 0, i64 3
  store i32 34, i32* %1, align 4
  call void @f(i32* getelementptr inbounds ([10 x i32], [10 x i32]* @x, i32 0, i32 0))
  %2 = getelementptr inbounds [10 x i32], [10 x i32]* %y, i32 0, i32 0
  call void @f(i32* %2)
  store i32 56, i32* getelementptr inbounds ([10 x i32], [10 x i32]* @x, i32 0, i64 3), align 4
  %3 = getelementptr inbounds [10 x i32], [10 x i32]* %y, i32 0, i64 3
  store i32 78, i32* %3, align 4
  call void @g(i32* getelementptr inbounds ([10 x i32], [10 x i32]* @x, i32 0, i32 0))
  %4 = getelementptr inbounds [10 x i32], [10 x i32]* %y, i32 0, i32 0
  call void @g(i32* %4)
  store i8 10, i8* getelementptr inbounds ([10 x i8], [10 x i8]* @xc, i32 0, i64 0), align 1
  store i8 65, i8* getelementptr inbounds ([10 x i8], [10 x i8]* @xc, i32 0, i64 1), align 1
  store i8 0, i8* getelementptr inbounds ([10 x i8], [10 x i8]* @xc, i32 0, i64 2), align 1
  %5 = getelementptr inbounds [10 x i8], [10 x i8]* %yc, i32 0, i64 0
  store i8 66, i8* %5, align 1
  %6 = getelementptr inbounds [10 x i8], [10 x i8]* %yc, i32 0, i64 1
  store i8 0, i8* %6, align 1
  call void @fc(i8* getelementptr inbounds ([10 x i8], [10 x i8]* @xc, i32 0, i32 0))
  %7 = getelementptr inbounds [10 x i8], [10 x i8]* %yc, i32 0, i32 0
  call void @fc(i8* %7)
  store i8 67, i8* getelementptr inbounds ([10 x i8], [10 x i8]* @xc, i32 0, i64 0), align 1
  store i8 0, i8* getelementptr inbounds ([10 x i8], [10 x i8]* @xc, i32 0, i64 1), align 1
  %8 = getelementptr inbounds [10 x i8], [10 x i8]* %yc, i32 0, i64 0
  store i8 68, i8* %8, align 1
  %9 = getelementptr inbounds [10 x i8], [10 x i8]* %yc, i32 0, i64 1
  store i8 10, i8* %9, align 1
  %10 = getelementptr inbounds [10 x i8], [10 x i8]* %yc, i32 0, i64 2
  store i8 0, i8* %10, align 1
  call void @gc(i8* getelementptr inbounds ([10 x i8], [10 x i8]* @xc, i32 0, i32 0))
  %11 = getelementptr inbounds [10 x i8], [10 x i8]* %yc, i32 0, i32 0
  call void @gc(i8* %11)
  ret i32 0
}

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
