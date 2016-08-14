; ModuleID = 'eval.c'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@p = common global i32 0, align 4
@s = common global [80 x i8] zeroinitializer, align 16
@zero = common global i32 0, align 4
@nine = common global i32 0, align 4
@plus = common global i32 0, align 4
@minus = common global i32 0, align 4
@times = common global i32 0, align 4
@div = common global i32 0, align 4
@lpar = common global i32 0, align 4
@rpar = common global i32 0, align 4
@bad_expression = common global [15 x i8] zeroinitializer, align 1
@cr = common global [2 x i8] zeroinitializer, align 1
@bad_number = common global [11 x i8] zeroinitializer, align 1
@test_data = common global [21 x i8] zeroinitializer, align 16

; Function Attrs: nounwind uwtable
define i32 @isNumber(i8 signext %c) #0 {
  %1 = alloca i8, align 1
  store i8 %c, i8* %1, align 1
  %2 = load i32, i32* @p, align 4
  %3 = sext i32 %2 to i64
  %4 = getelementptr inbounds [80 x i8], [80 x i8]* @s, i32 0, i64 %3
  %5 = load i8, i8* %4, align 1
  %6 = sext i8 %5 to i32
  %7 = load i32, i32* @zero, align 4
  %8 = icmp sge i32 %6, %7
  br i1 %8, label %9, label %17

; <label>:9                                       ; preds = %0
  %10 = load i32, i32* @p, align 4
  %11 = sext i32 %10 to i64
  %12 = getelementptr inbounds [80 x i8], [80 x i8]* @s, i32 0, i64 %11
  %13 = load i8, i8* %12, align 1
  %14 = sext i8 %13 to i32
  %15 = load i32, i32* @nine, align 4
  %16 = icmp sle i32 %14, %15
  br label %17

; <label>:17                                      ; preds = %9, %0
  %18 = phi i1 [ false, %0 ], [ %16, %9 ]
  %19 = zext i1 %18 to i32
  ret i32 %19
}

; Function Attrs: nounwind uwtable
define i32 @expr(i32 %l) #0 {
  %1 = alloca i32, align 4
  %2 = alloca i32, align 4
  %a = alloca i32, align 4
  %b = alloca i32, align 4
  store i32 %l, i32* %2, align 4
  %3 = load i32, i32* %2, align 4
  %4 = icmp eq i32 %3, 0
  br i1 %4, label %5, label %51

; <label>:5                                       ; preds = %0
  %6 = call i32 @expr(i32 1)
  store i32 %6, i32* %a, align 4
  br label %7

; <label>:7                                       ; preds = %48, %5
  %8 = load i32, i32* @p, align 4
  %9 = sext i32 %8 to i64
  %10 = getelementptr inbounds [80 x i8], [80 x i8]* @s, i32 0, i64 %9
  %11 = load i8, i8* %10, align 1
  %12 = sext i8 %11 to i32
  %13 = load i32, i32* @plus, align 4
  %14 = icmp ne i32 %12, %13
  br i1 %14, label %15, label %23

; <label>:15                                      ; preds = %7
  %16 = load i32, i32* @p, align 4
  %17 = sext i32 %16 to i64
  %18 = getelementptr inbounds [80 x i8], [80 x i8]* @s, i32 0, i64 %17
  %19 = load i8, i8* %18, align 1
  %20 = sext i8 %19 to i32
  %21 = load i32, i32* @minus, align 4
  %22 = icmp ne i32 %20, %21
  br label %23

; <label>:23                                      ; preds = %15, %7
  %24 = phi i1 [ false, %7 ], [ %22, %15 ]
  %25 = xor i1 %24, true
  br i1 %25, label %26, label %49

