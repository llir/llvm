; ModuleID = 'c4.ll'
target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
target triple = "x86_64-unknown-linux-gnu"

@p = common global i8* null, align 8
@tk = common global i32 0, align 4
@src = common global i32 0, align 4
@.str = private unnamed_addr constant [9 x i8] c"%d: %.*s\00", align 1
@line = common global i32 0, align 4
@lp = common global i8* null, align 8
@le = common global i32* null, align 8
@e = common global i32* null, align 8
@.str1 = private unnamed_addr constant [6 x i8] c"%8.4s\00", align 1
@.str2 = private unnamed_addr constant [191 x i8] c"LEA ,IMM ,JMP ,JSR ,BZ  ,BNZ ,ENT ,ADJ ,LEV ,LI  ,LC  ,SI  ,SC  ,PSH ,OR  ,XOR ,AND ,EQ  ,NE  ,LT  ,GT  ,LE  ,GE  ,SHL ,SHR ,ADD ,SUB ,MUL ,DIV ,MOD ,OPEN,READ,CLOS,PRTF,MALC,MSET,MCMP,EXIT,\00", align 1
@.str3 = private unnamed_addr constant [5 x i8] c" %d\0A\00", align 1
@.str4 = private unnamed_addr constant [2 x i8] c"\0A\00", align 1
@sym = common global i32* null, align 8
@id = common global i32* null, align 8
@ival = common global i32 0, align 4
@data = common global i8* null, align 8
@.str5 = private unnamed_addr constant [34 x i8] c"%d: unexpected eof in expression\0A\00", align 1
@ty = common global i32 0, align 4
@.str6 = private unnamed_addr constant [35 x i8] c"%d: open paren expected in sizeof\0A\00", align 1
@.str7 = private unnamed_addr constant [36 x i8] c"%d: close paren expected in sizeof\0A\00", align 1
@.str8 = private unnamed_addr constant [23 x i8] c"%d: bad function call\0A\00", align 1
@loc = common global i32 0, align 4
@.str9 = private unnamed_addr constant [24 x i8] c"%d: undefined variable\0A\00", align 1
@.str10 = private unnamed_addr constant [14 x i8] c"%d: bad cast\0A\00", align 1
@.str11 = private unnamed_addr constant [26 x i8] c"%d: close paren expected\0A\00", align 1
@.str12 = private unnamed_addr constant [21 x i8] c"%d: bad dereference\0A\00", align 1
@.str13 = private unnamed_addr constant [20 x i8] c"%d: bad address-of\0A\00", align 1
@.str14 = private unnamed_addr constant [33 x i8] c"%d: bad lvalue in pre-increment\0A\00", align 1
@.str15 = private unnamed_addr constant [20 x i8] c"%d: bad expression\0A\00", align 1
@.str16 = private unnamed_addr constant [30 x i8] c"%d: bad lvalue in assignment\0A\00", align 1
@.str17 = private unnamed_addr constant [31 x i8] c"%d: conditional missing colon\0A\00", align 1
@.str18 = private unnamed_addr constant [34 x i8] c"%d: bad lvalue in post-increment\0A\00", align 1
@.str19 = private unnamed_addr constant [28 x i8] c"%d: close bracket expected\0A\00", align 1
@.str20 = private unnamed_addr constant [27 x i8] c"%d: pointer type expected\0A\00", align 1
@.str21 = private unnamed_addr constant [26 x i8] c"%d: compiler error tk=%d\0A\00", align 1
@.str22 = private unnamed_addr constant [25 x i8] c"%d: open paren expected\0A\00", align 1
@.str23 = private unnamed_addr constant [24 x i8] c"%d: semicolon expected\0A\00", align 1
@debug = common global i32 0, align 4
@.str24 = private unnamed_addr constant [30 x i8] c"usage: c4 [-s] [-d] file ...\0A\00", align 1
@.str25 = private unnamed_addr constant [20 x i8] c"could not open(%s)\0A\00", align 1
@.str26 = private unnamed_addr constant [34 x i8] c"could not malloc(%d) symbol area\0A\00", align 1
@.str27 = private unnamed_addr constant [32 x i8] c"could not malloc(%d) text area\0A\00", align 1
@.str28 = private unnamed_addr constant [32 x i8] c"could not malloc(%d) data area\0A\00", align 1
@.str29 = private unnamed_addr constant [33 x i8] c"could not malloc(%d) stack area\0A\00", align 1
@.str30 = private unnamed_addr constant [101 x i8] c"char else enum if int return sizeof while open read close printf malloc memset memcmp exit void main\00", align 1
@.str31 = private unnamed_addr constant [34 x i8] c"could not malloc(%d) source area\0A\00", align 1
@.str32 = private unnamed_addr constant [20 x i8] c"read() returned %d\0A\00", align 1
@.str33 = private unnamed_addr constant [28 x i8] c"%d: bad enum identifier %d\0A\00", align 1
@.str34 = private unnamed_addr constant [26 x i8] c"%d: bad enum initializer\0A\00", align 1
@.str35 = private unnamed_addr constant [28 x i8] c"%d: bad global declaration\0A\00", align 1
@.str36 = private unnamed_addr constant [33 x i8] c"%d: duplicate global definition\0A\00", align 1
@.str37 = private unnamed_addr constant [31 x i8] c"%d: bad parameter declaration\0A\00", align 1
@.str38 = private unnamed_addr constant [36 x i8] c"%d: duplicate parameter definition\0A\00", align 1
@.str39 = private unnamed_addr constant [29 x i8] c"%d: bad function definition\0A\00", align 1
@.str40 = private unnamed_addr constant [27 x i8] c"%d: bad local declaration\0A\00", align 1
@.str41 = private unnamed_addr constant [32 x i8] c"%d: duplicate local definition\0A\00", align 1
@.str42 = private unnamed_addr constant [20 x i8] c"main() not defined\0A\00", align 1
@.str43 = private unnamed_addr constant [9 x i8] c"%d> %.4s\00", align 1
@.str44 = private unnamed_addr constant [21 x i8] c"exit(%d) cycle = %d\0A\00", align 1
@.str45 = private unnamed_addr constant [38 x i8] c"unknown instruction = %d! cycle = %d\0A\00", align 1

; Function Attrs: nounwind uwtable
define void @next() #0 {
  br label %1

; <label>:1                                       ; preds = %564, %0
  %2 = load i8** @p, align 8
  %3 = load i8* %2, align 1
  %4 = sext i8 %3 to i32
  store i32 %4, i32* @tk, align 4
  %5 = icmp ne i32 %4, 0
  br i1 %5, label %6, label %565

; <label>:6                                       ; preds = %1
  %7 = load i8** @p, align 8
  %8 = getelementptr inbounds i8* %7, i32 1
  store i8* %8, i8** @p, align 8
  %9 = load i32* @tk, align 4
  %10 = icmp eq i32 %9, 10
  br i1 %10, label %11, label %51

; <label>:11                                      ; preds = %6
  %12 = load i32* @src, align 4
  %13 = icmp ne i32 %12, 0
  br i1 %13, label %14, label %48

; <label>:14                                      ; preds = %11
  %15 = load i32* @line, align 4
  %16 = load i8** @p, align 8
  %17 = load i8** @lp, align 8
  %18 = ptrtoint i8* %16 to i64
  %19 = ptrtoint i8* %17 to i64
  %20 = sub i64 %18, %19
  %21 = load i8** @lp, align 8
  %22 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([9 x i8]* @.str, i32 0, i32 0), i32 %15, i64 %20, i8* %21)
  %23 = load i8** @p, align 8
  store i8* %23, i8** @lp, align 8
  br label %24

; <label>:24                                      ; preds = %46, %14
  %25 = load i32** @le, align 8
  %26 = load i32** @e, align 8
  %27 = icmp ult i32* %25, %26
  br i1 %27, label %28, label %47

; <label>:28                                      ; preds = %24
  %29 = load i32** @le, align 8
  %30 = getelementptr inbounds i32* %29, i32 1
  store i32* %30, i32** @le, align 8
  %31 = load i32* %30, align 4
  %32 = mul nsw i32 %31, 5
  %33 = sext i32 %32 to i64
  %34 = getelementptr inbounds [191 x i8]* @.str2, i32 0, i64 %33
  %35 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([6 x i8]* @.str1, i32 0, i32 0), i8* %34)
  %36 = load i32** @le, align 8
  %37 = load i32* %36, align 4
  %38 = icmp sle i32 %37, 7
  br i1 %38, label %39, label %44

; <label>:39                                      ; preds = %28
  %40 = load i32** @le, align 8
  %41 = getelementptr inbounds i32* %40, i32 1
  store i32* %41, i32** @le, align 8
  %42 = load i32* %41, align 4
  %43 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([5 x i8]* @.str3, i32 0, i32 0), i32 %42)
  br label %46

; <label>:44                                      ; preds = %28
  %45 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([2 x i8]* @.str4, i32 0, i32 0))
  br label %46

; <label>:46                                      ; preds = %44, %39
  br label %24

; <label>:47                                      ; preds = %24
  br label %48

; <label>:48                                      ; preds = %47, %11
  %49 = load i32* @line, align 4
  %50 = add nsw i32 %49, 1
  store i32 %50, i32* @line, align 4
  br label %564

; <label>:51                                      ; preds = %6
  %52 = load i32* @tk, align 4
  %53 = icmp eq i32 %52, 35
  br i1 %53, label %54, label %71

; <label>:54                                      ; preds = %51
  br label %55

; <label>:55                                      ; preds = %67, %54
  %56 = load i8** @p, align 8
  %57 = load i8* %56, align 1
  %58 = sext i8 %57 to i32
  %59 = icmp ne i32 %58, 0
  br i1 %59, label %60, label %65

; <label>:60                                      ; preds = %55
  %61 = load i8** @p, align 8
  %62 = load i8* %61, align 1
  %63 = sext i8 %62 to i32
  %64 = icmp ne i32 %63, 10
  br label %65

; <label>:65                                      ; preds = %60, %55
  %66 = phi i1 [ false, %55 ], [ %64, %60 ]
  br i1 %66, label %67, label %70

; <label>:67                                      ; preds = %65
  %68 = load i8** @p, align 8
  %69 = getelementptr inbounds i8* %68, i32 1
  store i8* %69, i8** @p, align 8
  br label %55

; <label>:70                                      ; preds = %65
  br label %563

; <label>:71                                      ; preds = %51
  %72 = load i32* @tk, align 4
  %73 = icmp sge i32 %72, 97
  br i1 %73, label %74, label %77

; <label>:74                                      ; preds = %71
  %75 = load i32* @tk, align 4
  %76 = icmp sle i32 %75, 122
  br i1 %76, label %86, label %77

; <label>:77                                      ; preds = %74, %71
  %78 = load i32* @tk, align 4
  %79 = icmp sge i32 %78, 65
  br i1 %79, label %80, label %83

; <label>:80                                      ; preds = %77
  %81 = load i32* @tk, align 4
  %82 = icmp sle i32 %81, 90
  br i1 %82, label %86, label %83

; <label>:83                                      ; preds = %80, %77
  %84 = load i32* @tk, align 4
  %85 = icmp eq i32 %84, 95
  br i1 %85, label %86, label %184

; <label>:86                                      ; preds = %83, %80, %74
  %87 = load i8** @p, align 8
  %88 = getelementptr inbounds i8* %87, i64 -1
  br label %89

; <label>:89                                      ; preds = %126, %86
  %90 = load i8** @p, align 8
  %91 = load i8* %90, align 1
  %92 = sext i8 %91 to i32
  %93 = icmp sge i32 %92, 97
  br i1 %93, label %94, label %99

; <label>:94                                      ; preds = %89
  %95 = load i8** @p, align 8
  %96 = load i8* %95, align 1
  %97 = sext i8 %96 to i32
  %98 = icmp sle i32 %97, 122
  br i1 %98, label %124, label %99

; <label>:99                                      ; preds = %94, %89
  %100 = load i8** @p, align 8
  %101 = load i8* %100, align 1
  %102 = sext i8 %101 to i32
  %103 = icmp sge i32 %102, 65
  br i1 %103, label %104, label %109

; <label>:104                                     ; preds = %99
  %105 = load i8** @p, align 8
  %106 = load i8* %105, align 1
  %107 = sext i8 %106 to i32
  %108 = icmp sle i32 %107, 90
  br i1 %108, label %124, label %109

; <label>:109                                     ; preds = %104, %99
  %110 = load i8** @p, align 8
  %111 = load i8* %110, align 1
  %112 = sext i8 %111 to i32
  %113 = icmp sge i32 %112, 48
  br i1 %113, label %114, label %119

; <label>:114                                     ; preds = %109
  %115 = load i8** @p, align 8
  %116 = load i8* %115, align 1
  %117 = sext i8 %116 to i32
  %118 = icmp sle i32 %117, 57
  br i1 %118, label %124, label %119

; <label>:119                                     ; preds = %114, %109
  %120 = load i8** @p, align 8
  %121 = load i8* %120, align 1
  %122 = sext i8 %121 to i32
  %123 = icmp eq i32 %122, 95
  br label %124

; <label>:124                                     ; preds = %119, %114, %104, %94
  %125 = phi i1 [ true, %114 ], [ true, %104 ], [ true, %94 ], [ %123, %119 ]
  br i1 %125, label %126, label %134

; <label>:126                                     ; preds = %124
  %127 = load i32* @tk, align 4
  %128 = mul nsw i32 %127, 147
  %129 = load i8** @p, align 8
  %130 = getelementptr inbounds i8* %129, i32 1
  store i8* %130, i8** @p, align 8
  %131 = load i8* %129, align 1
  %132 = sext i8 %131 to i32
  %133 = add nsw i32 %128, %132
  store i32 %133, i32* @tk, align 4
  br label %89

; <label>:134                                     ; preds = %124
  %135 = load i32* @tk, align 4
  %136 = shl i32 %135, 6
  %137 = sext i32 %136 to i64
  %138 = load i8** @p, align 8
  %139 = ptrtoint i8* %138 to i64
  %140 = ptrtoint i8* %88 to i64
  %141 = sub i64 %139, %140
  %142 = add nsw i64 %137, %141
  %143 = trunc i64 %142 to i32
  store i32 %143, i32* @tk, align 4
  %144 = load i32** @sym, align 8
  store i32* %144, i32** @id, align 8
  br label %145

; <label>:145                                     ; preds = %172, %134
  %146 = load i32** @id, align 8
  %147 = getelementptr inbounds i32* %146, i64 0
  %148 = load i32* %147, align 4
  %149 = icmp ne i32 %148, 0
  br i1 %149, label %150, label %175

; <label>:150                                     ; preds = %145
  %151 = load i32* @tk, align 4
  %152 = load i32** @id, align 8
  %153 = getelementptr inbounds i32* %152, i64 1
  %154 = load i32* %153, align 4
  %155 = icmp eq i32 %151, %154
  br i1 %155, label %156, label %172

; <label>:156                                     ; preds = %150
  %157 = load i32** @id, align 8
  %158 = getelementptr inbounds i32* %157, i64 2
  %159 = load i32* %158, align 4
  %160 = sext i32 %159 to i64
  %161 = inttoptr i64 %160 to i8*
  %162 = load i8** @p, align 8
  %163 = ptrtoint i8* %162 to i64
  %164 = ptrtoint i8* %88 to i64
  %165 = sub i64 %163, %164
  %166 = call i32 @memcmp(i8* %161, i8* %88, i64 %165) #6
  %167 = icmp ne i32 %166, 0
  br i1 %167, label %172, label %168

; <label>:168                                     ; preds = %156
  %169 = load i32** @id, align 8
  %170 = getelementptr inbounds i32* %169, i64 0
  %171 = load i32* %170, align 4
  store i32 %171, i32* @tk, align 4
  br label %565

; <label>:172                                     ; preds = %156, %150
  %173 = load i32** @id, align 8
  %174 = getelementptr inbounds i32* %173, i64 9
  store i32* %174, i32** @id, align 8
  br label %145

; <label>:175                                     ; preds = %145
  %176 = ptrtoint i8* %88 to i32
  %177 = load i32** @id, align 8
  %178 = getelementptr inbounds i32* %177, i64 2
  store i32 %176, i32* %178, align 4
  %179 = load i32* @tk, align 4
  %180 = load i32** @id, align 8
  %181 = getelementptr inbounds i32* %180, i64 1
  store i32 %179, i32* %181, align 4
  %182 = load i32** @id, align 8
  %183 = getelementptr inbounds i32* %182, i64 0
  store i32 133, i32* %183, align 4
  store i32 133, i32* @tk, align 4
  br label %565

; <label>:184                                     ; preds = %83
  %185 = load i32* @tk, align 4
  %186 = icmp sge i32 %185, 48
  br i1 %186, label %187, label %294

; <label>:187                                     ; preds = %184
  %188 = load i32* @tk, align 4
  %189 = icmp sle i32 %188, 57
  br i1 %189, label %190, label %294

; <label>:190                                     ; preds = %187
  %191 = load i32* @tk, align 4
  %192 = sub nsw i32 %191, 48
  store i32 %192, i32* @ival, align 4
  %193 = icmp ne i32 %192, 0
  br i1 %193, label %194, label %217

; <label>:194                                     ; preds = %190
  br label %195

; <label>:195                                     ; preds = %207, %194
  %196 = load i8** @p, align 8
  %197 = load i8* %196, align 1
  %198 = sext i8 %197 to i32
  %199 = icmp sge i32 %198, 48
  br i1 %199, label %200, label %205

; <label>:200                                     ; preds = %195
  %201 = load i8** @p, align 8
  %202 = load i8* %201, align 1
  %203 = sext i8 %202 to i32
  %204 = icmp sle i32 %203, 57
  br label %205

; <label>:205                                     ; preds = %200, %195
  %206 = phi i1 [ false, %195 ], [ %204, %200 ]
  br i1 %206, label %207, label %216

; <label>:207                                     ; preds = %205
  %208 = load i32* @ival, align 4
  %209 = mul nsw i32 %208, 10
  %210 = load i8** @p, align 8
  %211 = getelementptr inbounds i8* %210, i32 1
  store i8* %211, i8** @p, align 8
  %212 = load i8* %210, align 1
  %213 = sext i8 %212 to i32
  %214 = add nsw i32 %209, %213
  %215 = sub nsw i32 %214, 48
  store i32 %215, i32* @ival, align 4
  br label %195

; <label>:216                                     ; preds = %205
  br label %293

; <label>:217                                     ; preds = %190
  %218 = load i8** @p, align 8
  %219 = load i8* %218, align 1
  %220 = sext i8 %219 to i32
  %221 = icmp eq i32 %220, 120
  br i1 %221, label %227, label %222

; <label>:222                                     ; preds = %217
  %223 = load i8** @p, align 8
  %224 = load i8* %223, align 1
  %225 = sext i8 %224 to i32
  %226 = icmp eq i32 %225, 88
  br i1 %226, label %227, label %269

; <label>:227                                     ; preds = %222, %217
  br label %228

; <label>:228                                     ; preds = %258, %227
  %229 = load i8** @p, align 8
  %230 = getelementptr inbounds i8* %229, i32 1
  store i8* %230, i8** @p, align 8
  %231 = load i8* %230, align 1
  %232 = sext i8 %231 to i32
  store i32 %232, i32* @tk, align 4
  %233 = icmp ne i32 %232, 0
  br i1 %233, label %234, label %256

; <label>:234                                     ; preds = %228
  %235 = load i32* @tk, align 4
  %236 = icmp sge i32 %235, 48
  br i1 %236, label %237, label %240

; <label>:237                                     ; preds = %234
  %238 = load i32* @tk, align 4
  %239 = icmp sle i32 %238, 57
  br i1 %239, label %254, label %240

; <label>:240                                     ; preds = %237, %234
  %241 = load i32* @tk, align 4
  %242 = icmp sge i32 %241, 97
  br i1 %242, label %243, label %246

; <label>:243                                     ; preds = %240
  %244 = load i32* @tk, align 4
  %245 = icmp sle i32 %244, 102
  br i1 %245, label %254, label %246

; <label>:246                                     ; preds = %243, %240
  %247 = load i32* @tk, align 4
  %248 = icmp sge i32 %247, 65
  br i1 %248, label %249, label %252

