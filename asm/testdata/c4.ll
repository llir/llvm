; ModuleID = 'c4.c'
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
  %pp = alloca i8*, align 8
  br label %1

; <label>:1                                       ; preds = %569, %0
  %2 = load i8** @p, align 8
  %3 = load i8* %2, align 1
  %4 = sext i8 %3 to i32
  store i32 %4, i32* @tk, align 4
  %5 = icmp ne i32 %4, 0
  br i1 %5, label %6, label %570

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
  br label %569

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
  br label %568

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
  br i1 %85, label %86, label %188

; <label>:86                                      ; preds = %83, %80, %74
  %87 = load i8** @p, align 8
  %88 = getelementptr inbounds i8* %87, i64 -1
  store i8* %88, i8** %pp, align 8
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
  %139 = load i8** %pp, align 8
  %140 = ptrtoint i8* %138 to i64
  %141 = ptrtoint i8* %139 to i64
  %142 = sub i64 %140, %141
  %143 = add nsw i64 %137, %142
  %144 = trunc i64 %143 to i32
  store i32 %144, i32* @tk, align 4
  %145 = load i32** @sym, align 8
  store i32* %145, i32** @id, align 8
  br label %146

; <label>:146                                     ; preds = %175, %134
  %147 = load i32** @id, align 8
  %148 = getelementptr inbounds i32* %147, i64 0
  %149 = load i32* %148, align 4
  %150 = icmp ne i32 %149, 0
  br i1 %150, label %151, label %178

; <label>:151                                     ; preds = %146
  %152 = load i32* @tk, align 4
  %153 = load i32** @id, align 8
  %154 = getelementptr inbounds i32* %153, i64 1
  %155 = load i32* %154, align 4
  %156 = icmp eq i32 %152, %155
  br i1 %156, label %157, label %175

; <label>:157                                     ; preds = %151
  %158 = load i32** @id, align 8
  %159 = getelementptr inbounds i32* %158, i64 2
  %160 = load i32* %159, align 4
  %161 = sext i32 %160 to i64
  %162 = inttoptr i64 %161 to i8*
  %163 = load i8** %pp, align 8
  %164 = load i8** @p, align 8
  %165 = load i8** %pp, align 8
  %166 = ptrtoint i8* %164 to i64
  %167 = ptrtoint i8* %165 to i64
  %168 = sub i64 %166, %167
  %169 = call i32 @memcmp(i8* %162, i8* %163, i64 %168) #6
  %170 = icmp ne i32 %169, 0
  br i1 %170, label %175, label %171

; <label>:171                                     ; preds = %157
  %172 = load i32** @id, align 8
  %173 = getelementptr inbounds i32* %172, i64 0
  %174 = load i32* %173, align 4
  store i32 %174, i32* @tk, align 4
  br label %570

; <label>:175                                     ; preds = %157, %151
  %176 = load i32** @id, align 8
  %177 = getelementptr inbounds i32* %176, i64 9
  store i32* %177, i32** @id, align 8
  br label %146

; <label>:178                                     ; preds = %146
  %179 = load i8** %pp, align 8
  %180 = ptrtoint i8* %179 to i32
  %181 = load i32** @id, align 8
  %182 = getelementptr inbounds i32* %181, i64 2
  store i32 %180, i32* %182, align 4
  %183 = load i32* @tk, align 4
  %184 = load i32** @id, align 8
  %185 = getelementptr inbounds i32* %184, i64 1
  store i32 %183, i32* %185, align 4
  %186 = load i32** @id, align 8
  %187 = getelementptr inbounds i32* %186, i64 0
  store i32 133, i32* %187, align 4
  store i32 133, i32* @tk, align 4
  br label %570

; <label>:188                                     ; preds = %83
  %189 = load i32* @tk, align 4
  %190 = icmp sge i32 %189, 48
  br i1 %190, label %191, label %298

; <label>:191                                     ; preds = %188
  %192 = load i32* @tk, align 4
  %193 = icmp sle i32 %192, 57
  br i1 %193, label %194, label %298

; <label>:194                                     ; preds = %191
  %195 = load i32* @tk, align 4
  %196 = sub nsw i32 %195, 48
  store i32 %196, i32* @ival, align 4
  %197 = icmp ne i32 %196, 0
  br i1 %197, label %198, label %221

; <label>:198                                     ; preds = %194
  br label %199

; <label>:199                                     ; preds = %211, %198
  %200 = load i8** @p, align 8
  %201 = load i8* %200, align 1
  %202 = sext i8 %201 to i32
  %203 = icmp sge i32 %202, 48
  br i1 %203, label %204, label %209

; <label>:204                                     ; preds = %199
  %205 = load i8** @p, align 8
  %206 = load i8* %205, align 1
  %207 = sext i8 %206 to i32
  %208 = icmp sle i32 %207, 57
  br label %209

; <label>:209                                     ; preds = %204, %199
  %210 = phi i1 [ false, %199 ], [ %208, %204 ]
  br i1 %210, label %211, label %220

; <label>:211                                     ; preds = %209
  %212 = load i32* @ival, align 4
  %213 = mul nsw i32 %212, 10
  %214 = load i8** @p, align 8
  %215 = getelementptr inbounds i8* %214, i32 1
  store i8* %215, i8** @p, align 8
  %216 = load i8* %214, align 1
  %217 = sext i8 %216 to i32
  %218 = add nsw i32 %213, %217
  %219 = sub nsw i32 %218, 48
  store i32 %219, i32* @ival, align 4
  br label %199

; <label>:220                                     ; preds = %209
  br label %297

; <label>:221                                     ; preds = %194
  %222 = load i8** @p, align 8
  %223 = load i8* %222, align 1
  %224 = sext i8 %223 to i32
  %225 = icmp eq i32 %224, 120
  br i1 %225, label %231, label %226

; <label>:226                                     ; preds = %221
  %227 = load i8** @p, align 8
  %228 = load i8* %227, align 1
  %229 = sext i8 %228 to i32
  %230 = icmp eq i32 %229, 88
  br i1 %230, label %231, label %273

; <label>:231                                     ; preds = %226, %221
  br label %232

; <label>:232                                     ; preds = %262, %231
  %233 = load i8** @p, align 8
  %234 = getelementptr inbounds i8* %233, i32 1
  store i8* %234, i8** @p, align 8
  %235 = load i8* %234, align 1
  %236 = sext i8 %235 to i32
  store i32 %236, i32* @tk, align 4
  %237 = icmp ne i32 %236, 0
  br i1 %237, label %238, label %260

; <label>:238                                     ; preds = %232
  %239 = load i32* @tk, align 4
  %240 = icmp sge i32 %239, 48
  br i1 %240, label %241, label %244

; <label>:241                                     ; preds = %238
  %242 = load i32* @tk, align 4
  %243 = icmp sle i32 %242, 57
  br i1 %243, label %258, label %244

; <label>:244                                     ; preds = %241, %238
  %245 = load i32* @tk, align 4
  %246 = icmp sge i32 %245, 97
  br i1 %246, label %247, label %250

; <label>:247                                     ; preds = %244
  %248 = load i32* @tk, align 4
  %249 = icmp sle i32 %248, 102
  br i1 %249, label %258, label %250

; <label>:250                                     ; preds = %247, %244
  %251 = load i32* @tk, align 4
  %252 = icmp sge i32 %251, 65
  br i1 %252, label %253, label %256

; <label>:253                                     ; preds = %250
  %254 = load i32* @tk, align 4
  %255 = icmp sle i32 %254, 70
  br label %256

; <label>:256                                     ; preds = %253, %250
  %257 = phi i1 [ false, %250 ], [ %255, %253 ]
  br label %258

; <label>:258                                     ; preds = %256, %247, %241
  %259 = phi i1 [ true, %247 ], [ true, %241 ], [ %257, %256 ]
  br label %260

; <label>:260                                     ; preds = %258, %232
  %261 = phi i1 [ false, %232 ], [ %259, %258 ]
  br i1 %261, label %262, label %272

; <label>:262                                     ; preds = %260
  %263 = load i32* @ival, align 4
  %264 = mul nsw i32 %263, 16
  %265 = load i32* @tk, align 4
  %266 = and i32 %265, 15
  %267 = add nsw i32 %264, %266
  %268 = load i32* @tk, align 4
  %269 = icmp sge i32 %268, 65
  %270 = select i1 %269, i32 9, i32 0
  %271 = add nsw i32 %267, %270
  store i32 %271, i32* @ival, align 4
  br label %232

; <label>:272                                     ; preds = %260
  br label %296

; <label>:273                                     ; preds = %226
  br label %274

; <label>:274                                     ; preds = %286, %273
  %275 = load i8** @p, align 8
  %276 = load i8* %275, align 1
  %277 = sext i8 %276 to i32
  %278 = icmp sge i32 %277, 48
  br i1 %278, label %279, label %284

; <label>:279                                     ; preds = %274
  %280 = load i8** @p, align 8
  %281 = load i8* %280, align 1
  %282 = sext i8 %281 to i32
  %283 = icmp sle i32 %282, 55
  br label %284

; <label>:284                                     ; preds = %279, %274
  %285 = phi i1 [ false, %274 ], [ %283, %279 ]
  br i1 %285, label %286, label %295

; <label>:286                                     ; preds = %284
  %287 = load i32* @ival, align 4
  %288 = mul nsw i32 %287, 8
  %289 = load i8** @p, align 8
  %290 = getelementptr inbounds i8* %289, i32 1
  store i8* %290, i8** @p, align 8
  %291 = load i8* %289, align 1
  %292 = sext i8 %291 to i32
  %293 = add nsw i32 %288, %292
  %294 = sub nsw i32 %293, 48
  store i32 %294, i32* @ival, align 4
  br label %274

; <label>:295                                     ; preds = %284
  br label %296

; <label>:296                                     ; preds = %295, %272
  br label %297

; <label>:297                                     ; preds = %296, %220
  store i32 128, i32* @tk, align 4
  br label %570

; <label>:298                                     ; preds = %191, %188
  %299 = load i32* @tk, align 4
  %300 = icmp eq i32 %299, 47
  br i1 %300, label %301, label %327

; <label>:301                                     ; preds = %298
  %302 = load i8** @p, align 8
  %303 = load i8* %302, align 1
  %304 = sext i8 %303 to i32
  %305 = icmp eq i32 %304, 47
  br i1 %305, label %306, label %325

; <label>:306                                     ; preds = %301
  %307 = load i8** @p, align 8
  %308 = getelementptr inbounds i8* %307, i32 1
  store i8* %308, i8** @p, align 8
  br label %309

; <label>:309                                     ; preds = %321, %306
  %310 = load i8** @p, align 8
  %311 = load i8* %310, align 1
  %312 = sext i8 %311 to i32
  %313 = icmp ne i32 %312, 0
  br i1 %313, label %314, label %319

; <label>:314                                     ; preds = %309
  %315 = load i8** @p, align 8
  %316 = load i8* %315, align 1
  %317 = sext i8 %316 to i32
  %318 = icmp ne i32 %317, 10
  br label %319

; <label>:319                                     ; preds = %314, %309
  %320 = phi i1 [ false, %309 ], [ %318, %314 ]
  br i1 %320, label %321, label %324

; <label>:321                                     ; preds = %319
  %322 = load i8** @p, align 8
  %323 = getelementptr inbounds i8* %322, i32 1
  store i8* %323, i8** @p, align 8
  br label %309

; <label>:324                                     ; preds = %319
  br label %326

; <label>:325                                     ; preds = %301
  store i32 160, i32* @tk, align 4
  br label %570

; <label>:326                                     ; preds = %324
  br label %565

; <label>:327                                     ; preds = %298
  %328 = load i32* @tk, align 4
  %329 = icmp eq i32 %328, 39
  br i1 %329, label %333, label %330

; <label>:330                                     ; preds = %327
  %331 = load i32* @tk, align 4
  %332 = icmp eq i32 %331, 34
  br i1 %332, label %333, label %381

; <label>:333                                     ; preds = %330, %327
  %334 = load i8** @data, align 8
  store i8* %334, i8** %pp, align 8
  br label %335

; <label>:335                                     ; preds = %370, %333
  %336 = load i8** @p, align 8
  %337 = load i8* %336, align 1
  %338 = sext i8 %337 to i32
  %339 = icmp ne i32 %338, 0
  br i1 %339, label %340, label %346

; <label>:340                                     ; preds = %335
  %341 = load i8** @p, align 8
  %342 = load i8* %341, align 1
  %343 = sext i8 %342 to i32
  %344 = load i32* @tk, align 4
  %345 = icmp ne i32 %343, %344
  br label %346

; <label>:346                                     ; preds = %340, %335
  %347 = phi i1 [ false, %335 ], [ %345, %340 ]
  br i1 %347, label %348, label %371

; <label>:348                                     ; preds = %346
  %349 = load i8** @p, align 8
  %350 = getelementptr inbounds i8* %349, i32 1
  store i8* %350, i8** @p, align 8
  %351 = load i8* %349, align 1
  %352 = sext i8 %351 to i32
  store i32 %352, i32* @ival, align 4
  %353 = icmp eq i32 %352, 92
  br i1 %353, label %354, label %362

; <label>:354                                     ; preds = %348
  %355 = load i8** @p, align 8
  %356 = getelementptr inbounds i8* %355, i32 1
  store i8* %356, i8** @p, align 8
  %357 = load i8* %355, align 1
  %358 = sext i8 %357 to i32
  store i32 %358, i32* @ival, align 4
  %359 = icmp eq i32 %358, 110
  br i1 %359, label %360, label %361

; <label>:360                                     ; preds = %354
  store i32 10, i32* @ival, align 4
  br label %361

; <label>:361                                     ; preds = %360, %354
  br label %362

; <label>:362                                     ; preds = %361, %348
  %363 = load i32* @tk, align 4
  %364 = icmp eq i32 %363, 34
  br i1 %364, label %365, label %370

; <label>:365                                     ; preds = %362
  %366 = load i32* @ival, align 4
  %367 = trunc i32 %366 to i8
  %368 = load i8** @data, align 8
  %369 = getelementptr inbounds i8* %368, i32 1
  store i8* %369, i8** @data, align 8
  store i8 %367, i8* %368, align 1
  br label %370

; <label>:370                                     ; preds = %365, %362
  br label %335

; <label>:371                                     ; preds = %346
  %372 = load i8** @p, align 8
  %373 = getelementptr inbounds i8* %372, i32 1
  store i8* %373, i8** @p, align 8
  %374 = load i32* @tk, align 4
  %375 = icmp eq i32 %374, 34
  br i1 %375, label %376, label %379

; <label>:376                                     ; preds = %371
  %377 = load i8** %pp, align 8
  %378 = ptrtoint i8* %377 to i32
  store i32 %378, i32* @ival, align 4
  br label %380

; <label>:379                                     ; preds = %371
  store i32 128, i32* @tk, align 4
  br label %380

; <label>:380                                     ; preds = %379, %376
  br label %570

; <label>:381                                     ; preds = %330
  %382 = load i32* @tk, align 4
  %383 = icmp eq i32 %382, 61
  br i1 %383, label %384, label %394

; <label>:384                                     ; preds = %381
  %385 = load i8** @p, align 8
  %386 = load i8* %385, align 1
  %387 = sext i8 %386 to i32
  %388 = icmp eq i32 %387, 61
  br i1 %388, label %389, label %392

; <label>:389                                     ; preds = %384
  %390 = load i8** @p, align 8
  %391 = getelementptr inbounds i8* %390, i32 1
  store i8* %391, i8** @p, align 8
  store i32 149, i32* @tk, align 4
  br label %393

; <label>:392                                     ; preds = %384
  store i32 142, i32* @tk, align 4
  br label %393

; <label>:393                                     ; preds = %392, %389
  br label %570

; <label>:394                                     ; preds = %381
  %395 = load i32* @tk, align 4
  %396 = icmp eq i32 %395, 43
  br i1 %396, label %397, label %407

; <label>:397                                     ; preds = %394
  %398 = load i8** @p, align 8
  %399 = load i8* %398, align 1
  %400 = sext i8 %399 to i32
  %401 = icmp eq i32 %400, 43
  br i1 %401, label %402, label %405

; <label>:402                                     ; preds = %397
  %403 = load i8** @p, align 8
  %404 = getelementptr inbounds i8* %403, i32 1
  store i8* %404, i8** @p, align 8
  store i32 162, i32* @tk, align 4
  br label %406

; <label>:405                                     ; preds = %397
  store i32 157, i32* @tk, align 4
  br label %406

; <label>:406                                     ; preds = %405, %402
  br label %570

; <label>:407                                     ; preds = %394
  %408 = load i32* @tk, align 4
  %409 = icmp eq i32 %408, 45
  br i1 %409, label %410, label %420

; <label>:410                                     ; preds = %407
  %411 = load i8** @p, align 8
  %412 = load i8* %411, align 1
  %413 = sext i8 %412 to i32
  %414 = icmp eq i32 %413, 45
  br i1 %414, label %415, label %418

; <label>:415                                     ; preds = %410
  %416 = load i8** @p, align 8
  %417 = getelementptr inbounds i8* %416, i32 1
  store i8* %417, i8** @p, align 8
  store i32 163, i32* @tk, align 4
  br label %419

; <label>:418                                     ; preds = %410
  store i32 158, i32* @tk, align 4
  br label %419

; <label>:419                                     ; preds = %418, %415
  br label %570

; <label>:420                                     ; preds = %407
  %421 = load i32* @tk, align 4
  %422 = icmp eq i32 %421, 33
  br i1 %422, label %423, label %432

; <label>:423                                     ; preds = %420
  %424 = load i8** @p, align 8
  %425 = load i8* %424, align 1
  %426 = sext i8 %425 to i32
  %427 = icmp eq i32 %426, 61
  br i1 %427, label %428, label %431

; <label>:428                                     ; preds = %423
  %429 = load i8** @p, align 8
  %430 = getelementptr inbounds i8* %429, i32 1
  store i8* %430, i8** @p, align 8
  store i32 150, i32* @tk, align 4
  br label %431

; <label>:431                                     ; preds = %428, %423
  br label %570

; <label>:432                                     ; preds = %420
  %433 = load i32* @tk, align 4
  %434 = icmp eq i32 %433, 60
  br i1 %434, label %435, label %454

; <label>:435                                     ; preds = %432
  %436 = load i8** @p, align 8
  %437 = load i8* %436, align 1
  %438 = sext i8 %437 to i32
  %439 = icmp eq i32 %438, 61
  br i1 %439, label %440, label %443

; <label>:440                                     ; preds = %435
  %441 = load i8** @p, align 8
  %442 = getelementptr inbounds i8* %441, i32 1
  store i8* %442, i8** @p, align 8
  store i32 153, i32* @tk, align 4
  br label %453

; <label>:443                                     ; preds = %435
  %444 = load i8** @p, align 8
  %445 = load i8* %444, align 1
  %446 = sext i8 %445 to i32
  %447 = icmp eq i32 %446, 60
  br i1 %447, label %448, label %451