; <label>:26                                      ; preds = %23
  %27 = load i32, i32* @p, align 4
  %28 = sext i32 %27 to i64
  %29 = getelementptr inbounds [80 x i8], [80 x i8]* @s, i32 0, i64 %28
  %30 = load i8, i8* %29, align 1
  %31 = sext i8 %30 to i32
  %32 = load i32, i32* @plus, align 4
  %33 = icmp eq i32 %31, %32
  br i1 %33, label %34, label %41

; <label>:34                                      ; preds = %26
  %35 = load i32, i32* @p, align 4
  %36 = add nsw i32 %35, 1
  store i32 %36, i32* @p, align 4
  %37 = call i32 @expr(i32 1)
  store i32 %37, i32* %b, align 4
  %38 = load i32, i32* %a, align 4
  %39 = load i32, i32* %b, align 4
  %40 = add nsw i32 %38, %39
  store i32 %40, i32* %a, align 4
  br label %48

; <label>:41                                      ; preds = %26
  %42 = load i32, i32* @p, align 4
  %43 = add nsw i32 %42, 1
  store i32 %43, i32* @p, align 4
  %44 = call i32 @expr(i32 1)
  store i32 %44, i32* %b, align 4
  %45 = load i32, i32* %a, align 4
  %46 = load i32, i32* %b, align 4
  %47 = sub nsw i32 %45, %46
  store i32 %47, i32* %a, align 4
  br label %48

; <label>:48                                      ; preds = %41, %34
  br label %7

; <label>:49                                      ; preds = %23
  %50 = load i32, i32* %a, align 4
  store i32 %50, i32* %1
  br label %160

; <label>:51                                      ; preds = %0
  %52 = load i32, i32* %2, align 4
  %53 = icmp eq i32 %52, 1
  br i1 %53, label %54, label %100

; <label>:54                                      ; preds = %51
  %55 = call i32 @expr(i32 2)
  store i32 %55, i32* %a, align 4
  br label %56

; <label>:56                                      ; preds = %97, %54
  %57 = load i32, i32* @p, align 4
  %58 = sext i32 %57 to i64
  %59 = getelementptr inbounds [80 x i8], [80 x i8]* @s, i32 0, i64 %58
  %60 = load i8, i8* %59, align 1
  %61 = sext i8 %60 to i32
  %62 = load i32, i32* @times, align 4
  %63 = icmp ne i32 %61, %62
  br i1 %63, label %64, label %72

; <label>:64                                      ; preds = %56
  %65 = load i32, i32* @p, align 4
  %66 = sext i32 %65 to i64
  %67 = getelementptr inbounds [80 x i8], [80 x i8]* @s, i32 0, i64 %66
  %68 = load i8, i8* %67, align 1
  %69 = sext i8 %68 to i32
  %70 = load i32, i32* @div, align 4
  %71 = icmp ne i32 %69, %70
  br label %72

; <label>:72                                      ; preds = %64, %56
  %73 = phi i1 [ false, %56 ], [ %71, %64 ]
  %74 = xor i1 %73, true
  br i1 %74, label %75, label %98

; <label>:75                                      ; preds = %72
  %76 = load i32, i32* @p, align 4
  %77 = sext i32 %76 to i64
  %78 = getelementptr inbounds [80 x i8], [80 x i8]* @s, i32 0, i64 %77
  %79 = load i8, i8* %78, align 1
  %80 = sext i8 %79 to i32
  %81 = load i32, i32* @times, align 4
  %82 = icmp eq i32 %80, %81
  br i1 %82, label %83, label %90

; <label>:83                                      ; preds = %75
  %84 = load i32, i32* @p, align 4
  %85 = add nsw i32 %84, 1
  store i32 %85, i32* @p, align 4
  %86 = call i32 @expr(i32 2)
  store i32 %86, i32* %b, align 4
  %87 = load i32, i32* %a, align 4
  %88 = load i32, i32* %b, align 4
  %89 = mul nsw i32 %87, %88
  store i32 %89, i32* %a, align 4
  br label %97