; <label>:249                                     ; preds = %246
  %250 = load i32* @tk, align 4
  %251 = icmp sle i32 %250, 70
  br label %252

; <label>:252                                     ; preds = %249, %246
  %253 = phi i1 [ false, %246 ], [ %251, %249 ]
  br label %254

; <label>:254                                     ; preds = %252, %243, %237
  %255 = phi i1 [ true, %243 ], [ true, %237 ], [ %253, %252 ]
  br label %256

; <label>:256                                     ; preds = %254, %228
  %257 = phi i1 [ false, %228 ], [ %255, %254 ]
  br i1 %257, label %258, label %268

; <label>:258                                     ; preds = %256
  %259 = load i32* @ival, align 4
  %260 = mul nsw i32 %259, 16
  %261 = load i32* @tk, align 4
  %262 = and i32 %261, 15
  %263 = add nsw i32 %260, %262
  %264 = load i32* @tk, align 4
  %265 = icmp sge i32 %264, 65
  %266 = select i1 %265, i32 9, i32 0
  %267 = add nsw i32 %263, %266
  store i32 %267, i32* @ival, align 4
  br label %228

; <label>:268                                     ; preds = %256
  br label %292

; <label>:269                                     ; preds = %222
  br label %270

; <label>:270                                     ; preds = %282, %269
  %271 = load i8** @p, align 8
  %272 = load i8* %271, align 1
  %273 = sext i8 %272 to i32
  %274 = icmp sge i32 %273, 48
  br i1 %274, label %275, label %280

; <label>:275                                     ; preds = %270
  %276 = load i8** @p, align 8
  %277 = load i8* %276, align 1
  %278 = sext i8 %277 to i32
  %279 = icmp sle i32 %278, 55
  br label %280

; <label>:280                                     ; preds = %275, %270
  %281 = phi i1 [ false, %270 ], [ %279, %275 ]
  br i1 %281, label %282, label %291

; <label>:282                                     ; preds = %280
  %283 = load i32* @ival, align 4
  %284 = mul nsw i32 %283, 8
  %285 = load i8** @p, align 8
  %286 = getelementptr inbounds i8* %285, i32 1
  store i8* %286, i8** @p, align 8
  %287 = load i8* %285, align 1
  %288 = sext i8 %287 to i32
  %289 = add nsw i32 %284, %288
  %290 = sub nsw i32 %289, 48
  store i32 %290, i32* @ival, align 4
  br label %270

; <label>:291                                     ; preds = %280
  br label %292

; <label>:292                                     ; preds = %291, %268
  br label %293

; <label>:293                                     ; preds = %292, %216
  store i32 128, i32* @tk, align 4
  br label %565

; <label>:294                                     ; preds = %187, %184
  %295 = load i32* @tk, align 4
  %296 = icmp eq i32 %295, 47
  br i1 %296, label %297, label %323

; <label>:297                                     ; preds = %294
  %298 = load i8** @p, align 8
  %299 = load i8* %298, align 1
  %300 = sext i8 %299 to i32
  %301 = icmp eq i32 %300, 47
  br i1 %301, label %302, label %321

; <label>:302                                     ; preds = %297
  %303 = load i8** @p, align 8
  %304 = getelementptr inbounds i8* %303, i32 1
  store i8* %304, i8** @p, align 8
  br label %305

; <label>:305                                     ; preds = %317, %302
  %306 = load i8** @p, align 8
  %307 = load i8* %306, align 1
  %308 = sext i8 %307 to i32
  %309 = icmp ne i32 %308, 0
  br i1 %309, label %310, label %315

; <label>:310                                     ; preds = %305
  %311 = load i8** @p, align 8
  %312 = load i8* %311, align 1
  %313 = sext i8 %312 to i32
  %314 = icmp ne i32 %313, 10
  br label %315

; <label>:315                                     ; preds = %310, %305
  %316 = phi i1 [ false, %305 ], [ %314, %310 ]
  br i1 %316, label %317, label %320

; <label>:317                                     ; preds = %315
  %318 = load i8** @p, align 8
  %319 = getelementptr inbounds i8* %318, i32 1
  store i8* %319, i8** @p, align 8
  br label %305

; <label>:320                                     ; preds = %315
  br label %322

; <label>:321                                     ; preds = %297
  store i32 160, i32* @tk, align 4
  br label %565

; <label>:322                                     ; preds = %320
  br label %560

; <label>:323                                     ; preds = %294
  %324 = load i32* @tk, align 4
  %325 = icmp eq i32 %324, 39
  br i1 %325, label %329, label %326

; <label>:326                                     ; preds = %323
  %327 = load i32* @tk, align 4
  %328 = icmp eq i32 %327, 34
  br i1 %328, label %329, label %376

; <label>:329                                     ; preds = %326, %323
  %330 = load i8** @data, align 8
  br label %331

; <label>:331                                     ; preds = %366, %329
  %332 = load i8** @p, align 8
  %333 = load i8* %332, align 1
  %334 = sext i8 %333 to i32
  %335 = icmp ne i32 %334, 0
  br i1 %335, label %336, label %342

; <label>:336                                     ; preds = %331
  %337 = load i8** @p, align 8
  %338 = load i8* %337, align 1
  %339 = sext i8 %338 to i32
  %340 = load i32* @tk, align 4
  %341 = icmp ne i32 %339, %340
  br label %342

; <label>:342                                     ; preds = %336, %331
  %343 = phi i1 [ false, %331 ], [ %341, %336 ]
  br i1 %343, label %344, label %367

; <label>:344                                     ; preds = %342
  %345 = load i8** @p, align 8
  %346 = getelementptr inbounds i8* %345, i32 1
  store i8* %346, i8** @p, align 8
  %347 = load i8* %345, align 1
  %348 = sext i8 %347 to i32
  store i32 %348, i32* @ival, align 4
  %349 = icmp eq i32 %348, 92
  br i1 %349, label %350, label %358

; <label>:350                                     ; preds = %344
  %351 = load i8** @p, align 8
  %352 = getelementptr inbounds i8* %351, i32 1
  store i8* %352, i8** @p, align 8
  %353 = load i8* %351, align 1
  %354 = sext i8 %353 to i32
  store i32 %354, i32* @ival, align 4
  %355 = icmp eq i32 %354, 110
  br i1 %355, label %356, label %357

; <label>:356                                     ; preds = %350
  store i32 10, i32* @ival, align 4
  br label %357

; <label>:357                                     ; preds = %356, %350
  br label %358

; <label>:358                                     ; preds = %357, %344
  %359 = load i32* @tk, align 4
  %360 = icmp eq i32 %359, 34
  br i1 %360, label %361, label %366

; <label>:361                                     ; preds = %358
  %362 = load i32* @ival, align 4
  %363 = trunc i32 %362 to i8
  %364 = load i8** @data, align 8
  %365 = getelementptr inbounds i8* %364, i32 1
  store i8* %365, i8** @data, align 8
  store i8 %363, i8* %364, align 1
  br label %366

; <label>:366                                     ; preds = %361, %358
  br label %331

; <label>:367                                     ; preds = %342
  %368 = load i8** @p, align 8
  %369 = getelementptr inbounds i8* %368, i32 1
  store i8* %369, i8** @p, align 8
  %370 = load i32* @tk, align 4
  %371 = icmp eq i32 %370, 34
  br i1 %371, label %372, label %374

; <label>:372                                     ; preds = %367
  %373 = ptrtoint i8* %330 to i32
  store i32 %373, i32* @ival, align 4
  br label %375

; <label>:374                                     ; preds = %367
  store i32 128, i32* @tk, align 4
  br label %375

; <label>:375                                     ; preds = %374, %372
  br label %565

; <label>:376                                     ; preds = %326
  %377 = load i32* @tk, align 4
  %378 = icmp eq i32 %377, 61
  br i1 %378, label %379, label %389

; <label>:379                                     ; preds = %376
  %380 = load i8** @p, align 8
  %381 = load i8* %380, align 1
  %382 = sext i8 %381 to i32
  %383 = icmp eq i32 %382, 61
  br i1 %383, label %384, label %387

; <label>:384                                     ; preds = %379
  %385 = load i8** @p, align 8
  %386 = getelementptr inbounds i8* %385, i32 1
  store i8* %386, i8** @p, align 8
  store i32 149, i32* @tk, align 4
  br label %388

; <label>:387                                     ; preds = %379
  store i32 142, i32* @tk, align 4
  br label %388

; <label>:388                                     ; preds = %387, %384
  br label %565

; <label>:389                                     ; preds = %376
  %390 = load i32* @tk, align 4
  %391 = icmp eq i32 %390, 43
  br i1 %391, label %392, label %402

; <label>:392                                     ; preds = %389
  %393 = load i8** @p, align 8
  %394 = load i8* %393, align 1
  %395 = sext i8 %394 to i32
  %396 = icmp eq i32 %395, 43
  br i1 %396, label %397, label %400

; <label>:397                                     ; preds = %392
  %398 = load i8** @p, align 8
  %399 = getelementptr inbounds i8* %398, i32 1
  store i8* %399, i8** @p, align 8
  store i32 162, i32* @tk, align 4
  br label %401

; <label>:400                                     ; preds = %392
  store i32 157, i32* @tk, align 4
  br label %401

; <label>:401                                     ; preds = %400, %397
  br label %565

; <label>:402                                     ; preds = %389
  %403 = load i32* @tk, align 4
  %404 = icmp eq i32 %403, 45
  br i1 %404, label %405, label %415

; <label>:405                                     ; preds = %402
  %406 = load i8** @p, align 8
  %407 = load i8* %406, align 1
  %408 = sext i8 %407 to i32
  %409 = icmp eq i32 %408, 45
  br i1 %409, label %410, label %413

; <label>:410                                     ; preds = %405
  %411 = load i8** @p, align 8
  %412 = getelementptr inbounds i8* %411, i32 1
  store i8* %412, i8** @p, align 8
  store i32 163, i32* @tk, align 4
  br label %414

; <label>:413                                     ; preds = %405
  store i32 158, i32* @tk, align 4
  br label %414

; <label>:414                                     ; preds = %413, %410
  br label %565

; <label>:415                                     ; preds = %402
  %416 = load i32* @tk, align 4
  %417 = icmp eq i32 %416, 33
  br i1 %417, label %418, label %427

; <label>:418                                     ; preds = %415
  %419 = load i8** @p, align 8
  %420 = load i8* %419, align 1
  %421 = sext i8 %420 to i32
  %422 = icmp eq i32 %421, 61
  br i1 %422, label %423, label %426

; <label>:423                                     ; preds = %418
  %424 = load i8** @p, align 8
  %425 = getelementptr inbounds i8* %424, i32 1
  store i8* %425, i8** @p, align 8
  store i32 150, i32* @tk, align 4
  br label %426

; <label>:426                                     ; preds = %423, %418
  br label %565

; <label>:427                                     ; preds = %415
  %428 = load i32* @tk, align 4
  %429 = icmp eq i32 %428, 60
  br i1 %429, label %430, label %449

; <label>:430                                     ; preds = %427
  %431 = load i8** @p, align 8
  %432 = load i8* %431, align 1
  %433 = sext i8 %432 to i32
  %434 = icmp eq i32 %433, 61
  br i1 %434, label %435, label %438

; <label>:435                                     ; preds = %430
  %436 = load i8** @p, align 8
  %437 = getelementptr inbounds i8* %436, i32 1
  store i8* %437, i8** @p, align 8
  store i32 153, i32* @tk, align 4
  br label %448

; <label>:438                                     ; preds = %430
  %439 = load i8** @p, align 8
  %440 = load i8* %439, align 1
  %441 = sext i8 %440 to i32
  %442 = icmp eq i32 %441, 60
  br i1 %442, label %443, label %446

; <label>:443                                     ; preds = %438
  %444 = load i8** @p, align 8
  %445 = getelementptr inbounds i8* %444, i32 1
  store i8* %445, i8** @p, align 8
  store i32 155, i32* @tk, align 4
  br label %447

; <label>:446                                     ; preds = %438
  store i32 151, i32* @tk, align 4
  br label %447

; <label>:447                                     ; preds = %446, %443
  br label %448

; <label>:448                                     ; preds = %447, %435
  br label %565

; <label>:449                                     ; preds = %427
  %450 = load i32* @tk, align 4
  %451 = icmp eq i32 %450, 62
  br i1 %451, label %452, label %471

; <label>:452                                     ; preds = %449
  %453 = load i8** @p, align 8
  %454 = load i8* %453, align 1
  %455 = sext i8 %454 to i32
  %456 = icmp eq i32 %455, 61
  br i1 %456, label %457, label %460

; <label>:457                                     ; preds = %452
  %458 = load i8** @p, align 8
  %459 = getelementptr inbounds i8* %458, i32 1
  store i8* %459, i8** @p, align 8
  store i32 154, i32* @tk, align 4
  br label %470

; <label>:460                                     ; preds = %452
  %461 = load i8** @p, align 8
  %462 = load i8* %461, align 1
  %463 = sext i8 %462 to i32
  %464 = icmp eq i32 %463, 62
  br i1 %464, label %465, label %468

; <label>:465                                     ; preds = %460
  %466 = load i8** @p, align 8
  %467 = getelementptr inbounds i8* %466, i32 1
  store i8* %467, i8** @p, align 8
  store i32 156, i32* @tk, align 4
  br label %469

; <label>:468                                     ; preds = %460
  store i32 152, i32* @tk, align 4
  br label %469

; <label>:469                                     ; preds = %468, %465
  br label %470

; <label>:470                                     ; preds = %469, %457
  br label %565

; <label>:471                                     ; preds = %449
  %472 = load i32* @tk, align 4
  %473 = icmp eq i32 %472, 124
  br i1 %473, label %474, label %484

; <label>:474                                     ; preds = %471
  %475 = load i8** @p, align 8
  %476 = load i8* %475, align 1
  %477 = sext i8 %476 to i32
  %478 = icmp eq i32 %477, 124
  br i1 %478, label %479, label %482

; <label>:479                                     ; preds = %474
  %480 = load i8** @p, align 8
  %481 = getelementptr inbounds i8* %480, i32 1
  store i8* %481, i8** @p, align 8
  store i32 144, i32* @tk, align 4
  br label %483

; <label>:482                                     ; preds = %474
  store i32 146, i32* @tk, align 4
  br label %483

; <label>:483                                     ; preds = %482, %479
  br label %565

; <label>:484                                     ; preds = %471
  %485 = load i32* @tk, align 4
  %486 = icmp eq i32 %485, 38
  br i1 %486, label %487, label %497

; <label>:487                                     ; preds = %484
  %488 = load i8** @p, align 8
  %489 = load i8* %488, align 1
  %490 = sext i8 %489 to i32
  %491 = icmp eq i32 %490, 38
  br i1 %491, label %492, label %495

; <label>:492                                     ; preds = %487
  %493 = load i8** @p, align 8
  %494 = getelementptr inbounds i8* %493, i32 1
  store i8* %494, i8** @p, align 8
  store i32 145, i32* @tk, align 4
  br label %496

; <label>:495                                     ; preds = %487
  store i32 148, i32* @tk, align 4
  br label %496

; <label>:496                                     ; preds = %495, %492
  br label %565

; <label>:497                                     ; preds = %484
  %498 = load i32* @tk, align 4
  %499 = icmp eq i32 %498, 94
  br i1 %499, label %500, label %501

; <label>:500                                     ; preds = %497
  store i32 147, i32* @tk, align 4
  br label %565

; <label>:501                                     ; preds = %497
  %502 = load i32* @tk, align 4
  %503 = icmp eq i32 %502, 37
  br i1 %503, label %504, label %505

; <label>:504                                     ; preds = %501
  store i32 161, i32* @tk, align 4
  br label %565

; <label>:505                                     ; preds = %501
  %506 = load i32* @tk, align 4
  %507 = icmp eq i32 %506, 42
  br i1 %507, label %508, label %509

; <label>:508                                     ; preds = %505
  store i32 159, i32* @tk, align 4
  br label %565

; <label>:509                                     ; preds = %505
  %510 = load i32* @tk, align 4
  %511 = icmp eq i32 %510, 91
  br i1 %511, label %512, label %513

; <label>:512                                     ; preds = %509
  store i32 164, i32* @tk, align 4
  br label %565

; <label>:513                                     ; preds = %509
  %514 = load i32* @tk, align 4
  %515 = icmp eq i32 %514, 63
  br i1 %515, label %516, label %517

; <label>:516                                     ; preds = %513
  store i32 143, i32* @tk, align 4
  br label %565

; <label>:517                                     ; preds = %513
  %518 = load i32* @tk, align 4
  %519 = icmp eq i32 %518, 126
  br i1 %519, label %544, label %520

; <label>:520                                     ; preds = %517
  %521 = load i32* @tk, align 4
  %522 = icmp eq i32 %521, 59
  br i1 %522, label %544, label %523

; <label>:523                                     ; preds = %520
  %524 = load i32* @tk, align 4
  %525 = icmp eq i32 %524, 123
  br i1 %525, label %544, label %526

; <label>:526                                     ; preds = %523
  %527 = load i32* @tk, align 4
  %528 = icmp eq i32 %527, 125
  br i1 %528, label %544, label %529

; <label>:529                                     ; preds = %526
  %530 = load i32* @tk, align 4
  %531 = icmp eq i32 %530, 40
  br i1 %531, label %544, label %532

; <label>:532                                     ; preds = %529
  %533 = load i32* @tk, align 4
  %534 = icmp eq i32 %533, 41
  br i1 %534, label %544, label %535

; <label>:535                                     ; preds = %532
  %536 = load i32* @tk, align 4
  %537 = icmp eq i32 %536, 93
  br i1 %537, label %544, label %538

; <label>:538                                     ; preds = %535
  %539 = load i32* @tk, align 4
  %540 = icmp eq i32 %539, 44
  br i1 %540, label %544, label %541

; <label>:541                                     ; preds = %538
  %542 = load i32* @tk, align 4
  %543 = icmp eq i32 %542, 58
  br i1 %543, label %544, label %545

; <label>:544                                     ; preds = %541, %538, %535, %532, %529, %526, %523, %520, %517
  br label %565

; <label>:545                                     ; preds = %541
  br label %546

; <label>:546                                     ; preds = %545
  br label %547

; <label>:547                                     ; preds = %546
  br label %548

; <label>:548                                     ; preds = %547
  br label %549

; <label>:549                                     ; preds = %548
  br label %550

; <label>:550                                     ; preds = %549
  br label %551

; <label>:551                                     ; preds = %550
  br label %552

; <label>:552                                     ; preds = %551
  br label %553

; <label>:553                                     ; preds = %552
  br label %554

; <label>:554                                     ; preds = %553
  br label %555

; <label>:555                                     ; preds = %554
  br label %556

; <label>:556                                     ; preds = %555
  br label %557

; <label>:557                                     ; preds = %556
  br label %558

; <label>:558                                     ; preds = %557
  br label %559

; <label>:559                                     ; preds = %558
  br label %560

; <label>:560                                     ; preds = %559, %322
  br label %561

; <label>:561                                     ; preds = %560
  br label %562

; <label>:562                                     ; preds = %561
  br label %563

; <label>:563                                     ; preds = %562, %70
  br label %564

; <label>:564                                     ; preds = %563, %48
  br label %1

; <label>:565                                     ; preds = %544, %516, %512, %508, %504, %500, %496, %483, %470, %448, %426, %414, %401, %388, %375, %321, %293, %175, %168, %1
  ret void
}

declare i32 @printf(i8*, ...) #1

; Function Attrs: nounwind readonly
declare i32 @memcmp(i8*, i8*, i64) #2

; Function Attrs: nounwind uwtable
define void @expr(i32 %lev) #0 {
  %1 = load i32* @tk, align 4
  %2 = icmp ne i32 %1, 0
  br i1 %2, label %6, label %3

; <label>:3                                       ; preds = %0
  %4 = load i32* @line, align 4
  %5 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([34 x i8]* @.str5, i32 0, i32 0), i32 %4)
  call void @exit(i32 -1) #7
  unreachable

; <label>:6                                       ; preds = %0
  %7 = load i32* @tk, align 4
  %8 = icmp eq i32 %7, 128
  br i1 %8, label %9, label %15