; <label>:448                                     ; preds = %443
  %449 = load i8** @p, align 8
  %450 = getelementptr inbounds i8* %449, i32 1
  store i8* %450, i8** @p, align 8
  store i32 155, i32* @tk, align 4
  br label %452

; <label>:451                                     ; preds = %443
  store i32 151, i32* @tk, align 4
  br label %452

; <label>:452                                     ; preds = %451, %448
  br label %453

; <label>:453                                     ; preds = %452, %440
  br label %570

; <label>:454                                     ; preds = %432
  %455 = load i32* @tk, align 4
  %456 = icmp eq i32 %455, 62
  br i1 %456, label %457, label %476

; <label>:457                                     ; preds = %454
  %458 = load i8** @p, align 8
  %459 = load i8* %458, align 1
  %460 = sext i8 %459 to i32
  %461 = icmp eq i32 %460, 61
  br i1 %461, label %462, label %465

; <label>:462                                     ; preds = %457
  %463 = load i8** @p, align 8
  %464 = getelementptr inbounds i8* %463, i32 1
  store i8* %464, i8** @p, align 8
  store i32 154, i32* @tk, align 4
  br label %475

; <label>:465                                     ; preds = %457
  %466 = load i8** @p, align 8
  %467 = load i8* %466, align 1
  %468 = sext i8 %467 to i32
  %469 = icmp eq i32 %468, 62
  br i1 %469, label %470, label %473

; <label>:470                                     ; preds = %465
  %471 = load i8** @p, align 8
  %472 = getelementptr inbounds i8* %471, i32 1
  store i8* %472, i8** @p, align 8
  store i32 156, i32* @tk, align 4
  br label %474

; <label>:473                                     ; preds = %465
  store i32 152, i32* @tk, align 4
  br label %474

; <label>:474                                     ; preds = %473, %470
  br label %475

; <label>:475                                     ; preds = %474, %462
  br label %570

; <label>:476                                     ; preds = %454
  %477 = load i32* @tk, align 4
  %478 = icmp eq i32 %477, 124
  br i1 %478, label %479, label %489

; <label>:479                                     ; preds = %476
  %480 = load i8** @p, align 8
  %481 = load i8* %480, align 1
  %482 = sext i8 %481 to i32
  %483 = icmp eq i32 %482, 124
  br i1 %483, label %484, label %487

; <label>:484                                     ; preds = %479
  %485 = load i8** @p, align 8
  %486 = getelementptr inbounds i8* %485, i32 1
  store i8* %486, i8** @p, align 8
  store i32 144, i32* @tk, align 4
  br label %488

; <label>:487                                     ; preds = %479
  store i32 146, i32* @tk, align 4
  br label %488

; <label>:488                                     ; preds = %487, %484
  br label %570

; <label>:489                                     ; preds = %476
  %490 = load i32* @tk, align 4
  %491 = icmp eq i32 %490, 38
  br i1 %491, label %492, label %502

; <label>:492                                     ; preds = %489
  %493 = load i8** @p, align 8
  %494 = load i8* %493, align 1
  %495 = sext i8 %494 to i32
  %496 = icmp eq i32 %495, 38
  br i1 %496, label %497, label %500

; <label>:497                                     ; preds = %492
  %498 = load i8** @p, align 8
  %499 = getelementptr inbounds i8* %498, i32 1
  store i8* %499, i8** @p, align 8
  store i32 145, i32* @tk, align 4
  br label %501

; <label>:500                                     ; preds = %492
  store i32 148, i32* @tk, align 4
  br label %501

; <label>:501                                     ; preds = %500, %497
  br label %570

; <label>:502                                     ; preds = %489
  %503 = load i32* @tk, align 4
  %504 = icmp eq i32 %503, 94
  br i1 %504, label %505, label %506

; <label>:505                                     ; preds = %502
  store i32 147, i32* @tk, align 4
  br label %570

; <label>:506                                     ; preds = %502
  %507 = load i32* @tk, align 4
  %508 = icmp eq i32 %507, 37
  br i1 %508, label %509, label %510

; <label>:509                                     ; preds = %506
  store i32 161, i32* @tk, align 4
  br label %570

; <label>:510                                     ; preds = %506
  %511 = load i32* @tk, align 4
  %512 = icmp eq i32 %511, 42
  br i1 %512, label %513, label %514

; <label>:513                                     ; preds = %510
  store i32 159, i32* @tk, align 4
  br label %570

; <label>:514                                     ; preds = %510
  %515 = load i32* @tk, align 4
  %516 = icmp eq i32 %515, 91
  br i1 %516, label %517, label %518

; <label>:517                                     ; preds = %514
  store i32 164, i32* @tk, align 4
  br label %570

; <label>:518                                     ; preds = %514
  %519 = load i32* @tk, align 4
  %520 = icmp eq i32 %519, 63
  br i1 %520, label %521, label %522

; <label>:521                                     ; preds = %518
  store i32 143, i32* @tk, align 4
  br label %570

; <label>:522                                     ; preds = %518
  %523 = load i32* @tk, align 4
  %524 = icmp eq i32 %523, 126
  br i1 %524, label %549, label %525

; <label>:525                                     ; preds = %522
  %526 = load i32* @tk, align 4
  %527 = icmp eq i32 %526, 59
  br i1 %527, label %549, label %528

; <label>:528                                     ; preds = %525
  %529 = load i32* @tk, align 4
  %530 = icmp eq i32 %529, 123
  br i1 %530, label %549, label %531

; <label>:531                                     ; preds = %528
  %532 = load i32* @tk, align 4
  %533 = icmp eq i32 %532, 125
  br i1 %533, label %549, label %534

; <label>:534                                     ; preds = %531
  %535 = load i32* @tk, align 4
  %536 = icmp eq i32 %535, 40
  br i1 %536, label %549, label %537

; <label>:537                                     ; preds = %534
  %538 = load i32* @tk, align 4
  %539 = icmp eq i32 %538, 41
  br i1 %539, label %549, label %540

; <label>:540                                     ; preds = %537
  %541 = load i32* @tk, align 4
  %542 = icmp eq i32 %541, 93
  br i1 %542, label %549, label %543

; <label>:543                                     ; preds = %540
  %544 = load i32* @tk, align 4
  %545 = icmp eq i32 %544, 44
  br i1 %545, label %549, label %546

; <label>:546                                     ; preds = %543
  %547 = load i32* @tk, align 4
  %548 = icmp eq i32 %547, 58
  br i1 %548, label %549, label %550

; <label>:549                                     ; preds = %546, %543, %540, %537, %534, %531, %528, %525, %522
  br label %570

; <label>:550                                     ; preds = %546
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

; <label>:560                                     ; preds = %559
  br label %561

; <label>:561                                     ; preds = %560
  br label %562

; <label>:562                                     ; preds = %561
  br label %563

; <label>:563                                     ; preds = %562
  br label %564

; <label>:564                                     ; preds = %563
  br label %565

; <label>:565                                     ; preds = %564, %326
  br label %566

; <label>:566                                     ; preds = %565
  br label %567

; <label>:567                                     ; preds = %566
  br label %568

; <label>:568                                     ; preds = %567, %70
  br label %569

; <label>:569                                     ; preds = %568, %48
  br label %1

; <label>:570                                     ; preds = %171, %178, %297, %325, %380, %393, %406, %419, %431, %453, %475, %488, %501, %505, %509, %513, %517, %521, %549, %1
  ret void
}

declare i32 @printf(i8*, ...) #1

; Function Attrs: nounwind readonly
declare i32 @memcmp(i8*, i8*, i64) #2

; Function Attrs: nounwind uwtable
define void @expr(i32 %lev) #0 {
  %1 = alloca i32, align 4
  %t = alloca i32, align 4
  %d = alloca i32*, align 8
  store i32 %lev, i32* %1, align 4
  %2 = load i32* @tk, align 4
  %3 = icmp ne i32 %2, 0
  br i1 %3, label %7, label %4

; <label>:4                                       ; preds = %0
  %5 = load i32* @line, align 4
  %6 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([34 x i8]* @.str5, i32 0, i32 0), i32 %5)
  call void @exit(i32 -1) #7
  unreachable

; <label>:7                                       ; preds = %0
  %8 = load i32* @tk, align 4
  %9 = icmp eq i32 %8, 128
  br i1 %9, label %10, label %16

; <label>:10                                      ; preds = %7
  %11 = load i32** @e, align 8
  %12 = getelementptr inbounds i32* %11, i32 1
  store i32* %12, i32** @e, align 8
  store i32 1, i32* %12, align 4
  %13 = load i32* @ival, align 4
  %14 = load i32** @e, align 8
  %15 = getelementptr inbounds i32* %14, i32 1
  store i32* %15, i32** @e, align 8
  store i32 %13, i32* %15, align 4
  call void @next()
  store i32 1, i32* @ty, align 4
  br label %380

; <label>:16                                      ; preds = %7
  %17 = load i32* @tk, align 4
  %18 = icmp eq i32 %17, 34
  br i1 %18, label %19, label %36

; <label>:19                                      ; preds = %16
  %20 = load i32** @e, align 8
  %21 = getelementptr inbounds i32* %20, i32 1
  store i32* %21, i32** @e, align 8
  store i32 1, i32* %21, align 4
  %22 = load i32* @ival, align 4
  %23 = load i32** @e, align 8
  %24 = getelementptr inbounds i32* %23, i32 1
  store i32* %24, i32** @e, align 8
  store i32 %22, i32* %24, align 4
  call void @next()
  br label %25

; <label>:25                                      ; preds = %28, %19
  %26 = load i32* @tk, align 4
  %27 = icmp eq i32 %26, 34
  br i1 %27, label %28, label %29

; <label>:28                                      ; preds = %25
  call void @next()
  br label %25

; <label>:29                                      ; preds = %25
  %30 = load i8** @data, align 8
  %31 = ptrtoint i8* %30 to i32
  %32 = sext i32 %31 to i64
  %33 = add i64 %32, 4
  %34 = and i64 %33, -4
  %35 = inttoptr i64 %34 to i8*
  store i8* %35, i8** @data, align 8
  store i32 2, i32* @ty, align 4
  br label %379

; <label>:36                                      ; preds = %16
  %37 = load i32* @tk, align 4
  %38 = icmp eq i32 %37, 140
  br i1 %38, label %39, label %78

; <label>:39                                      ; preds = %36
  call void @next()
  %40 = load i32* @tk, align 4
  %41 = icmp eq i32 %40, 40
  br i1 %41, label %42, label %43

; <label>:42                                      ; preds = %39
  call void @next()
  br label %46

; <label>:43                                      ; preds = %39
  %44 = load i32* @line, align 4
  %45 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([35 x i8]* @.str6, i32 0, i32 0), i32 %44)
  call void @exit(i32 -1) #7
  unreachable

; <label>:46                                      ; preds = %42
  store i32 1, i32* @ty, align 4
  %47 = load i32* @tk, align 4
  %48 = icmp eq i32 %47, 138
  br i1 %48, label %49, label %50

; <label>:49                                      ; preds = %46
  call void @next()
  br label %55

; <label>:50                                      ; preds = %46
  %51 = load i32* @tk, align 4
  %52 = icmp eq i32 %51, 134
  br i1 %52, label %53, label %54

; <label>:53                                      ; preds = %50
  call void @next()
  store i32 0, i32* @ty, align 4
  br label %54

; <label>:54                                      ; preds = %53, %50
  br label %55

; <label>:55                                      ; preds = %54, %49
  br label %56

; <label>:56                                      ; preds = %59, %55
  %57 = load i32* @tk, align 4
  %58 = icmp eq i32 %57, 159
  br i1 %58, label %59, label %62

; <label>:59                                      ; preds = %56
  call void @next()
  %60 = load i32* @ty, align 4
  %61 = add nsw i32 %60, 2
  store i32 %61, i32* @ty, align 4
  br label %56

; <label>:62                                      ; preds = %56
  %63 = load i32* @tk, align 4
  %64 = icmp eq i32 %63, 41
  br i1 %64, label %65, label %66

; <label>:65                                      ; preds = %62
  call void @next()
  br label %69

; <label>:66                                      ; preds = %62
  %67 = load i32* @line, align 4
  %68 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([36 x i8]* @.str7, i32 0, i32 0), i32 %67)
  call void @exit(i32 -1) #7
  unreachable

; <label>:69                                      ; preds = %65
  %70 = load i32** @e, align 8
  %71 = getelementptr inbounds i32* %70, i32 1
  store i32* %71, i32** @e, align 8
  store i32 1, i32* %71, align 4
  %72 = load i32* @ty, align 4
  %73 = icmp eq i32 %72, 0
  %74 = select i1 %73, i64 1, i64 4
  %75 = trunc i64 %74 to i32
  %76 = load i32** @e, align 8
  %77 = getelementptr inbounds i32* %76, i32 1
  store i32* %77, i32** @e, align 8
  store i32 %75, i32* %77, align 4
  store i32 1, i32* @ty, align 4
  br label %378

; <label>:78                                      ; preds = %36
  %79 = load i32* @tk, align 4
  %80 = icmp eq i32 %79, 133
  br i1 %80, label %81, label %194

; <label>:81                                      ; preds = %78
  %82 = load i32** @id, align 8
  store i32* %82, i32** %d, align 8
  call void @next()
  %83 = load i32* @tk, align 4
  %84 = icmp eq i32 %83, 40
  br i1 %84, label %85, label %139

; <label>:85                                      ; preds = %81
  call void @next()
  store i32 0, i32* %t, align 4
  br label %86

; <label>:86                                      ; preds = %97, %85
  %87 = load i32* @tk, align 4
  %88 = icmp ne i32 %87, 41
  br i1 %88, label %89, label %98

; <label>:89                                      ; preds = %86
  call void @expr(i32 142)
  %90 = load i32** @e, align 8
  %91 = getelementptr inbounds i32* %90, i32 1
  store i32* %91, i32** @e, align 8
  store i32 13, i32* %91, align 4
  %92 = load i32* %t, align 4
  %93 = add nsw i32 %92, 1
  store i32 %93, i32* %t, align 4
  %94 = load i32* @tk, align 4
  %95 = icmp eq i32 %94, 44
  br i1 %95, label %96, label %97

; <label>:96                                      ; preds = %89
  call void @next()
  br label %97

; <label>:97                                      ; preds = %96, %89
  br label %86

; <label>:98                                      ; preds = %86
  call void @next()
  %99 = load i32** %d, align 8
  %100 = getelementptr inbounds i32* %99, i64 3
  %101 = load i32* %100, align 4
  %102 = icmp eq i32 %101, 130
  br i1 %102, label %103, label %109

; <label>:103                                     ; preds = %98
  %104 = load i32** %d, align 8
  %105 = getelementptr inbounds i32* %104, i64 5
  %106 = load i32* %105, align 4
  %107 = load i32** @e, align 8
  %108 = getelementptr inbounds i32* %107, i32 1
  store i32* %108, i32** @e, align 8
  store i32 %106, i32* %108, align 4
  br label %126

; <label>:109                                     ; preds = %98
  %110 = load i32** %d, align 8
  %111 = getelementptr inbounds i32* %110, i64 3
  %112 = load i32* %111, align 4
  %113 = icmp eq i32 %112, 129
  br i1 %113, label %114, label %122

; <label>:114                                     ; preds = %109
  %115 = load i32** @e, align 8
  %116 = getelementptr inbounds i32* %115, i32 1
  store i32* %116, i32** @e, align 8
  store i32 3, i32* %116, align 4
  %117 = load i32** %d, align 8
  %118 = getelementptr inbounds i32* %117, i64 5
  %119 = load i32* %118, align 4
  %120 = load i32** @e, align 8
  %121 = getelementptr inbounds i32* %120, i32 1
  store i32* %121, i32** @e, align 8
  store i32 %119, i32* %121, align 4
  br label %125

; <label>:122                                     ; preds = %109
  %123 = load i32* @line, align 4
  %124 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([23 x i8]* @.str8, i32 0, i32 0), i32 %123)
  call void @exit(i32 -1) #7
  unreachable

; <label>:125                                     ; preds = %114
  br label %126

; <label>:126                                     ; preds = %125, %103
  %127 = load i32* %t, align 4
  %128 = icmp ne i32 %127, 0
  br i1 %128, label %129, label %135

; <label>:129                                     ; preds = %126
  %130 = load i32** @e, align 8
  %131 = getelementptr inbounds i32* %130, i32 1
  store i32* %131, i32** @e, align 8
  store i32 7, i32* %131, align 4
  %132 = load i32* %t, align 4
  %133 = load i32** @e, align 8
  %134 = getelementptr inbounds i32* %133, i32 1
  store i32* %134, i32** @e, align 8
  store i32 %132, i32* %134, align 4
  br label %135

; <label>:135                                     ; preds = %129, %126
  %136 = load i32** %d, align 8
  %137 = getelementptr inbounds i32* %136, i64 4
  %138 = load i32* %137, align 4
  store i32 %138, i32* @ty, align 4
  br label %193

; <label>:139                                     ; preds = %81
  %140 = load i32** %d, align 8
  %141 = getelementptr inbounds i32* %140, i64 3
  %142 = load i32* %141, align 4
  %143 = icmp eq i32 %142, 128
  br i1 %143, label %144, label %152

; <label>:144                                     ; preds = %139
  %145 = load i32** @e, align 8
  %146 = getelementptr inbounds i32* %145, i32 1
  store i32* %146, i32** @e, align 8
  store i32 1, i32* %146, align 4
  %147 = load i32** %d, align 8
  %148 = getelementptr inbounds i32* %147, i64 5
  %149 = load i32* %148, align 4
  %150 = load i32** @e, align 8
  %151 = getelementptr inbounds i32* %150, i32 1
  store i32* %151, i32** @e, align 8
  store i32 %149, i32* %151, align 4
  store i32 1, i32* @ty, align 4
  br label %192

; <label>:152                                     ; preds = %139
  %153 = load i32** %d, align 8
  %154 = getelementptr inbounds i32* %153, i64 3
  %155 = load i32* %154, align 4
  %156 = icmp eq i32 %155, 132
  br i1 %156, label %157, label %167

; <label>:157                                     ; preds = %152
  %158 = load i32** @e, align 8
  %159 = getelementptr inbounds i32* %158, i32 1
  store i32* %159, i32** @e, align 8
  store i32 0, i32* %159, align 4
  %160 = load i32* @loc, align 4
  %161 = load i32** %d, align 8
  %162 = getelementptr inbounds i32* %161, i64 5
  %163 = load i32* %162, align 4
  %164 = sub nsw i32 %160, %163
  %165 = load i32** @e, align 8
  %166 = getelementptr inbounds i32* %165, i32 1
  store i32* %166, i32** @e, align 8
  store i32 %164, i32* %166, align 4
  br label %184

