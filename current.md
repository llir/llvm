# Version 0.x to be released (master branch)

This file tracks the implementation status of various LLVM IR language concepts, as covered by the master branch of [llir/llvm](https://github.com/llir/llvm).

The `asm` checkbox tracks read support of the language concept by the [llvm/asm](https://godoc.org/github.com/llir/llvm/asm) package.

The `ir` checkbox tracks support for an in-memory representation of the language concept by the [llvm/ir](https://godoc.org/github.com/llir/llvm/ir) package.

# Modules

* Source filename (ref [LangRef.html#source-filename](http://llvm.org/docs/LangRef.html#source-filename))
    - [x] asm
    - [ ] ir
* Target specifiers (ref [LangRef.html#data-layout](http://llvm.org/docs/LangRef.html#data-layout), [LangRef.html#target-triple](http://llvm.org/docs/LangRef.html#target-triple))
    - [x] asm
    - [x] ir [ir.Module.DataLayout](https://godoc.org/github.com/llir/llvm/ir#Module.DataLayout), [ir.Module.TargetTriple](https://godoc.org/github.com/llir/llvm/ir#Module.TargetTriple)
* Module-level inline assembly (ref [LangRef.html#module-level-inline-assembly](http://llvm.org/docs/LangRef.html#module-level-inline-assembly))
    - [x] asm
    - [ ] ir
* Type definitions (ref [LangRef.html#structure-types](http://llvm.org/docs/LangRef.html#structure-types))
    - [x] asm
    - [x] ir [ir.Module.Types](https://godoc.org/github.com/llir/llvm/ir#Module.Types)
* Comdat definitions (ref [LangRef.html#comdats](http://llvm.org/docs/LangRef.html#comdats))
    - [x] asm
    - [ ] ir
* Global variables (ref [LangRef.html#global-variables](http://llvm.org/docs/LangRef.html#global-variables))
    - [x] asm
    - [x] ir [ir.Module.Globals](https://godoc.org/github.com/llir/llvm/ir#Module.Globals)
* Functions (ref [LangRef.html#functions](http://llvm.org/docs/LangRef.html#functions))
    - [x] asm
    - [x] ir [ir.Module.Funcs](https://godoc.org/github.com/llir/llvm/ir#Module.Funcs)
* Attribute group definitions (ref [LangRef.html#attribute-groups](http://llvm.org/docs/LangRef.html#attribute-groups))
    - [x] asm
    - [ ] ir
* Metadata definitions (ref [LangRef.html#metadata](http://llvm.org/docs/LangRef.html#metadata))
    - [x] asm
    - [x] ir [ir.Module.NamedMetadata](https://godoc.org/github.com/llir/llvm/ir#Module.NamedMetadata), [ir.Module.Metadata](https://godoc.org/github.com/llir/llvm/ir#Module.Metadata)

# Types

* Void type (ref [LangRef.html#void-type](http://llvm.org/docs/LangRef.html#void-type))
    - [x] asm
    - [x] ir [ir/types.VoidType](https://godoc.org/github.com/llir/llvm/ir/types#VoidType)
* Function type (ref [LangRef.html#function-type](http://llvm.org/docs/LangRef.html#function-type))
    - [x] asm
    - [x] ir [ir/types.FuncType](https://godoc.org/github.com/llir/llvm/ir/types#FuncType)
* Integer type (ref [LangRef.html#integer-type](http://llvm.org/docs/LangRef.html#integer-type))
    - [x] asm
    - [x] ir [ir/types.IntType](https://godoc.org/github.com/llir/llvm/ir/types#IntType)
* Floating-point type (ref [LangRef.html#floating-point-types](http://llvm.org/docs/LangRef.html#floating-point-types))
    - [x] asm
    - [x] ir [ir/types.FloatType](https://godoc.org/github.com/llir/llvm/ir/types#FloatType)
* x86 MMX type (ref [LangRef.html#x86-mmx-type](http://llvm.org/docs/LangRef.html#x86-mmx-type))
    - [ ] asm
    - [ ] ir
* Pointer type (ref [LangRef.html#pointer-type](http://llvm.org/docs/LangRef.html#pointer-type))
    - [x] asm
    - [x] ir [ir/types.PointerType](https://godoc.org/github.com/llir/llvm/ir/types#PointerType)
* Vector type (ref [LangRef.html#vector-type](http://llvm.org/docs/LangRef.html#vector-type))
    - [x] asm
    - [x] ir [ir/types.VectorType](https://godoc.org/github.com/llir/llvm/ir/types#VectorType)
* Label type (ref [LangRef.html#label-type](http://llvm.org/docs/LangRef.html#label-type))
    - [x] asm
    - [x] ir [ir/types.LabelType](https://godoc.org/github.com/llir/llvm/ir/types#LabelType)
* Token type (ref [LangRef.html#token-type](http://llvm.org/docs/LangRef.html#token-type))
    - [ ] asm
    - [ ] ir
* Metadata type (ref [LangRef.html#metadata-type](http://llvm.org/docs/LangRef.html#metadata-type))
    - [x] asm
    - [x] ir [ir/types.MetadataType](https://godoc.org/github.com/llir/llvm/ir/types#MetadataType)
* Array type (ref [LangRef.html#array-type](http://llvm.org/docs/LangRef.html#array-type))
    - [x] asm
    - [x] ir [ir/types.ArrayType](https://godoc.org/github.com/llir/llvm/ir/types#ArrayType)
* Struct type (ref [LangRef.html#structure-type](http://llvm.org/docs/LangRef.html#structure-type))
    - [x] asm
    - [x] ir [ir/types.StructType](https://godoc.org/github.com/llir/llvm/ir/types#StructType)

# Constants

* Integer constant (ref [LangRef.html#simple-constants](http://llvm.org/docs/LangRef.html#simple-constants))
    - [x] asm
    - [x] ir [ir/constant.Int](https://godoc.org/github.com/llir/llvm/ir/constant#Int)
* Floating-point constant (ref [LangRef.html#simple-constants](http://llvm.org/docs/LangRef.html#simple-constants))
    - [x] asm
    - [x] ir [ir/constant.Float](https://godoc.org/github.com/llir/llvm/ir/constant#Float)
* Pointer constant (ref [LangRef.html#simple-constants](http://llvm.org/docs/LangRef.html#simple-constants), [LangRef.html#global-variable-and-function-addresses](http://llvm.org/docs/LangRef.html#global-variable-and-function-addresses))
    - [x] asm
    - [x] ir [ir/constant.Null](https://godoc.org/github.com/llir/llvm/ir/constant#Null), [ir.Global](https://godoc.org/github.com/llir/llvm/ir#Global), [ir.Function](https://godoc.org/github.com/llir/llvm/ir#Function)
* Token constant (ref [LangRef.html#simple-constants](http://llvm.org/docs/LangRef.html#simple-constants))
    - [ ] asm
    - [ ] ir
* Vector constant (ref [LangRef.html#complex-constants](http://llvm.org/docs/LangRef.html#complex-constants))
    - [x] asm
    - [x] ir [ir/constant.Vector](https://godoc.org/github.com/llir/llvm/ir/constant#Vector)
* Array constant (ref [LangRef.html#complex-constants](http://llvm.org/docs/LangRef.html#complex-constants))
    - [x] asm
    - [x] ir [ir/constant.Array](https://godoc.org/github.com/llir/llvm/ir/constant#Array)
* Struct constant (ref [LangRef.html#complex-constants](http://llvm.org/docs/LangRef.html#complex-constants))
    - [x] asm
    - [x] ir [ir/constant.Struct](https://godoc.org/github.com/llir/llvm/ir/constant#Struct)
* Zero initializer constant (ref [LangRef.html#complex-constants](http://llvm.org/docs/LangRef.html#complex-constants))
    - [x] asm
    - [x] ir [ir/constant.ZeroInitializer](https://godoc.org/github.com/llir/llvm/ir/constant#ZeroInitializer)
* Undefined value constant (ref [LangRef.html#undefined-values](http://llvm.org/docs/LangRef.html#undefined-values))
    - [x] asm
    - [x] ir [ir/constant.Undef](https://godoc.org/github.com/llir/llvm/ir/constant#Undef)
* Block address constant (ref [LangRef.html#addresses-of-basic-blocks](http://llvm.org/docs/LangRef.html#addresses-of-basic-blocks))
    - [ ] asm
    - [ ] ir

# Constant expressions

## Binary expressions

* add (ref [LangRef.html#add-instruction](http://llvm.org/docs/LangRef.html#add-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprAdd](https://godoc.org/github.com/llir/llvm/ir/constant#ExprAdd)
* fadd (ref [LangRef.html#fadd-instruction](http://llvm.org/docs/LangRef.html#fadd-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprFAdd](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFAdd)
* sub (ref [LangRef.html#sub-instruction](http://llvm.org/docs/LangRef.html#sub-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprSub](https://godoc.org/github.com/llir/llvm/ir/constant#ExprSub)
* fsub (ref [LangRef.html#fsub-instruction](http://llvm.org/docs/LangRef.html#fsub-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprFSub](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFSub)
* mul (ref [LangRef.html#mul-instruction](http://llvm.org/docs/LangRef.html#mul-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprMul](https://godoc.org/github.com/llir/llvm/ir/constant#ExprMul)
* fmul (ref [LangRef.html#fmul-instruction](http://llvm.org/docs/LangRef.html#fmul-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprFMul](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFMul)
* udiv (ref [LangRef.html#udiv-instruction](http://llvm.org/docs/LangRef.html#udiv-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprUDiv](https://godoc.org/github.com/llir/llvm/ir/constant#ExprUDiv)
* sdiv (ref [LangRef.html#sdiv-instruction](http://llvm.org/docs/LangRef.html#sdiv-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprSDiv](https://godoc.org/github.com/llir/llvm/ir/constant#ExprSDiv)
* fdiv (ref [LangRef.html#fdiv-instruction](http://llvm.org/docs/LangRef.html#fdiv-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprFDiv](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFDiv)
* urem (ref [LangRef.html#urem-instruction](http://llvm.org/docs/LangRef.html#urem-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprURem](https://godoc.org/github.com/llir/llvm/ir/constant#ExprURem)
* srem (ref [LangRef.html#srem-instruction](http://llvm.org/docs/LangRef.html#srem-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprSRem](https://godoc.org/github.com/llir/llvm/ir/constant#ExprSRem)
* frem (ref [LangRef.html#frem-instruction](http://llvm.org/docs/LangRef.html#frem-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprFRem](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFRem)

## Bitwise expressions

* shl (ref [LangRef.html#shl-instruction](http://llvm.org/docs/LangRef.html#shl-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprShl](https://godoc.org/github.com/llir/llvm/ir/constant#ExprShl)
* lshr (ref [LangRef.html#lshr-instruction](http://llvm.org/docs/LangRef.html#lshr-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprLShr](https://godoc.org/github.com/llir/llvm/ir/constant#ExprLShr)
* ashr (ref [LangRef.html#ashr-instruction](http://llvm.org/docs/LangRef.html#ashr-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprAShr](https://godoc.org/github.com/llir/llvm/ir/constant#ExprAShr)
* and (ref [LangRef.html#and-instruction](http://llvm.org/docs/LangRef.html#and-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprAnd](https://godoc.org/github.com/llir/llvm/ir/constant#ExprAnd)
* or (ref [LangRef.html#or-instruction](http://llvm.org/docs/LangRef.html#or-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprOr](https://godoc.org/github.com/llir/llvm/ir/constant#ExprOr)
* xor (ref [LangRef.html#xor-instruction](http://llvm.org/docs/LangRef.html#xor-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprXor](https://godoc.org/github.com/llir/llvm/ir/constant#ExprXor)

## Vector expressions

* extractelement (ref [LangRef.html#extractelement-instruction](http://llvm.org/docs/LangRef.html#extractelement-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprExtractElement](https://godoc.org/github.com/llir/llvm/ir/constant#ExprExtractElement)
* insertelement (ref [LangRef.html#insertelement-instruction](http://llvm.org/docs/LangRef.html#insertelement-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprInsertElement](https://godoc.org/github.com/llir/llvm/ir/constant#ExprInsertElement)
* shufflevector (ref [LangRef.html#shufflevector-instruction](http://llvm.org/docs/LangRef.html#shufflevector-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprShuffleVector](https://godoc.org/github.com/llir/llvm/ir/constant#ExprShuffleVector)

## Aggregate expressions

* extractvalue (ref [LangRef.html#extractvalue-instruction](http://llvm.org/docs/LangRef.html#extractvalue-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprExtractValue](https://godoc.org/github.com/llir/llvm/ir/constant#ExprExtractValue)
* insertvalue (ref [LangRef.html#insertvalue-instruction](http://llvm.org/docs/LangRef.html#insertvalue-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprInsertValue](https://godoc.org/github.com/llir/llvm/ir/constant#ExprInsertValue)

## Memory expressions

* getelementptr (ref [LangRef.html#getelementptr-instruction](http://llvm.org/docs/LangRef.html#getelementptr-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprGetElementPtr](https://godoc.org/github.com/llir/llvm/ir/constant#ExprGetElementPtr)

## Conversion expressions

* trunc (ref [LangRef.html#trunc-instruction](http://llvm.org/docs/LangRef.html#trunc-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprTrunc](https://godoc.org/github.com/llir/llvm/ir/constant#ExprTrunc)
* zext (ref [LangRef.html#zext-instruction](http://llvm.org/docs/LangRef.html#zext-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprZExt](https://godoc.org/github.com/llir/llvm/ir/constant#ExprZExt)
* sext (ref [LangRef.html#sext-instruction](http://llvm.org/docs/LangRef.html#sext-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprSExt](https://godoc.org/github.com/llir/llvm/ir/constant#ExprSExt)
* fptrunc (ref [LangRef.html#fptrunc-instruction](http://llvm.org/docs/LangRef.html#fptrunc-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprFPTrunc](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPTrunc)
* fpext (ref [LangRef.html#fpext-instruction](http://llvm.org/docs/LangRef.html#fpext-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprFPExt](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPExt)
* fptoui (ref [LangRef.html#fptoui-instruction](http://llvm.org/docs/LangRef.html#fptoui-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprFPToUI](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPToUI)
* fptosi (ref [LangRef.html#fptosi-instruction](http://llvm.org/docs/LangRef.html#fptosi-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprFPToSI](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFPToSI)
* uitofp (ref [LangRef.html#uitofp-instruction](http://llvm.org/docs/LangRef.html#uitofp-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprUIToFP](https://godoc.org/github.com/llir/llvm/ir/constant#ExprUIToFP)
* sitofp (ref [LangRef.html#sitofp-instruction](http://llvm.org/docs/LangRef.html#sitofp-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprSIToFP](https://godoc.org/github.com/llir/llvm/ir/constant#ExprSIToFP)
* ptrtoint (ref [LangRef.html#ptrtoint-instruction](http://llvm.org/docs/LangRef.html#ptrtoint-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprPtrToInt](https://godoc.org/github.com/llir/llvm/ir/constant#ExprPtrToInt)
* inttoptr (ref [LangRef.html#inttoptr-instruction](http://llvm.org/docs/LangRef.html#inttoptr-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprIntToPtr](https://godoc.org/github.com/llir/llvm/ir/constant#ExprIntToPtr)
* bitcast (ref [LangRef.html#bitcast-instruction](http://llvm.org/docs/LangRef.html#bitcast-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprBitCast](https://godoc.org/github.com/llir/llvm/ir/constant#ExprBitCast)
* addrspacecast (ref [LangRef.html#addrspacecast-instruction](http://llvm.org/docs/LangRef.html#addrspacecast-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprAddrSpaceCast](https://godoc.org/github.com/llir/llvm/ir/constant#ExprAddrSpaceCast)

## Other expressions

* icmp (ref [LangRef.html#icmp-instruction](http://llvm.org/docs/LangRef.html#icmp-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprICmp](https://godoc.org/github.com/llir/llvm/ir/constant#ExprICmp)
* fcmp (ref [LangRef.html#fcmp-instruction](http://llvm.org/docs/LangRef.html#fcmp-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprFCmp](https://godoc.org/github.com/llir/llvm/ir/constant#ExprFCmp)
* select (ref [LangRef.html#select-instruction](http://llvm.org/docs/LangRef.html#select-instruction))
    - [x] asm
    - [x] ir [ir/constant.ExprSelect](https://godoc.org/github.com/llir/llvm/ir/constant#ExprSelect)

# Global variables

Global variables (ref [[ir.Global]](https://godoc.org/github.com/llir/llvm/ir#Global))

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
    - [x] ir (ref [[ir.Global.Typ]](https://godoc.org/github.com/llir/llvm/ir#Global.Typ))
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
    - [x] ir (ref [[ir.Global.Metadata]](https://godoc.org/github.com/llir/llvm/ir#Global.Metadata))

# Functions

Functions (ref [[ir.Function]](https://godoc.org/github.com/llir/llvm/ir#Function))

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
    - [x] ir (ref [[ir.Function.CallConv]](https://godoc.org/github.com/llir/llvm/ir#Function.CallConv))
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
    - [x] ir (ref [[ir.Function.Metadata]](https://godoc.org/github.com/llir/llvm/ir#Function.Metadata))

# Instructions

## Binary instructions

* add (ref [LangRef.html#add-instruction](http://llvm.org/docs/LangRef.html#add-instruction))
    - [x] asm
    - [x] ir [ir.InstAdd](https://godoc.org/github.com/llir/llvm/ir#InstAdd)
* fadd (ref [LangRef.html#fadd-instruction](http://llvm.org/docs/LangRef.html#fadd-instruction))
    - [x] asm
    - [x] ir [ir.InstFAdd](https://godoc.org/github.com/llir/llvm/ir#InstFAdd)
* sub (ref [LangRef.html#sub-instruction](http://llvm.org/docs/LangRef.html#sub-instruction))
    - [x] asm
    - [x] ir [ir.InstSub](https://godoc.org/github.com/llir/llvm/ir#InstSub)
* fsub (ref [LangRef.html#fsub-instruction](http://llvm.org/docs/LangRef.html#fsub-instruction))
    - [x] asm
    - [x] ir [ir.InstFSub](https://godoc.org/github.com/llir/llvm/ir#InstFSub)
* mul (ref [LangRef.html#mul-instruction](http://llvm.org/docs/LangRef.html#mul-instruction))
    - [x] asm
    - [x] ir [ir.InstMul](https://godoc.org/github.com/llir/llvm/ir#InstMul)
* fmul (ref [LangRef.html#fmul-instruction](http://llvm.org/docs/LangRef.html#fmul-instruction))
    - [x] asm
    - [x] ir [ir.InstFMul](https://godoc.org/github.com/llir/llvm/ir#InstFMul)
* udiv (ref [LangRef.html#udiv-instruction](http://llvm.org/docs/LangRef.html#udiv-instruction))
    - [x] asm
    - [x] ir [ir.InstUDiv](https://godoc.org/github.com/llir/llvm/ir#InstUDiv)
* sdiv (ref [LangRef.html#sdiv-instruction](http://llvm.org/docs/LangRef.html#sdiv-instruction))
    - [x] asm
    - [x] ir [ir.InstSDiv](https://godoc.org/github.com/llir/llvm/ir#InstSDiv)
* fdiv (ref [LangRef.html#fdiv-instruction](http://llvm.org/docs/LangRef.html#fdiv-instruction))
    - [x] asm
    - [x] ir [ir.InstFDiv](https://godoc.org/github.com/llir/llvm/ir#InstFDiv)
* urem (ref [LangRef.html#urem-instruction](http://llvm.org/docs/LangRef.html#urem-instruction))
    - [x] asm
    - [x] ir [ir.InstURem](https://godoc.org/github.com/llir/llvm/ir#InstURem)
* srem (ref [LangRef.html#srem-instruction](http://llvm.org/docs/LangRef.html#srem-instruction))
    - [x] asm
    - [x] ir [ir.InstSRem](https://godoc.org/github.com/llir/llvm/ir#InstSRem)
* frem (ref [LangRef.html#frem-instruction](http://llvm.org/docs/LangRef.html#frem-instruction))
    - [x] asm
    - [x] ir [ir.InstFRem](https://godoc.org/github.com/llir/llvm/ir#InstFRem)

## Bitwise instructions

* shl (ref [LangRef.html#shl-instruction](http://llvm.org/docs/LangRef.html#shl-instruction))
    - [x] asm
    - [x] ir [ir.InstShl](https://godoc.org/github.com/llir/llvm/ir#InstShl)
* lshr (ref [LangRef.html#lshr-instruction](http://llvm.org/docs/LangRef.html#lshr-instruction))
    - [x] asm
    - [x] ir [ir.InstLShr](https://godoc.org/github.com/llir/llvm/ir#InstLShr)
* ashr (ref [LangRef.html#ashr-instruction](http://llvm.org/docs/LangRef.html#ashr-instruction))
    - [x] asm
    - [x] ir [ir.InstAShr](https://godoc.org/github.com/llir/llvm/ir#InstAShr)
* and (ref [LangRef.html#and-instruction](http://llvm.org/docs/LangRef.html#and-instruction))
    - [x] asm
    - [x] ir [ir.InstAnd](https://godoc.org/github.com/llir/llvm/ir#InstAnd)
* or (ref [LangRef.html#or-instruction](http://llvm.org/docs/LangRef.html#or-instruction))
    - [x] asm
    - [x] ir [ir.InstOr](https://godoc.org/github.com/llir/llvm/ir#InstOr)
* xor (ref [LangRef.html#xor-instruction](http://llvm.org/docs/LangRef.html#xor-instruction))
    - [x] asm
    - [x] ir [ir.InstXor](https://godoc.org/github.com/llir/llvm/ir#InstXor)

## Vector instructions

* extractelement (ref [LangRef.html#extractelement-instruction](http://llvm.org/docs/LangRef.html#extractelement-instruction))
    - [x] asm
    - [x] ir [ir.InstExtractElement](https://godoc.org/github.com/llir/llvm/ir#InstExtractElement)
* insertelement (ref [LangRef.html#insertelement-instruction](http://llvm.org/docs/LangRef.html#insertelement-instruction))
    - [x] asm
    - [x] ir [ir.InstInsertElement](https://godoc.org/github.com/llir/llvm/ir#InstInsertElement)
* shufflevector (ref [LangRef.html#shufflevector-instruction](http://llvm.org/docs/LangRef.html#shufflevector-instruction))
    - [x] asm
    - [x] ir [ir.InstShuffleVector](https://godoc.org/github.com/llir/llvm/ir#InstShuffleVector)

## Aggregate instructions

* extractvalue (ref [LangRef.html#extractvalue-instruction](http://llvm.org/docs/LangRef.html#extractvalue-instruction))
    - [x] asm
    - [x] ir [ir.InstExtractValue](https://godoc.org/github.com/llir/llvm/ir#InstExtractValue)
* insertvalue (ref [LangRef.html#insertvalue-instruction](http://llvm.org/docs/LangRef.html#insertvalue-instruction))
    - [x] asm
    - [x] ir [ir.InstInsertValue](https://godoc.org/github.com/llir/llvm/ir#InstInsertValue)

## Memory instructions

* alloca (ref [LangRef.html#alloca-instruction](http://llvm.org/docs/LangRef.html#alloca-instruction))
    - [x] asm
    - [x] ir [ir.InstAlloca](https://godoc.org/github.com/llir/llvm/ir#InstAlloca)
* load (ref [LangRef.html#load-instruction](http://llvm.org/docs/LangRef.html#load-instruction))
    - [x] asm
    - [x] ir [ir.InstLoad](https://godoc.org/github.com/llir/llvm/ir#InstLoad)
* store (ref [LangRef.html#store-instruction](http://llvm.org/docs/LangRef.html#store-instruction))
    - [x] asm
    - [x] ir [ir.InstStore](https://godoc.org/github.com/llir/llvm/ir#InstStore)
* fence (ref [LangRef.html#fence-instruction](http://llvm.org/docs/LangRef.html#fence-instruction))
    - [x] asm
    - [ ] ir
* cmpxchg (ref [LangRef.html#cmpxchg-instruction](http://llvm.org/docs/LangRef.html#cmpxchg-instruction))
    - [x] asm
    - [ ] ir
* atomicrmw (ref [LangRef.html#atomicrmw-instruction](http://llvm.org/docs/LangRef.html#atomicrmw-instruction))
    - [x] asm
    - [ ] ir
* getelementptr (ref [LangRef.html#getelementptr-instruction](http://llvm.org/docs/LangRef.html#getelementptr-instruction))
    - [x] asm
    - [x] ir [ir.InstGetElementPtr](https://godoc.org/github.com/llir/llvm/ir#InstGetElementPtr)

## Conversion instructions

* trunc (ref [LangRef.html#trunc-instruction](http://llvm.org/docs/LangRef.html#trunc-instruction))
    - [x] asm
    - [x] ir [ir.InstTrunc](https://godoc.org/github.com/llir/llvm/ir#InstTrunc)
* zext (ref [LangRef.html#zext-instruction](http://llvm.org/docs/LangRef.html#zext-instruction))
    - [x] asm
    - [x] ir [ir.InstZExt](https://godoc.org/github.com/llir/llvm/ir#InstZExt)
* sext (ref [LangRef.html#sext-instruction](http://llvm.org/docs/LangRef.html#sext-instruction))
    - [x] asm
    - [x] ir [ir.InstSExt](https://godoc.org/github.com/llir/llvm/ir#InstSExt)
* fptrunc (ref [LangRef.html#fptrunc-instruction](http://llvm.org/docs/LangRef.html#fptrunc-instruction))
    - [x] asm
    - [x] ir [ir.InstFPTrunc](https://godoc.org/github.com/llir/llvm/ir#InstFPTrunc)
* fpext (ref [LangRef.html#fpext-instruction](http://llvm.org/docs/LangRef.html#fpext-instruction))
    - [x] asm
    - [x] ir [ir.InstFPExt](https://godoc.org/github.com/llir/llvm/ir#InstFPExt)
* fptoui (ref [LangRef.html#fptoui-instruction](http://llvm.org/docs/LangRef.html#fptoui-instruction))
    - [x] asm
    - [x] ir [ir.InstFPToUI](https://godoc.org/github.com/llir/llvm/ir#InstFPToUI)
* fptosi (ref [LangRef.html#fptosi-instruction](http://llvm.org/docs/LangRef.html#fptosi-instruction))
    - [x] asm
    - [x] ir [ir.InstFPToSI](https://godoc.org/github.com/llir/llvm/ir#InstFPToSI)
* uitofp (ref [LangRef.html#uitofp-instruction](http://llvm.org/docs/LangRef.html#uitofp-instruction))
    - [x] asm
    - [x] ir [ir.InstUIToFP](https://godoc.org/github.com/llir/llvm/ir#InstUIToFP)
* sitofp (ref [LangRef.html#sitofp-instruction](http://llvm.org/docs/LangRef.html#sitofp-instruction))
    - [x] asm
    - [x] ir [ir.InstSIToFP](https://godoc.org/github.com/llir/llvm/ir#InstSIToFP)
* ptrtoint (ref [LangRef.html#ptrtoint-instruction](http://llvm.org/docs/LangRef.html#ptrtoint-instruction))
    - [x] asm
    - [x] ir [ir.InstPtrToInt](https://godoc.org/github.com/llir/llvm/ir#InstPtrToInt)
* inttoptr (ref [LangRef.html#inttoptr-instruction](http://llvm.org/docs/LangRef.html#inttoptr-instruction))
    - [x] asm
    - [x] ir [ir.InstIntToPtr](https://godoc.org/github.com/llir/llvm/ir#InstIntToPtr)
* bitcast (ref [LangRef.html#bitcast-instruction](http://llvm.org/docs/LangRef.html#bitcast-instruction))
    - [x] asm
    - [x] ir [ir.InstBitCast](https://godoc.org/github.com/llir/llvm/ir#InstBitCast)
* addrspacecast (ref [LangRef.html#addrspacecast-instruction](http://llvm.org/docs/LangRef.html#addrspacecast-instruction))
    - [x] asm
    - [x] ir [ir.InstAddrSpaceCast](https://godoc.org/github.com/llir/llvm/ir#InstAddrSpaceCast)

## Other instructions

* icmp (ref [LangRef.html#icmp-instruction](http://llvm.org/docs/LangRef.html#icmp-instruction))
    - [x] asm
    - [x] ir [ir.InstICmp](https://godoc.org/github.com/llir/llvm/ir#InstICmp)
* fcmp (ref [LangRef.html#fcmp-instruction](http://llvm.org/docs/LangRef.html#fcmp-instruction))
    - [x] asm
    - [x] ir [ir.InstFCmp](https://godoc.org/github.com/llir/llvm/ir#InstFCmp)
* phi (ref [LangRef.html#phi-instruction](http://llvm.org/docs/LangRef.html#phi-instruction))
    - [x] asm
    - [x] ir [ir.InstPhi](https://godoc.org/github.com/llir/llvm/ir#InstPhi)
* select (ref [LangRef.html#select-instruction](http://llvm.org/docs/LangRef.html#select-instruction))
    - [x] asm
    - [x] ir [ir.InstSelect](https://godoc.org/github.com/llir/llvm/ir#InstSelect)
* call (ref [LangRef.html#call-instruction](http://llvm.org/docs/LangRef.html#call-instruction))
    - [x] asm
    - [x] ir [ir.InstCall](https://godoc.org/github.com/llir/llvm/ir#InstCall)
* va_arg (ref [LangRef.html#va_arg-instruction](http://llvm.org/docs/LangRef.html#va_arg-instruction))
    - [x] asm
    - [ ] ir
* landingpad (ref [LangRef.html#landingpad-instruction](http://llvm.org/docs/LangRef.html#landingpad-instruction))
    - [x] asm
    - [ ] ir
* catchpad (ref [LangRef.html#catchpad-instruction](http://llvm.org/docs/LangRef.html#catchpad-instruction))
    - [x] asm
    - [ ] ir
* cleanuppad (ref [LangRef.html#cleanuppad-instruction](http://llvm.org/docs/LangRef.html#cleanuppad-instruction))
    - [x] asm
    - [ ] ir

# Terminators

* ret (ref [LangRef.html#ret-instruction](http://llvm.org/docs/LangRef.html#ret-instruction))
    - [x] asm
    - [x] ir [ir.TermRet](https://godoc.org/github.com/llir/llvm/ir#TermRet)
* br (ref [LangRef.html#br-instruction](http://llvm.org/docs/LangRef.html#br-instruction))
    - [x] asm
    - [x] ir [ir.TermBr](https://godoc.org/github.com/llir/llvm/ir#TermBr), [ir.TermCondBr](https://godoc.org/github.com/llir/llvm/ir#TermCondBr)
* switch (ref [LangRef.html#switch-instruction](http://llvm.org/docs/LangRef.html#switch-instruction))
    - [x] asm
    - [x] ir [ir.TermSwitch](https://godoc.org/github.com/llir/llvm/ir#TermSwitch)
* indirectbr (ref [LangRef.html#indirectbr-instruction](http://llvm.org/docs/LangRef.html#indirectbr-instruction))
    - [x] asm
    - [ ] ir
* invoke (ref [LangRef.html#invoke-instruction](http://llvm.org/docs/LangRef.html#invoke-instruction))
    - [x] asm
    - [ ] ir
* resume (ref [LangRef.html#resume-instruction](http://llvm.org/docs/LangRef.html#resume-instruction))
    - [x] asm
    - [ ] ir
* catchswitch (ref [LangRef.html#catchswitch-instruction](http://llvm.org/docs/LangRef.html#catchswitch-instruction))
    - [x] asm
    - [ ] ir
* catchret (ref [LangRef.html#catchret-instruction](http://llvm.org/docs/LangRef.html#catchret-instruction))
    - [x] asm
    - [ ] ir
* cleanupret (ref [LangRef.html#cleanupret-instruction](http://llvm.org/docs/LangRef.html#cleanupret-instruction))
    - [x] asm
    - [ ] ir
* unreachable (ref [LangRef.html#unreachable-instruction](http://llvm.org/docs/LangRef.html#unreachable-instruction))
    - [x] asm
    - [x] ir [[ir.TermUnreachable](https://godoc.org/github.com/llir/llvm/ir#TermUnreachable)]