; <label>:9                                       ; preds = %6
  %10 = load i32** @e, align 8
  %11 = getelementptr inbounds i32* %10, i32 1
  store i32* %11, i32** @e, align 8
  store i32 1, i32* %11, align 4
  %12 = load i32* @ival, align 4
  %13 = load i32** @e, align 8
  %14 = getelementptr inbounds i32* %13, i32 1
  store i32* %14, i32** @e, align 8
  store i32 %12, i32* %14, align 4
  call void @next()
  store i32 1, i32* @ty, align 4
  br label %361

; <label>:15                                      ; preds = %6
  %16 = load i32* @tk, align 4
  %17 = icmp eq i32 %16, 34
  br i1 %17, label %18, label %35

; <label>:18                                      ; preds = %15
  %19 = load i32** @e, align 8
  %20 = getelementptr inbounds i32* %19, i32 1
  store i32* %20, i32** @e, align 8
  store i32 1, i32* %20, align 4
  %21 = load i32* @ival, align 4
  %22 = load i32** @e, align 8
  %23 = getelementptr inbounds i32* %22, i32 1
  store i32* %23, i32** @e, align 8
  store i32 %21, i32* %23, align 4
  call void @next()
  br label %24

; <label>:24                                      ; preds = %27, %18
  %25 = load i32* @tk, align 4
  %26 = icmp eq i32 %25, 34
  br i1 %26, label %27, label %28

; <label>:27                                      ; preds = %24
  call void @next()
  br label %24

; <label>:28                                      ; preds = %24
  %29 = load i8** @data, align 8
  %30 = ptrtoint i8* %29 to i32
  %31 = sext i32 %30 to i64
  %32 = add i64 %31, 4
  %33 = and i64 %32, -4
  %34 = inttoptr i64 %33 to i8*
  store i8* %34, i8** @data, align 8
  store i32 2, i32* @ty, align 4
  br label %360

; <label>:35                                      ; preds = %15
  %36 = load i32* @tk, align 4
  %37 = icmp eq i32 %36, 140
  br i1 %37, label %38, label %77

; <label>:38                                      ; preds = %35
  call void @next()
  %39 = load i32* @tk, align 4
  %40 = icmp eq i32 %39, 40
  br i1 %40, label %41, label %42

; <label>:41                                      ; preds = %38
  call void @next()
  br label %45

; <label>:42                                      ; preds = %38
  %43 = load i32* @line, align 4
  %44 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([35 x i8]* @.str6, i32 0, i32 0), i32 %43)
  call void @exit(i32 -1) #7
  unreachable

; <label>:45                                      ; preds = %41
  store i32 1, i32* @ty, align 4
  %46 = load i32* @tk, align 4
  %47 = icmp eq i32 %46, 138
  br i1 %47, label %48, label %49

; <label>:48                                      ; preds = %45
  call void @next()
  br label %54

; <label>:49                                      ; preds = %45
  %50 = load i32* @tk, align 4
  %51 = icmp eq i32 %50, 134
  br i1 %51, label %52, label %53

; <label>:52                                      ; preds = %49
  call void @next()
  store i32 0, i32* @ty, align 4
  br label %53

; <label>:53                                      ; preds = %52, %49
  br label %54

; <label>:54                                      ; preds = %53, %48
  br label %55

; <label>:55                                      ; preds = %58, %54
  %56 = load i32* @tk, align 4
  %57 = icmp eq i32 %56, 159
  br i1 %57, label %58, label %61

; <label>:58                                      ; preds = %55
  call void @next()
  %59 = load i32* @ty, align 4
  %60 = add nsw i32 %59, 2
  store i32 %60, i32* @ty, align 4
  br label %55

; <label>:61                                      ; preds = %55
  %62 = load i32* @tk, align 4
  %63 = icmp eq i32 %62, 41
  br i1 %63, label %64, label %65

; <label>:64                                      ; preds = %61
  call void @next()
  br label %68

; <label>:65                                      ; preds = %61
  %66 = load i32* @line, align 4
  %67 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([36 x i8]* @.str7, i32 0, i32 0), i32 %66)
  call void @exit(i32 -1) #7
  unreachable

; <label>:68                                      ; preds = %64
  %69 = load i32** @e, align 8
  %70 = getelementptr inbounds i32* %69, i32 1
  store i32* %70, i32** @e, align 8
  store i32 1, i32* %70, align 4
  %71 = load i32* @ty, align 4
  %72 = icmp eq i32 %71, 0
  %73 = select i1 %72, i64 1, i64 4
  %74 = trunc i64 %73 to i32
  %75 = load i32** @e, align 8
  %76 = getelementptr inbounds i32* %75, i32 1
  store i32* %76, i32** @e, align 8
  store i32 %74, i32* %76, align 4
  store i32 1, i32* @ty, align 4
  br label %359

; <label>:77                                      ; preds = %35
  %78 = load i32* @tk, align 4
  %79 = icmp eq i32 %78, 133
  br i1 %79, label %80, label %178

; <label>:80                                      ; preds = %77
  %81 = load i32** @id, align 8
  call void @next()
  %82 = load i32* @tk, align 4
  %83 = icmp eq i32 %82, 40
  br i1 %83, label %84, label %130

; <label>:84                                      ; preds = %80
  call void @next()
  br label %85

; <label>:85                                      ; preds = %95, %84
  %t.0 = phi i32 [ 0, %84 ], [ %91, %95 ]
  %86 = load i32* @tk, align 4
  %87 = icmp ne i32 %86, 41
  br i1 %87, label %88, label %96

; <label>:88                                      ; preds = %85
  call void @expr(i32 142)
  %89 = load i32** @e, align 8
  %90 = getelementptr inbounds i32* %89, i32 1
  store i32* %90, i32** @e, align 8
  store i32 13, i32* %90, align 4
  %91 = add nsw i32 %t.0, 1
  %92 = load i32* @tk, align 4
  %93 = icmp eq i32 %92, 44
  br i1 %93, label %94, label %95

; <label>:94                                      ; preds = %88
  call void @next()
  br label %95

; <label>:95                                      ; preds = %94, %88
  br label %85

; <label>:96                                      ; preds = %85
  call void @next()
  %97 = getelementptr inbounds i32* %81, i64 3
  %98 = load i32* %97, align 4
  %99 = icmp eq i32 %98, 130
  br i1 %99, label %100, label %105

; <label>:100                                     ; preds = %96
  %101 = getelementptr inbounds i32* %81, i64 5
  %102 = load i32* %101, align 4
  %103 = load i32** @e, align 8
  %104 = getelementptr inbounds i32* %103, i32 1
  store i32* %104, i32** @e, align 8
  store i32 %102, i32* %104, align 4
  br label %120

; <label>:105                                     ; preds = %96
  %106 = getelementptr inbounds i32* %81, i64 3
  %107 = load i32* %106, align 4
  %108 = icmp eq i32 %107, 129
  br i1 %108, label %109, label %116

; <label>:109                                     ; preds = %105
  %110 = load i32** @e, align 8
  %111 = getelementptr inbounds i32* %110, i32 1
  store i32* %111, i32** @e, align 8
  store i32 3, i32* %111, align 4
  %112 = getelementptr inbounds i32* %81, i64 5
  %113 = load i32* %112, align 4
  %114 = load i32** @e, align 8
  %115 = getelementptr inbounds i32* %114, i32 1
  store i32* %115, i32** @e, align 8
  store i32 %113, i32* %115, align 4
  br label %119

; <label>:116                                     ; preds = %105
  %117 = load i32* @line, align 4
  %118 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([23 x i8]* @.str8, i32 0, i32 0), i32 %117)
  call void @exit(i32 -1) #7
  unreachable

; <label>:119                                     ; preds = %109
  br label %120

; <label>:120                                     ; preds = %119, %100
  %121 = icmp ne i32 %t.0, 0
  br i1 %121, label %122, label %127

; <label>:122                                     ; preds = %120
  %123 = load i32** @e, align 8
  %124 = getelementptr inbounds i32* %123, i32 1
  store i32* %124, i32** @e, align 8
  store i32 7, i32* %124, align 4
  %125 = load i32** @e, align 8
  %126 = getelementptr inbounds i32* %125, i32 1
  store i32* %126, i32** @e, align 8
  store i32 %t.0, i32* %126, align 4
  br label %127

; <label>:127                                     ; preds = %122, %120
  %128 = getelementptr inbounds i32* %81, i64 4
  %129 = load i32* %128, align 4
  store i32 %129, i32* @ty, align 4
  br label %177

; <label>:130                                     ; preds = %80
  %131 = getelementptr inbounds i32* %81, i64 3
  %132 = load i32* %131, align 4
  %133 = icmp eq i32 %132, 128
  br i1 %133, label %134, label %141

; <label>:134                                     ; preds = %130
  %135 = load i32** @e, align 8
  %136 = getelementptr inbounds i32* %135, i32 1
  store i32* %136, i32** @e, align 8
  store i32 1, i32* %136, align 4
  %137 = getelementptr inbounds i32* %81, i64 5
  %138 = load i32* %137, align 4
  %139 = load i32** @e, align 8
  %140 = getelementptr inbounds i32* %139, i32 1
  store i32* %140, i32** @e, align 8
  store i32 %138, i32* %140, align 4
  store i32 1, i32* @ty, align 4
  br label %176

; <label>:141                                     ; preds = %130
  %142 = getelementptr inbounds i32* %81, i64 3
  %143 = load i32* %142, align 4
  %144 = icmp eq i32 %143, 132
  br i1 %144, label %145, label %154

; <label>:145                                     ; preds = %141
  %146 = load i32** @e, align 8
  %147 = getelementptr inbounds i32* %146, i32 1
  store i32* %147, i32** @e, align 8
  store i32 0, i32* %147, align 4
  %148 = load i32* @loc, align 4
  %149 = getelementptr inbounds i32* %81, i64 5
  %150 = load i32* %149, align 4
  %151 = sub nsw i32 %148, %150
  %152 = load i32** @e, align 8
  %153 = getelementptr inbounds i32* %152, i32 1
  store i32* %153, i32** @e, align 8
  store i32 %151, i32* %153, align 4
  br label %169

; <label>:154                                     ; preds = %141
  %155 = getelementptr inbounds i32* %81, i64 3
  %156 = load i32* %155, align 4
  %157 = icmp eq i32 %156, 131
  br i1 %157, label %158, label %165

; <label>:158                                     ; preds = %154
  %159 = load i32** @e, align 8
  %160 = getelementptr inbounds i32* %159, i32 1
  store i32* %160, i32** @e, align 8
  store i32 1, i32* %160, align 4
  %161 = getelementptr inbounds i32* %81, i64 5
  %162 = load i32* %161, align 4
  %163 = load i32** @e, align 8
  %164 = getelementptr inbounds i32* %163, i32 1
  store i32* %164, i32** @e, align 8
  store i32 %162, i32* %164, align 4
  br label %168

; <label>:165                                     ; preds = %154
  %166 = load i32* @line, align 4
  %167 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([24 x i8]* @.str9, i32 0, i32 0), i32 %166)
  call void @exit(i32 -1) #7
  unreachable

; <label>:168                                     ; preds = %158
  br label %169

; <label>:169                                     ; preds = %168, %145
  %170 = getelementptr inbounds i32* %81, i64 4
  %171 = load i32* %170, align 4
  store i32 %171, i32* @ty, align 4
  %172 = icmp eq i32 %171, 0
  %173 = select i1 %172, i32 10, i32 9
  %174 = load i32** @e, align 8
  %175 = getelementptr inbounds i32* %174, i32 1
  store i32* %175, i32** @e, align 8
  store i32 %173, i32* %175, align 4
  br label %176

; <label>:176                                     ; preds = %169, %134
  br label %177

; <label>:177                                     ; preds = %176, %127
  br label %358

; <label>:178                                     ; preds = %77
  %179 = load i32* @tk, align 4
  %180 = icmp eq i32 %179, 40
  br i1 %180, label %181, label %213

; <label>:181                                     ; preds = %178
  call void @next()
  %182 = load i32* @tk, align 4
  %183 = icmp eq i32 %182, 138
  br i1 %183, label %187, label %184

; <label>:184                                     ; preds = %181
  %185 = load i32* @tk, align 4
  %186 = icmp eq i32 %185, 134
  br i1 %186, label %187, label %204

; <label>:187                                     ; preds = %184, %181
  %188 = load i32* @tk, align 4
  %189 = icmp eq i32 %188, 138
  %190 = select i1 %189, i32 1, i32 0
  call void @next()
  br label %191

; <label>:191                                     ; preds = %194, %187
  %t.1 = phi i32 [ %190, %187 ], [ %195, %194 ]
  %192 = load i32* @tk, align 4
  %193 = icmp eq i32 %192, 159
  br i1 %193, label %194, label %196

; <label>:194                                     ; preds = %191
  call void @next()
  %195 = add nsw i32 %t.1, 2
  br label %191

; <label>:196                                     ; preds = %191
  %197 = load i32* @tk, align 4
  %198 = icmp eq i32 %197, 41
  br i1 %198, label %199, label %200

; <label>:199                                     ; preds = %196
  call void @next()
  br label %203

; <label>:200                                     ; preds = %196
  %201 = load i32* @line, align 4
  %202 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([14 x i8]* @.str10, i32 0, i32 0), i32 %201)
  call void @exit(i32 -1) #7
  unreachable

; <label>:203                                     ; preds = %199
  call void @expr(i32 162)
  store i32 %t.1, i32* @ty, align 4
  br label %212

; <label>:204                                     ; preds = %184
  call void @expr(i32 142)
  %205 = load i32* @tk, align 4
  %206 = icmp eq i32 %205, 41
  br i1 %206, label %207, label %208

; <label>:207                                     ; preds = %204
  call void @next()
  br label %211

; <label>:208                                     ; preds = %204
  %209 = load i32* @line, align 4
  %210 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([26 x i8]* @.str11, i32 0, i32 0), i32 %209)
  call void @exit(i32 -1) #7
  unreachable

; <label>:211                                     ; preds = %207
  br label %212

; <label>:212                                     ; preds = %211, %203
  br label %357

; <label>:213                                     ; preds = %178
  %214 = load i32* @tk, align 4
  %215 = icmp eq i32 %214, 159
  br i1 %215, label %216, label %231

; <label>:216                                     ; preds = %213
  call void @next()
  call void @expr(i32 162)
  %217 = load i32* @ty, align 4
  %218 = icmp sgt i32 %217, 1
  br i1 %218, label %219, label %222

; <label>:219                                     ; preds = %216
  %220 = load i32* @ty, align 4
  %221 = sub nsw i32 %220, 2
  store i32 %221, i32* @ty, align 4
  br label %225

; <label>:222                                     ; preds = %216
  %223 = load i32* @line, align 4
  %224 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([21 x i8]* @.str12, i32 0, i32 0), i32 %223)
  call void @exit(i32 -1) #7
  unreachable

; <label>:225                                     ; preds = %219
  %226 = load i32* @ty, align 4
  %227 = icmp eq i32 %226, 0
  %228 = select i1 %227, i32 10, i32 9
  %229 = load i32** @e, align 8
  %230 = getelementptr inbounds i32* %229, i32 1
  store i32* %230, i32** @e, align 8
  store i32 %228, i32* %230, align 4
  br label %356

; <label>:231                                     ; preds = %213
  %232 = load i32* @tk, align 4
  %233 = icmp eq i32 %232, 148
  br i1 %233, label %234, label %251

; <label>:234                                     ; preds = %231
  call void @next()
  call void @expr(i32 162)
  %235 = load i32** @e, align 8
  %236 = load i32* %235, align 4
  %237 = icmp eq i32 %236, 10
  br i1 %237, label %242, label %238

; <label>:238                                     ; preds = %234
  %239 = load i32** @e, align 8
  %240 = load i32* %239, align 4
  %241 = icmp eq i32 %240, 9
  br i1 %241, label %242, label %245

; <label>:242                                     ; preds = %238, %234
  %243 = load i32** @e, align 8
  %244 = getelementptr inbounds i32* %243, i32 -1
  store i32* %244, i32** @e, align 8
  br label %248

; <label>:245                                     ; preds = %238
  %246 = load i32* @line, align 4
  %247 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([20 x i8]* @.str13, i32 0, i32 0), i32 %246)
  call void @exit(i32 -1) #7
  unreachable

; <label>:248                                     ; preds = %242
  %249 = load i32* @ty, align 4
  %250 = add nsw i32 %249, 2
  store i32 %250, i32* @ty, align 4
  br label %355

; <label>:251                                     ; preds = %231
  %252 = load i32* @tk, align 4
  %253 = icmp eq i32 %252, 33
  br i1 %253, label %254, label %263

; <label>:254                                     ; preds = %251
  call void @next()
  call void @expr(i32 162)
  %255 = load i32** @e, align 8
  %256 = getelementptr inbounds i32* %255, i32 1
  store i32* %256, i32** @e, align 8
  store i32 13, i32* %256, align 4
  %257 = load i32** @e, align 8
  %258 = getelementptr inbounds i32* %257, i32 1
  store i32* %258, i32** @e, align 8
  store i32 1, i32* %258, align 4
  %259 = load i32** @e, align 8
  %260 = getelementptr inbounds i32* %259, i32 1
  store i32* %260, i32** @e, align 8
  store i32 0, i32* %260, align 4
  %261 = load i32** @e, align 8
  %262 = getelementptr inbounds i32* %261, i32 1
  store i32* %262, i32** @e, align 8
  store i32 17, i32* %262, align 4
  store i32 1, i32* @ty, align 4
  br label %354

; <label>:263                                     ; preds = %251
  %264 = load i32* @tk, align 4
  %265 = icmp eq i32 %264, 126
  br i1 %265, label %266, label %275

; <label>:266                                     ; preds = %263
  call void @next()
  call void @expr(i32 162)
  %267 = load i32** @e, align 8
  %268 = getelementptr inbounds i32* %267, i32 1
  store i32* %268, i32** @e, align 8
  store i32 13, i32* %268, align 4
  %269 = load i32** @e, align 8
  %270 = getelementptr inbounds i32* %269, i32 1
  store i32* %270, i32** @e, align 8
  store i32 1, i32* %270, align 4
  %271 = load i32** @e, align 8
  %272 = getelementptr inbounds i32* %271, i32 1
  store i32* %272, i32** @e, align 8
  store i32 -1, i32* %272, align 4
  %273 = load i32** @e, align 8
  %274 = getelementptr inbounds i32* %273, i32 1
  store i32* %274, i32** @e, align 8
  store i32 15, i32* %274, align 4
  store i32 1, i32* @ty, align 4
  br label %353

; <label>:275                                     ; preds = %263
  %276 = load i32* @tk, align 4
  %277 = icmp eq i32 %276, 157
  br i1 %277, label %278, label %279

; <label>:278                                     ; preds = %275
  call void @next()
  call void @expr(i32 162)
  store i32 1, i32* @ty, align 4
  br label %352

; <label>:279                                     ; preds = %275
  %280 = load i32* @tk, align 4
  %281 = icmp eq i32 %280, 158
  br i1 %281, label %282, label %300

; <label>:282                                     ; preds = %279
  call void @next()
  %283 = load i32** @e, align 8
  %284 = getelementptr inbounds i32* %283, i32 1
  store i32* %284, i32** @e, align 8
  store i32 1, i32* %284, align 4
  %285 = load i32* @tk, align 4
  %286 = icmp eq i32 %285, 128
  br i1 %286, label %287, label %292

; <label>:287                                     ; preds = %282
  %288 = load i32* @ival, align 4
  %289 = sub nsw i32 0, %288
  %290 = load i32** @e, align 8
  %291 = getelementptr inbounds i32* %290, i32 1
  store i32* %291, i32** @e, align 8
  store i32 %289, i32* %291, align 4
  call void @next()
  br label %299

; <label>:292                                     ; preds = %282
  %293 = load i32** @e, align 8
  %294 = getelementptr inbounds i32* %293, i32 1
  store i32* %294, i32** @e, align 8
  store i32 -1, i32* %294, align 4
  %295 = load i32** @e, align 8
  %296 = getelementptr inbounds i32* %295, i32 1
  store i32* %296, i32** @e, align 8
  store i32 13, i32* %296, align 4
  call void @expr(i32 162)
  %297 = load i32** @e, align 8
  %298 = getelementptr inbounds i32* %297, i32 1
  store i32* %298, i32** @e, align 8
  store i32 27, i32* %298, align 4
  br label %299