; <label>:167                                     ; preds = %152
  %168 = load i32** %d, align 8
  %169 = getelementptr inbounds i32* %168, i64 3
  %170 = load i32* %169, align 4
  %171 = icmp eq i32 %170, 131
  br i1 %171, label %172, label %180

; <label>:172                                     ; preds = %167
  %173 = load i32** @e, align 8
  %174 = getelementptr inbounds i32* %173, i32 1
  store i32* %174, i32** @e, align 8
  store i32 1, i32* %174, align 4
  %175 = load i32** %d, align 8
  %176 = getelementptr inbounds i32* %175, i64 5
  %177 = load i32* %176, align 4
  %178 = load i32** @e, align 8
  %179 = getelementptr inbounds i32* %178, i32 1
  store i32* %179, i32** @e, align 8
  store i32 %177, i32* %179, align 4
  br label %183

; <label>:180                                     ; preds = %167
  %181 = load i32* @line, align 4
  %182 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([24 x i8]* @.str9, i32 0, i32 0), i32 %181)
  call void @exit(i32 -1) #7
  unreachable

; <label>:183                                     ; preds = %172
  br label %184

; <label>:184                                     ; preds = %183, %157
  %185 = load i32** %d, align 8
  %186 = getelementptr inbounds i32* %185, i64 4
  %187 = load i32* %186, align 4
  store i32 %187, i32* @ty, align 4
  %188 = icmp eq i32 %187, 0
  %189 = select i1 %188, i32 10, i32 9
  %190 = load i32** @e, align 8
  %191 = getelementptr inbounds i32* %190, i32 1
  store i32* %191, i32** @e, align 8
  store i32 %189, i32* %191, align 4
  br label %192

; <label>:192                                     ; preds = %184, %144
  br label %193

; <label>:193                                     ; preds = %192, %135
  br label %377

; <label>:194                                     ; preds = %78
  %195 = load i32* @tk, align 4
  %196 = icmp eq i32 %195, 40
  br i1 %196, label %197, label %231

; <label>:197                                     ; preds = %194
  call void @next()
  %198 = load i32* @tk, align 4
  %199 = icmp eq i32 %198, 138
  br i1 %199, label %203, label %200

; <label>:200                                     ; preds = %197
  %201 = load i32* @tk, align 4
  %202 = icmp eq i32 %201, 134
  br i1 %202, label %203, label %222

; <label>:203                                     ; preds = %200, %197
  %204 = load i32* @tk, align 4
  %205 = icmp eq i32 %204, 138
  %206 = select i1 %205, i32 1, i32 0
  store i32 %206, i32* %t, align 4
  call void @next()
  br label %207

; <label>:207                                     ; preds = %210, %203
  %208 = load i32* @tk, align 4
  %209 = icmp eq i32 %208, 159
  br i1 %209, label %210, label %213

; <label>:210                                     ; preds = %207
  call void @next()
  %211 = load i32* %t, align 4
  %212 = add nsw i32 %211, 2
  store i32 %212, i32* %t, align 4
  br label %207

; <label>:213                                     ; preds = %207
  %214 = load i32* @tk, align 4
  %215 = icmp eq i32 %214, 41
  br i1 %215, label %216, label %217

; <label>:216                                     ; preds = %213
  call void @next()
  br label %220

; <label>:217                                     ; preds = %213
  %218 = load i32* @line, align 4
  %219 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([14 x i8]* @.str10, i32 0, i32 0), i32 %218)
  call void @exit(i32 -1) #7
  unreachable

; <label>:220                                     ; preds = %216
  call void @expr(i32 162)
  %221 = load i32* %t, align 4
  store i32 %221, i32* @ty, align 4
  br label %230

; <label>:222                                     ; preds = %200
  call void @expr(i32 142)
  %223 = load i32* @tk, align 4
  %224 = icmp eq i32 %223, 41
  br i1 %224, label %225, label %226

; <label>:225                                     ; preds = %222
  call void @next()
  br label %229

; <label>:226                                     ; preds = %222
  %227 = load i32* @line, align 4
  %228 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([26 x i8]* @.str11, i32 0, i32 0), i32 %227)
  call void @exit(i32 -1) #7
  unreachable

; <label>:229                                     ; preds = %225
  br label %230

; <label>:230                                     ; preds = %229, %220
  br label %376

; <label>:231                                     ; preds = %194
  %232 = load i32* @tk, align 4
  %233 = icmp eq i32 %232, 159
  br i1 %233, label %234, label %249

; <label>:234                                     ; preds = %231
  call void @next()
  call void @expr(i32 162)
  %235 = load i32* @ty, align 4
  %236 = icmp sgt i32 %235, 1
  br i1 %236, label %237, label %240

; <label>:237                                     ; preds = %234
  %238 = load i32* @ty, align 4
  %239 = sub nsw i32 %238, 2
  store i32 %239, i32* @ty, align 4
  br label %243

; <label>:240                                     ; preds = %234
  %241 = load i32* @line, align 4
  %242 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([21 x i8]* @.str12, i32 0, i32 0), i32 %241)
  call void @exit(i32 -1) #7
  unreachable

; <label>:243                                     ; preds = %237
  %244 = load i32* @ty, align 4
  %245 = icmp eq i32 %244, 0
  %246 = select i1 %245, i32 10, i32 9
  %247 = load i32** @e, align 8
  %248 = getelementptr inbounds i32* %247, i32 1
  store i32* %248, i32** @e, align 8
  store i32 %246, i32* %248, align 4
  br label %375

; <label>:249                                     ; preds = %231
  %250 = load i32* @tk, align 4
  %251 = icmp eq i32 %250, 148
  br i1 %251, label %252, label %269

; <label>:252                                     ; preds = %249
  call void @next()
  call void @expr(i32 162)
  %253 = load i32** @e, align 8
  %254 = load i32* %253, align 4
  %255 = icmp eq i32 %254, 10
  br i1 %255, label %260, label %256

; <label>:256                                     ; preds = %252
  %257 = load i32** @e, align 8
  %258 = load i32* %257, align 4
  %259 = icmp eq i32 %258, 9
  br i1 %259, label %260, label %263

; <label>:260                                     ; preds = %256, %252
  %261 = load i32** @e, align 8
  %262 = getelementptr inbounds i32* %261, i32 -1
  store i32* %262, i32** @e, align 8
  br label %266

; <label>:263                                     ; preds = %256
  %264 = load i32* @line, align 4
  %265 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([20 x i8]* @.str13, i32 0, i32 0), i32 %264)
  call void @exit(i32 -1) #7
  unreachable

; <label>:266                                     ; preds = %260
  %267 = load i32* @ty, align 4
  %268 = add nsw i32 %267, 2
  store i32 %268, i32* @ty, align 4
  br label %374

; <label>:269                                     ; preds = %249
  %270 = load i32* @tk, align 4
  %271 = icmp eq i32 %270, 33
  br i1 %271, label %272, label %281

; <label>:272                                     ; preds = %269
  call void @next()
  call void @expr(i32 162)
  %273 = load i32** @e, align 8
  %274 = getelementptr inbounds i32* %273, i32 1
  store i32* %274, i32** @e, align 8
  store i32 13, i32* %274, align 4
  %275 = load i32** @e, align 8
  %276 = getelementptr inbounds i32* %275, i32 1
  store i32* %276, i32** @e, align 8
  store i32 1, i32* %276, align 4
  %277 = load i32** @e, align 8
  %278 = getelementptr inbounds i32* %277, i32 1
  store i32* %278, i32** @e, align 8
  store i32 0, i32* %278, align 4
  %279 = load i32** @e, align 8
  %280 = getelementptr inbounds i32* %279, i32 1
  store i32* %280, i32** @e, align 8
  store i32 17, i32* %280, align 4
  store i32 1, i32* @ty, align 4
  br label %373

; <label>:281                                     ; preds = %269
  %282 = load i32* @tk, align 4
  %283 = icmp eq i32 %282, 126
  br i1 %283, label %284, label %293

; <label>:284                                     ; preds = %281
  call void @next()
  call void @expr(i32 162)
  %285 = load i32** @e, align 8
  %286 = getelementptr inbounds i32* %285, i32 1
  store i32* %286, i32** @e, align 8
  store i32 13, i32* %286, align 4
  %287 = load i32** @e, align 8
  %288 = getelementptr inbounds i32* %287, i32 1
  store i32* %288, i32** @e, align 8
  store i32 1, i32* %288, align 4
  %289 = load i32** @e, align 8
  %290 = getelementptr inbounds i32* %289, i32 1
  store i32* %290, i32** @e, align 8
  store i32 -1, i32* %290, align 4
  %291 = load i32** @e, align 8
  %292 = getelementptr inbounds i32* %291, i32 1
  store i32* %292, i32** @e, align 8
  store i32 15, i32* %292, align 4
  store i32 1, i32* @ty, align 4
  br label %372

; <label>:293                                     ; preds = %281
  %294 = load i32* @tk, align 4
  %295 = icmp eq i32 %294, 157
  br i1 %295, label %296, label %297

; <label>:296                                     ; preds = %293
  call void @next()
  call void @expr(i32 162)
  store i32 1, i32* @ty, align 4
  br label %371

; <label>:297                                     ; preds = %293
  %298 = load i32* @tk, align 4
  %299 = icmp eq i32 %298, 158
  br i1 %299, label %300, label %318

; <label>:300                                     ; preds = %297
  call void @next()
  %301 = load i32** @e, align 8
  %302 = getelementptr inbounds i32* %301, i32 1
  store i32* %302, i32** @e, align 8
  store i32 1, i32* %302, align 4
  %303 = load i32* @tk, align 4
  %304 = icmp eq i32 %303, 128
  br i1 %304, label %305, label %310

; <label>:305                                     ; preds = %300
  %306 = load i32* @ival, align 4
  %307 = sub nsw i32 0, %306
  %308 = load i32** @e, align 8
  %309 = getelementptr inbounds i32* %308, i32 1
  store i32* %309, i32** @e, align 8
  store i32 %307, i32* %309, align 4
  call void @next()
  br label %317

; <label>:310                                     ; preds = %300
  %311 = load i32** @e, align 8
  %312 = getelementptr inbounds i32* %311, i32 1
  store i32* %312, i32** @e, align 8
  store i32 -1, i32* %312, align 4
  %313 = load i32** @e, align 8
  %314 = getelementptr inbounds i32* %313, i32 1
  store i32* %314, i32** @e, align 8
  store i32 13, i32* %314, align 4
  call void @expr(i32 162)
  %315 = load i32** @e, align 8
  %316 = getelementptr inbounds i32* %315, i32 1
  store i32* %316, i32** @e, align 8
  store i32 27, i32* %316, align 4
  br label %317

; <label>:317                                     ; preds = %310, %305
  store i32 1, i32* @ty, align 4
  br label %370

; <label>:318                                     ; preds = %297
  %319 = load i32* @tk, align 4
  %320 = icmp eq i32 %319, 162
  br i1 %320, label %324, label %321

; <label>:321                                     ; preds = %318
  %322 = load i32* @tk, align 4
  %323 = icmp eq i32 %322, 163
  br i1 %323, label %324, label %366

; <label>:324                                     ; preds = %321, %318
  %325 = load i32* @tk, align 4
  store i32 %325, i32* %t, align 4
  call void @next()
  call void @expr(i32 162)
  %326 = load i32** @e, align 8
  %327 = load i32* %326, align 4
  %328 = icmp eq i32 %327, 10
  br i1 %328, label %329, label %333

; <label>:329                                     ; preds = %324
  %330 = load i32** @e, align 8
  store i32 13, i32* %330, align 4
  %331 = load i32** @e, align 8
  %332 = getelementptr inbounds i32* %331, i32 1
  store i32* %332, i32** @e, align 8
  store i32 10, i32* %332, align 4
  br label %345

; <label>:333                                     ; preds = %324
  %334 = load i32** @e, align 8
  %335 = load i32* %334, align 4
  %336 = icmp eq i32 %335, 9
  br i1 %336, label %337, label %341

; <label>:337                                     ; preds = %333
  %338 = load i32** @e, align 8
  store i32 13, i32* %338, align 4
  %339 = load i32** @e, align 8
  %340 = getelementptr inbounds i32* %339, i32 1
  store i32* %340, i32** @e, align 8
  store i32 9, i32* %340, align 4
  br label %344

; <label>:341                                     ; preds = %333
  %342 = load i32* @line, align 4
  %343 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([33 x i8]* @.str14, i32 0, i32 0), i32 %342)
  call void @exit(i32 -1) #7
  unreachable

; <label>:344                                     ; preds = %337
  br label %345

; <label>:345                                     ; preds = %344, %329
  %346 = load i32** @e, align 8
  %347 = getelementptr inbounds i32* %346, i32 1
  store i32* %347, i32** @e, align 8
  store i32 13, i32* %347, align 4
  %348 = load i32** @e, align 8
  %349 = getelementptr inbounds i32* %348, i32 1
  store i32* %349, i32** @e, align 8
  store i32 1, i32* %349, align 4
  %350 = load i32* @ty, align 4
  %351 = icmp sgt i32 %350, 2
  %352 = select i1 %351, i64 4, i64 1
  %353 = trunc i64 %352 to i32
  %354 = load i32** @e, align 8
  %355 = getelementptr inbounds i32* %354, i32 1
  store i32* %355, i32** @e, align 8
  store i32 %353, i32* %355, align 4
  %356 = load i32* %t, align 4
  %357 = icmp eq i32 %356, 162
  %358 = select i1 %357, i32 25, i32 26
  %359 = load i32** @e, align 8
  %360 = getelementptr inbounds i32* %359, i32 1
  store i32* %360, i32** @e, align 8
  store i32 %358, i32* %360, align 4
  %361 = load i32* @ty, align 4
  %362 = icmp eq i32 %361, 0
  %363 = select i1 %362, i32 12, i32 11
  %364 = load i32** @e, align 8
  %365 = getelementptr inbounds i32* %364, i32 1
  store i32* %365, i32** @e, align 8
  store i32 %363, i32* %365, align 4
  br label %369

; <label>:366                                     ; preds = %321
  %367 = load i32* @line, align 4
  %368 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([20 x i8]* @.str15, i32 0, i32 0), i32 %367)
  call void @exit(i32 -1) #7
  unreachable

; <label>:369                                     ; preds = %345
  br label %370

; <label>:370                                     ; preds = %369, %317
  br label %371

; <label>:371                                     ; preds = %370, %296
  br label %372

; <label>:372                                     ; preds = %371, %284
  br label %373

; <label>:373                                     ; preds = %372, %272
  br label %374

; <label>:374                                     ; preds = %373, %266
  br label %375

; <label>:375                                     ; preds = %374, %243
  br label %376

; <label>:376                                     ; preds = %375, %230
  br label %377

; <label>:377                                     ; preds = %376, %193
  br label %378

; <label>:378                                     ; preds = %377, %69
  br label %379

; <label>:379                                     ; preds = %378, %29
  br label %380

; <label>:380                                     ; preds = %379, %10
  br label %381

; <label>:381                                     ; preds = %380
  br label %382

; <label>:382                                     ; preds = %761, %381
  %383 = load i32* @tk, align 4
  %384 = load i32* %1, align 4
  %385 = icmp sge i32 %383, %384
  br i1 %385, label %386, label %762

; <label>:386                                     ; preds = %382
  %387 = load i32* @ty, align 4
  store i32 %387, i32* %t, align 4
  %388 = load i32* @tk, align 4
  %389 = icmp eq i32 %388, 142
  br i1 %389, label %390, label %409

; <label>:390                                     ; preds = %386
  call void @next()
  %391 = load i32** @e, align 8
  %392 = load i32* %391, align 4
  %393 = icmp eq i32 %392, 10
  br i1 %393, label %398, label %394

; <label>:394                                     ; preds = %390
  %395 = load i32** @e, align 8
  %396 = load i32* %395, align 4
  %397 = icmp eq i32 %396, 9
  br i1 %397, label %398, label %400

; <label>:398                                     ; preds = %394, %390
  %399 = load i32** @e, align 8
  store i32 13, i32* %399, align 4
  br label %403

; <label>:400                                     ; preds = %394
  %401 = load i32* @line, align 4
  %402 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([30 x i8]* @.str16, i32 0, i32 0), i32 %401)
  call void @exit(i32 -1) #7
  unreachable

; <label>:403                                     ; preds = %398
  call void @expr(i32 142)
  %404 = load i32* %t, align 4
  store i32 %404, i32* @ty, align 4
  %405 = icmp eq i32 %404, 0
  %406 = select i1 %405, i32 12, i32 11
  %407 = load i32** @e, align 8
  %408 = getelementptr inbounds i32* %407, i32 1
  store i32* %408, i32** @e, align 8
  store i32 %406, i32* %408, align 4
  br label %761

; <label>:409                                     ; preds = %386
  %410 = load i32* @tk, align 4
  %411 = icmp eq i32 %410, 143
  br i1 %411, label %412, label %436

; <label>:412                                     ; preds = %409
  call void @next()
  %413 = load i32** @e, align 8
  %414 = getelementptr inbounds i32* %413, i32 1
  store i32* %414, i32** @e, align 8
  store i32 4, i32* %414, align 4
  %415 = load i32** @e, align 8
  %416 = getelementptr inbounds i32* %415, i32 1
  store i32* %416, i32** @e, align 8
  store i32* %416, i32** %d, align 8
  call void @expr(i32 142)
  %417 = load i32* @tk, align 4
  %418 = icmp eq i32 %417, 58
  br i1 %418, label %419, label %420

; <label>:419                                     ; preds = %412
  call void @next()
  br label %423

; <label>:420                                     ; preds = %412
  %421 = load i32* @line, align 4
  %422 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([31 x i8]* @.str17, i32 0, i32 0), i32 %421)
  call void @exit(i32 -1) #7
  unreachable

; <label>:423                                     ; preds = %419
  %424 = load i32** @e, align 8
  %425 = getelementptr inbounds i32* %424, i64 3
  %426 = ptrtoint i32* %425 to i32
  %427 = load i32** %d, align 8
  store i32 %426, i32* %427, align 4
  %428 = load i32** @e, align 8
  %429 = getelementptr inbounds i32* %428, i32 1
  store i32* %429, i32** @e, align 8
  store i32 2, i32* %429, align 4
  %430 = load i32** @e, align 8
  %431 = getelementptr inbounds i32* %430, i32 1
  store i32* %431, i32** @e, align 8
  store i32* %431, i32** %d, align 8
  call void @expr(i32 143)
  %432 = load i32** @e, align 8
  %433 = getelementptr inbounds i32* %432, i64 1
  %434 = ptrtoint i32* %433 to i32
  %435 = load i32** %d, align 8
  store i32 %434, i32* %435, align 4
  br label %760

; <label>:436                                     ; preds = %409
  %437 = load i32* @tk, align 4
  %438 = icmp eq i32 %437, 144
  br i1 %438, label %439, label %448

