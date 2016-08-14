; ModuleID = 'sim11.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define void @f(i32 %i, i8* %v) #0 {
  %1 = alloca i32, align 4
  %2 = alloca i8*, align 8
  store i32 %i, i32* %1, align 4
  store i8* %v, i8** %2, align 8
  %3 = load i32, i32* %1, align 4
  %4 = trunc i32 %3 to i8
  %5 = load i8*, i8** %2, align 8
  %6 = getelementptr inbounds i8, i8* %5, i64 0
  store i8 %4, i8* %6, align 1
  ret void
}

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %1 = alloca i32, align 4
  %t = alloca [2 x i8], align 1
  %b = alloca i32, align 4
  store i32 0, i32* %1
  store i32 10, i32* %b, align 4
  %2 = getelementptr inbounds [2 x i8], [2 x i8]* %t, i32 0, i64 1
  store i8 0, i8* %2, align 1
  br label %3

; <label>:3                                       ; preds = %6, %0
  %4 = load i32, i32* %b, align 4
  %5 = icmp ne i32 %4, 0
  br i1 %5, label %6, label %14

; <label>:6                                       ; preds = %3
  %7 = load i32, i32* %b, align 4
  %8 = add nsw i32 48, %7
  %9 = sub nsw i32 %8, 1
  %10 = getelementptr inbounds [2 x i8], [2 x i8]* %t, i32 0, i32 0
  call void @f(i32 %9, i8* %10)
  %11 = getelementptr inbounds [2 x i8], [2 x i8]* %t, i32 0, i32 0
  call void @putstring(i8* %11)
  %12 = load i32, i32* %b, align 4
  %13 = sub nsw i32 %12, 1
  store i32 %13, i32* %b, align 4
  br label %3

; <label>:14                                      ; preds = %3
  %15 = getelementptr inbounds [2 x i8], [2 x i8]* %t, i32 0, i32 0
  call void @f(i32 10, i8* %15)
  %16 = getelementptr inbounds [2 x i8], [2 x i8]* %t, i32 0, i32 0
  call void @putstring(i8* %16)
  %17 = load i32, i32* %1
  ret i32 %17
}

declare void @putstring(i8*) #1

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