; <label>:299                                     ; preds = %292, %287
  store i32 1, i32* @ty, align 4
  br label %351

; <label>:300                                     ; preds = %279
  %301 = load i32* @tk, align 4
  %302 = icmp eq i32 %301, 162
  br i1 %302, label %306, label %303

; <label>:303                                     ; preds = %300
  %304 = load i32* @tk, align 4
  %305 = icmp eq i32 %304, 163
  br i1 %305, label %306, label %347

; <label>:306                                     ; preds = %303, %300
  %307 = load i32* @tk, align 4
  call void @next()
  call void @expr(i32 162)
  %308 = load i32** @e, align 8
  %309 = load i32* %308, align 4
  %310 = icmp eq i32 %309, 10
  br i1 %310, label %311, label %315

; <label>:311                                     ; preds = %306
  %312 = load i32** @e, align 8
  store i32 13, i32* %312, align 4
  %313 = load i32** @e, align 8
  %314 = getelementptr inbounds i32* %313, i32 1
  store i32* %314, i32** @e, align 8
  store i32 10, i32* %314, align 4
  br label %327

; <label>:315                                     ; preds = %306
  %316 = load i32** @e, align 8
  %317 = load i32* %316, align 4
  %318 = icmp eq i32 %317, 9
  br i1 %318, label %319, label %323

; <label>:319                                     ; preds = %315
  %320 = load i32** @e, align 8
  store i32 13, i32* %320, align 4
  %321 = load i32** @e, align 8
  %322 = getelementptr inbounds i32* %321, i32 1
  store i32* %322, i32** @e, align 8
  store i32 9, i32* %322, align 4
  br label %326

; <label>:323                                     ; preds = %315
  %324 = load i32* @line, align 4
  %325 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([33 x i8]* @.str14, i32 0, i32 0), i32 %324)
  call void @exit(i32 -1) #7
  unreachable

; <label>:326                                     ; preds = %319
  br label %327

; <label>:327                                     ; preds = %326, %311
  %328 = load i32** @e, align 8
  %329 = getelementptr inbounds i32* %328, i32 1
  store i32* %329, i32** @e, align 8
  store i32 13, i32* %329, align 4
  %330 = load i32** @e, align 8
  %331 = getelementptr inbounds i32* %330, i32 1
  store i32* %331, i32** @e, align 8
  store i32 1, i32* %331, align 4
  %332 = load i32* @ty, align 4
  %333 = icmp sgt i32 %332, 2
  %334 = select i1 %333, i64 4, i64 1
  %335 = trunc i64 %334 to i32
  %336 = load i32** @e, align 8
  %337 = getelementptr inbounds i32* %336, i32 1
  store i32* %337, i32** @e, align 8
  store i32 %335, i32* %337, align 4
  %338 = icmp eq i32 %307, 162
  %339 = select i1 %338, i32 25, i32 26
  %340 = load i32** @e, align 8
  %341 = getelementptr inbounds i32* %340, i32 1
  store i32* %341, i32** @e, align 8
  store i32 %339, i32* %341, align 4
  %342 = load i32* @ty, align 4
  %343 = icmp eq i32 %342, 0
  %344 = select i1 %343, i32 12, i32 11
  %345 = load i32** @e, align 8
  %346 = getelementptr inbounds i32* %345, i32 1
  store i32* %346, i32** @e, align 8
  store i32 %344, i32* %346, align 4
  br label %350

; <label>:347                                     ; preds = %303
  %348 = load i32* @line, align 4
  %349 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([20 x i8]* @.str15, i32 0, i32 0), i32 %348)
  call void @exit(i32 -1) #7
  unreachable

; <label>:350                                     ; preds = %327
  br label %351

; <label>:351                                     ; preds = %350, %299
  br label %352

; <label>:352                                     ; preds = %351, %278
  br label %353

; <label>:353                                     ; preds = %352, %266
  br label %354

; <label>:354                                     ; preds = %353, %254
  br label %355

; <label>:355                                     ; preds = %354, %248
  br label %356

; <label>:356                                     ; preds = %355, %225
  br label %357

; <label>:357                                     ; preds = %356, %212
  br label %358

; <label>:358                                     ; preds = %357, %177
  br label %359

; <label>:359                                     ; preds = %358, %68
  br label %360

; <label>:360                                     ; preds = %359, %28
  br label %361

; <label>:361                                     ; preds = %360, %9
  br label %362

; <label>:362                                     ; preds = %361
  br label %363

; <label>:363                                     ; preds = %729, %362
  %364 = load i32* @tk, align 4
  %365 = icmp sge i32 %364, %lev
  br i1 %365, label %366, label %730

; <label>:366                                     ; preds = %363
  %367 = load i32* @ty, align 4
  %368 = load i32* @tk, align 4
  %369 = icmp eq i32 %368, 142
  br i1 %369, label %370, label %388

; <label>:370                                     ; preds = %366
  call void @next()
  %371 = load i32** @e, align 8
  %372 = load i32* %371, align 4
  %373 = icmp eq i32 %372, 10
  br i1 %373, label %378, label %374

; <label>:374                                     ; preds = %370
  %375 = load i32** @e, align 8
  %376 = load i32* %375, align 4
  %377 = icmp eq i32 %376, 9
  br i1 %377, label %378, label %380

; <label>:378                                     ; preds = %374, %370
  %379 = load i32** @e, align 8
  store i32 13, i32* %379, align 4
  br label %383

; <label>:380                                     ; preds = %374
  %381 = load i32* @line, align 4
  %382 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([30 x i8]* @.str16, i32 0, i32 0), i32 %381)
  call void @exit(i32 -1) #7
  unreachable

; <label>:383                                     ; preds = %378
  call void @expr(i32 142)
  store i32 %367, i32* @ty, align 4
  %384 = icmp eq i32 %367, 0
  %385 = select i1 %384, i32 12, i32 11
  %386 = load i32** @e, align 8
  %387 = getelementptr inbounds i32* %386, i32 1
  store i32* %387, i32** @e, align 8
  store i32 %385, i32* %387, align 4
  br label %729

; <label>:388                                     ; preds = %366
  %389 = load i32* @tk, align 4
  %390 = icmp eq i32 %389, 143
  br i1 %390, label %391, label %413

; <label>:391                                     ; preds = %388
  call void @next()
  %392 = load i32** @e, align 8
  %393 = getelementptr inbounds i32* %392, i32 1
  store i32* %393, i32** @e, align 8
  store i32 4, i32* %393, align 4
  %394 = load i32** @e, align 8
  %395 = getelementptr inbounds i32* %394, i32 1
  store i32* %395, i32** @e, align 8
  call void @expr(i32 142)
  %396 = load i32* @tk, align 4
  %397 = icmp eq i32 %396, 58
  br i1 %397, label %398, label %399

; <label>:398                                     ; preds = %391
  call void @next()
  br label %402

; <label>:399                                     ; preds = %391
  %400 = load i32* @line, align 4
  %401 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([31 x i8]* @.str17, i32 0, i32 0), i32 %400)
  call void @exit(i32 -1) #7
  unreachable

; <label>:402                                     ; preds = %398
  %403 = load i32** @e, align 8
  %404 = getelementptr inbounds i32* %403, i64 3
  %405 = ptrtoint i32* %404 to i32
  store i32 %405, i32* %395, align 4
  %406 = load i32** @e, align 8
  %407 = getelementptr inbounds i32* %406, i32 1
  store i32* %407, i32** @e, align 8
  store i32 2, i32* %407, align 4
  %408 = load i32** @e, align 8
  %409 = getelementptr inbounds i32* %408, i32 1
  store i32* %409, i32** @e, align 8
  call void @expr(i32 143)
  %410 = load i32** @e, align 8
  %411 = getelementptr inbounds i32* %410, i64 1
  %412 = ptrtoint i32* %411 to i32
  store i32 %412, i32* %409, align 4
  br label %728

; <label>:413                                     ; preds = %388
  %414 = load i32* @tk, align 4
  %415 = icmp eq i32 %414, 144
  br i1 %415, label %416, label %424

; <label>:416                                     ; preds = %413
  call void @next()
  %417 = load i32** @e, align 8
  %418 = getelementptr inbounds i32* %417, i32 1
  store i32* %418, i32** @e, align 8
  store i32 5, i32* %418, align 4
  %419 = load i32** @e, align 8
  %420 = getelementptr inbounds i32* %419, i32 1
  store i32* %420, i32** @e, align 8
  call void @expr(i32 145)
  %421 = load i32** @e, align 8
  %422 = getelementptr inbounds i32* %421, i64 1
  %423 = ptrtoint i32* %422 to i32
  store i32 %423, i32* %420, align 4
  store i32 1, i32* @ty, align 4
  br label %727

; <label>:424                                     ; preds = %413
  %425 = load i32* @tk, align 4
  %426 = icmp eq i32 %425, 145
  br i1 %426, label %427, label %435

; <label>:427                                     ; preds = %424
  call void @next()
  %428 = load i32** @e, align 8
  %429 = getelementptr inbounds i32* %428, i32 1
  store i32* %429, i32** @e, align 8
  store i32 4, i32* %429, align 4
  %430 = load i32** @e, align 8
  %431 = getelementptr inbounds i32* %430, i32 1
  store i32* %431, i32** @e, align 8
  call void @expr(i32 146)
  %432 = load i32** @e, align 8
  %433 = getelementptr inbounds i32* %432, i64 1
  %434 = ptrtoint i32* %433 to i32
  store i32 %434, i32* %431, align 4
  store i32 1, i32* @ty, align 4
  br label %726

; <label>:435                                     ; preds = %424
  %436 = load i32* @tk, align 4
  %437 = icmp eq i32 %436, 146
  br i1 %437, label %438, label %443

; <label>:438                                     ; preds = %435
  call void @next()
  %439 = load i32** @e, align 8
  %440 = getelementptr inbounds i32* %439, i32 1
  store i32* %440, i32** @e, align 8
  store i32 13, i32* %440, align 4
  call void @expr(i32 147)
  %441 = load i32** @e, align 8
  %442 = getelementptr inbounds i32* %441, i32 1
  store i32* %442, i32** @e, align 8
  store i32 14, i32* %442, align 4
  store i32 1, i32* @ty, align 4
  br label %725

; <label>:443                                     ; preds = %435
  %444 = load i32* @tk, align 4
  %445 = icmp eq i32 %444, 147
  br i1 %445, label %446, label %451

; <label>:446                                     ; preds = %443
  call void @next()
  %447 = load i32** @e, align 8
  %448 = getelementptr inbounds i32* %447, i32 1
  store i32* %448, i32** @e, align 8
  store i32 13, i32* %448, align 4
  call void @expr(i32 148)
  %449 = load i32** @e, align 8
  %450 = getelementptr inbounds i32* %449, i32 1
  store i32* %450, i32** @e, align 8
  store i32 15, i32* %450, align 4
  store i32 1, i32* @ty, align 4
  br label %724

; <label>:451                                     ; preds = %443
  %452 = load i32* @tk, align 4
  %453 = icmp eq i32 %452, 148
  br i1 %453, label %454, label %459

; <label>:454                                     ; preds = %451
  call void @next()
  %455 = load i32** @e, align 8
  %456 = getelementptr inbounds i32* %455, i32 1
  store i32* %456, i32** @e, align 8
  store i32 13, i32* %456, align 4
  call void @expr(i32 149)
  %457 = load i32** @e, align 8
  %458 = getelementptr inbounds i32* %457, i32 1
  store i32* %458, i32** @e, align 8
  store i32 16, i32* %458, align 4
  store i32 1, i32* @ty, align 4
  br label %723

; <label>:459                                     ; preds = %451
  %460 = load i32* @tk, align 4
  %461 = icmp eq i32 %460, 149
  br i1 %461, label %462, label %467

; <label>:462                                     ; preds = %459
  call void @next()
  %463 = load i32** @e, align 8
  %464 = getelementptr inbounds i32* %463, i32 1
  store i32* %464, i32** @e, align 8
  store i32 13, i32* %464, align 4
  call void @expr(i32 151)
  %465 = load i32** @e, align 8
  %466 = getelementptr inbounds i32* %465, i32 1
  store i32* %466, i32** @e, align 8
  store i32 17, i32* %466, align 4
  store i32 1, i32* @ty, align 4
  br label %722

; <label>:467                                     ; preds = %459
  %468 = load i32* @tk, align 4
  %469 = icmp eq i32 %468, 150
  br i1 %469, label %470, label %475

; <label>:470                                     ; preds = %467
  call void @next()
  %471 = load i32** @e, align 8
  %472 = getelementptr inbounds i32* %471, i32 1
  store i32* %472, i32** @e, align 8
  store i32 13, i32* %472, align 4
  call void @expr(i32 151)
  %473 = load i32** @e, align 8
  %474 = getelementptr inbounds i32* %473, i32 1
  store i32* %474, i32** @e, align 8
  store i32 18, i32* %474, align 4
  store i32 1, i32* @ty, align 4
  br label %721

; <label>:475                                     ; preds = %467
  %476 = load i32* @tk, align 4
  %477 = icmp eq i32 %476, 151
  br i1 %477, label %478, label %483

; <label>:478                                     ; preds = %475
  call void @next()
  %479 = load i32** @e, align 8
  %480 = getelementptr inbounds i32* %479, i32 1
  store i32* %480, i32** @e, align 8
  store i32 13, i32* %480, align 4
  call void @expr(i32 155)
  %481 = load i32** @e, align 8
  %482 = getelementptr inbounds i32* %481, i32 1
  store i32* %482, i32** @e, align 8
  store i32 19, i32* %482, align 4
  store i32 1, i32* @ty, align 4
  br label %720

; <label>:483                                     ; preds = %475
  %484 = load i32* @tk, align 4
  %485 = icmp eq i32 %484, 152
  br i1 %485, label %486, label %491

; <label>:486                                     ; preds = %483
  call void @next()
  %487 = load i32** @e, align 8
  %488 = getelementptr inbounds i32* %487, i32 1
  store i32* %488, i32** @e, align 8
  store i32 13, i32* %488, align 4
  call void @expr(i32 155)
  %489 = load i32** @e, align 8
  %490 = getelementptr inbounds i32* %489, i32 1
  store i32* %490, i32** @e, align 8
  store i32 20, i32* %490, align 4
  store i32 1, i32* @ty, align 4
  br label %719

; <label>:491                                     ; preds = %483
  %492 = load i32* @tk, align 4
  %493 = icmp eq i32 %492, 153
  br i1 %493, label %494, label %499

; <label>:494                                     ; preds = %491
  call void @next()
  %495 = load i32** @e, align 8
  %496 = getelementptr inbounds i32* %495, i32 1
  store i32* %496, i32** @e, align 8
  store i32 13, i32* %496, align 4
  call void @expr(i32 155)
  %497 = load i32** @e, align 8
  %498 = getelementptr inbounds i32* %497, i32 1
  store i32* %498, i32** @e, align 8
  store i32 21, i32* %498, align 4
  store i32 1, i32* @ty, align 4
  br label %718

; <label>:499                                     ; preds = %491
  %500 = load i32* @tk, align 4
  %501 = icmp eq i32 %500, 154
  br i1 %501, label %502, label %507

; <label>:502                                     ; preds = %499
  call void @next()
  %503 = load i32** @e, align 8
  %504 = getelementptr inbounds i32* %503, i32 1
  store i32* %504, i32** @e, align 8
  store i32 13, i32* %504, align 4
  call void @expr(i32 155)
  %505 = load i32** @e, align 8
  %506 = getelementptr inbounds i32* %505, i32 1
  store i32* %506, i32** @e, align 8
  store i32 22, i32* %506, align 4
  store i32 1, i32* @ty, align 4
  br label %717

; <label>:507                                     ; preds = %499
  %508 = load i32* @tk, align 4
  %509 = icmp eq i32 %508, 155
  br i1 %509, label %510, label %515

; <label>:510                                     ; preds = %507
  call void @next()
  %511 = load i32** @e, align 8
  %512 = getelementptr inbounds i32* %511, i32 1
  store i32* %512, i32** @e, align 8
  store i32 13, i32* %512, align 4
  call void @expr(i32 157)
  %513 = load i32** @e, align 8
  %514 = getelementptr inbounds i32* %513, i32 1
  store i32* %514, i32** @e, align 8
  store i32 23, i32* %514, align 4
  store i32 1, i32* @ty, align 4
  br label %716

; <label>:515                                     ; preds = %507
  %516 = load i32* @tk, align 4
  %517 = icmp eq i32 %516, 156
  br i1 %517, label %518, label %523

; <label>:518                                     ; preds = %515
  call void @next()
  %519 = load i32** @e, align 8
  %520 = getelementptr inbounds i32* %519, i32 1
  store i32* %520, i32** @e, align 8
  store i32 13, i32* %520, align 4
  call void @expr(i32 157)
  %521 = load i32** @e, align 8
  %522 = getelementptr inbounds i32* %521, i32 1
  store i32* %522, i32** @e, align 8
  store i32 24, i32* %522, align 4
  store i32 1, i32* @ty, align 4
  br label %715

; <label>:523                                     ; preds = %515
  %524 = load i32* @tk, align 4
  %525 = icmp eq i32 %524, 157
  br i1 %525, label %526, label %542

; <label>:526                                     ; preds = %523
  call void @next()
  %527 = load i32** @e, align 8
  %528 = getelementptr inbounds i32* %527, i32 1
  store i32* %528, i32** @e, align 8
  store i32 13, i32* %528, align 4
  call void @expr(i32 159)
  store i32 %367, i32* @ty, align 4
  %529 = icmp sgt i32 %367, 2
  br i1 %529, label %530, label %539

; <label>:530                                     ; preds = %526
  %531 = load i32** @e, align 8
  %532 = getelementptr inbounds i32* %531, i32 1
  store i32* %532, i32** @e, align 8
  store i32 13, i32* %532, align 4
  %533 = load i32** @e, align 8
  %534 = getelementptr inbounds i32* %533, i32 1
  store i32* %534, i32** @e, align 8
  store i32 1, i32* %534, align 4
  %535 = load i32** @e, align 8
  %536 = getelementptr inbounds i32* %535, i32 1
  store i32* %536, i32** @e, align 8
  store i32 4, i32* %536, align 4
  %537 = load i32** @e, align 8
  %538 = getelementptr inbounds i32* %537, i32 1
  store i32* %538, i32** @e, align 8
  store i32 27, i32* %538, align 4
  br label %539

; <label>:539                                     ; preds = %530, %526
  %540 = load i32** @e, align 8
  %541 = getelementptr inbounds i32* %540, i32 1
  store i32* %541, i32** @e, align 8
  store i32 25, i32* %541, align 4
  br label %714

; <label>:542                                     ; preds = %523
  %543 = load i32* @tk, align 4
  %544 = icmp eq i32 %543, 158
  br i1 %544, label %545, label %581

; <label>:545                                     ; preds = %542
  call void @next()
  %546 = load i32** @e, align 8
  %547 = getelementptr inbounds i32* %546, i32 1
  store i32* %547, i32** @e, align 8
  store i32 13, i32* %547, align 4
  call void @expr(i32 159)
  %548 = icmp sgt i32 %367, 2
  br i1 %548, label %549, label %563

; <label>:549                                     ; preds = %545
  %550 = load i32* @ty, align 4
  %551 = icmp eq i32 %367, %550
  br i1 %551, label %552, label %563

; <label>:552                                     ; preds = %549
  %553 = load i32** @e, align 8
  %554 = getelementptr inbounds i32* %553, i32 1
  store i32* %554, i32** @e, align 8
  store i32 26, i32* %554, align 4
  %555 = load i32** @e, align 8
  %556 = getelementptr inbounds i32* %555, i32 1
  store i32* %556, i32** @e, align 8
  store i32 13, i32* %556, align 4
  %557 = load i32** @e, align 8
  %558 = getelementptr inbounds i32* %557, i32 1
  store i32* %558, i32** @e, align 8
  store i32 1, i32* %558, align 4
  %559 = load i32** @e, align 8
  %560 = getelementptr inbounds i32* %559, i32 1
  store i32* %560, i32** @e, align 8
  store i32 4, i32* %560, align 4
  %561 = load i32** @e, align 8
  %562 = getelementptr inbounds i32* %561, i32 1
  store i32* %562, i32** @e, align 8
  store i32 28, i32* %562, align 4
  store i32 1, i32* @ty, align 4
  br label %580