; <label>:439                                     ; preds = %436
  call void @next()
  %440 = load i32** @e, align 8
  %441 = getelementptr inbounds i32* %440, i32 1
  store i32* %441, i32** @e, align 8
  store i32 5, i32* %441, align 4
  %442 = load i32** @e, align 8
  %443 = getelementptr inbounds i32* %442, i32 1
  store i32* %443, i32** @e, align 8
  store i32* %443, i32** %d, align 8
  call void @expr(i32 145)
  %444 = load i32** @e, align 8
  %445 = getelementptr inbounds i32* %444, i64 1
  %446 = ptrtoint i32* %445 to i32
  %447 = load i32** %d, align 8
  store i32 %446, i32* %447, align 4
  store i32 1, i32* @ty, align 4
  br label %759

; <label>:448                                     ; preds = %436
  %449 = load i32* @tk, align 4
  %450 = icmp eq i32 %449, 145
  br i1 %450, label %451, label %460

; <label>:451                                     ; preds = %448
  call void @next()
  %452 = load i32** @e, align 8
  %453 = getelementptr inbounds i32* %452, i32 1
  store i32* %453, i32** @e, align 8
  store i32 4, i32* %453, align 4
  %454 = load i32** @e, align 8
  %455 = getelementptr inbounds i32* %454, i32 1
  store i32* %455, i32** @e, align 8
  store i32* %455, i32** %d, align 8
  call void @expr(i32 146)
  %456 = load i32** @e, align 8
  %457 = getelementptr inbounds i32* %456, i64 1
  %458 = ptrtoint i32* %457 to i32
  %459 = load i32** %d, align 8
  store i32 %458, i32* %459, align 4
  store i32 1, i32* @ty, align 4
  br label %758

; <label>:460                                     ; preds = %448
  %461 = load i32* @tk, align 4
  %462 = icmp eq i32 %461, 146
  br i1 %462, label %463, label %468

; <label>:463                                     ; preds = %460
  call void @next()
  %464 = load i32** @e, align 8
  %465 = getelementptr inbounds i32* %464, i32 1
  store i32* %465, i32** @e, align 8
  store i32 13, i32* %465, align 4
  call void @expr(i32 147)
  %466 = load i32** @e, align 8
  %467 = getelementptr inbounds i32* %466, i32 1
  store i32* %467, i32** @e, align 8
  store i32 14, i32* %467, align 4
  store i32 1, i32* @ty, align 4
  br label %757

; <label>:468                                     ; preds = %460
  %469 = load i32* @tk, align 4
  %470 = icmp eq i32 %469, 147
  br i1 %470, label %471, label %476

; <label>:471                                     ; preds = %468
  call void @next()
  %472 = load i32** @e, align 8
  %473 = getelementptr inbounds i32* %472, i32 1
  store i32* %473, i32** @e, align 8
  store i32 13, i32* %473, align 4
  call void @expr(i32 148)
  %474 = load i32** @e, align 8
  %475 = getelementptr inbounds i32* %474, i32 1
  store i32* %475, i32** @e, align 8
  store i32 15, i32* %475, align 4
  store i32 1, i32* @ty, align 4
  br label %756

; <label>:476                                     ; preds = %468
  %477 = load i32* @tk, align 4
  %478 = icmp eq i32 %477, 148
  br i1 %478, label %479, label %484

; <label>:479                                     ; preds = %476
  call void @next()
  %480 = load i32** @e, align 8
  %481 = getelementptr inbounds i32* %480, i32 1
  store i32* %481, i32** @e, align 8
  store i32 13, i32* %481, align 4
  call void @expr(i32 149)
  %482 = load i32** @e, align 8
  %483 = getelementptr inbounds i32* %482, i32 1
  store i32* %483, i32** @e, align 8
  store i32 16, i32* %483, align 4
  store i32 1, i32* @ty, align 4
  br label %755

; <label>:484                                     ; preds = %476
  %485 = load i32* @tk, align 4
  %486 = icmp eq i32 %485, 149
  br i1 %486, label %487, label %492

; <label>:487                                     ; preds = %484
  call void @next()
  %488 = load i32** @e, align 8
  %489 = getelementptr inbounds i32* %488, i32 1
  store i32* %489, i32** @e, align 8
  store i32 13, i32* %489, align 4
  call void @expr(i32 151)
  %490 = load i32** @e, align 8
  %491 = getelementptr inbounds i32* %490, i32 1
  store i32* %491, i32** @e, align 8
  store i32 17, i32* %491, align 4
  store i32 1, i32* @ty, align 4
  br label %754

; <label>:492                                     ; preds = %484
  %493 = load i32* @tk, align 4
  %494 = icmp eq i32 %493, 150
  br i1 %494, label %495, label %500

; <label>:495                                     ; preds = %492
  call void @next()
  %496 = load i32** @e, align 8
  %497 = getelementptr inbounds i32* %496, i32 1
  store i32* %497, i32** @e, align 8
  store i32 13, i32* %497, align 4
  call void @expr(i32 151)
  %498 = load i32** @e, align 8
  %499 = getelementptr inbounds i32* %498, i32 1
  store i32* %499, i32** @e, align 8
  store i32 18, i32* %499, align 4
  store i32 1, i32* @ty, align 4
  br label %753

; <label>:500                                     ; preds = %492
  %501 = load i32* @tk, align 4
  %502 = icmp eq i32 %501, 151
  br i1 %502, label %503, label %508

; <label>:503                                     ; preds = %500
  call void @next()
  %504 = load i32** @e, align 8
  %505 = getelementptr inbounds i32* %504, i32 1
  store i32* %505, i32** @e, align 8
  store i32 13, i32* %505, align 4
  call void @expr(i32 155)
  %506 = load i32** @e, align 8
  %507 = getelementptr inbounds i32* %506, i32 1
  store i32* %507, i32** @e, align 8
  store i32 19, i32* %507, align 4
  store i32 1, i32* @ty, align 4
  br label %752

; <label>:508                                     ; preds = %500
  %509 = load i32* @tk, align 4
  %510 = icmp eq i32 %509, 152
  br i1 %510, label %511, label %516

; <label>:511                                     ; preds = %508
  call void @next()
  %512 = load i32** @e, align 8
  %513 = getelementptr inbounds i32* %512, i32 1
  store i32* %513, i32** @e, align 8
  store i32 13, i32* %513, align 4
  call void @expr(i32 155)
  %514 = load i32** @e, align 8
  %515 = getelementptr inbounds i32* %514, i32 1
  store i32* %515, i32** @e, align 8
  store i32 20, i32* %515, align 4
  store i32 1, i32* @ty, align 4
  br label %751

; <label>:516                                     ; preds = %508
  %517 = load i32* @tk, align 4
  %518 = icmp eq i32 %517, 153
  br i1 %518, label %519, label %524

; <label>:519                                     ; preds = %516
  call void @next()
  %520 = load i32** @e, align 8
  %521 = getelementptr inbounds i32* %520, i32 1
  store i32* %521, i32** @e, align 8
  store i32 13, i32* %521, align 4
  call void @expr(i32 155)
  %522 = load i32** @e, align 8
  %523 = getelementptr inbounds i32* %522, i32 1
  store i32* %523, i32** @e, align 8
  store i32 21, i32* %523, align 4
  store i32 1, i32* @ty, align 4
  br label %750

; <label>:524                                     ; preds = %516
  %525 = load i32* @tk, align 4
  %526 = icmp eq i32 %525, 154
  br i1 %526, label %527, label %532

; <label>:527                                     ; preds = %524
  call void @next()
  %528 = load i32** @e, align 8
  %529 = getelementptr inbounds i32* %528, i32 1
  store i32* %529, i32** @e, align 8
  store i32 13, i32* %529, align 4
  call void @expr(i32 155)
  %530 = load i32** @e, align 8
  %531 = getelementptr inbounds i32* %530, i32 1
  store i32* %531, i32** @e, align 8
  store i32 22, i32* %531, align 4
  store i32 1, i32* @ty, align 4
  br label %749

; <label>:532                                     ; preds = %524
  %533 = load i32* @tk, align 4
  %534 = icmp eq i32 %533, 155
  br i1 %534, label %535, label %540

; <label>:535                                     ; preds = %532
  call void @next()
  %536 = load i32** @e, align 8
  %537 = getelementptr inbounds i32* %536, i32 1
  store i32* %537, i32** @e, align 8
  store i32 13, i32* %537, align 4
  call void @expr(i32 157)
  %538 = load i32** @e, align 8
  %539 = getelementptr inbounds i32* %538, i32 1
  store i32* %539, i32** @e, align 8
  store i32 23, i32* %539, align 4
  store i32 1, i32* @ty, align 4
  br label %748

; <label>:540                                     ; preds = %532
  %541 = load i32* @tk, align 4
  %542 = icmp eq i32 %541, 156
  br i1 %542, label %543, label %548

; <label>:543                                     ; preds = %540
  call void @next()
  %544 = load i32** @e, align 8
  %545 = getelementptr inbounds i32* %544, i32 1
  store i32* %545, i32** @e, align 8
  store i32 13, i32* %545, align 4
  call void @expr(i32 157)
  %546 = load i32** @e, align 8
  %547 = getelementptr inbounds i32* %546, i32 1
  store i32* %547, i32** @e, align 8
  store i32 24, i32* %547, align 4
  store i32 1, i32* @ty, align 4
  br label %747

; <label>:548                                     ; preds = %540
  %549 = load i32* @tk, align 4
  %550 = icmp eq i32 %549, 157
  br i1 %550, label %551, label %568

; <label>:551                                     ; preds = %548
  call void @next()
  %552 = load i32** @e, align 8
  %553 = getelementptr inbounds i32* %552, i32 1
  store i32* %553, i32** @e, align 8
  store i32 13, i32* %553, align 4
  call void @expr(i32 159)
  %554 = load i32* %t, align 4
  store i32 %554, i32* @ty, align 4
  %555 = icmp sgt i32 %554, 2
  br i1 %555, label %556, label %565

; <label>:556                                     ; preds = %551
  %557 = load i32** @e, align 8
  %558 = getelementptr inbounds i32* %557, i32 1
  store i32* %558, i32** @e, align 8
  store i32 13, i32* %558, align 4
  %559 = load i32** @e, align 8
  %560 = getelementptr inbounds i32* %559, i32 1
  store i32* %560, i32** @e, align 8
  store i32 1, i32* %560, align 4
  %561 = load i32** @e, align 8
  %562 = getelementptr inbounds i32* %561, i32 1
  store i32* %562, i32** @e, align 8
  store i32 4, i32* %562, align 4
  %563 = load i32** @e, align 8
  %564 = getelementptr inbounds i32* %563, i32 1
  store i32* %564, i32** @e, align 8
  store i32 27, i32* %564, align 4
  br label %565

; <label>:565                                     ; preds = %556, %551
  %566 = load i32** @e, align 8
  %567 = getelementptr inbounds i32* %566, i32 1
  store i32* %567, i32** @e, align 8
  store i32 25, i32* %567, align 4
  br label %746

; <label>:568                                     ; preds = %548
  %569 = load i32* @tk, align 4
  %570 = icmp eq i32 %569, 158
  br i1 %570, label %571, label %610

; <label>:571                                     ; preds = %568
  call void @next()
  %572 = load i32** @e, align 8
  %573 = getelementptr inbounds i32* %572, i32 1
  store i32* %573, i32** @e, align 8
  store i32 13, i32* %573, align 4
  call void @expr(i32 159)
  %574 = load i32* %t, align 4
  %575 = icmp sgt i32 %574, 2
  br i1 %575, label %576, label %591

; <label>:576                                     ; preds = %571
  %577 = load i32* %t, align 4
  %578 = load i32* @ty, align 4
  %579 = icmp eq i32 %577, %578
  br i1 %579, label %580, label %591

; <label>:580                                     ; preds = %576
  %581 = load i32** @e, align 8
  %582 = getelementptr inbounds i32* %581, i32 1
  store i32* %582, i32** @e, align 8
  store i32 26, i32* %582, align 4
  %583 = load i32** @e, align 8
  %584 = getelementptr inbounds i32* %583, i32 1
  store i32* %584, i32** @e, align 8
  store i32 13, i32* %584, align 4
  %585 = load i32** @e, align 8
  %586 = getelementptr inbounds i32* %585, i32 1
  store i32* %586, i32** @e, align 8
  store i32 1, i32* %586, align 4
  %587 = load i32** @e, align 8
  %588 = getelementptr inbounds i32* %587, i32 1
  store i32* %588, i32** @e, align 8
  store i32 4, i32* %588, align 4
  %589 = load i32** @e, align 8
  %590 = getelementptr inbounds i32* %589, i32 1
  store i32* %590, i32** @e, align 8
  store i32 28, i32* %590, align 4
  store i32 1, i32* @ty, align 4
  br label %609

; <label>:591                                     ; preds = %576, %571
  %592 = load i32* %t, align 4
  store i32 %592, i32* @ty, align 4
  %593 = icmp sgt i32 %592, 2
  br i1 %593, label %594, label %605

; <label>:594                                     ; preds = %591
  %595 = load i32** @e, align 8
  %596 = getelementptr inbounds i32* %595, i32 1
  store i32* %596, i32** @e, align 8
  store i32 13, i32* %596, align 4
  %597 = load i32** @e, align 8
  %598 = getelementptr inbounds i32* %597, i32 1
  store i32* %598, i32** @e, align 8
  store i32 1, i32* %598, align 4
  %599 = load i32** @e, align 8
  %600 = getelementptr inbounds i32* %599, i32 1
  store i32* %600, i32** @e, align 8
  store i32 4, i32* %600, align 4
  %601 = load i32** @e, align 8
  %602 = getelementptr inbounds i32* %601, i32 1
  store i32* %602, i32** @e, align 8
  store i32 27, i32* %602, align 4
  %603 = load i32** @e, align 8
  %604 = getelementptr inbounds i32* %603, i32 1
  store i32* %604, i32** @e, align 8
  store i32 26, i32* %604, align 4
  br label %608

; <label>:605                                     ; preds = %591
  %606 = load i32** @e, align 8
  %607 = getelementptr inbounds i32* %606, i32 1
  store i32* %607, i32** @e, align 8
  store i32 26, i32* %607, align 4
  br label %608

; <label>:608                                     ; preds = %605, %594
  br label %609

; <label>:609                                     ; preds = %608, %580
  br label %745

; <label>:610                                     ; preds = %568
  %611 = load i32* @tk, align 4
  %612 = icmp eq i32 %611, 159
  br i1 %612, label %613, label %618

; <label>:613                                     ; preds = %610
  call void @next()
  %614 = load i32** @e, align 8
  %615 = getelementptr inbounds i32* %614, i32 1
  store i32* %615, i32** @e, align 8
  store i32 13, i32* %615, align 4
  call void @expr(i32 162)
  %616 = load i32** @e, align 8
  %617 = getelementptr inbounds i32* %616, i32 1
  store i32* %617, i32** @e, align 8
  store i32 27, i32* %617, align 4
  store i32 1, i32* @ty, align 4
  br label %744

; <label>:618                                     ; preds = %610
  %619 = load i32* @tk, align 4
  %620 = icmp eq i32 %619, 160
  br i1 %620, label %621, label %626

; <label>:621                                     ; preds = %618
  call void @next()
  %622 = load i32** @e, align 8
  %623 = getelementptr inbounds i32* %622, i32 1
  store i32* %623, i32** @e, align 8
  store i32 13, i32* %623, align 4
  call void @expr(i32 162)
  %624 = load i32** @e, align 8
  %625 = getelementptr inbounds i32* %624, i32 1
  store i32* %625, i32** @e, align 8
  store i32 28, i32* %625, align 4
  store i32 1, i32* @ty, align 4
  br label %743

; <label>:626                                     ; preds = %618
  %627 = load i32* @tk, align 4
  %628 = icmp eq i32 %627, 161
  br i1 %628, label %629, label %634

; <label>:629                                     ; preds = %626
  call void @next()
  %630 = load i32** @e, align 8
  %631 = getelementptr inbounds i32* %630, i32 1
  store i32* %631, i32** @e, align 8
  store i32 13, i32* %631, align 4
  call void @expr(i32 162)
  %632 = load i32** @e, align 8
  %633 = getelementptr inbounds i32* %632, i32 1
  store i32* %633, i32** @e, align 8
  store i32 29, i32* %633, align 4
  store i32 1, i32* @ty, align 4
  br label %742

; <label>:634                                     ; preds = %626
  %635 = load i32* @tk, align 4
  %636 = icmp eq i32 %635, 162
  br i1 %636, label %640, label %637

; <label>:637                                     ; preds = %634
  %638 = load i32* @tk, align 4
  %639 = icmp eq i32 %638, 163
  br i1 %639, label %640, label %696

; <label>:640                                     ; preds = %637, %634
  %641 = load i32** @e, align 8
  %642 = load i32* %641, align 4
  %643 = icmp eq i32 %642, 10
  br i1 %643, label %644, label %648

; <label>:644                                     ; preds = %640
  %645 = load i32** @e, align 8
  store i32 13, i32* %645, align 4
  %646 = load i32** @e, align 8
  %647 = getelementptr inbounds i32* %646, i32 1
  store i32* %647, i32** @e, align 8
  store i32 10, i32* %647, align 4
  br label %660

; <label>:648                                     ; preds = %640
  %649 = load i32** @e, align 8
  %650 = load i32* %649, align 4
  %651 = icmp eq i32 %650, 9
  br i1 %651, label %652, label %656

; <label>:652                                     ; preds = %648
  %653 = load i32** @e, align 8
  store i32 13, i32* %653, align 4
  %654 = load i32** @e, align 8
  %655 = getelementptr inbounds i32* %654, i32 1
  store i32* %655, i32** @e, align 8
  store i32 9, i32* %655, align 4
  br label %659

; <label>:656                                     ; preds = %648
  %657 = load i32* @line, align 4
  %658 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([34 x i8]* @.str18, i32 0, i32 0), i32 %657)
  call void @exit(i32 -1) #7
  unreachable

; <label>:659                                     ; preds = %652
  br label %660