; <label>:90                                      ; preds = %75
  %91 = load i32, i32* @p, align 4
  %92 = add nsw i32 %91, 1
  store i32 %92, i32* @p, align 4
  %93 = call i32 @expr(i32 2)
  store i32 %93, i32* %b, align 4
  %94 = load i32, i32* %a, align 4
  %95 = load i32, i32* %b, align 4
  %96 = sdiv i32 %94, %95
  store i32 %96, i32* %a, align 4
  br label %97

; <label>:97                                      ; preds = %90, %83
  br label %56

; <label>:98                                      ; preds = %72
  %99 = load i32, i32* %a, align 4
  store i32 %99, i32* %1
  br label %160

; <label>:100                                     ; preds = %51
  %101 = load i32, i32* %2, align 4
  %102 = icmp eq i32 %101, 2
  br i1 %102, label %103, label %158

; <label>:103                                     ; preds = %100
  %104 = load i32, i32* @p, align 4
  %105 = sext i32 %104 to i64
  %106 = getelementptr inbounds [80 x i8], [80 x i8]* @s, i32 0, i64 %105
  %107 = load i8, i8* %106, align 1
  %108 = sext i8 %107 to i32
  %109 = load i32, i32* @lpar, align 4
  %110 = icmp eq i32 %108, %109
  br i1 %110, label %111, label %127

; <label>:111                                     ; preds = %103
  %112 = load i32, i32* @p, align 4
  %113 = add nsw i32 %112, 1
  store i32 %113, i32* @p, align 4
  %114 = call i32 @expr(i32 0)
  store i32 %114, i32* %a, align 4
  %115 = load i32, i32* @p, align 4
  %116 = sext i32 %115 to i64
  %117 = getelementptr inbounds [80 x i8], [80 x i8]* @s, i32 0, i64 %116
  %118 = load i8, i8* %117, align 1
  %119 = sext i8 %118 to i32
  %120 = load i32, i32* @rpar, align 4
  %121 = icmp ne i32 %119, %120
  br i1 %121, label %122, label %123

; <label>:122                                     ; preds = %111
  call void @putstring(i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i32 0))
  call void @putstring(i8* getelementptr inbounds ([2 x i8], [2 x i8]* @cr, i32 0, i32 0))
  br label %123

; <label>:123                                     ; preds = %122, %111
  %124 = load i32, i32* @p, align 4
  %125 = add nsw i32 %124, 1
  store i32 %125, i32* @p, align 4
  %126 = load i32, i32* %a, align 4
  store i32 %126, i32* %1
  br label %160

; <label>:127                                     ; preds = %103
  %128 = load i32, i32* @p, align 4
  %129 = sext i32 %128 to i64
  %130 = getelementptr inbounds [80 x i8], [80 x i8]* @s, i32 0, i64 %129
  %131 = load i8, i8* %130, align 1
  %132 = call i32 @isNumber(i8 signext %131)
  %133 = icmp ne i32 %132, 0
  br i1 %133, label %135, label %134

; <label>:134                                     ; preds = %127
  call void @putstring(i8* getelementptr inbounds ([11 x i8], [11 x i8]* @bad_number, i32 0, i32 0))
  call void @putstring(i8* getelementptr inbounds ([2 x i8], [2 x i8]* @cr, i32 0, i32 0))
  store i32 0, i32* %1
  br label %160

; <label>:135                                     ; preds = %127
  store i32 0, i32* %a, align 4
  br label %136

; <label>:136                                     ; preds = %143, %135
  %137 = load i32, i32* @p, align 4
  %138 = sext i32 %137 to i64
  %139 = getelementptr inbounds [80 x i8], [80 x i8]* @s, i32 0, i64 %138
  %140 = load i8, i8* %139, align 1
  %141 = call i32 @isNumber(i8 signext %140)
  %142 = icmp ne i32 %141, 0
  br i1 %142, label %143, label %156