; <label>:563                                     ; preds = %549, %545
  store i32 %367, i32* @ty, align 4
  %564 = icmp sgt i32 %367, 2
  br i1 %564, label %565, label %576

; <label>:565                                     ; preds = %563
  %566 = load i32** @e, align 8
  %567 = getelementptr inbounds i32* %566, i32 1
  store i32* %567, i32** @e, align 8
  store i32 13, i32* %567, align 4
  %568 = load i32** @e, align 8
  %569 = getelementptr inbounds i32* %568, i32 1
  store i32* %569, i32** @e, align 8
  store i32 1, i32* %569, align 4
  %570 = load i32** @e, align 8
  %571 = getelementptr inbounds i32* %570, i32 1
  store i32* %571, i32** @e, align 8
  store i32 4, i32* %571, align 4
  %572 = load i32** @e, align 8
  %573 = getelementptr inbounds i32* %572, i32 1
  store i32* %573, i32** @e, align 8
  store i32 27, i32* %573, align 4
  %574 = load i32** @e, align 8
  %575 = getelementptr inbounds i32* %574, i32 1
  store i32* %575, i32** @e, align 8
  store i32 26, i32* %575, align 4
  br label %579

; <label>:576                                     ; preds = %563
  %577 = load i32** @e, align 8
  %578 = getelementptr inbounds i32* %577, i32 1
  store i32* %578, i32** @e, align 8
  store i32 26, i32* %578, align 4
  br label %579

; <label>:579                                     ; preds = %576, %565
  br label %580

; <label>:580                                     ; preds = %579, %552
  br label %713

; <label>:581                                     ; preds = %542
  %582 = load i32* @tk, align 4
  %583 = icmp eq i32 %582, 159
  br i1 %583, label %584, label %589

; <label>:584                                     ; preds = %581
  call void @next()
  %585 = load i32** @e, align 8
  %586 = getelementptr inbounds i32* %585, i32 1
  store i32* %586, i32** @e, align 8
  store i32 13, i32* %586, align 4
  call void @expr(i32 162)
  %587 = load i32** @e, align 8
  %588 = getelementptr inbounds i32* %587, i32 1
  store i32* %588, i32** @e, align 8
  store i32 27, i32* %588, align 4
  store i32 1, i32* @ty, align 4
  br label %712

; <label>:589                                     ; preds = %581
  %590 = load i32* @tk, align 4
  %591 = icmp eq i32 %590, 160
  br i1 %591, label %592, label %597

; <label>:592                                     ; preds = %589
  call void @next()
  %593 = load i32** @e, align 8
  %594 = getelementptr inbounds i32* %593, i32 1
  store i32* %594, i32** @e, align 8
  store i32 13, i32* %594, align 4
  call void @expr(i32 162)
  %595 = load i32** @e, align 8
  %596 = getelementptr inbounds i32* %595, i32 1
  store i32* %596, i32** @e, align 8
  store i32 28, i32* %596, align 4
  store i32 1, i32* @ty, align 4
  br label %711

; <label>:597                                     ; preds = %589
  %598 = load i32* @tk, align 4
  %599 = icmp eq i32 %598, 161
  br i1 %599, label %600, label %605

; <label>:600                                     ; preds = %597
  call void @next()
  %601 = load i32** @e, align 8
  %602 = getelementptr inbounds i32* %601, i32 1
  store i32* %602, i32** @e, align 8
  store i32 13, i32* %602, align 4
  call void @expr(i32 162)
  %603 = load i32** @e, align 8
  %604 = getelementptr inbounds i32* %603, i32 1
  store i32* %604, i32** @e, align 8
  store i32 29, i32* %604, align 4
  store i32 1, i32* @ty, align 4
  br label %710

; <label>:605                                     ; preds = %597
  %606 = load i32* @tk, align 4
  %607 = icmp eq i32 %606, 162
  br i1 %607, label %611, label %608

; <label>:608                                     ; preds = %605
  %609 = load i32* @tk, align 4
  %610 = icmp eq i32 %609, 163
  br i1 %610, label %611, label %667

; <label>:611                                     ; preds = %608, %605
  %612 = load i32** @e, align 8
  %613 = load i32* %612, align 4
  %614 = icmp eq i32 %613, 10
  br i1 %614, label %615, label %619

; <label>:615                                     ; preds = %611
  %616 = load i32** @e, align 8
  store i32 13, i32* %616, align 4
  %617 = load i32** @e, align 8
  %618 = getelementptr inbounds i32* %617, i32 1
  store i32* %618, i32** @e, align 8
  store i32 10, i32* %618, align 4
  br label %631

; <label>:619                                     ; preds = %611
  %620 = load i32** @e, align 8
  %621 = load i32* %620, align 4
  %622 = icmp eq i32 %621, 9
  br i1 %622, label %623, label %627

; <label>:623                                     ; preds = %619
  %624 = load i32** @e, align 8
  store i32 13, i32* %624, align 4
  %625 = load i32** @e, align 8
  %626 = getelementptr inbounds i32* %625, i32 1
  store i32* %626, i32** @e, align 8
  store i32 9, i32* %626, align 4
  br label %630

; <label>:627                                     ; preds = %619
  %628 = load i32* @line, align 4
  %629 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([34 x i8]* @.str18, i32 0, i32 0), i32 %628)
  call void @exit(i32 -1) #7
  unreachable

; <label>:630                                     ; preds = %623
  br label %631

; <label>:631                                     ; preds = %630, %615
  %632 = load i32** @e, align 8
  %633 = getelementptr inbounds i32* %632, i32 1
  store i32* %633, i32** @e, align 8
  store i32 13, i32* %633, align 4
  %634 = load i32** @e, align 8
  %635 = getelementptr inbounds i32* %634, i32 1
  store i32* %635, i32** @e, align 8
  store i32 1, i32* %635, align 4
  %636 = load i32* @ty, align 4
  %637 = icmp sgt i32 %636, 2
  %638 = select i1 %637, i64 4, i64 1
  %639 = trunc i64 %638 to i32
  %640 = load i32** @e, align 8
  %641 = getelementptr inbounds i32* %640, i32 1
  store i32* %641, i32** @e, align 8
  store i32 %639, i32* %641, align 4
  %642 = load i32* @tk, align 4
  %643 = icmp eq i32 %642, 162
  %644 = select i1 %643, i32 25, i32 26
  %645 = load i32** @e, align 8
  %646 = getelementptr inbounds i32* %645, i32 1
  store i32* %646, i32** @e, align 8
  store i32 %644, i32* %646, align 4
  %647 = load i32* @ty, align 4
  %648 = icmp eq i32 %647, 0
  %649 = select i1 %648, i32 12, i32 11
  %650 = load i32** @e, align 8
  %651 = getelementptr inbounds i32* %650, i32 1
  store i32* %651, i32** @e, align 8
  store i32 %649, i32* %651, align 4
  %652 = load i32** @e, align 8
  %653 = getelementptr inbounds i32* %652, i32 1
  store i32* %653, i32** @e, align 8
  store i32 13, i32* %653, align 4
  %654 = load i32** @e, align 8
  %655 = getelementptr inbounds i32* %654, i32 1
  store i32* %655, i32** @e, align 8
  store i32 1, i32* %655, align 4
  %656 = load i32* @ty, align 4
  %657 = icmp sgt i32 %656, 2
  %658 = select i1 %657, i64 4, i64 1
  %659 = trunc i64 %658 to i32
  %660 = load i32** @e, align 8
  %661 = getelementptr inbounds i32* %660, i32 1
  store i32* %661, i32** @e, align 8
  store i32 %659, i32* %661, align 4
  %662 = load i32* @tk, align 4
  %663 = icmp eq i32 %662, 162
  %664 = select i1 %663, i32 26, i32 25
  %665 = load i32** @e, align 8
  %666 = getelementptr inbounds i32* %665, i32 1
  store i32* %666, i32** @e, align 8
  store i32 %664, i32* %666, align 4
  call void @next()
  br label %709

; <label>:667                                     ; preds = %608
  %668 = load i32* @tk, align 4
  %669 = icmp eq i32 %668, 164
  br i1 %669, label %670, label %704

; <label>:670                                     ; preds = %667
  call void @next()
  %671 = load i32** @e, align 8
  %672 = getelementptr inbounds i32* %671, i32 1
  store i32* %672, i32** @e, align 8
  store i32 13, i32* %672, align 4
  call void @expr(i32 142)
  %673 = load i32* @tk, align 4
  %674 = icmp eq i32 %673, 93
  br i1 %674, label %675, label %676

; <label>:675                                     ; preds = %670
  call void @next()
  br label %679

; <label>:676                                     ; preds = %670
  %677 = load i32* @line, align 4
  %678 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([28 x i8]* @.str19, i32 0, i32 0), i32 %677)
  call void @exit(i32 -1) #7
  unreachable

; <label>:679                                     ; preds = %675
  %680 = icmp sgt i32 %367, 2
  br i1 %680, label %681, label %690

; <label>:681                                     ; preds = %679
  %682 = load i32** @e, align 8
  %683 = getelementptr inbounds i32* %682, i32 1
  store i32* %683, i32** @e, align 8
  store i32 13, i32* %683, align 4
  %684 = load i32** @e, align 8
  %685 = getelementptr inbounds i32* %684, i32 1
  store i32* %685, i32** @e, align 8
  store i32 1, i32* %685, align 4
  %686 = load i32** @e, align 8
  %687 = getelementptr inbounds i32* %686, i32 1
  store i32* %687, i32** @e, align 8
  store i32 4, i32* %687, align 4
  %688 = load i32** @e, align 8
  %689 = getelementptr inbounds i32* %688, i32 1
  store i32* %689, i32** @e, align 8
  store i32 27, i32* %689, align 4
  br label %696

; <label>:690                                     ; preds = %679
  %691 = icmp slt i32 %367, 2
  br i1 %691, label %692, label %695

; <label>:692                                     ; preds = %690
  %693 = load i32* @line, align 4
  %694 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([27 x i8]* @.str20, i32 0, i32 0), i32 %693)
  call void @exit(i32 -1) #7
  unreachable

; <label>:695                                     ; preds = %690
  br label %696

; <label>:696                                     ; preds = %695, %681
  %697 = load i32** @e, align 8
  %698 = getelementptr inbounds i32* %697, i32 1
  store i32* %698, i32** @e, align 8
  store i32 25, i32* %698, align 4
  %699 = sub nsw i32 %367, 2
  store i32 %699, i32* @ty, align 4
  %700 = icmp eq i32 %699, 0
  %701 = select i1 %700, i32 10, i32 9
  %702 = load i32** @e, align 8
  %703 = getelementptr inbounds i32* %702, i32 1
  store i32* %703, i32** @e, align 8
  store i32 %701, i32* %703, align 4
  br label %708

; <label>:704                                     ; preds = %667
  %705 = load i32* @line, align 4
  %706 = load i32* @tk, align 4
  %707 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([26 x i8]* @.str21, i32 0, i32 0), i32 %705, i32 %706)
  call void @exit(i32 -1) #7
  unreachable

; <label>:708                                     ; preds = %696
  br label %709

; <label>:709                                     ; preds = %708, %631
  br label %710

; <label>:710                                     ; preds = %709, %600
  br label %711

; <label>:711                                     ; preds = %710, %592
  br label %712

; <label>:712                                     ; preds = %711, %584
  br label %713

; <label>:713                                     ; preds = %712, %580
  br label %714

; <label>:714                                     ; preds = %713, %539
  br label %715

; <label>:715                                     ; preds = %714, %518
  br label %716

; <label>:716                                     ; preds = %715, %510
  br label %717

; <label>:717                                     ; preds = %716, %502
  br label %718

; <label>:718                                     ; preds = %717, %494
  br label %719

; <label>:719                                     ; preds = %718, %486
  br label %720

; <label>:720                                     ; preds = %719, %478
  br label %721

; <label>:721                                     ; preds = %720, %470
  br label %722

; <label>:722                                     ; preds = %721, %462
  br label %723

; <label>:723                                     ; preds = %722, %454
  br label %724

; <label>:724                                     ; preds = %723, %446
  br label %725

; <label>:725                                     ; preds = %724, %438
  br label %726

; <label>:726                                     ; preds = %725, %427
  br label %727

; <label>:727                                     ; preds = %726, %416
  br label %728

; <label>:728                                     ; preds = %727, %402
  br label %729

; <label>:729                                     ; preds = %728, %383
  br label %363

; <label>:730                                     ; preds = %363
  ret void
}

; Function Attrs: noreturn nounwind
declare void @exit(i32) #3

; Function Attrs: nounwind uwtable
define void @stmt() #0 {
  %1 = load i32* @tk, align 4
  %2 = icmp eq i32 %1, 137
  br i1 %2, label %3, label %36

; <label>:3                                       ; preds = %0
  call void @next()
  %4 = load i32* @tk, align 4
  %5 = icmp eq i32 %4, 40
  br i1 %5, label %6, label %7

; <label>:6                                       ; preds = %3
  call void @next()
  br label %10

; <label>:7                                       ; preds = %3
  %8 = load i32* @line, align 4
  %9 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([25 x i8]* @.str22, i32 0, i32 0), i32 %8)
  call void @exit(i32 -1) #7
  unreachable

; <label>:10                                      ; preds = %6
  call void @expr(i32 142)
  %11 = load i32* @tk, align 4
  %12 = icmp eq i32 %11, 41
  br i1 %12, label %13, label %14

; <label>:13                                      ; preds = %10
  call void @next()
  br label %17

; <label>:14                                      ; preds = %10
  %15 = load i32* @line, align 4
  %16 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([26 x i8]* @.str11, i32 0, i32 0), i32 %15)
  call void @exit(i32 -1) #7
  unreachable

; <label>:17                                      ; preds = %13
  %18 = load i32** @e, align 8
  %19 = getelementptr inbounds i32* %18, i32 1
  store i32* %19, i32** @e, align 8
  store i32 4, i32* %19, align 4
  %20 = load i32** @e, align 8
  %21 = getelementptr inbounds i32* %20, i32 1
  store i32* %21, i32** @e, align 8
  call void @stmt()
  %22 = load i32* @tk, align 4
  %23 = icmp eq i32 %22, 135
  br i1 %23, label %24, label %32

; <label>:24                                      ; preds = %17
  %25 = load i32** @e, align 8
  %26 = getelementptr inbounds i32* %25, i64 3
  %27 = ptrtoint i32* %26 to i32
  store i32 %27, i32* %21, align 4
  %28 = load i32** @e, align 8
  %29 = getelementptr inbounds i32* %28, i32 1
  store i32* %29, i32** @e, align 8
  store i32 2, i32* %29, align 4
  %30 = load i32** @e, align 8
  %31 = getelementptr inbounds i32* %30, i32 1
  store i32* %31, i32** @e, align 8
  call void @next()
  call void @stmt()
  br label %32

; <label>:32                                      ; preds = %24, %17
  %b.0 = phi i32* [ %31, %24 ], [ %21, %17 ]
  %33 = load i32** @e, align 8
  %34 = getelementptr inbounds i32* %33, i64 1
  %35 = ptrtoint i32* %34 to i32
  store i32 %35, i32* %b.0, align 4
  br label %110

; <label>:36                                      ; preds = %0
  %37 = load i32* @tk, align 4
  %38 = icmp eq i32 %37, 141
  br i1 %38, label %39, label %68

; <label>:39                                      ; preds = %36
  call void @next()
  %40 = load i32** @e, align 8
  %41 = getelementptr inbounds i32* %40, i64 1
  %42 = load i32* @tk, align 4
  %43 = icmp eq i32 %42, 40
  br i1 %43, label %44, label %45

; <label>:44                                      ; preds = %39
  call void @next()
  br label %48

; <label>:45                                      ; preds = %39
  %46 = load i32* @line, align 4
  %47 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([25 x i8]* @.str22, i32 0, i32 0), i32 %46)
  call void @exit(i32 -1) #7
  unreachable

; <label>:48                                      ; preds = %44
  call void @expr(i32 142)
  %49 = load i32* @tk, align 4
  %50 = icmp eq i32 %49, 41
  br i1 %50, label %51, label %52

; <label>:51                                      ; preds = %48
  call void @next()
  br label %55

; <label>:52                                      ; preds = %48
  %53 = load i32* @line, align 4
  %54 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([26 x i8]* @.str11, i32 0, i32 0), i32 %53)
  call void @exit(i32 -1) #7
  unreachable

; <label>:55                                      ; preds = %51
  %56 = load i32** @e, align 8
  %57 = getelementptr inbounds i32* %56, i32 1
  store i32* %57, i32** @e, align 8
  store i32 4, i32* %57, align 4
  %58 = load i32** @e, align 8
  %59 = getelementptr inbounds i32* %58, i32 1
  store i32* %59, i32** @e, align 8
  call void @stmt()
  %60 = load i32** @e, align 8
  %61 = getelementptr inbounds i32* %60, i32 1
  store i32* %61, i32** @e, align 8
  store i32 2, i32* %61, align 4
  %62 = ptrtoint i32* %41 to i32
  %63 = load i32** @e, align 8
  %64 = getelementptr inbounds i32* %63, i32 1
  store i32* %64, i32** @e, align 8
  store i32 %62, i32* %64, align 4
  %65 = load i32** @e, align 8
  %66 = getelementptr inbounds i32* %65, i64 1
  %67 = ptrtoint i32* %66 to i32
  store i32 %67, i32* %59, align 4
  br label %109

; <label>:68                                      ; preds = %36
  %69 = load i32* @tk, align 4
  %70 = icmp eq i32 %69, 139
  br i1 %70, label %71, label %85

; <label>:71                                      ; preds = %68
  call void @next()
  %72 = load i32* @tk, align 4
  %73 = icmp ne i32 %72, 59
  br i1 %73, label %74, label %75

; <label>:74                                      ; preds = %71
  call void @expr(i32 142)
  br label %75

; <label>:75                                      ; preds = %74, %71
  %76 = load i32** @e, align 8
  %77 = getelementptr inbounds i32* %76, i32 1
  store i32* %77, i32** @e, align 8
  store i32 8, i32* %77, align 4
  %78 = load i32* @tk, align 4
  %79 = icmp eq i32 %78, 59
  br i1 %79, label %80, label %81

; <label>:80                                      ; preds = %75
  call void @next()
  br label %84

; <label>:81                                      ; preds = %75
  %82 = load i32* @line, align 4
  %83 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([24 x i8]* @.str23, i32 0, i32 0), i32 %82)
  call void @exit(i32 -1) #7
  unreachable

; <label>:84                                      ; preds = %80
  br label %108

; <label>:85                                      ; preds = %68
  %86 = load i32* @tk, align 4
  %87 = icmp eq i32 %86, 123
  br i1 %87, label %88, label %94

; <label>:88                                      ; preds = %85
  call void @next()
  br label %89

; <label>:89                                      ; preds = %92, %88
  %90 = load i32* @tk, align 4
  %91 = icmp ne i32 %90, 125
  br i1 %91, label %92, label %93

; <label>:92                                      ; preds = %89
  call void @stmt()
  br label %89

; <label>:93                                      ; preds = %89
  call void @next()
  br label %107

; <label>:94                                      ; preds = %85
  %95 = load i32* @tk, align 4
  %96 = icmp eq i32 %95, 59
  br i1 %96, label %97, label %98

; <label>:97                                      ; preds = %94
  call void @next()
  br label %106

; <label>:98                                      ; preds = %94
  call void @expr(i32 142)
  %99 = load i32* @tk, align 4
  %100 = icmp eq i32 %99, 59
  br i1 %100, label %101, label %102

; <label>:101                                     ; preds = %98
  call void @next()
  br label %105

; <label>:102                                     ; preds = %98
  %103 = load i32* @line, align 4
  %104 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([24 x i8]* @.str23, i32 0, i32 0), i32 %103)
  call void @exit(i32 -1) #7
  unreachable

; <label>:105                                     ; preds = %101
  br label %106