; <label>:660                                     ; preds = %659, %644
  %661 = load i32** @e, align 8
  %662 = getelementptr inbounds i32* %661, i32 1
  store i32* %662, i32** @e, align 8
  store i32 13, i32* %662, align 4
  %663 = load i32** @e, align 8
  %664 = getelementptr inbounds i32* %663, i32 1
  store i32* %664, i32** @e, align 8
  store i32 1, i32* %664, align 4
  %665 = load i32* @ty, align 4
  %666 = icmp sgt i32 %665, 2
  %667 = select i1 %666, i64 4, i64 1
  %668 = trunc i64 %667 to i32
  %669 = load i32** @e, align 8
  %670 = getelementptr inbounds i32* %669, i32 1
  store i32* %670, i32** @e, align 8
  store i32 %668, i32* %670, align 4
  %671 = load i32* @tk, align 4
  %672 = icmp eq i32 %671, 162
  %673 = select i1 %672, i32 25, i32 26
  %674 = load i32** @e, align 8
  %675 = getelementptr inbounds i32* %674, i32 1
  store i32* %675, i32** @e, align 8
  store i32 %673, i32* %675, align 4
  %676 = load i32* @ty, align 4
  %677 = icmp eq i32 %676, 0
  %678 = select i1 %677, i32 12, i32 11
  %679 = load i32** @e, align 8
  %680 = getelementptr inbounds i32* %679, i32 1
  store i32* %680, i32** @e, align 8
  store i32 %678, i32* %680, align 4
  %681 = load i32** @e, align 8
  %682 = getelementptr inbounds i32* %681, i32 1
  store i32* %682, i32** @e, align 8
  store i32 13, i32* %682, align 4
  %683 = load i32** @e, align 8
  %684 = getelementptr inbounds i32* %683, i32 1
  store i32* %684, i32** @e, align 8
  store i32 1, i32* %684, align 4
  %685 = load i32* @ty, align 4
  %686 = icmp sgt i32 %685, 2
  %687 = select i1 %686, i64 4, i64 1
  %688 = trunc i64 %687 to i32
  %689 = load i32** @e, align 8
  %690 = getelementptr inbounds i32* %689, i32 1
  store i32* %690, i32** @e, align 8
  store i32 %688, i32* %690, align 4
  %691 = load i32* @tk, align 4
  %692 = icmp eq i32 %691, 162
  %693 = select i1 %692, i32 26, i32 25
  %694 = load i32** @e, align 8
  %695 = getelementptr inbounds i32* %694, i32 1
  store i32* %695, i32** @e, align 8
  store i32 %693, i32* %695, align 4
  call void @next()
  br label %741

; <label>:696                                     ; preds = %637
  %697 = load i32* @tk, align 4
  %698 = icmp eq i32 %697, 164
  br i1 %698, label %699, label %736

; <label>:699                                     ; preds = %696
  call void @next()
  %700 = load i32** @e, align 8
  %701 = getelementptr inbounds i32* %700, i32 1
  store i32* %701, i32** @e, align 8
  store i32 13, i32* %701, align 4
  call void @expr(i32 142)
  %702 = load i32* @tk, align 4
  %703 = icmp eq i32 %702, 93
  br i1 %703, label %704, label %705

; <label>:704                                     ; preds = %699
  call void @next()
  br label %708

; <label>:705                                     ; preds = %699
  %706 = load i32* @line, align 4
  %707 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([28 x i8]* @.str19, i32 0, i32 0), i32 %706)
  call void @exit(i32 -1) #7
  unreachable

; <label>:708                                     ; preds = %704
  %709 = load i32* %t, align 4
  %710 = icmp sgt i32 %709, 2
  br i1 %710, label %711, label %720

; <label>:711                                     ; preds = %708
  %712 = load i32** @e, align 8
  %713 = getelementptr inbounds i32* %712, i32 1
  store i32* %713, i32** @e, align 8
  store i32 13, i32* %713, align 4
  %714 = load i32** @e, align 8
  %715 = getelementptr inbounds i32* %714, i32 1
  store i32* %715, i32** @e, align 8
  store i32 1, i32* %715, align 4
  %716 = load i32** @e, align 8
  %717 = getelementptr inbounds i32* %716, i32 1
  store i32* %717, i32** @e, align 8
  store i32 4, i32* %717, align 4
  %718 = load i32** @e, align 8
  %719 = getelementptr inbounds i32* %718, i32 1
  store i32* %719, i32** @e, align 8
  store i32 27, i32* %719, align 4
  br label %727

; <label>:720                                     ; preds = %708
  %721 = load i32* %t, align 4
  %722 = icmp slt i32 %721, 2
  br i1 %722, label %723, label %726

; <label>:723                                     ; preds = %720
  %724 = load i32* @line, align 4
  %725 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([27 x i8]* @.str20, i32 0, i32 0), i32 %724)
  call void @exit(i32 -1) #7
  unreachable

; <label>:726                                     ; preds = %720
  br label %727

; <label>:727                                     ; preds = %726, %711
  %728 = load i32** @e, align 8
  %729 = getelementptr inbounds i32* %728, i32 1
  store i32* %729, i32** @e, align 8
  store i32 25, i32* %729, align 4
  %730 = load i32* %t, align 4
  %731 = sub nsw i32 %730, 2
  store i32 %731, i32* @ty, align 4
  %732 = icmp eq i32 %731, 0
  %733 = select i1 %732, i32 10, i32 9
  %734 = load i32** @e, align 8
  %735 = getelementptr inbounds i32* %734, i32 1
  store i32* %735, i32** @e, align 8
  store i32 %733, i32* %735, align 4
  br label %740

; <label>:736                                     ; preds = %696
  %737 = load i32* @line, align 4
  %738 = load i32* @tk, align 4
  %739 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([26 x i8]* @.str21, i32 0, i32 0), i32 %737, i32 %738)
  call void @exit(i32 -1) #7
  unreachable

; <label>:740                                     ; preds = %727
  br label %741

; <label>:741                                     ; preds = %740, %660
  br label %742

; <label>:742                                     ; preds = %741, %629
  br label %743

; <label>:743                                     ; preds = %742, %621
  br label %744

; <label>:744                                     ; preds = %743, %613
  br label %745

; <label>:745                                     ; preds = %744, %609
  br label %746

; <label>:746                                     ; preds = %745, %565
  br label %747

; <label>:747                                     ; preds = %746, %543
  br label %748

; <label>:748                                     ; preds = %747, %535
  br label %749

; <label>:749                                     ; preds = %748, %527
  br label %750

; <label>:750                                     ; preds = %749, %519
  br label %751

; <label>:751                                     ; preds = %750, %511
  br label %752

; <label>:752                                     ; preds = %751, %503
  br label %753

; <label>:753                                     ; preds = %752, %495
  br label %754

; <label>:754                                     ; preds = %753, %487
  br label %755

; <label>:755                                     ; preds = %754, %479
  br label %756

; <label>:756                                     ; preds = %755, %471
  br label %757

; <label>:757                                     ; preds = %756, %463
  br label %758

; <label>:758                                     ; preds = %757, %451
  br label %759

; <label>:759                                     ; preds = %758, %439
  br label %760

; <label>:760                                     ; preds = %759, %423
  br label %761

; <label>:761                                     ; preds = %760, %403
  br label %382

; <label>:762                                     ; preds = %382
  ret void
}

; Function Attrs: noreturn nounwind
declare void @exit(i32) #3

; Function Attrs: nounwind uwtable
define void @stmt() #0 {
  %a = alloca i32*, align 8
  %b = alloca i32*, align 8
  %1 = load i32* @tk, align 4
  %2 = icmp eq i32 %1, 137
  br i1 %2, label %3, label %38

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
  store i32* %21, i32** %b, align 8
  call void @stmt()
  %22 = load i32* @tk, align 4
  %23 = icmp eq i32 %22, 135
  br i1 %23, label %24, label %33

; <label>:24                                      ; preds = %17
  %25 = load i32** @e, align 8
  %26 = getelementptr inbounds i32* %25, i64 3
  %27 = ptrtoint i32* %26 to i32
  %28 = load i32** %b, align 8
  store i32 %27, i32* %28, align 4
  %29 = load i32** @e, align 8
  %30 = getelementptr inbounds i32* %29, i32 1
  store i32* %30, i32** @e, align 8
  store i32 2, i32* %30, align 4
  %31 = load i32** @e, align 8
  %32 = getelementptr inbounds i32* %31, i32 1
  store i32* %32, i32** @e, align 8
  store i32* %32, i32** %b, align 8
  call void @next()
  call void @stmt()
  br label %33

; <label>:33                                      ; preds = %24, %17
  %34 = load i32** @e, align 8
  %35 = getelementptr inbounds i32* %34, i64 1
  %36 = ptrtoint i32* %35 to i32
  %37 = load i32** %b, align 8
  store i32 %36, i32* %37, align 4
  br label %114

; <label>:38                                      ; preds = %0
  %39 = load i32* @tk, align 4
  %40 = icmp eq i32 %39, 141
  br i1 %40, label %41, label %72

; <label>:41                                      ; preds = %38
  call void @next()
  %42 = load i32** @e, align 8
  %43 = getelementptr inbounds i32* %42, i64 1
  store i32* %43, i32** %a, align 8
  %44 = load i32* @tk, align 4
  %45 = icmp eq i32 %44, 40
  br i1 %45, label %46, label %47

; <label>:46                                      ; preds = %41
  call void @next()
  br label %50

; <label>:47                                      ; preds = %41
  %48 = load i32* @line, align 4
  %49 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([25 x i8]* @.str22, i32 0, i32 0), i32 %48)
  call void @exit(i32 -1) #7
  unreachable

; <label>:50                                      ; preds = %46
  call void @expr(i32 142)
  %51 = load i32* @tk, align 4
  %52 = icmp eq i32 %51, 41
  br i1 %52, label %53, label %54

; <label>:53                                      ; preds = %50
  call void @next()
  br label %57

; <label>:54                                      ; preds = %50
  %55 = load i32* @line, align 4
  %56 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([26 x i8]* @.str11, i32 0, i32 0), i32 %55)
  call void @exit(i32 -1) #7
  unreachable

; <label>:57                                      ; preds = %53
  %58 = load i32** @e, align 8
  %59 = getelementptr inbounds i32* %58, i32 1
  store i32* %59, i32** @e, align 8
  store i32 4, i32* %59, align 4
  %60 = load i32** @e, align 8
  %61 = getelementptr inbounds i32* %60, i32 1
  store i32* %61, i32** @e, align 8
  store i32* %61, i32** %b, align 8
  call void @stmt()
  %62 = load i32** @e, align 8
  %63 = getelementptr inbounds i32* %62, i32 1
  store i32* %63, i32** @e, align 8
  store i32 2, i32* %63, align 4
  %64 = load i32** %a, align 8
  %65 = ptrtoint i32* %64 to i32
  %66 = load i32** @e, align 8
  %67 = getelementptr inbounds i32* %66, i32 1
  store i32* %67, i32** @e, align 8
  store i32 %65, i32* %67, align 4
  %68 = load i32** @e, align 8
  %69 = getelementptr inbounds i32* %68, i64 1
  %70 = ptrtoint i32* %69 to i32
  %71 = load i32** %b, align 8
  store i32 %70, i32* %71, align 4
  br label %113

; <label>:72                                      ; preds = %38
  %73 = load i32* @tk, align 4
  %74 = icmp eq i32 %73, 139
  br i1 %74, label %75, label %89

; <label>:75                                      ; preds = %72
  call void @next()
  %76 = load i32* @tk, align 4
  %77 = icmp ne i32 %76, 59
  br i1 %77, label %78, label %79

; <label>:78                                      ; preds = %75
  call void @expr(i32 142)
  br label %79

; <label>:79                                      ; preds = %78, %75
  %80 = load i32** @e, align 8
  %81 = getelementptr inbounds i32* %80, i32 1
  store i32* %81, i32** @e, align 8
  store i32 8, i32* %81, align 4
  %82 = load i32* @tk, align 4
  %83 = icmp eq i32 %82, 59
  br i1 %83, label %84, label %85

; <label>:84                                      ; preds = %79
  call void @next()
  br label %88

; <label>:85                                      ; preds = %79
  %86 = load i32* @line, align 4
  %87 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([24 x i8]* @.str23, i32 0, i32 0), i32 %86)
  call void @exit(i32 -1) #7
  unreachable

; <label>:88                                      ; preds = %84
  br label %112

; <label>:89                                      ; preds = %72
  %90 = load i32* @tk, align 4
  %91 = icmp eq i32 %90, 123
  br i1 %91, label %92, label %98

; <label>:92                                      ; preds = %89
  call void @next()
  br label %93

; <label>:93                                      ; preds = %96, %92
  %94 = load i32* @tk, align 4
  %95 = icmp ne i32 %94, 125
  br i1 %95, label %96, label %97

; <label>:96                                      ; preds = %93
  call void @stmt()
  br label %93

; <label>:97                                      ; preds = %93
  call void @next()
  br label %111

; <label>:98                                      ; preds = %89
  %99 = load i32* @tk, align 4
  %100 = icmp eq i32 %99, 59
  br i1 %100, label %101, label %102

; <label>:101                                     ; preds = %98
  call void @next()
  br label %110

; <label>:102                                     ; preds = %98
  call void @expr(i32 142)
  %103 = load i32* @tk, align 4
  %104 = icmp eq i32 %103, 59
  br i1 %104, label %105, label %106

; <label>:105                                     ; preds = %102
  call void @next()
  br label %109

; <label>:106                                     ; preds = %102
  %107 = load i32* @line, align 4
  %108 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([24 x i8]* @.str23, i32 0, i32 0), i32 %107)
  call void @exit(i32 -1) #7
  unreachable

; <label>:109                                     ; preds = %105
  br label %110

; <label>:110                                     ; preds = %109, %101
  br label %111

; <label>:111                                     ; preds = %110, %97
  br label %112

; <label>:112                                     ; preds = %111, %88
  br label %113

; <label>:113                                     ; preds = %112, %57
  br label %114

; <label>:114                                     ; preds = %113, %33
  ret void
}

; Function Attrs: nounwind uwtable
define i32 @main(i32 %argc, i8** %argv) #0 {
  %1 = alloca i32, align 4
  %2 = alloca i32, align 4
  %3 = alloca i8**, align 8
  %fd = alloca i32, align 4
  %bt = alloca i32, align 4
  %ty = alloca i32, align 4
  %poolsz = alloca i32, align 4
  %idmain = alloca i32*, align 8
  %pc = alloca i32*, align 8
  %sp = alloca i32*, align 8
  %bp = alloca i32*, align 8
  %a = alloca i32, align 4
  %cycle = alloca i32, align 4
  %i = alloca i32, align 4
  %t = alloca i32*, align 8
  store i32 0, i32* %1
  store i32 %argc, i32* %2, align 4
  store i8** %argv, i8*** %3, align 8
  %4 = load i32* %2, align 4
  %5 = add nsw i32 %4, -1
  store i32 %5, i32* %2, align 4
  %6 = load i8*** %3, align 8
  %7 = getelementptr inbounds i8** %6, i32 1
  store i8** %7, i8*** %3, align 8
  %8 = load i32* %2, align 4
  %9 = icmp sgt i32 %8, 0
  br i1 %9, label %10, label %28

; <label>:10                                      ; preds = %0
  %11 = load i8*** %3, align 8
  %12 = load i8** %11, align 8
  %13 = load i8* %12, align 1
  %14 = sext i8 %13 to i32
  %15 = icmp eq i32 %14, 45
  br i1 %15, label %16, label %28

; <label>:16                                      ; preds = %10
  %17 = load i8*** %3, align 8
  %18 = load i8** %17, align 8
  %19 = getelementptr inbounds i8* %18, i64 1
  %20 = load i8* %19, align 1
  %21 = sext i8 %20 to i32
  %22 = icmp eq i32 %21, 115
  br i1 %22, label %23, label %28

; <label>:23                                      ; preds = %16
  store i32 1, i32* @src, align 4
  %24 = load i32* %2, align 4
  %25 = add nsw i32 %24, -1
  store i32 %25, i32* %2, align 4
  %26 = load i8*** %3, align 8
  %27 = getelementptr inbounds i8** %26, i32 1
  store i8** %27, i8*** %3, align 8
  br label %28

; <label>:28                                      ; preds = %23, %16, %10, %0
  %29 = load i32* %2, align 4
  %30 = icmp sgt i32 %29, 0
  br i1 %30, label %31, label %49

; <label>:31                                      ; preds = %28
  %32 = load i8*** %3, align 8
  %33 = load i8** %32, align 8
  %34 = load i8* %33, align 1
  %35 = sext i8 %34 to i32
  %36 = icmp eq i32 %35, 45
  br i1 %36, label %37, label %49

; <label>:37                                      ; preds = %31
  %38 = load i8*** %3, align 8
  %39 = load i8** %38, align 8
  %40 = getelementptr inbounds i8* %39, i64 1
  %41 = load i8* %40, align 1
  %42 = sext i8 %41 to i32
  %43 = icmp eq i32 %42, 100
  br i1 %43, label %44, label %49

; <label>:44                                      ; preds = %37
  store i32 1, i32* @debug, align 4
  %45 = load i32* %2, align 4
  %46 = add nsw i32 %45, -1
  store i32 %46, i32* %2, align 4
  %47 = load i8*** %3, align 8
  %48 = getelementptr inbounds i8** %47, i32 1
  store i8** %48, i8*** %3, align 8
  br label %49

; <label>:49                                      ; preds = %44, %37, %31, %28
  %50 = load i32* %2, align 4
  %51 = icmp slt i32 %50, 1
  br i1 %51, label %52, label %54

; <label>:52                                      ; preds = %49
  %53 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([30 x i8]* @.str24, i32 0, i32 0))
  store i32 -1, i32* %1
  br label %995

; <label>:54                                      ; preds = %49
  %55 = load i8*** %3, align 8
  %56 = load i8** %55, align 8
  %57 = call i32 (i8*, i32, ...)* bitcast (i32 (...)* @open to i32 (i8*, i32, ...)*)(i8* %56, i32 0)
  store i32 %57, i32* %fd, align 4
  %58 = icmp slt i32 %57, 0
  br i1 %58, label %59, label %63

; <label>:59                                      ; preds = %54
  %60 = load i8*** %3, align 8
  %61 = load i8** %60, align 8
  %62 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([20 x i8]* @.str25, i32 0, i32 0), i8* %61)
  store i32 -1, i32* %1
  br label %995

; <label>:63                                      ; preds = %54
  store i32 262144, i32* %poolsz, align 4
  %64 = load i32* %poolsz, align 4
  %65 = sext i32 %64 to i64
  %66 = call noalias i8* @malloc(i64 %65) #5
  %67 = bitcast i8* %66 to i32*
  store i32* %67, i32** @sym, align 8
  %68 = icmp ne i32* %67, null
  br i1 %68, label %72, label %69

; <label>:69                                      ; preds = %63
  %70 = load i32* %poolsz, align 4
  %71 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([34 x i8]* @.str26, i32 0, i32 0), i32 %70)
  store i32 -1, i32* %1
  br label %995

; <label>:72                                      ; preds = %63
  %73 = load i32* %poolsz, align 4
  %74 = sext i32 %73 to i64
  %75 = call noalias i8* @malloc(i64 %74) #5
  %76 = bitcast i8* %75 to i32*
  store i32* %76, i32** @e, align 8
  store i32* %76, i32** @le, align 8
  %77 = icmp ne i32* %76, null
  br i1 %77, label %81, label %78

; <label>:78                                      ; preds = %72
  %79 = load i32* %poolsz, align 4
  %80 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([32 x i8]* @.str27, i32 0, i32 0), i32 %79)
  store i32 -1, i32* %1
  br label %995

; <label>:81                                      ; preds = %72
  %82 = load i32* %poolsz, align 4
  %83 = sext i32 %82 to i64
  %84 = call noalias i8* @malloc(i64 %83) #5
  store i8* %84, i8** @data, align 8
  %85 = icmp ne i8* %84, null
  br i1 %85, label %89, label %86