; <label>:143                                     ; preds = %136
  %144 = load i32, i32* %a, align 4
  %145 = mul nsw i32 %144, 10
  %146 = load i32, i32* @p, align 4
  %147 = sext i32 %146 to i64
  %148 = getelementptr inbounds [80 x i8], [80 x i8]* @s, i32 0, i64 %147
  %149 = load i8, i8* %148, align 1
  %150 = sext i8 %149 to i32
  %151 = load i32, i32* @zero, align 4
  %152 = sub nsw i32 %150, %151
  %153 = add nsw i32 %145, %152
  store i32 %153, i32* %a, align 4
  %154 = load i32, i32* @p, align 4
  %155 = add nsw i32 %154, 1
  store i32 %155, i32* @p, align 4
  br label %136

; <label>:156                                     ; preds = %136
  %157 = load i32, i32* %a, align 4
  store i32 %157, i32* %1
  br label %160

; <label>:158                                     ; preds = %100
  br label %159

; <label>:159                                     ; preds = %158
  br label %160

; <label>:160                                     ; preds = %49, %98, %123, %134, %156, %159
  %161 = load i32, i32* %1
  ret i32 %161
}

declare void @putstring(i8*) #1

; Function Attrs: nounwind uwtable
define i32 @main() #0 {
  %1 = alloca i32, align 4
  store i32 0, i32* %1
  store i32 48, i32* @zero, align 4
  store i32 57, i32* @nine, align 4
  store i32 43, i32* @plus, align 4
  store i32 45, i32* @minus, align 4
  store i32 42, i32* @times, align 4
  store i32 47, i32* @div, align 4
  store i32 40, i32* @lpar, align 4
  store i32 41, i32* @rpar, align 4
  store i8 66, i8* getelementptr inbounds ([11 x i8], [11 x i8]* @bad_number, i32 0, i64 0), align 1
  store i8 97, i8* getelementptr inbounds ([11 x i8], [11 x i8]* @bad_number, i32 0, i64 1), align 1
  store i8 100, i8* getelementptr inbounds ([11 x i8], [11 x i8]* @bad_number, i32 0, i64 2), align 1
  store i8 32, i8* getelementptr inbounds ([11 x i8], [11 x i8]* @bad_number, i32 0, i64 3), align 1
  store i8 110, i8* getelementptr inbounds ([11 x i8], [11 x i8]* @bad_number, i32 0, i64 4), align 1
  store i8 117, i8* getelementptr inbounds ([11 x i8], [11 x i8]* @bad_number, i32 0, i64 5), align 1
  store i8 109, i8* getelementptr inbounds ([11 x i8], [11 x i8]* @bad_number, i32 0, i64 6), align 1
  store i8 98, i8* getelementptr inbounds ([11 x i8], [11 x i8]* @bad_number, i32 0, i64 7), align 1
  store i8 101, i8* getelementptr inbounds ([11 x i8], [11 x i8]* @bad_number, i32 0, i64 8), align 1
  store i8 114, i8* getelementptr inbounds ([11 x i8], [11 x i8]* @bad_number, i32 0, i64 9), align 1
  store i8 0, i8* getelementptr inbounds ([11 x i8], [11 x i8]* @bad_number, i32 0, i64 10), align 1
  store i8 66, i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i64 0), align 1
  store i8 97, i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i64 1), align 1
  store i8 100, i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i64 2), align 1
  store i8 32, i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i64 3), align 1
  store i8 101, i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i64 4), align 1
  store i8 120, i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i64 5), align 1
  store i8 112, i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i64 6), align 1
  store i8 114, i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i64 7), align 1
  store i8 101, i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i64 8), align 1
  store i8 115, i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i64 9), align 1
  store i8 115, i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i64 10), align 1
  store i8 105, i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i64 11), align 1
  store i8 111, i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i64 12), align 1
  store i8 110, i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i64 13), align 1
  store i8 0, i8* getelementptr inbounds ([15 x i8], [15 x i8]* @bad_expression, i32 0, i64 14), align 1
  store i8 10, i8* getelementptr inbounds ([2 x i8], [2 x i8]* @cr, i32 0, i64 0), align 1
  store i8 0, i8* getelementptr inbounds ([2 x i8], [2 x i8]* @cr, i32 0, i64 1), align 1
  store i8 40, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 0), align 1
  store i8 49, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 1), align 1
  store i8 50, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 2), align 1
  store i8 45, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 3), align 1
  store i8 52, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 4), align 1
  store i8 41, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 5), align 1
  store i8 43, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 6), align 1
  store i8 40, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 7), align 1
  store i8 57, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 8), align 1
  store i8 57, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 9), align 1
  store i8 45, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 10), align 1
  store i8 49, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 11), align 1
  store i8 49, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 12), align 1
  store i8 43, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 13), align 1
  store i8 49, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 14), align 1
  store i8 54, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 15), align 1
  store i8 41, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 16), align 1
  store i8 42, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 17), align 1
  store i8 49, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 18), align 1
  store i8 57, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 19), align 1
  store i8 0, i8* getelementptr inbounds ([21 x i8], [21 x i8]* @test_data, i32 0, i64 20), align 1
  call void @getstring(i8* getelementptr inbounds ([80 x i8], [80 x i8]* @s, i32 0, i32 0))
  %2 = load i8, i8* getelementptr inbounds ([80 x i8], [80 x i8]* @s, i32 0, i64 0), align 1
  %3 = sext i8 %2 to i32
  %4 = icmp eq i32 %3, 116
  br i1 %4, label %5, label %27