; <label>:106                                     ; preds = %105, %97
  br label %107

; <label>:107                                     ; preds = %106, %93
  br label %108

; <label>:108                                     ; preds = %107, %84
  br label %109

; <label>:109                                     ; preds = %108, %55
  br label %110

; <label>:110                                     ; preds = %109, %32
  ret void
}

; Function Attrs: nounwind uwtable
define i32 @main(i32 %argc, i8** %argv) #0 {
  %1 = add nsw i32 %argc, -1
  %2 = getelementptr inbounds i8** %argv, i32 1
  %3 = icmp sgt i32 %1, 0
  br i1 %3, label %4, label %18

; <label>:4                                       ; preds = %0
  %5 = load i8** %2, align 8
  %6 = load i8* %5, align 1
  %7 = sext i8 %6 to i32
  %8 = icmp eq i32 %7, 45
  br i1 %8, label %9, label %18

; <label>:9                                       ; preds = %4
  %10 = load i8** %2, align 8
  %11 = getelementptr inbounds i8* %10, i64 1
  %12 = load i8* %11, align 1
  %13 = sext i8 %12 to i32
  %14 = icmp eq i32 %13, 115
  br i1 %14, label %15, label %18

; <label>:15                                      ; preds = %9
  store i32 1, i32* @src, align 4
  %16 = add nsw i32 %1, -1
  %17 = getelementptr inbounds i8** %2, i32 1
  br label %18

; <label>:18                                      ; preds = %15, %9, %4, %0
  %.02 = phi i8** [ %17, %15 ], [ %2, %9 ], [ %2, %4 ], [ %2, %0 ]
  %.01 = phi i32 [ %16, %15 ], [ %1, %9 ], [ %1, %4 ], [ %1, %0 ]
  %19 = icmp sgt i32 %.01, 0
  br i1 %19, label %20, label %34

; <label>:20                                      ; preds = %18
  %21 = load i8** %.02, align 8
  %22 = load i8* %21, align 1
  %23 = sext i8 %22 to i32
  %24 = icmp eq i32 %23, 45
  br i1 %24, label %25, label %34

; <label>:25                                      ; preds = %20
  %26 = load i8** %.02, align 8
  %27 = getelementptr inbounds i8* %26, i64 1
  %28 = load i8* %27, align 1
  %29 = sext i8 %28 to i32
  %30 = icmp eq i32 %29, 100
  br i1 %30, label %31, label %34

; <label>:31                                      ; preds = %25
  store i32 1, i32* @debug, align 4
  %32 = add nsw i32 %.01, -1
  %33 = getelementptr inbounds i8** %.02, i32 1
  br label %34

; <label>:34                                      ; preds = %31, %25, %20, %18
  %.13 = phi i8** [ %33, %31 ], [ %.02, %25 ], [ %.02, %20 ], [ %.02, %18 ]
  %.1 = phi i32 [ %32, %31 ], [ %.01, %25 ], [ %.01, %20 ], [ %.01, %18 ]
  %35 = icmp slt i32 %.1, 1
  br i1 %35, label %36, label %38

; <label>:36                                      ; preds = %34
  %37 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([30 x i8]* @.str24, i32 0, i32 0))
  br label %797

; <label>:38                                      ; preds = %34
  %39 = load i8** %.13, align 8
  %40 = call i32 (i8*, i32, ...)* bitcast (i32 (...)* @open to i32 (i8*, i32, ...)*)(i8* %39, i32 0)
  %41 = icmp slt i32 %40, 0
  br i1 %41, label %42, label %45

; <label>:42                                      ; preds = %38
  %43 = load i8** %.13, align 8
  %44 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([20 x i8]* @.str25, i32 0, i32 0), i8* %43)
  br label %797

; <label>:45                                      ; preds = %38
  %46 = sext i32 262144 to i64
  %47 = call noalias i8* @malloc(i64 %46) #5
  %48 = bitcast i8* %47 to i32*
  store i32* %48, i32** @sym, align 8
  %49 = icmp ne i32* %48, null
  br i1 %49, label %52, label %50

; <label>:50                                      ; preds = %45
  %51 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([34 x i8]* @.str26, i32 0, i32 0), i32 262144)
  br label %797

; <label>:52                                      ; preds = %45
  %53 = sext i32 262144 to i64
  %54 = call noalias i8* @malloc(i64 %53) #5
  %55 = bitcast i8* %54 to i32*
  store i32* %55, i32** @e, align 8
  store i32* %55, i32** @le, align 8
  %56 = icmp ne i32* %55, null
  br i1 %56, label %59, label %57

; <label>:57                                      ; preds = %52
  %58 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([32 x i8]* @.str27, i32 0, i32 0), i32 262144)
  br label %797

; <label>:59                                      ; preds = %52
  %60 = sext i32 262144 to i64
  %61 = call noalias i8* @malloc(i64 %60) #5
  store i8* %61, i8** @data, align 8
  %62 = icmp ne i8* %61, null
  br i1 %62, label %65, label %63

; <label>:63                                      ; preds = %59
  %64 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([32 x i8]* @.str28, i32 0, i32 0), i32 262144)
  br label %797

; <label>:65                                      ; preds = %59
  %66 = sext i32 262144 to i64
  %67 = call noalias i8* @malloc(i64 %66) #5
  %68 = bitcast i8* %67 to i32*
  %69 = icmp ne i32* %68, null
  br i1 %69, label %72, label %70

; <label>:70                                      ; preds = %65
  %71 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([33 x i8]* @.str29, i32 0, i32 0), i32 262144)
  br label %797

; <label>:72                                      ; preds = %65
  %73 = load i32** @sym, align 8
  %74 = bitcast i32* %73 to i8*
  %75 = sext i32 262144 to i64
  call void @llvm.memset.p0i8.i64(i8* %74, i8 0, i64 %75, i32 4, i1 false)
  %76 = load i32** @e, align 8
  %77 = bitcast i32* %76 to i8*
  %78 = sext i32 262144 to i64
  call void @llvm.memset.p0i8.i64(i8* %77, i8 0, i64 %78, i32 4, i1 false)
  %79 = load i8** @data, align 8
  %80 = sext i32 262144 to i64
  call void @llvm.memset.p0i8.i64(i8* %79, i8 0, i64 %80, i32 1, i1 false)
  store i8* getelementptr inbounds ([101 x i8]* @.str30, i32 0, i32 0), i8** @p, align 8
  br label %81

; <label>:81                                      ; preds = %83, %72
  %i.0 = phi i32 [ 134, %72 ], [ %84, %83 ]
  %82 = icmp sle i32 %i.0, 141
  br i1 %82, label %83, label %87

; <label>:83                                      ; preds = %81
  call void @next()
  %84 = add nsw i32 %i.0, 1
  %85 = load i32** @id, align 8
  %86 = getelementptr inbounds i32* %85, i64 0
  store i32 %i.0, i32* %86, align 4
  br label %81

; <label>:87                                      ; preds = %81
  br label %88

; <label>:88                                      ; preds = %90, %87
  %i.1 = phi i32 [ 30, %87 ], [ %95, %90 ]
  %89 = icmp sle i32 %i.1, 37
  br i1 %89, label %90, label %98

; <label>:90                                      ; preds = %88
  call void @next()
  %91 = load i32** @id, align 8
  %92 = getelementptr inbounds i32* %91, i64 3
  store i32 130, i32* %92, align 4
  %93 = load i32** @id, align 8
  %94 = getelementptr inbounds i32* %93, i64 4
  store i32 1, i32* %94, align 4
  %95 = add nsw i32 %i.1, 1
  %96 = load i32** @id, align 8
  %97 = getelementptr inbounds i32* %96, i64 5
  store i32 %i.1, i32* %97, align 4
  br label %88

; <label>:98                                      ; preds = %88
  call void @next()
  %99 = load i32** @id, align 8
  %100 = getelementptr inbounds i32* %99, i64 0
  store i32 134, i32* %100, align 4
  call void @next()
  %101 = load i32** @id, align 8
  %102 = sext i32 262144 to i64
  %103 = call noalias i8* @malloc(i64 %102) #5
  store i8* %103, i8** @p, align 8
  store i8* %103, i8** @lp, align 8
  %104 = icmp ne i8* %103, null
  br i1 %104, label %107, label %105

; <label>:105                                     ; preds = %98
  %106 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([34 x i8]* @.str31, i32 0, i32 0), i32 262144)
  br label %797

; <label>:107                                     ; preds = %98
  %108 = load i8** @p, align 8
  %109 = sub nsw i32 262144, 1
  %110 = sext i32 %109 to i64
  %111 = call i64 @read(i32 %40, i8* %108, i64 %110)
  %112 = trunc i64 %111 to i32
  %113 = icmp sle i32 %112, 0
  br i1 %113, label %114, label %116

; <label>:114                                     ; preds = %107
  %115 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([20 x i8]* @.str32, i32 0, i32 0), i32 %112)
  br label %797

; <label>:116                                     ; preds = %107
  %117 = sext i32 %112 to i64
  %118 = load i8** @p, align 8
  %119 = getelementptr inbounds i8* %118, i64 %117
  store i8 0, i8* %119, align 1
  %120 = call i32 @close(i32 %40)
  store i32 1, i32* @line, align 4
  call void @next()
  br label %121

; <label>:121                                     ; preds = %411, %116
  %122 = load i32* @tk, align 4
  %123 = icmp ne i32 %122, 0
  br i1 %123, label %124, label %412

; <label>:124                                     ; preds = %121
  %125 = load i32* @tk, align 4
  %126 = icmp eq i32 %125, 138
  br i1 %126, label %127, label %128

; <label>:127                                     ; preds = %124
  call void @next()
  br label %180

; <label>:128                                     ; preds = %124
  %129 = load i32* @tk, align 4
  %130 = icmp eq i32 %129, 134
  br i1 %130, label %131, label %132

; <label>:131                                     ; preds = %128
  call void @next()
  br label %179

; <label>:132                                     ; preds = %128
  %133 = load i32* @tk, align 4
  %134 = icmp eq i32 %133, 136
  br i1 %134, label %135, label %178

; <label>:135                                     ; preds = %132
  call void @next()
  %136 = load i32* @tk, align 4
  %137 = icmp ne i32 %136, 123
  br i1 %137, label %138, label %139

; <label>:138                                     ; preds = %135
  call void @next()
  br label %139

; <label>:139                                     ; preds = %138, %135
  %140 = load i32* @tk, align 4
  %141 = icmp eq i32 %140, 123
  br i1 %141, label %142, label %177

; <label>:142                                     ; preds = %139
  call void @next()
  br label %143

; <label>:143                                     ; preds = %175, %142
  %i.2 = phi i32 [ 0, %142 ], [ %169, %175 ]
  %144 = load i32* @tk, align 4
  %145 = icmp ne i32 %144, 125
  br i1 %145, label %146, label %176

; <label>:146                                     ; preds = %143
  %147 = load i32* @tk, align 4
  %148 = icmp ne i32 %147, 133
  br i1 %148, label %149, label %153

; <label>:149                                     ; preds = %146
  %150 = load i32* @line, align 4
  %151 = load i32* @tk, align 4
  %152 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([28 x i8]* @.str33, i32 0, i32 0), i32 %150, i32 %151)
  br label %797

; <label>:153                                     ; preds = %146
  call void @next()
  %154 = load i32* @tk, align 4
  %155 = icmp eq i32 %154, 142
  br i1 %155, label %156, label %164

; <label>:156                                     ; preds = %153
  call void @next()
  %157 = load i32* @tk, align 4
  %158 = icmp ne i32 %157, 128
  br i1 %158, label %159, label %162

; <label>:159                                     ; preds = %156
  %160 = load i32* @line, align 4
  %161 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([26 x i8]* @.str34, i32 0, i32 0), i32 %160)
  br label %797

; <label>:162                                     ; preds = %156
  %163 = load i32* @ival, align 4
  call void @next()
  br label %164

; <label>:164                                     ; preds = %162, %153
  %i.3 = phi i32 [ %163, %162 ], [ %i.2, %153 ]
  %165 = load i32** @id, align 8
  %166 = getelementptr inbounds i32* %165, i64 3
  store i32 128, i32* %166, align 4
  %167 = load i32** @id, align 8
  %168 = getelementptr inbounds i32* %167, i64 4
  store i32 1, i32* %168, align 4
  %169 = add nsw i32 %i.3, 1
  %170 = load i32** @id, align 8
  %171 = getelementptr inbounds i32* %170, i64 5
  store i32 %i.3, i32* %171, align 4
  %172 = load i32* @tk, align 4
  %173 = icmp eq i32 %172, 44
  br i1 %173, label %174, label %175

; <label>:174                                     ; preds = %164
  call void @next()
  br label %175

; <label>:175                                     ; preds = %174, %164
  br label %143

; <label>:176                                     ; preds = %143
  call void @next()
  br label %177

; <label>:177                                     ; preds = %176, %139
  br label %178

; <label>:178                                     ; preds = %177, %132
  br label %179

; <label>:179                                     ; preds = %178, %131
  %bt.0 = phi i32 [ 0, %131 ], [ 1, %178 ]
  br label %180

; <label>:180                                     ; preds = %179, %127
  %bt.1 = phi i32 [ 1, %127 ], [ %bt.0, %179 ]
  br label %181

; <label>:181                                     ; preds = %410, %180
  %bt.2 = phi i32 [ %bt.1, %180 ], [ %bt.4, %410 ]
  %182 = load i32* @tk, align 4
  %183 = icmp ne i32 %182, 59
  br i1 %183, label %184, label %187

; <label>:184                                     ; preds = %181
  %185 = load i32* @tk, align 4
  %186 = icmp ne i32 %185, 125
  br label %187

; <label>:187                                     ; preds = %184, %181
  %188 = phi i1 [ false, %181 ], [ %186, %184 ]
  br i1 %188, label %189, label %411

; <label>:189                                     ; preds = %187
  br label %190

; <label>:190                                     ; preds = %193, %189
  %ty.0 = phi i32 [ %bt.2, %189 ], [ %194, %193 ]
  %191 = load i32* @tk, align 4
  %192 = icmp eq i32 %191, 159
  br i1 %192, label %193, label %195

; <label>:193                                     ; preds = %190
  call void @next()
  %194 = add nsw i32 %ty.0, 2
  br label %190

; <label>:195                                     ; preds = %190
  %196 = load i32* @tk, align 4
  %197 = icmp ne i32 %196, 133
  br i1 %197, label %198, label %201

; <label>:198                                     ; preds = %195
  %199 = load i32* @line, align 4
  %200 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([28 x i8]* @.str35, i32 0, i32 0), i32 %199)
  br label %797

; <label>:201                                     ; preds = %195
  %202 = load i32** @id, align 8
  %203 = getelementptr inbounds i32* %202, i64 3
  %204 = load i32* %203, align 4
  %205 = icmp ne i32 %204, 0
  br i1 %205, label %206, label %209

; <label>:206                                     ; preds = %201
  %207 = load i32* @line, align 4
  %208 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([33 x i8]* @.str36, i32 0, i32 0), i32 %207)
  br label %797

; <label>:209                                     ; preds = %201
  call void @next()
  %210 = load i32** @id, align 8
  %211 = getelementptr inbounds i32* %210, i64 4
  store i32 %ty.0, i32* %211, align 4
  %212 = load i32* @tk, align 4
  %213 = icmp eq i32 %212, 40
  br i1 %213, label %214, label %397

; <label>:214                                     ; preds = %209
  %215 = load i32** @id, align 8
  %216 = getelementptr inbounds i32* %215, i64 3
  store i32 129, i32* %216, align 4
  %217 = load i32** @e, align 8
  %218 = getelementptr inbounds i32* %217, i64 1
  %219 = ptrtoint i32* %218 to i32
  %220 = load i32** @id, align 8
  %221 = getelementptr inbounds i32* %220, i64 5
  store i32 %219, i32* %221, align 4
  call void @next()
  br label %222

; <label>:222                                     ; preds = %280, %214
  %i.4 = phi i32 [ 0, %214 ], [ %274, %280 ]
  %223 = load i32* @tk, align 4
  %224 = icmp ne i32 %223, 41
  br i1 %224, label %225, label %281

; <label>:225                                     ; preds = %222
  %226 = load i32* @tk, align 4
  %227 = icmp eq i32 %226, 138
  br i1 %227, label %228, label %229

; <label>:228                                     ; preds = %225
  call void @next()
  br label %234

; <label>:229                                     ; preds = %225
  %230 = load i32* @tk, align 4
  %231 = icmp eq i32 %230, 134
  br i1 %231, label %232, label %233

; <label>:232                                     ; preds = %229
  call void @next()
  br label %233

; <label>:233                                     ; preds = %232, %229
  %ty.1 = phi i32 [ 0, %232 ], [ 1, %229 ]
  br label %234

; <label>:234                                     ; preds = %233, %228
  %ty.2 = phi i32 [ 1, %228 ], [ %ty.1, %233 ]
  br label %235

; <label>:235                                     ; preds = %238, %234
  %ty.3 = phi i32 [ %ty.2, %234 ], [ %239, %238 ]
  %236 = load i32* @tk, align 4
  %237 = icmp eq i32 %236, 159
  br i1 %237, label %238, label %240

; <label>:238                                     ; preds = %235
  call void @next()
  %239 = add nsw i32 %ty.3, 2
  br label %235

; <label>:240                                     ; preds = %235
  %241 = load i32* @tk, align 4
  %242 = icmp ne i32 %241, 133
  br i1 %242, label %243, label %246

; <label>:243                                     ; preds = %240
  %244 = load i32* @line, align 4
  %245 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([31 x i8]* @.str37, i32 0, i32 0), i32 %244)
  br label %797

; <label>:246                                     ; preds = %240
  %247 = load i32** @id, align 8
  %248 = getelementptr inbounds i32* %247, i64 3
  %249 = load i32* %248, align 4
  %250 = icmp eq i32 %249, 132
  br i1 %250, label %251, label %254

; <label>:251                                     ; preds = %246
  %252 = load i32* @line, align 4
  %253 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([36 x i8]* @.str38, i32 0, i32 0), i32 %252)
  br label %797

; <label>:254                                     ; preds = %246
  %255 = load i32** @id, align 8
  %256 = getelementptr inbounds i32* %255, i64 3
  %257 = load i32* %256, align 4
  %258 = load i32** @id, align 8
  %259 = getelementptr inbounds i32* %258, i64 6
  store i32 %257, i32* %259, align 4
  %260 = load i32** @id, align 8
  %261 = getelementptr inbounds i32* %260, i64 3
  store i32 132, i32* %261, align 4
  %262 = load i32** @id, align 8
  %263 = getelementptr inbounds i32* %262, i64 4
  %264 = load i32* %263, align 4
  %265 = load i32** @id, align 8
  %266 = getelementptr inbounds i32* %265, i64 7
  store i32 %264, i32* %266, align 4
  %267 = load i32** @id, align 8
  %268 = getelementptr inbounds i32* %267, i64 4
  store i32 %ty.3, i32* %268, align 4
  %269 = load i32** @id, align 8
  %270 = getelementptr inbounds i32* %269, i64 5
  %271 = load i32* %270, align 4
  %272 = load i32** @id, align 8
  %273 = getelementptr inbounds i32* %272, i64 8
  store i32 %271, i32* %273, align 4
  %274 = add nsw i32 %i.4, 1
  %275 = load i32** @id, align 8
  %276 = getelementptr inbounds i32* %275, i64 5
  store i32 %i.4, i32* %276, align 4
  call void @next()
  %277 = load i32* @tk, align 4
  %278 = icmp eq i32 %277, 44
  br i1 %278, label %279, label %280

; <label>:279                                     ; preds = %254
  call void @next()
  br label %280

; <label>:280                                     ; preds = %279, %254
  br label %222

; <label>:281                                     ; preds = %222
  call void @next()
  %282 = load i32* @tk, align 4
  %283 = icmp ne i32 %282, 123
  br i1 %283, label %284, label %287

; <label>:284                                     ; preds = %281
  %285 = load i32* @line, align 4
  %286 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([29 x i8]* @.str39, i32 0, i32 0), i32 %285)
  br label %797

; <label>:287                                     ; preds = %281
  %288 = add nsw i32 %i.4, 1
  store i32 %288, i32* @loc, align 4
  call void @next()
  br label %289