; <label>:86                                      ; preds = %81
  %87 = load i32* %poolsz, align 4
  %88 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([32 x i8]* @.str28, i32 0, i32 0), i32 %87)
  store i32 -1, i32* %1
  br label %995

; <label>:89                                      ; preds = %81
  %90 = load i32* %poolsz, align 4
  %91 = sext i32 %90 to i64
  %92 = call noalias i8* @malloc(i64 %91) #5
  %93 = bitcast i8* %92 to i32*
  store i32* %93, i32** %sp, align 8
  %94 = icmp ne i32* %93, null
  br i1 %94, label %98, label %95

; <label>:95                                      ; preds = %89
  %96 = load i32* %poolsz, align 4
  %97 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([33 x i8]* @.str29, i32 0, i32 0), i32 %96)
  store i32 -1, i32* %1
  br label %995

; <label>:98                                      ; preds = %89
  %99 = load i32** @sym, align 8
  %100 = bitcast i32* %99 to i8*
  %101 = load i32* %poolsz, align 4
  %102 = sext i32 %101 to i64
  call void @llvm.memset.p0i8.i64(i8* %100, i8 0, i64 %102, i32 4, i1 false)
  %103 = load i32** @e, align 8
  %104 = bitcast i32* %103 to i8*
  %105 = load i32* %poolsz, align 4
  %106 = sext i32 %105 to i64
  call void @llvm.memset.p0i8.i64(i8* %104, i8 0, i64 %106, i32 4, i1 false)
  %107 = load i8** @data, align 8
  %108 = load i32* %poolsz, align 4
  %109 = sext i32 %108 to i64
  call void @llvm.memset.p0i8.i64(i8* %107, i8 0, i64 %109, i32 1, i1 false)
  store i8* getelementptr inbounds ([101 x i8]* @.str30, i32 0, i32 0), i8** @p, align 8
  store i32 134, i32* %i, align 4
  br label %110

; <label>:110                                     ; preds = %113, %98
  %111 = load i32* %i, align 4
  %112 = icmp sle i32 %111, 141
  br i1 %112, label %113, label %118

; <label>:113                                     ; preds = %110
  call void @next()
  %114 = load i32* %i, align 4
  %115 = add nsw i32 %114, 1
  store i32 %115, i32* %i, align 4
  %116 = load i32** @id, align 8
  %117 = getelementptr inbounds i32* %116, i64 0
  store i32 %114, i32* %117, align 4
  br label %110

; <label>:118                                     ; preds = %110
  store i32 30, i32* %i, align 4
  br label %119

; <label>:119                                     ; preds = %122, %118
  %120 = load i32* %i, align 4
  %121 = icmp sle i32 %120, 37
  br i1 %121, label %122, label %131

; <label>:122                                     ; preds = %119
  call void @next()
  %123 = load i32** @id, align 8
  %124 = getelementptr inbounds i32* %123, i64 3
  store i32 130, i32* %124, align 4
  %125 = load i32** @id, align 8
  %126 = getelementptr inbounds i32* %125, i64 4
  store i32 1, i32* %126, align 4
  %127 = load i32* %i, align 4
  %128 = add nsw i32 %127, 1
  store i32 %128, i32* %i, align 4
  %129 = load i32** @id, align 8
  %130 = getelementptr inbounds i32* %129, i64 5
  store i32 %127, i32* %130, align 4
  br label %119

; <label>:131                                     ; preds = %119
  call void @next()
  %132 = load i32** @id, align 8
  %133 = getelementptr inbounds i32* %132, i64 0
  store i32 134, i32* %133, align 4
  call void @next()
  %134 = load i32** @id, align 8
  store i32* %134, i32** %idmain, align 8
  %135 = load i32* %poolsz, align 4
  %136 = sext i32 %135 to i64
  %137 = call noalias i8* @malloc(i64 %136) #5
  store i8* %137, i8** @p, align 8
  store i8* %137, i8** @lp, align 8
  %138 = icmp ne i8* %137, null
  br i1 %138, label %142, label %139

; <label>:139                                     ; preds = %131
  %140 = load i32* %poolsz, align 4
  %141 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([34 x i8]* @.str31, i32 0, i32 0), i32 %140)
  store i32 -1, i32* %1
  br label %995

; <label>:142                                     ; preds = %131
  %143 = load i32* %fd, align 4
  %144 = load i8** @p, align 8
  %145 = load i32* %poolsz, align 4
  %146 = sub nsw i32 %145, 1
  %147 = sext i32 %146 to i64
  %148 = call i64 @read(i32 %143, i8* %144, i64 %147)
  %149 = trunc i64 %148 to i32
  store i32 %149, i32* %i, align 4
  %150 = icmp sle i32 %149, 0
  br i1 %150, label %151, label %154

; <label>:151                                     ; preds = %142
  %152 = load i32* %i, align 4
  %153 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([20 x i8]* @.str32, i32 0, i32 0), i32 %152)
  store i32 -1, i32* %1
  br label %995

; <label>:154                                     ; preds = %142
  %155 = load i32* %i, align 4
  %156 = sext i32 %155 to i64
  %157 = load i8** @p, align 8
  %158 = getelementptr inbounds i8* %157, i64 %156
  store i8 0, i8* %158, align 1
  %159 = load i32* %fd, align 4
  %160 = call i32 @close(i32 %159)
  store i32 1, i32* @line, align 4
  call void @next()
  br label %161

; <label>:161                                     ; preds = %464, %154
  %162 = load i32* @tk, align 4
  %163 = icmp ne i32 %162, 0
  br i1 %163, label %164, label %465

; <label>:164                                     ; preds = %161
  store i32 1, i32* %bt, align 4
  %165 = load i32* @tk, align 4
  %166 = icmp eq i32 %165, 138
  br i1 %166, label %167, label %168

; <label>:167                                     ; preds = %164
  call void @next()
  br label %221

; <label>:168                                     ; preds = %164
  %169 = load i32* @tk, align 4
  %170 = icmp eq i32 %169, 134
  br i1 %170, label %171, label %172

; <label>:171                                     ; preds = %168
  call void @next()
  store i32 0, i32* %bt, align 4
  br label %220

; <label>:172                                     ; preds = %168
  %173 = load i32* @tk, align 4
  %174 = icmp eq i32 %173, 136
  br i1 %174, label %175, label %219

; <label>:175                                     ; preds = %172
  call void @next()
  %176 = load i32* @tk, align 4
  %177 = icmp ne i32 %176, 123
  br i1 %177, label %178, label %179

; <label>:178                                     ; preds = %175
  call void @next()
  br label %179

; <label>:179                                     ; preds = %178, %175
  %180 = load i32* @tk, align 4
  %181 = icmp eq i32 %180, 123
  br i1 %181, label %182, label %218

; <label>:182                                     ; preds = %179
  call void @next()
  store i32 0, i32* %i, align 4
  br label %183

; <label>:183                                     ; preds = %216, %182
  %184 = load i32* @tk, align 4
  %185 = icmp ne i32 %184, 125
  br i1 %185, label %186, label %217

; <label>:186                                     ; preds = %183
  %187 = load i32* @tk, align 4
  %188 = icmp ne i32 %187, 133
  br i1 %188, label %189, label %193

; <label>:189                                     ; preds = %186
  %190 = load i32* @line, align 4
  %191 = load i32* @tk, align 4
  %192 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([28 x i8]* @.str33, i32 0, i32 0), i32 %190, i32 %191)
  store i32 -1, i32* %1
  br label %995

; <label>:193                                     ; preds = %186
  call void @next()
  %194 = load i32* @tk, align 4
  %195 = icmp eq i32 %194, 142
  br i1 %195, label %196, label %204

; <label>:196                                     ; preds = %193
  call void @next()
  %197 = load i32* @tk, align 4
  %198 = icmp ne i32 %197, 128
  br i1 %198, label %199, label %202

; <label>:199                                     ; preds = %196
  %200 = load i32* @line, align 4
  %201 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([26 x i8]* @.str34, i32 0, i32 0), i32 %200)
  store i32 -1, i32* %1
  br label %995

; <label>:202                                     ; preds = %196
  %203 = load i32* @ival, align 4
  store i32 %203, i32* %i, align 4
  call void @next()
  br label %204

; <label>:204                                     ; preds = %202, %193
  %205 = load i32** @id, align 8
  %206 = getelementptr inbounds i32* %205, i64 3
  store i32 128, i32* %206, align 4
  %207 = load i32** @id, align 8
  %208 = getelementptr inbounds i32* %207, i64 4
  store i32 1, i32* %208, align 4
  %209 = load i32* %i, align 4
  %210 = add nsw i32 %209, 1
  store i32 %210, i32* %i, align 4
  %211 = load i32** @id, align 8
  %212 = getelementptr inbounds i32* %211, i64 5
  store i32 %209, i32* %212, align 4
  %213 = load i32* @tk, align 4
  %214 = icmp eq i32 %213, 44
  br i1 %214, label %215, label %216

; <label>:215                                     ; preds = %204
  call void @next()
  br label %216

; <label>:216                                     ; preds = %215, %204
  br label %183

; <label>:217                                     ; preds = %183
  call void @next()
  br label %218

; <label>:218                                     ; preds = %217, %179
  br label %219

; <label>:219                                     ; preds = %218, %172
  br label %220

; <label>:220                                     ; preds = %219, %171
  br label %221

; <label>:221                                     ; preds = %220, %167
  br label %222

; <label>:222                                     ; preds = %463, %221
  %223 = load i32* @tk, align 4
  %224 = icmp ne i32 %223, 59
  br i1 %224, label %225, label %228

; <label>:225                                     ; preds = %222
  %226 = load i32* @tk, align 4
  %227 = icmp ne i32 %226, 125
  br label %228

; <label>:228                                     ; preds = %225, %222
  %229 = phi i1 [ false, %222 ], [ %227, %225 ]
  br i1 %229, label %230, label %464

; <label>:230                                     ; preds = %228
  %231 = load i32* %bt, align 4
  store i32 %231, i32* %ty, align 4
  br label %232

; <label>:232                                     ; preds = %235, %230
  %233 = load i32* @tk, align 4
  %234 = icmp eq i32 %233, 159
  br i1 %234, label %235, label %238

; <label>:235                                     ; preds = %232
  call void @next()
  %236 = load i32* %ty, align 4
  %237 = add nsw i32 %236, 2
  store i32 %237, i32* %ty, align 4
  br label %232

; <label>:238                                     ; preds = %232
  %239 = load i32* @tk, align 4
  %240 = icmp ne i32 %239, 133
  br i1 %240, label %241, label %244

; <label>:241                                     ; preds = %238
  %242 = load i32* @line, align 4
  %243 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([28 x i8]* @.str35, i32 0, i32 0), i32 %242)
  store i32 -1, i32* %1
  br label %995

; <label>:244                                     ; preds = %238
  %245 = load i32** @id, align 8
  %246 = getelementptr inbounds i32* %245, i64 3
  %247 = load i32* %246, align 4
  %248 = icmp ne i32 %247, 0
  br i1 %248, label %249, label %252

; <label>:249                                     ; preds = %244
  %250 = load i32* @line, align 4
  %251 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([33 x i8]* @.str36, i32 0, i32 0), i32 %250)
  store i32 -1, i32* %1
  br label %995

; <label>:252                                     ; preds = %244
  call void @next()
  %253 = load i32* %ty, align 4
  %254 = load i32** @id, align 8
  %255 = getelementptr inbounds i32* %254, i64 4
  store i32 %253, i32* %255, align 4
  %256 = load i32* @tk, align 4
  %257 = icmp eq i32 %256, 40
  br i1 %257, label %258, label %450

; <label>:258                                     ; preds = %252
  %259 = load i32** @id, align 8
  %260 = getelementptr inbounds i32* %259, i64 3
  store i32 129, i32* %260, align 4
  %261 = load i32** @e, align 8
  %262 = getelementptr inbounds i32* %261, i64 1
  %263 = ptrtoint i32* %262 to i32
  %264 = load i32** @id, align 8
  %265 = getelementptr inbounds i32* %264, i64 5
  store i32 %263, i32* %265, align 4
  call void @next()
  store i32 0, i32* %i, align 4
  br label %266

; <label>:266                                     ; preds = %327, %258
  %267 = load i32* @tk, align 4
  %268 = icmp ne i32 %267, 41
  br i1 %268, label %269, label %328

; <label>:269                                     ; preds = %266
  store i32 1, i32* %ty, align 4
  %270 = load i32* @tk, align 4
  %271 = icmp eq i32 %270, 138
  br i1 %271, label %272, label %273

; <label>:272                                     ; preds = %269
  call void @next()
  br label %278

; <label>:273                                     ; preds = %269
  %274 = load i32* @tk, align 4
  %275 = icmp eq i32 %274, 134
  br i1 %275, label %276, label %277

; <label>:276                                     ; preds = %273
  call void @next()
  store i32 0, i32* %ty, align 4
  br label %277

; <label>:277                                     ; preds = %276, %273
  br label %278

; <label>:278                                     ; preds = %277, %272
  br label %279

; <label>:279                                     ; preds = %282, %278
  %280 = load i32* @tk, align 4
  %281 = icmp eq i32 %280, 159
  br i1 %281, label %282, label %285

; <label>:282                                     ; preds = %279
  call void @next()
  %283 = load i32* %ty, align 4
  %284 = add nsw i32 %283, 2
  store i32 %284, i32* %ty, align 4
  br label %279

; <label>:285                                     ; preds = %279
  %286 = load i32* @tk, align 4
  %287 = icmp ne i32 %286, 133
  br i1 %287, label %288, label %291

; <label>:288                                     ; preds = %285
  %289 = load i32* @line, align 4
  %290 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([31 x i8]* @.str37, i32 0, i32 0), i32 %289)
  store i32 -1, i32* %1
  br label %995

; <label>:291                                     ; preds = %285
  %292 = load i32** @id, align 8
  %293 = getelementptr inbounds i32* %292, i64 3
  %294 = load i32* %293, align 4
  %295 = icmp eq i32 %294, 132
  br i1 %295, label %296, label %299

; <label>:296                                     ; preds = %291
  %297 = load i32* @line, align 4
  %298 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([36 x i8]* @.str38, i32 0, i32 0), i32 %297)
  store i32 -1, i32* %1
  br label %995

; <label>:299                                     ; preds = %291
  %300 = load i32** @id, align 8
  %301 = getelementptr inbounds i32* %300, i64 3
  %302 = load i32* %301, align 4
  %303 = load i32** @id, align 8
  %304 = getelementptr inbounds i32* %303, i64 6
  store i32 %302, i32* %304, align 4
  %305 = load i32** @id, align 8
  %306 = getelementptr inbounds i32* %305, i64 3
  store i32 132, i32* %306, align 4
  %307 = load i32** @id, align 8
  %308 = getelementptr inbounds i32* %307, i64 4
  %309 = load i32* %308, align 4
  %310 = load i32** @id, align 8
  %311 = getelementptr inbounds i32* %310, i64 7
  store i32 %309, i32* %311, align 4
  %312 = load i32* %ty, align 4
  %313 = load i32** @id, align 8
  %314 = getelementptr inbounds i32* %313, i64 4
  store i32 %312, i32* %314, align 4
  %315 = load i32** @id, align 8
  %316 = getelementptr inbounds i32* %315, i64 5
  %317 = load i32* %316, align 4
  %318 = load i32** @id, align 8
  %319 = getelementptr inbounds i32* %318, i64 8
  store i32 %317, i32* %319, align 4
  %320 = load i32* %i, align 4
  %321 = add nsw i32 %320, 1
  store i32 %321, i32* %i, align 4
  %322 = load i32** @id, align 8
  %323 = getelementptr inbounds i32* %322, i64 5
  store i32 %320, i32* %323, align 4
  call void @next()
  %324 = load i32* @tk, align 4
  %325 = icmp eq i32 %324, 44
  br i1 %325, label %326, label %327

; <label>:326                                     ; preds = %299
  call void @next()
  br label %327

; <label>:327                                     ; preds = %326, %299
  br label %266

; <label>:328                                     ; preds = %266
  call void @next()
  %329 = load i32* @tk, align 4
  %330 = icmp ne i32 %329, 123
  br i1 %330, label %331, label %334

; <label>:331                                     ; preds = %328
  %332 = load i32* @line, align 4
  %333 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([29 x i8]* @.str39, i32 0, i32 0), i32 %332)
  store i32 -1, i32* %1
  br label %995

; <label>:334                                     ; preds = %328
  %335 = load i32* %i, align 4
  %336 = add nsw i32 %335, 1
  store i32 %336, i32* %i, align 4
  store i32 %336, i32* @loc, align 4
  call void @next()
  br label %337

; <label>:337                                     ; preds = %403, %334
  %338 = load i32* @tk, align 4
  %339 = icmp eq i32 %338, 138
  br i1 %339, label %343, label %340

; <label>:340                                     ; preds = %337
  %341 = load i32* @tk, align 4
  %342 = icmp eq i32 %341, 134
  br label %343

; <label>:343                                     ; preds = %340, %337
  %344 = phi i1 [ true, %337 ], [ %342, %340 ]
  br i1 %344, label %345, label %404

; <label>:345                                     ; preds = %343
  %346 = load i32* @tk, align 4
  %347 = icmp eq i32 %346, 138
  %348 = select i1 %347, i32 1, i32 0
  store i32 %348, i32* %bt, align 4
  call void @next()
  br label %349

; <label>:349                                     ; preds = %402, %345
  %350 = load i32* @tk, align 4
  %351 = icmp ne i32 %350, 59
  br i1 %351, label %352, label %403

; <label>:352                                     ; preds = %349
  %353 = load i32* %bt, align 4
  store i32 %353, i32* %ty, align 4
  br label %354

; <label>:354                                     ; preds = %357, %352
  %355 = load i32* @tk, align 4
  %356 = icmp eq i32 %355, 159
  br i1 %356, label %357, label %360

; <label>:357                                     ; preds = %354
  call void @next()
  %358 = load i32* %ty, align 4
  %359 = add nsw i32 %358, 2
  store i32 %359, i32* %ty, align 4
  br label %354

; <label>:360                                     ; preds = %354
  %361 = load i32* @tk, align 4
  %362 = icmp ne i32 %361, 133
  br i1 %362, label %363, label %366

; <label>:363                                     ; preds = %360
  %364 = load i32* @line, align 4
  %365 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([27 x i8]* @.str40, i32 0, i32 0), i32 %364)
  store i32 -1, i32* %1
  br label %995

; <label>:366                                     ; preds = %360
  %367 = load i32** @id, align 8
  %368 = getelementptr inbounds i32* %367, i64 3
  %369 = load i32* %368, align 4
  %370 = icmp eq i32 %369, 132
  br i1 %370, label %371, label %374

