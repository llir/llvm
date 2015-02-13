; ModuleID = 'for.ll'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

; Function Attrs: nounwind uwtable
define i32 @main(i32 %argc, i8** %argv) #0 {
  br label %1

; <label>:1                                       ; preds = %5, %0
  %status.0 = phi i32 [ 0, %0 ], [ %4, %5 ]
  %i.0 = phi i32 [ 3, %0 ], [ %6, %5 ]
  %2 = icmp slt i32 %i.0, 10
  br i1 %2, label %3, label %7

; <label>:3                                       ; preds = %1
  %4 = add nsw i32 %status.0, %i.0
  br label %5

; <label>:5                                       ; preds = %3
  %6 = add nsw i32 %i.0, 1
  br label %1

; <label>:7                                       ; preds = %1
  ret i32 %status.0
}

attributes #0 = { nounwind uwtable "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = metadata !{metadata !"clang version 3.5.1 (tags/RELEASE_351/final)"}