; <label>:289                                     ; preds = %351, %287
  %i.5 = phi i32 [ %288, %287 ], [ %i.6, %351 ]
  %bt.3 = phi i32 [ %bt.2, %287 ], [ %300, %351 ]
  %290 = load i32* @tk, align 4
  %291 = icmp eq i32 %290, 138
  br i1 %291, label %295, label %292

; <label>:292                                     ; preds = %289
  %293 = load i32* @tk, align 4
  %294 = icmp eq i32 %293, 134
  br label %295

; <label>:295                                     ; preds = %292, %289
  %296 = phi i1 [ true, %289 ], [ %294, %292 ]
  br i1 %296, label %297, label %352

; <label>:297                                     ; preds = %295
  %298 = load i32* @tk, align 4
  %299 = icmp eq i32 %298, 138
  %300 = select i1 %299, i32 1, i32 0
  call void @next()
  br label %301

; <label>:301                                     ; preds = %350, %297
  %i.6 = phi i32 [ %i.5, %297 ], [ %344, %350 ]
  %302 = load i32* @tk, align 4
  %303 = icmp ne i32 %302, 59
  br i1 %303, label %304, label %351

; <label>:304                                     ; preds = %301
  br label %305

; <label>:305                                     ; preds = %308, %304
  %ty.4 = phi i32 [ %300, %304 ], [ %309, %308 ]
  %306 = load i32* @tk, align 4
  %307 = icmp eq i32 %306, 159
  br i1 %307, label %308, label %310

; <label>:308                                     ; preds = %305
  call void @next()
  %309 = add nsw i32 %ty.4, 2
  br label %305

; <label>:310                                     ; preds = %305
  %311 = load i32* @tk, align 4
  %312 = icmp ne i32 %311, 133
  br i1 %312, label %313, label %316

; <label>:313                                     ; preds = %310
  %314 = load i32* @line, align 4
  %315 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([27 x i8]* @.str40, i32 0, i32 0), i32 %314)
  br label %797

; <label>:316                                     ; preds = %310
  %317 = load i32** @id, align 8
  %318 = getelementptr inbounds i32* %317, i64 3
  %319 = load i32* %318, align 4
  %320 = icmp eq i32 %319, 132
  br i1 %320, label %321, label %324

; <label>:321                                     ; preds = %316
  %322 = load i32* @line, align 4
  %323 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([32 x i8]* @.str41, i32 0, i32 0), i32 %322)
  br label %797

; <label>:324                                     ; preds = %316
  %325 = load i32** @id, align 8
  %326 = getelementptr inbounds i32* %325, i64 3
  %327 = load i32* %326, align 4
  %328 = load i32** @id, align 8
  %329 = getelementptr inbounds i32* %328, i64 6
  store i32 %327, i32* %329, align 4
  %330 = load i32** @id, align 8
  %331 = getelementptr inbounds i32* %330, i64 3
  store i32 132, i32* %331, align 4
  %332 = load i32** @id, align 8
  %333 = getelementptr inbounds i32* %332, i64 4
  %334 = load i32* %333, align 4
  %335 = load i32** @id, align 8
  %336 = getelementptr inbounds i32* %335, i64 7
  store i32 %334, i32* %336, align 4
  %337 = load i32** @id, align 8
  %338 = getelementptr inbounds i32* %337, i64 4
  store i32 %ty.4, i32* %338, align 4
  %339 = load i32** @id, align 8
  %340 = getelementptr inbounds i32* %339, i64 5
  %341 = load i32* %340, align 4
  %342 = load i32** @id, align 8
  %343 = getelementptr inbounds i32* %342, i64 8
  store i32 %341, i32* %343, align 4
  %344 = add nsw i32 %i.6, 1
  %345 = load i32** @id, align 8
  %346 = getelementptr inbounds i32* %345, i64 5
  store i32 %344, i32* %346, align 4
  call void @next()
  %347 = load i32* @tk, align 4
  %348 = icmp eq i32 %347, 44
  br i1 %348, label %349, label %350

; <label>:349                                     ; preds = %324
  call void @next()
  br label %350

; <label>:350                                     ; preds = %349, %324
  br label %301

; <label>:351                                     ; preds = %301
  call void @next()
  br label %289

; <label>:352                                     ; preds = %295
  %353 = load i32** @e, align 8
  %354 = getelementptr inbounds i32* %353, i32 1
  store i32* %354, i32** @e, align 8
  store i32 6, i32* %354, align 4
  %355 = load i32* @loc, align 4
  %356 = sub nsw i32 %i.5, %355
  %357 = load i32** @e, align 8
  %358 = getelementptr inbounds i32* %357, i32 1
  store i32* %358, i32** @e, align 8
  store i32 %356, i32* %358, align 4
  br label %359

; <label>:359                                     ; preds = %362, %352
  %360 = load i32* @tk, align 4
  %361 = icmp ne i32 %360, 125
  br i1 %361, label %362, label %363

; <label>:362                                     ; preds = %359
  call void @stmt()
  br label %359

; <label>:363                                     ; preds = %359
  %364 = load i32** @e, align 8
  %365 = getelementptr inbounds i32* %364, i32 1
  store i32* %365, i32** @e, align 8
  store i32 8, i32* %365, align 4
  %366 = load i32** @sym, align 8
  store i32* %366, i32** @id, align 8
  br label %367

; <label>:367                                     ; preds = %393, %363
  %368 = load i32** @id, align 8
  %369 = getelementptr inbounds i32* %368, i64 0
  %370 = load i32* %369, align 4
  %371 = icmp ne i32 %370, 0
  br i1 %371, label %372, label %396

; <label>:372                                     ; preds = %367
  %373 = load i32** @id, align 8
  %374 = getelementptr inbounds i32* %373, i64 3
  %375 = load i32* %374, align 4
  %376 = icmp eq i32 %375, 132
  br i1 %376, label %377, label %393

; <label>:377                                     ; preds = %372
  %378 = load i32** @id, align 8
  %379 = getelementptr inbounds i32* %378, i64 6
  %380 = load i32* %379, align 4
  %381 = load i32** @id, align 8
  %382 = getelementptr inbounds i32* %381, i64 3
  store i32 %380, i32* %382, align 4
  %383 = load i32** @id, align 8
  %384 = getelementptr inbounds i32* %383, i64 7
  %385 = load i32* %384, align 4
  %386 = load i32** @id, align 8
  %387 = getelementptr inbounds i32* %386, i64 4
  store i32 %385, i32* %387, align 4
  %388 = load i32** @id, align 8
  %389 = getelementptr inbounds i32* %388, i64 8
  %390 = load i32* %389, align 4
  %391 = load i32** @id, align 8
  %392 = getelementptr inbounds i32* %391, i64 5
  store i32 %390, i32* %392, align 4
  br label %393

; <label>:393                                     ; preds = %377, %372
  %394 = load i32** @id, align 8
  %395 = getelementptr inbounds i32* %394, i64 9
  store i32* %395, i32** @id, align 8
  br label %367

; <label>:396                                     ; preds = %367
  br label %406

; <label>:397                                     ; preds = %209
  %398 = load i32** @id, align 8
  %399 = getelementptr inbounds i32* %398, i64 3
  store i32 131, i32* %399, align 4
  %400 = load i8** @data, align 8
  %401 = ptrtoint i8* %400 to i32
  %402 = load i32** @id, align 8
  %403 = getelementptr inbounds i32* %402, i64 5
  store i32 %401, i32* %403, align 4
  %404 = load i8** @data, align 8
  %405 = getelementptr inbounds i8* %404, i64 4
  store i8* %405, i8** @data, align 8
  br label %406

; <label>:406                                     ; preds = %397, %396
  %bt.4 = phi i32 [ %bt.3, %396 ], [ %bt.2, %397 ]
  %407 = load i32* @tk, align 4
  %408 = icmp eq i32 %407, 44
  br i1 %408, label %409, label %410

; <label>:409                                     ; preds = %406
  call void @next()
  br label %410

; <label>:410                                     ; preds = %409, %406
  br label %181

; <label>:411                                     ; preds = %187
  call void @next()
  br label %121

; <label>:412                                     ; preds = %121
  %413 = getelementptr inbounds i32* %101, i64 5
  %414 = load i32* %413, align 4
  %415 = sext i32 %414 to i64
  %416 = inttoptr i64 %415 to i32*
  %417 = icmp ne i32* %416, null
  br i1 %417, label %420, label %418

; <label>:418                                     ; preds = %412
  %419 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([20 x i8]* @.str42, i32 0, i32 0))
  br label %797

; <label>:420                                     ; preds = %412
  %421 = load i32* @src, align 4
  %422 = icmp ne i32 %421, 0
  br i1 %422, label %423, label %424

; <label>:423                                     ; preds = %420
  br label %797

; <label>:424                                     ; preds = %420
  %425 = ptrtoint i32* %68 to i32
  %426 = add nsw i32 %425, 262144
  %427 = sext i32 %426 to i64
  %428 = inttoptr i64 %427 to i32*
  %429 = getelementptr inbounds i32* %428, i32 -1
  store i32 37, i32* %429, align 4
  %430 = getelementptr inbounds i32* %429, i32 -1
  store i32 13, i32* %430, align 4
  %431 = getelementptr inbounds i32* %430, i32 -1
  store i32 %.1, i32* %431, align 4
  %432 = ptrtoint i8** %.13 to i32
  %433 = getelementptr inbounds i32* %431, i32 -1
  store i32 %432, i32* %433, align 4
  %434 = ptrtoint i32* %430 to i32
  %435 = getelementptr inbounds i32* %433, i32 -1
  store i32 %434, i32* %435, align 4
  br label %436

; <label>:436                                     ; preds = %796, %424
  %a.0 = phi i32 [ undef, %424 ], [ %a.36, %796 ]
  %bp.0 = phi i32* [ undef, %424 ], [ %bp.9, %796 ]
  %sp.0 = phi i32* [ %435, %424 ], [ %sp.30, %796 ]
  %pc.0 = phi i32* [ %416, %424 ], [ %pc.9, %796 ]
  %cycle.0 = phi i32 [ 0, %424 ], [ %439, %796 ]
  %437 = getelementptr inbounds i32* %pc.0, i32 1
  %438 = load i32* %pc.0, align 4
  %439 = add nsw i32 %cycle.0, 1
  %440 = load i32* @debug, align 4
  %441 = icmp ne i32 %440, 0
  br i1 %441, label %442, label %454

; <label>:442                                     ; preds = %436
  %443 = mul nsw i32 %438, 5
  %444 = sext i32 %443 to i64
  %445 = getelementptr inbounds [191 x i8]* @.str2, i32 0, i64 %444
  %446 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([9 x i8]* @.str43, i32 0, i32 0), i32 %439, i8* %445)
  %447 = icmp sle i32 %438, 7
  br i1 %447, label %448, label %451

; <label>:448                                     ; preds = %442
  %449 = load i32* %437, align 4
  %450 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([5 x i8]* @.str3, i32 0, i32 0), i32 %449)
  br label %453

; <label>:451                                     ; preds = %442
  %452 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([2 x i8]* @.str4, i32 0, i32 0))
  br label %453

; <label>:453                                     ; preds = %451, %448
  br label %454

; <label>:454                                     ; preds = %453, %436
  %455 = icmp eq i32 %438, 0
  br i1 %455, label %456, label %462

; <label>:456                                     ; preds = %454
  %457 = getelementptr inbounds i32* %437, i32 1
  %458 = load i32* %437, align 4
  %459 = sext i32 %458 to i64
  %460 = getelementptr inbounds i32* %bp.0, i64 %459
  %461 = ptrtoint i32* %460 to i32
  br label %796

; <label>:462                                     ; preds = %454
  %463 = icmp eq i32 %438, 1
  br i1 %463, label %464, label %467

; <label>:464                                     ; preds = %462
  %465 = getelementptr inbounds i32* %437, i32 1
  %466 = load i32* %437, align 4
  br label %795

; <label>:467                                     ; preds = %462
  %468 = icmp eq i32 %438, 2
  br i1 %468, label %469, label %473

; <label>:469                                     ; preds = %467
  %470 = load i32* %437, align 4
  %471 = sext i32 %470 to i64
  %472 = inttoptr i64 %471 to i32*
  br label %794

; <label>:473                                     ; preds = %467
  %474 = icmp eq i32 %438, 3
  br i1 %474, label %475, label %482

; <label>:475                                     ; preds = %473
  %476 = getelementptr inbounds i32* %437, i64 1
  %477 = ptrtoint i32* %476 to i32
  %478 = getelementptr inbounds i32* %sp.0, i32 -1
  store i32 %477, i32* %478, align 4
  %479 = load i32* %437, align 4
  %480 = sext i32 %479 to i64
  %481 = inttoptr i64 %480 to i32*
  br label %793

; <label>:482                                     ; preds = %473
  %483 = icmp eq i32 %438, 4
  br i1 %483, label %484, label %494

; <label>:484                                     ; preds = %482
  %485 = icmp ne i32 %a.0, 0
  br i1 %485, label %486, label %488

; <label>:486                                     ; preds = %484
  %487 = getelementptr inbounds i32* %437, i64 1
  br label %492

; <label>:488                                     ; preds = %484
  %489 = load i32* %437, align 4
  %490 = sext i32 %489 to i64
  %491 = inttoptr i64 %490 to i32*
  br label %492

; <label>:492                                     ; preds = %488, %486
  %493 = phi i32* [ %487, %486 ], [ %491, %488 ]
  br label %792

; <label>:494                                     ; preds = %482
  %495 = icmp eq i32 %438, 5
  br i1 %495, label %496, label %506

; <label>:496                                     ; preds = %494
  %497 = icmp ne i32 %a.0, 0
  br i1 %497, label %498, label %502

; <label>:498                                     ; preds = %496
  %499 = load i32* %437, align 4
  %500 = sext i32 %499 to i64
  %501 = inttoptr i64 %500 to i32*
  br label %504

; <label>:502                                     ; preds = %496
  %503 = getelementptr inbounds i32* %437, i64 1
  br label %504

; <label>:504                                     ; preds = %502, %498
  %505 = phi i32* [ %501, %498 ], [ %503, %502 ]
  br label %791

; <label>:506                                     ; preds = %494
  %507 = icmp eq i32 %438, 6
  br i1 %507, label %508, label %516

; <label>:508                                     ; preds = %506
  %509 = ptrtoint i32* %bp.0 to i32
  %510 = getelementptr inbounds i32* %sp.0, i32 -1
  store i32 %509, i32* %510, align 4
  %511 = getelementptr inbounds i32* %437, i32 1
  %512 = load i32* %437, align 4
  %513 = sext i32 %512 to i64
  %514 = sub i64 0, %513
  %515 = getelementptr inbounds i32* %510, i64 %514
  br label %790

; <label>:516                                     ; preds = %506
  %517 = icmp eq i32 %438, 7
  br i1 %517, label %518, label %523

; <label>:518                                     ; preds = %516
  %519 = getelementptr inbounds i32* %437, i32 1
  %520 = load i32* %437, align 4
  %521 = sext i32 %520 to i64
  %522 = getelementptr inbounds i32* %sp.0, i64 %521
  br label %789

; <label>:523                                     ; preds = %516
  %524 = icmp eq i32 %438, 8
  br i1 %524, label %525, label %534

; <label>:525                                     ; preds = %523
  %526 = getelementptr inbounds i32* %bp.0, i32 1
  %527 = load i32* %bp.0, align 4
  %528 = sext i32 %527 to i64
  %529 = inttoptr i64 %528 to i32*
  %530 = getelementptr inbounds i32* %526, i32 1
  %531 = load i32* %526, align 4
  %532 = sext i32 %531 to i64
  %533 = inttoptr i64 %532 to i32*
  br label %788

; <label>:534                                     ; preds = %523
  %535 = icmp eq i32 %438, 9
  br i1 %535, label %536, label %540

; <label>:536                                     ; preds = %534
  %537 = sext i32 %a.0 to i64
  %538 = inttoptr i64 %537 to i32*
  %539 = load i32* %538, align 4
  br label %787

; <label>:540                                     ; preds = %534
  %541 = icmp eq i32 %438, 10
  br i1 %541, label %542, label %547

; <label>:542                                     ; preds = %540
  %543 = sext i32 %a.0 to i64
  %544 = inttoptr i64 %543 to i8*
  %545 = load i8* %544, align 1
  %546 = sext i8 %545 to i32
  br label %786

; <label>:547                                     ; preds = %540
  %548 = icmp eq i32 %438, 11
  br i1 %548, label %549, label %554

; <label>:549                                     ; preds = %547
  %550 = getelementptr inbounds i32* %sp.0, i32 1
  %551 = load i32* %sp.0, align 4
  %552 = sext i32 %551 to i64
  %553 = inttoptr i64 %552 to i32*
  store i32 %a.0, i32* %553, align 4
  br label %785

; <label>:554                                     ; preds = %547
  %555 = icmp eq i32 %438, 12
  br i1 %555, label %556, label %563

; <label>:556                                     ; preds = %554
  %557 = trunc i32 %a.0 to i8
  %558 = getelementptr inbounds i32* %sp.0, i32 1
  %559 = load i32* %sp.0, align 4
  %560 = sext i32 %559 to i64
  %561 = inttoptr i64 %560 to i8*
  store i8 %557, i8* %561, align 1
  %562 = sext i8 %557 to i32
  br label %784

; <label>:563                                     ; preds = %554
  %564 = icmp eq i32 %438, 13
  br i1 %564, label %565, label %567

; <label>:565                                     ; preds = %563
  %566 = getelementptr inbounds i32* %sp.0, i32 -1
  store i32 %a.0, i32* %566, align 4
  br label %783

; <label>:567                                     ; preds = %563
  %568 = icmp eq i32 %438, 14
  br i1 %568, label %569, label %573

; <label>:569                                     ; preds = %567
  %570 = getelementptr inbounds i32* %sp.0, i32 1
  %571 = load i32* %sp.0, align 4
  %572 = or i32 %571, %a.0
  br label %782

; <label>:573                                     ; preds = %567
  %574 = icmp eq i32 %438, 15
  br i1 %574, label %575, label %579

; <label>:575                                     ; preds = %573
  %576 = getelementptr inbounds i32* %sp.0, i32 1
  %577 = load i32* %sp.0, align 4
  %578 = xor i32 %577, %a.0
  br label %781

; <label>:579                                     ; preds = %573
  %580 = icmp eq i32 %438, 16
  br i1 %580, label %581, label %585

; <label>:581                                     ; preds = %579
  %582 = getelementptr inbounds i32* %sp.0, i32 1
  %583 = load i32* %sp.0, align 4
  %584 = and i32 %583, %a.0
  br label %780

; <label>:585                                     ; preds = %579
  %586 = icmp eq i32 %438, 17
  br i1 %586, label %587, label %592

; <label>:587                                     ; preds = %585
  %588 = getelementptr inbounds i32* %sp.0, i32 1
  %589 = load i32* %sp.0, align 4
  %590 = icmp eq i32 %589, %a.0
  %591 = zext i1 %590 to i32
  br label %779

; <label>:592                                     ; preds = %585
  %593 = icmp eq i32 %438, 18
  br i1 %593, label %594, label %599

; <label>:594                                     ; preds = %592
  %595 = getelementptr inbounds i32* %sp.0, i32 1
  %596 = load i32* %sp.0, align 4
  %597 = icmp ne i32 %596, %a.0
  %598 = zext i1 %597 to i32
  br label %778

; <label>:599                                     ; preds = %592
  %600 = icmp eq i32 %438, 19
  br i1 %600, label %601, label %606

; <label>:601                                     ; preds = %599
  %602 = getelementptr inbounds i32* %sp.0, i32 1
  %603 = load i32* %sp.0, align 4
  %604 = icmp slt i32 %603, %a.0
  %605 = zext i1 %604 to i32
  br label %777

; <label>:606                                     ; preds = %599
  %607 = icmp eq i32 %438, 20
  br i1 %607, label %608, label %613

; <label>:608                                     ; preds = %606
  %609 = getelementptr inbounds i32* %sp.0, i32 1
  %610 = load i32* %sp.0, align 4
  %611 = icmp sgt i32 %610, %a.0
  %612 = zext i1 %611 to i32
  br label %776

