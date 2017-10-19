# Version 0.x to be released (master branch)

This file tracks the implementation status of various LLVM IR language concepts, as covered by the master branch of [llir/llvm](https://github.com/llir/llvm).

The `asm` checkbox tracks read support of the language concept by the [llvm/asm](https://godoc.org/github.com/llir/llvm/asm) package.

The `ir` checkbox tracks support for an in-memory representation of the language concept by the [llvm/ir](https://godoc.org/github.com/llir/llvm/ir) package.

# Modules

* Source filename (ref [[1](http://llvm.org/docs/LangRef.html#source-filename)])
    - [x] asm
    - [ ] ir
* Target specifiers (ref [[1](http://llvm.org/docs/LangRef.html#data-layout), [2](http://llvm.org/docs/LangRef.html#target-triple)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#Module.DataLayout), [2](https://godoc.org/github.com/llir/llvm/ir#Module.TargetTriple)]
* Module-level inline assembly (ref [[1](http://llvm.org/docs/LangRef.html#module-level-inline-assembly)])
    - [x] asm
    - [ ] ir
* Type definitions (ref [[1](http://llvm.org/docs/LangRef.html#structure-types)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#Module.Types)]
* Comdat definitions (ref [[1](http://llvm.org/docs/LangRef.html#comdats)])
    - [x] asm
    - [ ] ir
* Global variables (ref [[1](http://llvm.org/docs/LangRef.html#global-variables)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#Module.Globals)]
* Functions (ref [[1](http://llvm.org/docs/LangRef.html#functions)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#Module.Funcs)]
* Attribute group definitions (ref [[1](http://llvm.org/docs/LangRef.html#attribute-groups)])
    - [x] asm
    - [ ] ir
* Metadata definitions (ref [[1](http://llvm.org/docs/LangRef.html#metadata)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#Module.NamedMetadata), [2](https://godoc.org/github.com/llir/llvm/ir#Module.Metadata)]

# Types

* Void type (ref [[1](http://llvm.org/docs/LangRef.html#void-type)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/types#VoidType)]
* Function type (ref [[1](http://llvm.org/docs/LangRef.html#function-type)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/types#FuncType)]
* Integer type (ref [[1](http://llvm.org/docs/LangRef.html#integer-type)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/types#IntType)]
* Floating-point type (ref [[1](http://llvm.org/docs/LangRef.html#floating-point-types)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/types#FloatType)]
* x86 MMX type (ref [[1](http://llvm.org/docs/LangRef.html#x86-mmx-type)])
    - [ ] asm
    - [ ] ir
* Pointer type (ref [[1](http://llvm.org/docs/LangRef.html#pointer-type)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/types#PointerType)]
* Vector type (ref [[1](http://llvm.org/docs/LangRef.html#vector-type)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/types#VectorType)]
* Label type (ref [[1](http://llvm.org/docs/LangRef.html#label-type)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/types#LabelType)]
* Token type (ref [[1](http://llvm.org/docs/LangRef.html#token-type)])
    - [ ] asm
    - [ ] ir
* Metadata type (ref [[1](http://llvm.org/docs/LangRef.html#metadata-type)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/types#MetadataType)]
* Array type (ref [[1](http://llvm.org/docs/LangRef.html#array-type)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/types#ArrayType)]
* Struct type (ref [[1](http://llvm.org/docs/LangRef.html#structure-type)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/types#StructType)]

# Constants

* Integer constant (ref [[1](http://llvm.org/docs/LangRef.html#simple-constants)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#Int)]
* Floating-point constant (ref [[1](http://llvm.org/docs/LangRef.html#simple-constants)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#Float)]
* Pointer constant (ref [[1](http://llvm.org/docs/LangRef.html#simple-constants), [2](http://llvm.org/docs/LangRef.html#global-variable-and-function-addresses)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#Null), [2](https://godoc.org/github.com/llir/llvm/ir#Global), [3](https://godoc.org/github.com/llir/llvm/ir#Function)]
* Token constant (ref [[1](http://llvm.org/docs/LangRef.html#simple-constants)])
    - [ ] asm
    - [ ] ir
* Vector constant (ref [[1](http://llvm.org/docs/LangRef.html#complex-constants)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#Vector)]
* Array constant (ref [[1](http://llvm.org/docs/LangRef.html#complex-constants)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#Array)]
* Struct constant (ref [[1](http://llvm.org/docs/LangRef.html#complex-constants)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#Struct)]
* Zero initializer constant (ref [[1](http://llvm.org/docs/LangRef.html#complex-constants)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ZeroInitializer)]
* Undefined value constant (ref [[1](http://llvm.org/docs/LangRef.html#undefined-values)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#Undef)]
* Block address constant (ref [[1](http://llvm.org/docs/LangRef.html#addresses-of-basic-blocks)])
    - [ ] asm
    - [ ] ir

# Constant expressions

## Binary expressions

* add (ref [[1](http://llvm.org/docs/LangRef.html#add-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprAdd)]
* fadd (ref [[1](http://llvm.org/docs/LangRef.html#fadd-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFAdd)]
* sub (ref [[1](http://llvm.org/docs/LangRef.html#sub-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprSub)]
* fsub (ref [[1](http://llvm.org/docs/LangRef.html#fsub-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFSub)]
* mul (ref [[1](http://llvm.org/docs/LangRef.html#mul-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprMul)]
* fmul (ref [[1](http://llvm.org/docs/LangRef.html#fmul-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFMul)]
* udiv (ref [[1](http://llvm.org/docs/LangRef.html#udiv-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprUDiv)]
* sdiv (ref [[1](http://llvm.org/docs/LangRef.html#sdiv-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprSDiv)]
* fdiv (ref [[1](http://llvm.org/docs/LangRef.html#fdiv-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFDiv)]
* urem (ref [[1](http://llvm.org/docs/LangRef.html#urem-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprURem)]
* srem (ref [[1](http://llvm.org/docs/LangRef.html#srem-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprSRem)]
* frem (ref [[1](http://llvm.org/docs/LangRef.html#frem-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFRem)]

## Bitwise expressions

* shl (ref [[1](http://llvm.org/docs/LangRef.html#shl-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprShl)]
* lshr (ref [[1](http://llvm.org/docs/LangRef.html#lshr-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprLShr)]
* ashr (ref [[1](http://llvm.org/docs/LangRef.html#ashr-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprAShr)]
* and (ref [[1](http://llvm.org/docs/LangRef.html#and-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprAnd)]
* or (ref [[1](http://llvm.org/docs/LangRef.html#or-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprOr)]
* xor (ref [[1](http://llvm.org/docs/LangRef.html#xor-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprXor)]

## Vector expressions

* extractelement (ref [[1](http://llvm.org/docs/LangRef.html#extractelement-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprExtractElement)]
* insertelement (ref [[1](http://llvm.org/docs/LangRef.html#insertelement-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprInsertElement)]
* shufflevector (ref [[1](http://llvm.org/docs/LangRef.html#shufflevector-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprShuffleVector)]

## Aggregate expressions

* extractvalue (ref [[1](http://llvm.org/docs/LangRef.html#extractvalue-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprExtractValue)]
* insertvalue (ref [[1](http://llvm.org/docs/LangRef.html#insertvalue-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprInsertValue)]

## Memory expressions

* getelementptr (ref [[1](http://llvm.org/docs/LangRef.html#getelementptr-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprGetElementPtr)]

## Conversion expressions

* trunc (ref [[1](http://llvm.org/docs/LangRef.html#trunc-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprTrunc)]
* zext (ref [[1](http://llvm.org/docs/LangRef.html#zext-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprZExt)]
* sext (ref [[1](http://llvm.org/docs/LangRef.html#sext-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprSExt)]
* fptrunc (ref [[1](http://llvm.org/docs/LangRef.html#fptrunc-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPTrunc)]
* fpext (ref [[1](http://llvm.org/docs/LangRef.html#fpext-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPExt)]
* fptoui (ref [[1](http://llvm.org/docs/LangRef.html#fptoui-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPToUI)]
* fptosi (ref [[1](http://llvm.org/docs/LangRef.html#fptosi-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPToSI)]
* uitofp (ref [[1](http://llvm.org/docs/LangRef.html#uitofp-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprUIToFP)]
* sitofp (ref [[1](http://llvm.org/docs/LangRef.html#sitofp-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprSIToFP)]
* ptrtoint (ref [[1](http://llvm.org/docs/LangRef.html#ptrtoint-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprPtrToInt)]
* inttoptr (ref [[1](http://llvm.org/docs/LangRef.html#inttoptr-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprIntToPtr)]
* bitcast (ref [[1](http://llvm.org/docs/LangRef.html#bitcast-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprBitCast)]
* addrspacecast (ref [[1](http://llvm.org/docs/LangRef.html#addrspacecast-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprAddrSpaceCast)]

## Other expressions

* icmp (ref [[1](http://llvm.org/docs/LangRef.html#icmp-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprICmp)]
* fcmp (ref [[1](http://llvm.org/docs/LangRef.html#fcmp-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFCmp)]
* select (ref [[1](http://llvm.org/docs/LangRef.html#select-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir/constant#ExprSelect)]

# Global variables

Global variables (ref [[1]](https://godoc.org/github.com/llir/llvm/ir#Global))

* Linkage type
    - [x] asm
    - [ ] ir
* Visibility style
    - [x] asm
    - [ ] ir
* DLL storage class
    - [x] asm
    - [ ] ir
* Thread local storage model
    - [x] asm
    - [ ] ir
* Unnamed address
    - [x] asm
    - [ ] ir
* Address space
    - [x] asm
    - [x] ir (ref [[1]](https://godoc.org/github.com/llir/llvm/ir#Global.Typ))
* Externally initialized
    - [x] asm
    - [ ] ir
* Section name
    - [x] asm
    - [ ] ir
* COMDAT name
    - [x] asm
    - [ ] ir
* Alignment
    - [x] asm
    - [ ] ir
* Attached metadata
    - [x] asm
    - [x] ir (ref [[1]](https://godoc.org/github.com/llir/llvm/ir#Global.Metadata))

# Functions

Functions (ref [[1]](https://godoc.org/github.com/llir/llvm/ir#Function))

* Linkage type
    - [x] asm
    - [ ] ir
* Visibility style
    - [x] asm
    - [ ] ir
* DLL storage class
    - [x] asm
    - [ ] ir
* Calling convention
    - [x] asm
    - [x] ir (ref [[1]](https://godoc.org/github.com/llir/llvm/ir#Function.CallConv))
* Return type parameter attributes
    - [x] asm
    - [ ] ir
* Argument parameter attributes
    - [x] asm
    - [ ] ir
* Unnamed address
    - [x] asm
    - [ ] ir
* Function attributes
    - [x] asm
    - [ ] ir
* Section name
    - [x] asm
    - [ ] ir
* COMDAT name
    - [x] asm
    - [ ] ir
* Alignment
    - [x] asm
    - [ ] ir
* Garbage collector name
    - [x] asm
    - [ ] ir
* Prefix data
    - [x] asm
    - [ ] ir
* Prologue data
    - [x] asm
    - [ ] ir
* Personality function data
    - [x] asm
    - [ ] ir
* Attached metadata
    - [x] asm
    - [x] ir (ref [[1]](https://godoc.org/github.com/llir/llvm/ir#Function.Metadata))

# Instructions

## Binary instructions

* add (ref [[1](http://llvm.org/docs/LangRef.html#add-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstAdd)]
* fadd (ref [[1](http://llvm.org/docs/LangRef.html#fadd-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstFAdd)]
* sub (ref [[1](http://llvm.org/docs/LangRef.html#sub-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstSub)]
* fsub (ref [[1](http://llvm.org/docs/LangRef.html#fsub-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstFSub)]
* mul (ref [[1](http://llvm.org/docs/LangRef.html#mul-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstMul)]
* fmul (ref [[1](http://llvm.org/docs/LangRef.html#fmul-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstFMul)]
* udiv (ref [[1](http://llvm.org/docs/LangRef.html#udiv-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstUDiv)]
* sdiv (ref [[1](http://llvm.org/docs/LangRef.html#sdiv-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstSDiv)]
* fdiv (ref [[1](http://llvm.org/docs/LangRef.html#fdiv-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstFDiv)]
* urem (ref [[1](http://llvm.org/docs/LangRef.html#urem-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstURem)]
* srem (ref [[1](http://llvm.org/docs/LangRef.html#srem-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstSRem)]
* frem (ref [[1](http://llvm.org/docs/LangRef.html#frem-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstFRem)]

## Bitwise instructions

* shl (ref [[1](http://llvm.org/docs/LangRef.html#shl-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstShl)]
* lshr (ref [[1](http://llvm.org/docs/LangRef.html#lshr-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstLShr)]
* ashr (ref [[1](http://llvm.org/docs/LangRef.html#ashr-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstAShr)]
* and (ref [[1](http://llvm.org/docs/LangRef.html#and-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstAnd)]
* or (ref [[1](http://llvm.org/docs/LangRef.html#or-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstOr)]
* xor (ref [[1](http://llvm.org/docs/LangRef.html#xor-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstXor)]

## Vector instructions

* extractelement (ref [[1](http://llvm.org/docs/LangRef.html#extractelement-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstExtractElement)]
* insertelement (ref [[1](http://llvm.org/docs/LangRef.html#insertelement-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstInsertElement)]
* shufflevector (ref [[1](http://llvm.org/docs/LangRef.html#shufflevector-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstShuffleVector)]

## Aggregate instructions

* extractvalue (ref [[1](http://llvm.org/docs/LangRef.html#extractvalue-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstExtractValue)]
* insertvalue (ref [[1](http://llvm.org/docs/LangRef.html#insertvalue-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstInsertValue)]

## Memory instructions

* alloca (ref [[1](http://llvm.org/docs/LangRef.html#alloca-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstAlloca)]
* load (ref [[1](http://llvm.org/docs/LangRef.html#load-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstLoad)]
* store (ref [[1](http://llvm.org/docs/LangRef.html#store-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstStore)]
* fence (ref [[1](http://llvm.org/docs/LangRef.html#fence-instruction)])
    - [x] asm
    - [ ] ir
* cmpxchg (ref [[1](http://llvm.org/docs/LangRef.html#cmpxchg-instruction)])
    - [x] asm
    - [ ] ir
* atomicrmw (ref [[1](http://llvm.org/docs/LangRef.html#atomicrmw-instruction)])
    - [x] asm
    - [ ] ir
* getelementptr (ref [[1](http://llvm.org/docs/LangRef.html#getelementptr-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstGetElementPtr)]

## Conversion instructions

* trunc (ref [[1](http://llvm.org/docs/LangRef.html#trunc-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstTrunc)]
* zext (ref [[1](http://llvm.org/docs/LangRef.html#zext-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstZExt)]
* sext (ref [[1](http://llvm.org/docs/LangRef.html#sext-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstSExt)]
* fptrunc (ref [[1](http://llvm.org/docs/LangRef.html#fptrunc-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstFPTrunc)]
* fpext (ref [[1](http://llvm.org/docs/LangRef.html#fpext-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstFPExt)]
* fptoui (ref [[1](http://llvm.org/docs/LangRef.html#fptoui-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstFPToUI)]
* fptosi (ref [[1](http://llvm.org/docs/LangRef.html#fptosi-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstFPToSI)]
* uitofp (ref [[1](http://llvm.org/docs/LangRef.html#uitofp-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstUIToFP)]
* sitofp (ref [[1](http://llvm.org/docs/LangRef.html#sitofp-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstSIToFP)]
* ptrtoint (ref [[1](http://llvm.org/docs/LangRef.html#ptrtoint-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstPtrToInt)]
* inttoptr (ref [[1](http://llvm.org/docs/LangRef.html#inttoptr-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstIntToPtr)]
* bitcast (ref [[1](http://llvm.org/docs/LangRef.html#bitcast-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstBitCast)]
* addrspacecast (ref [[1](http://llvm.org/docs/LangRef.html#addrspacecast-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstAddrSpaceCast)]

## Other instructions

* icmp (ref [[1](http://llvm.org/docs/LangRef.html#icmp-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstICmp)]
* fcmp (ref [[1](http://llvm.org/docs/LangRef.html#fcmp-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstFCmp)]
* phi (ref [[1](http://llvm.org/docs/LangRef.html#phi-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstPhi)]
* select (ref [[1](http://llvm.org/docs/LangRef.html#select-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstSelect)]
* call (ref [[1](http://llvm.org/docs/LangRef.html#call-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#InstCall)]
* va_arg (ref [[1](http://llvm.org/docs/LangRef.html#va_arg-instruction)])
    - [x] asm
    - [ ] ir
* landingpad (ref [[1](http://llvm.org/docs/LangRef.html#landingpad-instruction)])
    - [x] asm
    - [ ] ir
* catchpad (ref [[1](http://llvm.org/docs/LangRef.html#catchpad-instruction)])
    - [x] asm
    - [ ] ir
* cleanuppad (ref [[1](http://llvm.org/docs/LangRef.html#cleanuppad-instruction)])
    - [x] asm
    - [ ] ir

# Terminators

* ret (ref [[1](http://llvm.org/docs/LangRef.html#ret-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#TermRet)]
* br (ref [[1](http://llvm.org/docs/LangRef.html#br-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#TermBr), [1](https://godoc.org/github.com/llir/llvm/ir#TermCondBr)]
* switch (ref [[1](http://llvm.org/docs/LangRef.html#switch-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#TermSwitch)]
* indirectbr (ref [[1](http://llvm.org/docs/LangRef.html#indirectbr-instruction)])
    - [x] asm
    - [ ] ir
* invoke (ref [[1](http://llvm.org/docs/LangRef.html#invoke-instruction)])
    - [x] asm
    - [ ] ir
* resume (ref [[1](http://llvm.org/docs/LangRef.html#resume-instruction)])
    - [x] asm
    - [ ] ir
* catchswitch (ref [[1](http://llvm.org/docs/LangRef.html#catchswitch-instruction)])
    - [x] asm
    - [ ] ir
* catchret (ref [[1](http://llvm.org/docs/LangRef.html#catchret-instruction)])
    - [x] asm
    - [ ] ir
* cleanupret (ref [[1](http://llvm.org/docs/LangRef.html#cleanupret-instruction)])
    - [x] asm
    - [ ] ir
* unreachable (ref [[1](http://llvm.org/docs/LangRef.html#unreachable-instruction)])
    - [x] asm
    - [x] ir [[1](https://godoc.org/github.com/llir/llvm/ir#TermUnreachable)]
