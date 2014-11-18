; ModuleID = 'for.ll'

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