; <label>:613                                     ; preds = %606
  %614 = icmp eq i32 %438, 21
  br i1 %614, label %615, label %620

; <label>:615                                     ; preds = %613
  %616 = getelementptr inbounds i32* %sp.0, i32 1
  %617 = load i32* %sp.0, align 4
  %618 = icmp sle i32 %617, %a.0
  %619 = zext i1 %618 to i32
  br label %775

; <label>:620                                     ; preds = %613
  %621 = icmp eq i32 %438, 22
  br i1 %621, label %622, label %627

; <label>:622                                     ; preds = %620
  %623 = getelementptr inbounds i32* %sp.0, i32 1
  %624 = load i32* %sp.0, align 4
  %625 = icmp sge i32 %624, %a.0
  %626 = zext i1 %625 to i32
  br label %774

; <label>:627                                     ; preds = %620
  %628 = icmp eq i32 %438, 23
  br i1 %628, label %629, label %633

; <label>:629                                     ; preds = %627
  %630 = getelementptr inbounds i32* %sp.0, i32 1
  %631 = load i32* %sp.0, align 4
  %632 = shl i32 %631, %a.0
  br label %773

; <label>:633                                     ; preds = %627
  %634 = icmp eq i32 %438, 24
  br i1 %634, label %635, label %639

; <label>:635                                     ; preds = %633
  %636 = getelementptr inbounds i32* %sp.0, i32 1
  %637 = load i32* %sp.0, align 4
  %638 = ashr i32 %637, %a.0
  br label %772

; <label>:639                                     ; preds = %633
  %640 = icmp eq i32 %438, 25
  br i1 %640, label %641, label %645

; <label>:641                                     ; preds = %639
  %642 = getelementptr inbounds i32* %sp.0, i32 1
  %643 = load i32* %sp.0, align 4
  %644 = add nsw i32 %643, %a.0
  br label %771

; <label>:645                                     ; preds = %639
  %646 = icmp eq i32 %438, 26
  br i1 %646, label %647, label %651

; <label>:647                                     ; preds = %645
  %648 = getelementptr inbounds i32* %sp.0, i32 1
  %649 = load i32* %sp.0, align 4
  %650 = sub nsw i32 %649, %a.0
  br label %770

; <label>:651                                     ; preds = %645
  %652 = icmp eq i32 %438, 27
  br i1 %652, label %653, label %657

; <label>:653                                     ; preds = %651
  %654 = getelementptr inbounds i32* %sp.0, i32 1
  %655 = load i32* %sp.0, align 4
  %656 = mul nsw i32 %655, %a.0
  br label %769

; <label>:657                                     ; preds = %651
  %658 = icmp eq i32 %438, 28
  br i1 %658, label %659, label %663

; <label>:659                                     ; preds = %657
  %660 = getelementptr inbounds i32* %sp.0, i32 1
  %661 = load i32* %sp.0, align 4
  %662 = sdiv i32 %661, %a.0
  br label %768

; <label>:663                                     ; preds = %657
  %664 = icmp eq i32 %438, 29
  br i1 %664, label %665, label %669

; <label>:665                                     ; preds = %663
  %666 = getelementptr inbounds i32* %sp.0, i32 1
  %667 = load i32* %sp.0, align 4
  %668 = srem i32 %667, %a.0
  br label %767

; <label>:669                                     ; preds = %663
  %670 = icmp eq i32 %438, 30
  br i1 %670, label %671, label %678

; <label>:671                                     ; preds = %669
  %672 = getelementptr inbounds i32* %sp.0, i64 1
  %673 = load i32* %672, align 4
  %674 = sext i32 %673 to i64
  %675 = inttoptr i64 %674 to i8*
  %676 = load i32* %sp.0, align 4
  %677 = call i32 (i8*, i32, ...)* bitcast (i32 (...)* @open to i32 (i8*, i32, ...)*)(i8* %675, i32 %676)
  br label %766

; <label>:678                                     ; preds = %669
  %679 = icmp eq i32 %438, 31
  br i1 %679, label %680, label %691

; <label>:680                                     ; preds = %678
  %681 = getelementptr inbounds i32* %sp.0, i64 2
  %682 = load i32* %681, align 4
  %683 = getelementptr inbounds i32* %sp.0, i64 1
  %684 = load i32* %683, align 4
  %685 = sext i32 %684 to i64
  %686 = inttoptr i64 %685 to i8*
  %687 = load i32* %sp.0, align 4
  %688 = sext i32 %687 to i64
  %689 = call i64 @read(i32 %682, i8* %686, i64 %688)
  %690 = trunc i64 %689 to i32
  br label %765

; <label>:691                                     ; preds = %678
  %692 = icmp eq i32 %438, 32
  br i1 %692, label %693, label %696

; <label>:693                                     ; preds = %691
  %694 = load i32* %sp.0, align 4
  %695 = call i32 @close(i32 %694)
  br label %764

; <label>:696                                     ; preds = %691
  %697 = icmp eq i32 %438, 33
  br i1 %697, label %698, label %718

; <label>:698                                     ; preds = %696
  %699 = getelementptr inbounds i32* %437, i64 1
  %700 = load i32* %699, align 4
  %701 = sext i32 %700 to i64
  %702 = getelementptr inbounds i32* %sp.0, i64 %701
  %703 = getelementptr inbounds i32* %702, i64 -1
  %704 = load i32* %703, align 4
  %705 = sext i32 %704 to i64
  %706 = inttoptr i64 %705 to i8*
  %707 = getelementptr inbounds i32* %702, i64 -2
  %708 = load i32* %707, align 4
  %709 = getelementptr inbounds i32* %702, i64 -3
  %710 = load i32* %709, align 4
  %711 = getelementptr inbounds i32* %702, i64 -4
  %712 = load i32* %711, align 4
  %713 = getelementptr inbounds i32* %702, i64 -5
  %714 = load i32* %713, align 4
  %715 = getelementptr inbounds i32* %702, i64 -6
  %716 = load i32* %715, align 4
  %717 = call i32 (i8*, ...)* @printf(i8* %706, i32 %708, i32 %710, i32 %712, i32 %714, i32 %716)
  br label %763

; <label>:718                                     ; preds = %696
  %719 = icmp eq i32 %438, 34
  br i1 %719, label %720, label %725

; <label>:720                                     ; preds = %718
  %721 = load i32* %sp.0, align 4
  %722 = sext i32 %721 to i64
  %723 = call noalias i8* @malloc(i64 %722) #5
  %724 = ptrtoint i8* %723 to i32
  br label %762

; <label>:725                                     ; preds = %718
  %726 = icmp eq i32 %438, 35
  br i1 %726, label %727, label %738

; <label>:727                                     ; preds = %725
  %728 = getelementptr inbounds i32* %sp.0, i64 2
  %729 = load i32* %728, align 4
  %730 = sext i32 %729 to i64
  %731 = inttoptr i64 %730 to i8*
  %732 = getelementptr inbounds i32* %sp.0, i64 1
  %733 = load i32* %732, align 4
  %734 = trunc i32 %733 to i8
  %735 = load i32* %sp.0, align 4
  %736 = sext i32 %735 to i64
  call void @llvm.memset.p0i8.i64(i8* %731, i8 %734, i64 %736, i32 1, i1 false)
  %737 = ptrtoint i8* %731 to i32
  br label %761

; <label>:738                                     ; preds = %725
  %739 = icmp eq i32 %438, 36
  br i1 %739, label %740, label %752

; <label>:740                                     ; preds = %738
  %741 = getelementptr inbounds i32* %sp.0, i64 2
  %742 = load i32* %741, align 4
  %743 = sext i32 %742 to i64
  %744 = inttoptr i64 %743 to i8*
  %745 = getelementptr inbounds i32* %sp.0, i64 1
  %746 = load i32* %745, align 4
  %747 = sext i32 %746 to i64
  %748 = inttoptr i64 %747 to i8*
  %749 = load i32* %sp.0, align 4
  %750 = sext i32 %749 to i64
  %751 = call i32 @memcmp(i8* %744, i8* %748, i64 %750) #6
  br label %760

; <label>:752                                     ; preds = %738
  %753 = icmp eq i32 %438, 37
  br i1 %753, label %754, label %758

; <label>:754                                     ; preds = %752
  %755 = load i32* %sp.0, align 4
  %756 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([21 x i8]* @.str44, i32 0, i32 0), i32 %755, i32 %439)
  %757 = load i32* %sp.0, align 4
  br label %797

; <label>:758                                     ; preds = %752
  %759 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([38 x i8]* @.str45, i32 0, i32 0), i32 %438, i32 %439)
  br label %797

; <label>:760                                     ; preds = %740
  br label %761

; <label>:761                                     ; preds = %760, %727
  %a.1 = phi i32 [ %737, %727 ], [ %751, %760 ]
  br label %762

; <label>:762                                     ; preds = %761, %720
  %a.2 = phi i32 [ %724, %720 ], [ %a.1, %761 ]
  br label %763

; <label>:763                                     ; preds = %762, %698
  %a.3 = phi i32 [ %717, %698 ], [ %a.2, %762 ]
  br label %764

; <label>:764                                     ; preds = %763, %693
  %a.4 = phi i32 [ %695, %693 ], [ %a.3, %763 ]
  br label %765

; <label>:765                                     ; preds = %764, %680
  %a.5 = phi i32 [ %690, %680 ], [ %a.4, %764 ]
  br label %766

; <label>:766                                     ; preds = %765, %671
  %a.6 = phi i32 [ %677, %671 ], [ %a.5, %765 ]
  br label %767

; <label>:767                                     ; preds = %766, %665
  %a.7 = phi i32 [ %668, %665 ], [ %a.6, %766 ]
  %sp.1 = phi i32* [ %666, %665 ], [ %sp.0, %766 ]
  br label %768

; <label>:768                                     ; preds = %767, %659
  %a.8 = phi i32 [ %662, %659 ], [ %a.7, %767 ]
  %sp.2 = phi i32* [ %660, %659 ], [ %sp.1, %767 ]
  br label %769

; <label>:769                                     ; preds = %768, %653
  %a.9 = phi i32 [ %656, %653 ], [ %a.8, %768 ]
  %sp.3 = phi i32* [ %654, %653 ], [ %sp.2, %768 ]
  br label %770

; <label>:770                                     ; preds = %769, %647
  %a.10 = phi i32 [ %650, %647 ], [ %a.9, %769 ]
  %sp.4 = phi i32* [ %648, %647 ], [ %sp.3, %769 ]
  br label %771

; <label>:771                                     ; preds = %770, %641
  %a.11 = phi i32 [ %644, %641 ], [ %a.10, %770 ]
  %sp.5 = phi i32* [ %642, %641 ], [ %sp.4, %770 ]
  br label %772

; <label>:772                                     ; preds = %771, %635
  %a.12 = phi i32 [ %638, %635 ], [ %a.11, %771 ]
  %sp.6 = phi i32* [ %636, %635 ], [ %sp.5, %771 ]
  br label %773

; <label>:773                                     ; preds = %772, %629
  %a.13 = phi i32 [ %632, %629 ], [ %a.12, %772 ]
  %sp.7 = phi i32* [ %630, %629 ], [ %sp.6, %772 ]
  br label %774

; <label>:774                                     ; preds = %773, %622
  %a.14 = phi i32 [ %626, %622 ], [ %a.13, %773 ]
  %sp.8 = phi i32* [ %623, %622 ], [ %sp.7, %773 ]
  br label %775

; <label>:775                                     ; preds = %774, %615
  %a.15 = phi i32 [ %619, %615 ], [ %a.14, %774 ]
  %sp.9 = phi i32* [ %616, %615 ], [ %sp.8, %774 ]
  br label %776

; <label>:776                                     ; preds = %775, %608
  %a.16 = phi i32 [ %612, %608 ], [ %a.15, %775 ]
  %sp.10 = phi i32* [ %609, %608 ], [ %sp.9, %775 ]
  br label %777

; <label>:777                                     ; preds = %776, %601
  %a.17 = phi i32 [ %605, %601 ], [ %a.16, %776 ]
  %sp.11 = phi i32* [ %602, %601 ], [ %sp.10, %776 ]
  br label %778

; <label>:778                                     ; preds = %777, %594
  %a.18 = phi i32 [ %598, %594 ], [ %a.17, %777 ]
  %sp.12 = phi i32* [ %595, %594 ], [ %sp.11, %777 ]
  br label %779

; <label>:779                                     ; preds = %778, %587
  %a.19 = phi i32 [ %591, %587 ], [ %a.18, %778 ]
  %sp.13 = phi i32* [ %588, %587 ], [ %sp.12, %778 ]
  br label %780

; <label>:780                                     ; preds = %779, %581
  %a.20 = phi i32 [ %584, %581 ], [ %a.19, %779 ]
  %sp.14 = phi i32* [ %582, %581 ], [ %sp.13, %779 ]
  br label %781

; <label>:781                                     ; preds = %780, %575
  %a.21 = phi i32 [ %578, %575 ], [ %a.20, %780 ]
  %sp.15 = phi i32* [ %576, %575 ], [ %sp.14, %780 ]
  br label %782

; <label>:782                                     ; preds = %781, %569
  %a.22 = phi i32 [ %572, %569 ], [ %a.21, %781 ]
  %sp.16 = phi i32* [ %570, %569 ], [ %sp.15, %781 ]
  br label %783

; <label>:783                                     ; preds = %782, %565
  %a.23 = phi i32 [ %a.0, %565 ], [ %a.22, %782 ]
  %sp.17 = phi i32* [ %566, %565 ], [ %sp.16, %782 ]
  br label %784

; <label>:784                                     ; preds = %783, %556
  %a.24 = phi i32 [ %562, %556 ], [ %a.23, %783 ]
  %sp.18 = phi i32* [ %558, %556 ], [ %sp.17, %783 ]
  br label %785

; <label>:785                                     ; preds = %784, %549
  %a.25 = phi i32 [ %a.0, %549 ], [ %a.24, %784 ]
  %sp.19 = phi i32* [ %550, %549 ], [ %sp.18, %784 ]
  br label %786

; <label>:786                                     ; preds = %785, %542
  %a.26 = phi i32 [ %546, %542 ], [ %a.25, %785 ]
  %sp.20 = phi i32* [ %sp.0, %542 ], [ %sp.19, %785 ]
  br label %787

; <label>:787                                     ; preds = %786, %536
  %a.27 = phi i32 [ %539, %536 ], [ %a.26, %786 ]
  %sp.21 = phi i32* [ %sp.0, %536 ], [ %sp.20, %786 ]
  br label %788

; <label>:788                                     ; preds = %787, %525
  %a.28 = phi i32 [ %a.0, %525 ], [ %a.27, %787 ]
  %bp.1 = phi i32* [ %529, %525 ], [ %bp.0, %787 ]
  %sp.22 = phi i32* [ %530, %525 ], [ %sp.21, %787 ]
  %pc.1 = phi i32* [ %533, %525 ], [ %437, %787 ]
  br label %789

; <label>:789                                     ; preds = %788, %518
  %a.29 = phi i32 [ %a.0, %518 ], [ %a.28, %788 ]
  %bp.2 = phi i32* [ %bp.0, %518 ], [ %bp.1, %788 ]
  %sp.23 = phi i32* [ %522, %518 ], [ %sp.22, %788 ]
  %pc.2 = phi i32* [ %519, %518 ], [ %pc.1, %788 ]
  br label %790

; <label>:790                                     ; preds = %789, %508
  %a.30 = phi i32 [ %a.0, %508 ], [ %a.29, %789 ]
  %bp.3 = phi i32* [ %510, %508 ], [ %bp.2, %789 ]
  %sp.24 = phi i32* [ %515, %508 ], [ %sp.23, %789 ]
  %pc.3 = phi i32* [ %511, %508 ], [ %pc.2, %789 ]
  br label %791

; <label>:791                                     ; preds = %790, %504
  %a.31 = phi i32 [ %a.0, %504 ], [ %a.30, %790 ]
  %bp.4 = phi i32* [ %bp.0, %504 ], [ %bp.3, %790 ]
  %sp.25 = phi i32* [ %sp.0, %504 ], [ %sp.24, %790 ]
  %pc.4 = phi i32* [ %505, %504 ], [ %pc.3, %790 ]
  br label %792

; <label>:792                                     ; preds = %791, %492
  %a.32 = phi i32 [ %a.0, %492 ], [ %a.31, %791 ]
  %bp.5 = phi i32* [ %bp.0, %492 ], [ %bp.4, %791 ]
  %sp.26 = phi i32* [ %sp.0, %492 ], [ %sp.25, %791 ]
  %pc.5 = phi i32* [ %493, %492 ], [ %pc.4, %791 ]
  br label %793

; <label>:793                                     ; preds = %792, %475
  %a.33 = phi i32 [ %a.0, %475 ], [ %a.32, %792 ]
  %bp.6 = phi i32* [ %bp.0, %475 ], [ %bp.5, %792 ]
  %sp.27 = phi i32* [ %478, %475 ], [ %sp.26, %792 ]
  %pc.6 = phi i32* [ %481, %475 ], [ %pc.5, %792 ]
  br label %794

; <label>:794                                     ; preds = %793, %469
  %a.34 = phi i32 [ %a.0, %469 ], [ %a.33, %793 ]
  %bp.7 = phi i32* [ %bp.0, %469 ], [ %bp.6, %793 ]
  %sp.28 = phi i32* [ %sp.0, %469 ], [ %sp.27, %793 ]
  %pc.7 = phi i32* [ %472, %469 ], [ %pc.6, %793 ]
  br label %795

; <label>:795                                     ; preds = %794, %464
  %a.35 = phi i32 [ %466, %464 ], [ %a.34, %794 ]
  %bp.8 = phi i32* [ %bp.0, %464 ], [ %bp.7, %794 ]
  %sp.29 = phi i32* [ %sp.0, %464 ], [ %sp.28, %794 ]
  %pc.8 = phi i32* [ %465, %464 ], [ %pc.7, %794 ]
  br label %796

; <label>:796                                     ; preds = %795, %456
  %a.36 = phi i32 [ %461, %456 ], [ %a.35, %795 ]
  %bp.9 = phi i32* [ %bp.0, %456 ], [ %bp.8, %795 ]
  %sp.30 = phi i32* [ %sp.0, %456 ], [ %sp.29, %795 ]
  %pc.9 = phi i32* [ %457, %456 ], [ %pc.8, %795 ]
  br label %436

; <label>:797                                     ; preds = %758, %754, %423, %418, %321, %313, %284, %251, %243, %206, %198, %159, %149, %114, %105, %70, %63, %57, %50, %42, %36
  %.0 = phi i32 [ -1, %36 ], [ -1, %42 ], [ -1, %114 ], [ -1, %198 ], [ -1, %206 ], [ -1, %243 ], [ -1, %251 ], [ -1, %284 ], [ -1, %313 ], [ -1, %321 ], [ -1, %149 ], [ -1, %159 ], [ 0, %423 ], [ %757, %754 ], [ -1, %758 ], [ -1, %418 ], [ -1, %105 ], [ -1, %70 ], [ -1, %63 ], [ -1, %57 ], [ -1, %50 ]
  ret i32 %.0
}

declare i32 @open(...) #1

; Function Attrs: nounwind
declare noalias i8* @malloc(i64) #4

; Function Attrs: nounwind
declare void @llvm.memset.p0i8.i64(i8* nocapture, i8, i64, i32, i1) #5

declare i64 @read(i32, i8*, i64) #1

declare i32 @close(i32) #1

attributes #0 = { nounwind uwtable "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #1 = { "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #2 = { nounwind readonly "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #3 = { noreturn nounwind "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #4 = { nounwind "less-precise-fpmad"="false" "no-frame-pointer-elim"="true" "no-frame-pointer-elim-non-leaf" "no-infs-fp-math"="false" "no-nans-fp-math"="false" "stack-protector-buffer-size"="8" "unsafe-fp-math"="false" "use-soft-float"="false" }
attributes #5 = { nounwind }
attributes #6 = { nounwind readonly }
attributes #7 = { noreturn nounwind }

!llvm.ident = !{!0}

!0 = !{!"clang version 3.6.0 (tags/RELEASE_360/final)"}