; <label>:371                                     ; preds = %366
  %372 = load i32* @line, align 4
  %373 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([32 x i8]* @.str41, i32 0, i32 0), i32 %372)
  store i32 -1, i32* %1
  br label %995

; <label>:374                                     ; preds = %366
  %375 = load i32** @id, align 8
  %376 = getelementptr inbounds i32* %375, i64 3
  %377 = load i32* %376, align 4
  %378 = load i32** @id, align 8
  %379 = getelementptr inbounds i32* %378, i64 6
  store i32 %377, i32* %379, align 4
  %380 = load i32** @id, align 8
  %381 = getelementptr inbounds i32* %380, i64 3
  store i32 132, i32* %381, align 4
  %382 = load i32** @id, align 8
  %383 = getelementptr inbounds i32* %382, i64 4
  %384 = load i32* %383, align 4
  %385 = load i32** @id, align 8
  %386 = getelementptr inbounds i32* %385, i64 7
  store i32 %384, i32* %386, align 4
  %387 = load i32* %ty, align 4
  %388 = load i32** @id, align 8
  %389 = getelementptr inbounds i32* %388, i64 4
  store i32 %387, i32* %389, align 4
  %390 = load i32** @id, align 8
  %391 = getelementptr inbounds i32* %390, i64 5
  %392 = load i32* %391, align 4
  %393 = load i32** @id, align 8
  %394 = getelementptr inbounds i32* %393, i64 8
  store i32 %392, i32* %394, align 4
  %395 = load i32* %i, align 4
  %396 = add nsw i32 %395, 1
  store i32 %396, i32* %i, align 4
  %397 = load i32** @id, align 8
  %398 = getelementptr inbounds i32* %397, i64 5
  store i32 %396, i32* %398, align 4
  call void @next()
  %399 = load i32* @tk, align 4
  %400 = icmp eq i32 %399, 44
  br i1 %400, label %401, label %402

; <label>:401                                     ; preds = %374
  call void @next()
  br label %402

; <label>:402                                     ; preds = %401, %374
  br label %349

; <label>:403                                     ; preds = %349
  call void @next()
  br label %337

; <label>:404                                     ; preds = %343
  %405 = load i32** @e, align 8
  %406 = getelementptr inbounds i32* %405, i32 1
  store i32* %406, i32** @e, align 8
  store i32 6, i32* %406, align 4
  %407 = load i32* %i, align 4
  %408 = load i32* @loc, align 4
  %409 = sub nsw i32 %407, %408
  %410 = load i32** @e, align 8
  %411 = getelementptr inbounds i32* %410, i32 1
  store i32* %411, i32** @e, align 8
  store i32 %409, i32* %411, align 4
  br label %412

; <label>:412                                     ; preds = %415, %404
  %413 = load i32* @tk, align 4
  %414 = icmp ne i32 %413, 125
  br i1 %414, label %415, label %416

; <label>:415                                     ; preds = %412
  call void @stmt()
  br label %412

; <label>:416                                     ; preds = %412
  %417 = load i32** @e, align 8
  %418 = getelementptr inbounds i32* %417, i32 1
  store i32* %418, i32** @e, align 8
  store i32 8, i32* %418, align 4
  %419 = load i32** @sym, align 8
  store i32* %419, i32** @id, align 8
  br label %420

; <label>:420                                     ; preds = %446, %416
  %421 = load i32** @id, align 8
  %422 = getelementptr inbounds i32* %421, i64 0
  %423 = load i32* %422, align 4
  %424 = icmp ne i32 %423, 0
  br i1 %424, label %425, label %449

; <label>:425                                     ; preds = %420
  %426 = load i32** @id, align 8
  %427 = getelementptr inbounds i32* %426, i64 3
  %428 = load i32* %427, align 4
  %429 = icmp eq i32 %428, 132
  br i1 %429, label %430, label %446

; <label>:430                                     ; preds = %425
  %431 = load i32** @id, align 8
  %432 = getelementptr inbounds i32* %431, i64 6
  %433 = load i32* %432, align 4
  %434 = load i32** @id, align 8
  %435 = getelementptr inbounds i32* %434, i64 3
  store i32 %433, i32* %435, align 4
  %436 = load i32** @id, align 8
  %437 = getelementptr inbounds i32* %436, i64 7
  %438 = load i32* %437, align 4
  %439 = load i32** @id, align 8
  %440 = getelementptr inbounds i32* %439, i64 4
  store i32 %438, i32* %440, align 4
  %441 = load i32** @id, align 8
  %442 = getelementptr inbounds i32* %441, i64 8
  %443 = load i32* %442, align 4
  %444 = load i32** @id, align 8
  %445 = getelementptr inbounds i32* %444, i64 5
  store i32 %443, i32* %445, align 4
  br label %446

; <label>:446                                     ; preds = %430, %425
  %447 = load i32** @id, align 8
  %448 = getelementptr inbounds i32* %447, i64 9
  store i32* %448, i32** @id, align 8
  br label %420

; <label>:449                                     ; preds = %420
  br label %459

; <label>:450                                     ; preds = %252
  %451 = load i32** @id, align 8
  %452 = getelementptr inbounds i32* %451, i64 3
  store i32 131, i32* %452, align 4
  %453 = load i8** @data, align 8
  %454 = ptrtoint i8* %453 to i32
  %455 = load i32** @id, align 8
  %456 = getelementptr inbounds i32* %455, i64 5
  store i32 %454, i32* %456, align 4
  %457 = load i8** @data, align 8
  %458 = getelementptr inbounds i8* %457, i64 4
  store i8* %458, i8** @data, align 8
  br label %459

; <label>:459                                     ; preds = %450, %449
  %460 = load i32* @tk, align 4
  %461 = icmp eq i32 %460, 44
  br i1 %461, label %462, label %463

; <label>:462                                     ; preds = %459
  call void @next()
  br label %463

; <label>:463                                     ; preds = %462, %459
  br label %222

; <label>:464                                     ; preds = %228
  call void @next()
  br label %161

; <label>:465                                     ; preds = %161
  %466 = load i32** %idmain, align 8
  %467 = getelementptr inbounds i32* %466, i64 5
  %468 = load i32* %467, align 4
  %469 = sext i32 %468 to i64
  %470 = inttoptr i64 %469 to i32*
  store i32* %470, i32** %pc, align 8
  %471 = icmp ne i32* %470, null
  br i1 %471, label %474, label %472

; <label>:472                                     ; preds = %465
  %473 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([20 x i8]* @.str42, i32 0, i32 0))
  store i32 -1, i32* %1
  br label %995

; <label>:474                                     ; preds = %465
  %475 = load i32* @src, align 4
  %476 = icmp ne i32 %475, 0
  br i1 %476, label %477, label %478

; <label>:477                                     ; preds = %474
  store i32 0, i32* %1
  br label %995

; <label>:478                                     ; preds = %474
  %479 = load i32** %sp, align 8
  %480 = ptrtoint i32* %479 to i32
  %481 = load i32* %poolsz, align 4
  %482 = add nsw i32 %480, %481
  %483 = sext i32 %482 to i64
  %484 = inttoptr i64 %483 to i32*
  store i32* %484, i32** %sp, align 8
  %485 = load i32** %sp, align 8
  %486 = getelementptr inbounds i32* %485, i32 -1
  store i32* %486, i32** %sp, align 8
  store i32 37, i32* %486, align 4
  %487 = load i32** %sp, align 8
  %488 = getelementptr inbounds i32* %487, i32 -1
  store i32* %488, i32** %sp, align 8
  store i32 13, i32* %488, align 4
  %489 = load i32** %sp, align 8
  store i32* %489, i32** %t, align 8
  %490 = load i32* %2, align 4
  %491 = load i32** %sp, align 8
  %492 = getelementptr inbounds i32* %491, i32 -1
  store i32* %492, i32** %sp, align 8
  store i32 %490, i32* %492, align 4
  %493 = load i8*** %3, align 8
  %494 = ptrtoint i8** %493 to i32
  %495 = load i32** %sp, align 8
  %496 = getelementptr inbounds i32* %495, i32 -1
  store i32* %496, i32** %sp, align 8
  store i32 %494, i32* %496, align 4
  %497 = load i32** %t, align 8
  %498 = ptrtoint i32* %497 to i32
  %499 = load i32** %sp, align 8
  %500 = getelementptr inbounds i32* %499, i32 -1
  store i32* %500, i32** %sp, align 8
  store i32 %498, i32* %500, align 4
  store i32 0, i32* %cycle, align 4
  br label %501

; <label>:501                                     ; preds = %478, %994
  %502 = load i32** %pc, align 8
  %503 = getelementptr inbounds i32* %502, i32 1
  store i32* %503, i32** %pc, align 8
  %504 = load i32* %502, align 4
  store i32 %504, i32* %i, align 4
  %505 = load i32* %cycle, align 4
  %506 = add nsw i32 %505, 1
  store i32 %506, i32* %cycle, align 4
  %507 = load i32* @debug, align 4
  %508 = icmp ne i32 %507, 0
  br i1 %508, label %509, label %525

; <label>:509                                     ; preds = %501
  %510 = load i32* %cycle, align 4
  %511 = load i32* %i, align 4
  %512 = mul nsw i32 %511, 5
  %513 = sext i32 %512 to i64
  %514 = getelementptr inbounds [191 x i8]* @.str2, i32 0, i64 %513
  %515 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([9 x i8]* @.str43, i32 0, i32 0), i32 %510, i8* %514)
  %516 = load i32* %i, align 4
  %517 = icmp sle i32 %516, 7
  br i1 %517, label %518, label %522

; <label>:518                                     ; preds = %509
  %519 = load i32** %pc, align 8
  %520 = load i32* %519, align 4
  %521 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([5 x i8]* @.str3, i32 0, i32 0), i32 %520)
  br label %524

; <label>:522                                     ; preds = %509
  %523 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([2 x i8]* @.str4, i32 0, i32 0))
  br label %524

; <label>:524                                     ; preds = %522, %518
  br label %525

; <label>:525                                     ; preds = %524, %501
  %526 = load i32* %i, align 4
  %527 = icmp eq i32 %526, 0
  br i1 %527, label %528, label %536

; <label>:528                                     ; preds = %525
  %529 = load i32** %bp, align 8
  %530 = load i32** %pc, align 8
  %531 = getelementptr inbounds i32* %530, i32 1
  store i32* %531, i32** %pc, align 8
  %532 = load i32* %530, align 4
  %533 = sext i32 %532 to i64
  %534 = getelementptr inbounds i32* %529, i64 %533
  %535 = ptrtoint i32* %534 to i32
  store i32 %535, i32* %a, align 4
  br label %994

; <label>:536                                     ; preds = %525
  %537 = load i32* %i, align 4
  %538 = icmp eq i32 %537, 1
  br i1 %538, label %539, label %543

; <label>:539                                     ; preds = %536
  %540 = load i32** %pc, align 8
  %541 = getelementptr inbounds i32* %540, i32 1
  store i32* %541, i32** %pc, align 8
  %542 = load i32* %540, align 4
  store i32 %542, i32* %a, align 4
  br label %993

; <label>:543                                     ; preds = %536
  %544 = load i32* %i, align 4
  %545 = icmp eq i32 %544, 2
  br i1 %545, label %546, label %551

; <label>:546                                     ; preds = %543
  %547 = load i32** %pc, align 8
  %548 = load i32* %547, align 4
  %549 = sext i32 %548 to i64
  %550 = inttoptr i64 %549 to i32*
  store i32* %550, i32** %pc, align 8
  br label %992

; <label>:551                                     ; preds = %543
  %552 = load i32* %i, align 4
  %553 = icmp eq i32 %552, 3
  br i1 %553, label %554, label %564

; <label>:554                                     ; preds = %551
  %555 = load i32** %pc, align 8
  %556 = getelementptr inbounds i32* %555, i64 1
  %557 = ptrtoint i32* %556 to i32
  %558 = load i32** %sp, align 8
  %559 = getelementptr inbounds i32* %558, i32 -1
  store i32* %559, i32** %sp, align 8
  store i32 %557, i32* %559, align 4
  %560 = load i32** %pc, align 8
  %561 = load i32* %560, align 4
  %562 = sext i32 %561 to i64
  %563 = inttoptr i64 %562 to i32*
  store i32* %563, i32** %pc, align 8
  br label %991

; <label>:564                                     ; preds = %551
  %565 = load i32* %i, align 4
  %566 = icmp eq i32 %565, 4
  br i1 %566, label %567, label %580

; <label>:567                                     ; preds = %564
  %568 = load i32* %a, align 4
  %569 = icmp ne i32 %568, 0
  br i1 %569, label %570, label %573

; <label>:570                                     ; preds = %567
  %571 = load i32** %pc, align 8
  %572 = getelementptr inbounds i32* %571, i64 1
  br label %578

; <label>:573                                     ; preds = %567
  %574 = load i32** %pc, align 8
  %575 = load i32* %574, align 4
  %576 = sext i32 %575 to i64
  %577 = inttoptr i64 %576 to i32*
  br label %578

; <label>:578                                     ; preds = %573, %570
  %579 = phi i32* [ %572, %570 ], [ %577, %573 ]
  store i32* %579, i32** %pc, align 8
  br label %990

; <label>:580                                     ; preds = %564
  %581 = load i32* %i, align 4
  %582 = icmp eq i32 %581, 5
  br i1 %582, label %583, label %596

; <label>:583                                     ; preds = %580
  %584 = load i32* %a, align 4
  %585 = icmp ne i32 %584, 0
  br i1 %585, label %586, label %591

; <label>:586                                     ; preds = %583
  %587 = load i32** %pc, align 8
  %588 = load i32* %587, align 4
  %589 = sext i32 %588 to i64
  %590 = inttoptr i64 %589 to i32*
  br label %594

; <label>:591                                     ; preds = %583
  %592 = load i32** %pc, align 8
  %593 = getelementptr inbounds i32* %592, i64 1
  br label %594

; <label>:594                                     ; preds = %591, %586
  %595 = phi i32* [ %590, %586 ], [ %593, %591 ]
  store i32* %595, i32** %pc, align 8
  br label %989

; <label>:596                                     ; preds = %580
  %597 = load i32* %i, align 4
  %598 = icmp eq i32 %597, 6
  br i1 %598, label %599, label %612

; <label>:599                                     ; preds = %596
  %600 = load i32** %bp, align 8
  %601 = ptrtoint i32* %600 to i32
  %602 = load i32** %sp, align 8
  %603 = getelementptr inbounds i32* %602, i32 -1
  store i32* %603, i32** %sp, align 8
  store i32 %601, i32* %603, align 4
  %604 = load i32** %sp, align 8
  store i32* %604, i32** %bp, align 8
  %605 = load i32** %sp, align 8
  %606 = load i32** %pc, align 8
  %607 = getelementptr inbounds i32* %606, i32 1
  store i32* %607, i32** %pc, align 8
  %608 = load i32* %606, align 4
  %609 = sext i32 %608 to i64
  %610 = sub i64 0, %609
  %611 = getelementptr inbounds i32* %605, i64 %610
  store i32* %611, i32** %sp, align 8
  br label %988

; <label>:612                                     ; preds = %596
  %613 = load i32* %i, align 4
  %614 = icmp eq i32 %613, 7
  br i1 %614, label %615, label %622

; <label>:615                                     ; preds = %612
  %616 = load i32** %sp, align 8
  %617 = load i32** %pc, align 8
  %618 = getelementptr inbounds i32* %617, i32 1
  store i32* %618, i32** %pc, align 8
  %619 = load i32* %617, align 4
  %620 = sext i32 %619 to i64
  %621 = getelementptr inbounds i32* %616, i64 %620
  store i32* %621, i32** %sp, align 8
  br label %987

; <label>:622                                     ; preds = %612
  %623 = load i32* %i, align 4
  %624 = icmp eq i32 %623, 8
  br i1 %624, label %625, label %637

; <label>:625                                     ; preds = %622
  %626 = load i32** %bp, align 8
  store i32* %626, i32** %sp, align 8
  %627 = load i32** %sp, align 8
  %628 = getelementptr inbounds i32* %627, i32 1
  store i32* %628, i32** %sp, align 8
  %629 = load i32* %627, align 4
  %630 = sext i32 %629 to i64
  %631 = inttoptr i64 %630 to i32*
  store i32* %631, i32** %bp, align 8
  %632 = load i32** %sp, align 8
  %633 = getelementptr inbounds i32* %632, i32 1
  store i32* %633, i32** %sp, align 8
  %634 = load i32* %632, align 4
  %635 = sext i32 %634 to i64
  %636 = inttoptr i64 %635 to i32*
  store i32* %636, i32** %pc, align 8
  br label %986

; <label>:637                                     ; preds = %622
  %638 = load i32* %i, align 4
  %639 = icmp eq i32 %638, 9
  br i1 %639, label %640, label %645

; <label>:640                                     ; preds = %637
  %641 = load i32* %a, align 4
  %642 = sext i32 %641 to i64
  %643 = inttoptr i64 %642 to i32*
  %644 = load i32* %643, align 4
  store i32 %644, i32* %a, align 4
  br label %985

; <label>:645                                     ; preds = %637
  %646 = load i32* %i, align 4
  %647 = icmp eq i32 %646, 10
  br i1 %647, label %648, label %654

; <label>:648                                     ; preds = %645
  %649 = load i32* %a, align 4
  %650 = sext i32 %649 to i64
  %651 = inttoptr i64 %650 to i8*
  %652 = load i8* %651, align 1
  %653 = sext i8 %652 to i32
  store i32 %653, i32* %a, align 4
  br label %984

; <label>:654                                     ; preds = %645
  %655 = load i32* %i, align 4
  %656 = icmp eq i32 %655, 11
  br i1 %656, label %657, label %664

; <label>:657                                     ; preds = %654
  %658 = load i32* %a, align 4
  %659 = load i32** %sp, align 8
  %660 = getelementptr inbounds i32* %659, i32 1
  store i32* %660, i32** %sp, align 8
  %661 = load i32* %659, align 4
  %662 = sext i32 %661 to i64
  %663 = inttoptr i64 %662 to i32*
  store i32 %658, i32* %663, align 4
  br label %983

; <label>:664                                     ; preds = %654
  %665 = load i32* %i, align 4
  %666 = icmp eq i32 %665, 12
  br i1 %666, label %667, label %676

; <label>:667                                     ; preds = %664
  %668 = load i32* %a, align 4
  %669 = trunc i32 %668 to i8
  %670 = load i32** %sp, align 8
  %671 = getelementptr inbounds i32* %670, i32 1
  store i32* %671, i32** %sp, align 8
  %672 = load i32* %670, align 4
  %673 = sext i32 %672 to i64
  %674 = inttoptr i64 %673 to i8*
  store i8 %669, i8* %674, align 1
  %675 = sext i8 %669 to i32
  store i32 %675, i32* %a, align 4
  br label %982

; <label>:676                                     ; preds = %664
  %677 = load i32* %i, align 4
  %678 = icmp eq i32 %677, 13
  br i1 %678, label %679, label %683