; <label>:5                                       ; preds = %0
  store i32 0, i32* @p, align 4
  br label %6

; <label>:6                                       ; preds = %13, %5
  %7 = load i32, i32* @p, align 4
  %8 = sext i32 %7 to i64
  %9 = getelementptr inbounds [21 x i8], [21 x i8]* @test_data, i32 0, i64 %8
  %10 = load i8, i8* %9, align 1
  %11 = sext i8 %10 to i32
  %12 = icmp ne i32 %11, 0
  br i1 %12, label %13, label %23

; <label>:13                                      ; preds = %6
  %14 = load i32, i32* @p, align 4
  %15 = sext i32 %14 to i64
  %16 = getelementptr inbounds [21 x i8], [21 x i8]* @test_data, i32 0, i64 %15
  %17 = load i8, i8* %16, align 1
  %18 = load i32, i32* @p, align 4
  %19 = sext i32 %18 to i64
  %20 = getelementptr inbounds [80 x i8], [80 x i8]* @s, i32 0, i64 %19
  store i8 %17, i8* %20, align 1
  %21 = load i32, i32* @p, align 4
  %22 = add nsw i32 %21, 1
  store i32 %22, i32* @p, align 4
  br label %6

; <label>:23                                      ; preds = %6
  %24 = load i32, i32* @p, align 4
  %25 = sext i32 %24 to i64
  %26 = getelementptr inbounds [80 x i8], [80 x i8]* @s, i32 0, i64 %25
  store i8 0, i8* %26, align 1
  br label %27

; <label>:27                                      ; preds = %23, %0
  store i32 0, i32* @p, align 4
  %28 = call i32 @expr(i32 0)
  call void @putint(i32 %28)
  call void @putstring(i8* getelementptr inbounds ([2 x i8], [2 x i8]* @cr, i32 0, i32 0))
  %29 = load i32, i32* %1
  ret i32 %29
}

declare void @getstring(i8*) #1

declare void @putint(i32) #1

attributes #0 = { nounwind uwtable "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "disable-tail-calls"="false" "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "target-cpu"="x86-64" "target-features"="+sse,+sse2" "unsafe-fp-math"="false" "use-soft-float"="false" }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.7.1 (tags/RELEASE_371/final)"}