; <label>:679                                     ; preds = %676
  %680 = load i32* %a, align 4
  %681 = load i32** %sp, align 8
  %682 = getelementptr inbounds i32* %681, i32 -1
  store i32* %682, i32** %sp, align 8
  store i32 %680, i32* %682, align 4
  br label %981

; <label>:683                                     ; preds = %676
  %684 = load i32* %i, align 4
  %685 = icmp eq i32 %684, 14
  br i1 %685, label %686, label %692

; <label>:686                                     ; preds = %683
  %687 = load i32** %sp, align 8
  %688 = getelementptr inbounds i32* %687, i32 1
  store i32* %688, i32** %sp, align 8
  %689 = load i32* %687, align 4
  %690 = load i32* %a, align 4
  %691 = or i32 %689, %690
  store i32 %691, i32* %a, align 4
  br label %980

; <label>:692                                     ; preds = %683
  %693 = load i32* %i, align 4
  %694 = icmp eq i32 %693, 15
  br i1 %694, label %695, label %701

; <label>:695                                     ; preds = %692
  %696 = load i32** %sp, align 8
  %697 = getelementptr inbounds i32* %696, i32 1
  store i32* %697, i32** %sp, align 8
  %698 = load i32* %696, align 4
  %699 = load i32* %a, align 4
  %700 = xor i32 %698, %699
  store i32 %700, i32* %a, align 4
  br label %979

; <label>:701                                     ; preds = %692
  %702 = load i32* %i, align 4
  %703 = icmp eq i32 %702, 16
  br i1 %703, label %704, label %710

; <label>:704                                     ; preds = %701
  %705 = load i32** %sp, align 8
  %706 = getelementptr inbounds i32* %705, i32 1
  store i32* %706, i32** %sp, align 8
  %707 = load i32* %705, align 4
  %708 = load i32* %a, align 4
  %709 = and i32 %707, %708
  store i32 %709, i32* %a, align 4
  br label %978

; <label>:710                                     ; preds = %701
  %711 = load i32* %i, align 4
  %712 = icmp eq i32 %711, 17
  br i1 %712, label %713, label %720

; <label>:713                                     ; preds = %710
  %714 = load i32** %sp, align 8
  %715 = getelementptr inbounds i32* %714, i32 1
  store i32* %715, i32** %sp, align 8
  %716 = load i32* %714, align 4
  %717 = load i32* %a, align 4
  %718 = icmp eq i32 %716, %717
  %719 = zext i1 %718 to i32
  store i32 %719, i32* %a, align 4
  br label %977

; <label>:720                                     ; preds = %710
  %721 = load i32* %i, align 4
  %722 = icmp eq i32 %721, 18
  br i1 %722, label %723, label %730

; <label>:723                                     ; preds = %720
  %724 = load i32** %sp, align 8
  %725 = getelementptr inbounds i32* %724, i32 1
  store i32* %725, i32** %sp, align 8
  %726 = load i32* %724, align 4
  %727 = load i32* %a, align 4
  %728 = icmp ne i32 %726, %727
  %729 = zext i1 %728 to i32
  store i32 %729, i32* %a, align 4
  br label %976

; <label>:730                                     ; preds = %720
  %731 = load i32* %i, align 4
  %732 = icmp eq i32 %731, 19
  br i1 %732, label %733, label %740

; <label>:733                                     ; preds = %730
  %734 = load i32** %sp, align 8
  %735 = getelementptr inbounds i32* %734, i32 1
  store i32* %735, i32** %sp, align 8
  %736 = load i32* %734, align 4
  %737 = load i32* %a, align 4
  %738 = icmp slt i32 %736, %737
  %739 = zext i1 %738 to i32
  store i32 %739, i32* %a, align 4
  br label %975

; <label>:740                                     ; preds = %730
  %741 = load i32* %i, align 4
  %742 = icmp eq i32 %741, 20
  br i1 %742, label %743, label %750

; <label>:743                                     ; preds = %740
  %744 = load i32** %sp, align 8
  %745 = getelementptr inbounds i32* %744, i32 1
  store i32* %745, i32** %sp, align 8
  %746 = load i32* %744, align 4
  %747 = load i32* %a, align 4
  %748 = icmp sgt i32 %746, %747
  %749 = zext i1 %748 to i32
  store i32 %749, i32* %a, align 4
  br label %974

; <label>:750                                     ; preds = %740
  %751 = load i32* %i, align 4
  %752 = icmp eq i32 %751, 21
  br i1 %752, label %753, label %760

; <label>:753                                     ; preds = %750
  %754 = load i32** %sp, align 8
  %755 = getelementptr inbounds i32* %754, i32 1
  store i32* %755, i32** %sp, align 8
  %756 = load i32* %754, align 4
  %757 = load i32* %a, align 4
  %758 = icmp sle i32 %756, %757
  %759 = zext i1 %758 to i32
  store i32 %759, i32* %a, align 4
  br label %973

; <label>:760                                     ; preds = %750
  %761 = load i32* %i, align 4
  %762 = icmp eq i32 %761, 22
  br i1 %762, label %763, label %770

; <label>:763                                     ; preds = %760
  %764 = load i32** %sp, align 8
  %765 = getelementptr inbounds i32* %764, i32 1
  store i32* %765, i32** %sp, align 8
  %766 = load i32* %764, align 4
  %767 = load i32* %a, align 4
  %768 = icmp sge i32 %766, %767
  %769 = zext i1 %768 to i32
  store i32 %769, i32* %a, align 4
  br label %972

; <label>:770                                     ; preds = %760
  %771 = load i32* %i, align 4
  %772 = icmp eq i32 %771, 23
  br i1 %772, label %773, label %779

; <label>:773                                     ; preds = %770
  %774 = load i32** %sp, align 8
  %775 = getelementptr inbounds i32* %774, i32 1
  store i32* %775, i32** %sp, align 8
  %776 = load i32* %774, align 4
  %777 = load i32* %a, align 4
  %778 = shl i32 %776, %777
  store i32 %778, i32* %a, align 4
  br label %971

; <label>:779                                     ; preds = %770
  %780 = load i32* %i, align 4
  %781 = icmp eq i32 %780, 24
  br i1 %781, label %782, label %788

; <label>:782                                     ; preds = %779
  %783 = load i32** %sp, align 8
  %784 = getelementptr inbounds i32* %783, i32 1
  store i32* %784, i32** %sp, align 8
  %785 = load i32* %783, align 4
  %786 = load i32* %a, align 4
  %787 = ashr i32 %785, %786
  store i32 %787, i32* %a, align 4
  br label %970

; <label>:788                                     ; preds = %779
  %789 = load i32* %i, align 4
  %790 = icmp eq i32 %789, 25
  br i1 %790, label %791, label %797

; <label>:791                                     ; preds = %788
  %792 = load i32** %sp, align 8
  %793 = getelementptr inbounds i32* %792, i32 1
  store i32* %793, i32** %sp, align 8
  %794 = load i32* %792, align 4
  %795 = load i32* %a, align 4
  %796 = add nsw i32 %794, %795
  store i32 %796, i32* %a, align 4
  br label %969

; <label>:797                                     ; preds = %788
  %798 = load i32* %i, align 4
  %799 = icmp eq i32 %798, 26
  br i1 %799, label %800, label %806

; <label>:800                                     ; preds = %797
  %801 = load i32** %sp, align 8
  %802 = getelementptr inbounds i32* %801, i32 1
  store i32* %802, i32** %sp, align 8
  %803 = load i32* %801, align 4
  %804 = load i32* %a, align 4
  %805 = sub nsw i32 %803, %804
  store i32 %805, i32* %a, align 4
  br label %968

; <label>:806                                     ; preds = %797
  %807 = load i32* %i, align 4
  %808 = icmp eq i32 %807, 27
  br i1 %808, label %809, label %815

; <label>:809                                     ; preds = %806
  %810 = load i32** %sp, align 8
  %811 = getelementptr inbounds i32* %810, i32 1
  store i32* %811, i32** %sp, align 8
  %812 = load i32* %810, align 4
  %813 = load i32* %a, align 4
  %814 = mul nsw i32 %812, %813
  store i32 %814, i32* %a, align 4
  br label %967

; <label>:815                                     ; preds = %806
  %816 = load i32* %i, align 4
  %817 = icmp eq i32 %816, 28
  br i1 %817, label %818, label %824

; <label>:818                                     ; preds = %815
  %819 = load i32** %sp, align 8
  %820 = getelementptr inbounds i32* %819, i32 1
  store i32* %820, i32** %sp, align 8
  %821 = load i32* %819, align 4
  %822 = load i32* %a, align 4
  %823 = sdiv i32 %821, %822
  store i32 %823, i32* %a, align 4
  br label %966

; <label>:824                                     ; preds = %815
  %825 = load i32* %i, align 4
  %826 = icmp eq i32 %825, 29
  br i1 %826, label %827, label %833

; <label>:827                                     ; preds = %824
  %828 = load i32** %sp, align 8
  %829 = getelementptr inbounds i32* %828, i32 1
  store i32* %829, i32** %sp, align 8
  %830 = load i32* %828, align 4
  %831 = load i32* %a, align 4
  %832 = srem i32 %830, %831
  store i32 %832, i32* %a, align 4
  br label %965

; <label>:833                                     ; preds = %824
  %834 = load i32* %i, align 4
  %835 = icmp eq i32 %834, 30
  br i1 %835, label %836, label %845

; <label>:836                                     ; preds = %833
  %837 = load i32** %sp, align 8
  %838 = getelementptr inbounds i32* %837, i64 1
  %839 = load i32* %838, align 4
  %840 = sext i32 %839 to i64
  %841 = inttoptr i64 %840 to i8*
  %842 = load i32** %sp, align 8
  %843 = load i32* %842, align 4
  %844 = call i32 (i8*, i32, ...)* bitcast (i32 (...)* @open to i32 (i8*, i32, ...)*)(i8* %841, i32 %843)
  store i32 %844, i32* %a, align 4
  br label %964

; <label>:845                                     ; preds = %833
  %846 = load i32* %i, align 4
  %847 = icmp eq i32 %846, 31
  br i1 %847, label %848, label %862

; <label>:848                                     ; preds = %845
  %849 = load i32** %sp, align 8
  %850 = getelementptr inbounds i32* %849, i64 2
  %851 = load i32* %850, align 4
  %852 = load i32** %sp, align 8
  %853 = getelementptr inbounds i32* %852, i64 1
  %854 = load i32* %853, align 4
  %855 = sext i32 %854 to i64
  %856 = inttoptr i64 %855 to i8*
  %857 = load i32** %sp, align 8
  %858 = load i32* %857, align 4
  %859 = sext i32 %858 to i64
  %860 = call i64 @read(i32 %851, i8* %856, i64 %859)
  %861 = trunc i64 %860 to i32
  store i32 %861, i32* %a, align 4
  br label %963

; <label>:862                                     ; preds = %845
  %863 = load i32* %i, align 4
  %864 = icmp eq i32 %863, 32
  br i1 %864, label %865, label %869

; <label>:865                                     ; preds = %862
  %866 = load i32** %sp, align 8
  %867 = load i32* %866, align 4
  %868 = call i32 @close(i32 %867)
  store i32 %868, i32* %a, align 4
  br label %962

; <label>:869                                     ; preds = %862
  %870 = load i32* %i, align 4
  %871 = icmp eq i32 %870, 33
  br i1 %871, label %872, label %900

; <label>:872                                     ; preds = %869
  %873 = load i32** %sp, align 8
  %874 = load i32** %pc, align 8
  %875 = getelementptr inbounds i32* %874, i64 1
  %876 = load i32* %875, align 4
  %877 = sext i32 %876 to i64
  %878 = getelementptr inbounds i32* %873, i64 %877
  store i32* %878, i32** %t, align 8
  %879 = load i32** %t, align 8
  %880 = getelementptr inbounds i32* %879, i64 -1
  %881 = load i32* %880, align 4
  %882 = sext i32 %881 to i64
  %883 = inttoptr i64 %882 to i8*
  %884 = load i32** %t, align 8
  %885 = getelementptr inbounds i32* %884, i64 -2
  %886 = load i32* %885, align 4
  %887 = load i32** %t, align 8
  %888 = getelementptr inbounds i32* %887, i64 -3
  %889 = load i32* %888, align 4
  %890 = load i32** %t, align 8
  %891 = getelementptr inbounds i32* %890, i64 -4
  %892 = load i32* %891, align 4
  %893 = load i32** %t, align 8
  %894 = getelementptr inbounds i32* %893, i64 -5
  %895 = load i32* %894, align 4
  %896 = load i32** %t, align 8
  %897 = getelementptr inbounds i32* %896, i64 -6
  %898 = load i32* %897, align 4
  %899 = call i32 (i8*, ...)* @printf(i8* %883, i32 %886, i32 %889, i32 %892, i32 %895, i32 %898)
  store i32 %899, i32* %a, align 4
  br label %961

; <label>:900                                     ; preds = %869
  %901 = load i32* %i, align 4
  %902 = icmp eq i32 %901, 34
  br i1 %902, label %903, label %909

; <label>:903                                     ; preds = %900
  %904 = load i32** %sp, align 8
  %905 = load i32* %904, align 4
  %906 = sext i32 %905 to i64
  %907 = call noalias i8* @malloc(i64 %906) #5
  %908 = ptrtoint i8* %907 to i32
  store i32 %908, i32* %a, align 4
  br label %960

; <label>:909                                     ; preds = %900
  %910 = load i32* %i, align 4
  %911 = icmp eq i32 %910, 35
  br i1 %911, label %912, label %926

; <label>:912                                     ; preds = %909
  %913 = load i32** %sp, align 8
  %914 = getelementptr inbounds i32* %913, i64 2
  %915 = load i32* %914, align 4
  %916 = sext i32 %915 to i64
  %917 = inttoptr i64 %916 to i8*
  %918 = load i32** %sp, align 8
  %919 = getelementptr inbounds i32* %918, i64 1
  %920 = load i32* %919, align 4
  %921 = trunc i32 %920 to i8
  %922 = load i32** %sp, align 8
  %923 = load i32* %922, align 4
  %924 = sext i32 %923 to i64
  call void @llvm.memset.p0i8.i64(i8* %917, i8 %921, i64 %924, i32 1, i1 false)
  %925 = ptrtoint i8* %917 to i32
  store i32 %925, i32* %a, align 4
  br label %959

; <label>:926                                     ; preds = %909
  %927 = load i32* %i, align 4
  %928 = icmp eq i32 %927, 36
  br i1 %928, label %929, label %944

; <label>:929                                     ; preds = %926
  %930 = load i32** %sp, align 8
  %931 = getelementptr inbounds i32* %930, i64 2
  %932 = load i32* %931, align 4
  %933 = sext i32 %932 to i64
  %934 = inttoptr i64 %933 to i8*
  %935 = load i32** %sp, align 8
  %936 = getelementptr inbounds i32* %935, i64 1
  %937 = load i32* %936, align 4
  %938 = sext i32 %937 to i64
  %939 = inttoptr i64 %938 to i8*
  %940 = load i32** %sp, align 8
  %941 = load i32* %940, align 4
  %942 = sext i32 %941 to i64
  %943 = call i32 @memcmp(i8* %934, i8* %939, i64 %942) #6
  store i32 %943, i32* %a, align 4
  br label %958

; <label>:944                                     ; preds = %926
  %945 = load i32* %i, align 4
  %946 = icmp eq i32 %945, 37
  br i1 %946, label %947, label %954

; <label>:947                                     ; preds = %944
  %948 = load i32** %sp, align 8
  %949 = load i32* %948, align 4
  %950 = load i32* %cycle, align 4
  %951 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([21 x i8]* @.str44, i32 0, i32 0), i32 %949, i32 %950)
  %952 = load i32** %sp, align 8
  %953 = load i32* %952, align 4
  store i32 %953, i32* %1
  br label %995

; <label>:954                                     ; preds = %944
  %955 = load i32* %i, align 4
  %956 = load i32* %cycle, align 4
  %957 = call i32 (i8*, ...)* @printf(i8* getelementptr inbounds ([38 x i8]* @.str45, i32 0, i32 0), i32 %955, i32 %956)
  store i32 -1, i32* %1
  br label %995

; <label>:958                                     ; preds = %929
  br label %959

; <label>:959                                     ; preds = %958, %912
  br label %960

; <label>:960                                     ; preds = %959, %903
  br label %961

; <label>:961                                     ; preds = %960, %872
  br label %962

; <label>:962                                     ; preds = %961, %865
  br label %963

; <label>:963                                     ; preds = %962, %848
  br label %964

; <label>:964                                     ; preds = %963, %836
  br label %965

; <label>:965                                     ; preds = %964, %827
  br label %966

; <label>:966                                     ; preds = %965, %818
  br label %967

; <label>:967                                     ; preds = %966, %809
  br label %968

; <label>:968                                     ; preds = %967, %800
  br label %969

; <label>:969                                     ; preds = %968, %791
  br label %970

; <label>:970                                     ; preds = %969, %782
  br label %971

; <label>:971                                     ; preds = %970, %773
  br label %972

; <label>:972                                     ; preds = %971, %763
  br label %973

; <label>:973                                     ; preds = %972, %753
  br label %974

; <label>:974                                     ; preds = %973, %743
  br label %975

; <label>:975                                     ; preds = %974, %733
  br label %976

; <label>:976                                     ; preds = %975, %723
  br label %977

; <label>:977                                     ; preds = %976, %713
  br label %978

; <label>:978                                     ; preds = %977, %704
  br label %979

; <label>:979                                     ; preds = %978, %695
  br label %980

; <label>:980                                     ; preds = %979, %686
  br label %981

; <label>:981                                     ; preds = %980, %679
  br label %982

; <label>:982                                     ; preds = %981, %667
  br label %983

; <label>:983                                     ; preds = %982, %657
  br label %984

; <label>:984                                     ; preds = %983, %648
  br label %985

; <label>:985                                     ; preds = %984, %640
  br label %986

; <label>:986                                     ; preds = %985, %625
  br label %987

; <label>:987                                     ; preds = %986, %615
  br label %988

; <label>:988                                     ; preds = %987, %599
  br label %989

; <label>:989                                     ; preds = %988, %594
  br label %990

; <label>:990                                     ; preds = %989, %578
  br label %991

; <label>:991                                     ; preds = %990, %554
  br label %992

; <label>:992                                     ; preds = %991, %546
  br label %993

; <label>:993                                     ; preds = %992, %539
  br label %994

; <label>:994                                     ; preds = %993, %528
  br label %501

; <label>:995                                     ; preds = %954, %947, %477, %472, %371, %363, %331, %296, %288, %249, %241, %199, %189, %151, %139, %95, %86, %78, %69, %59, %52
  %996 = load i32* %1
  ret i32 %996
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

!0 = metadata !{metadata !"clang version 3.5.1 (tags/RELEASE_351/final)"}
