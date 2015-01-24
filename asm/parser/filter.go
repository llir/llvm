package parser

import (
	"log"

	"github.com/mewlang/llvm/asm/token"
)

// filter filters out token types which are not yet handled by the parser.
func filter(tokens []token.Token) []token.Token {
	subset := make([]token.Token, 0, len(tokens))
	for _, token := range tokens {
		if valid[token.Kind] {
			subset = append(subset, token)
		} else {
			log.Printf("filter: token type %v not yet handled by the parser.\n", token.Kind)
		}
	}
	return subset
}

// valid specifies the subset of tokens which the parser is currently able to
// handle.
var valid = map[token.Kind]bool{
	// Special tokens.
	token.EOF: true,

	// Top-level entities.
	token.KwDeclare: true, // declare
	token.KwDefine:  true, // define

	// Instructions.
	// Terminator instructions.
	token.KwRet:         true, // ret
	token.KwBr:          true, // br
	token.KwSwitch:      true, // switch
	token.KwIndirectbr:  true, // indirectbr
	token.KwInvoke:      true, // invoke
	token.KwResume:      true, // resume
	token.KwUnreachable: true, // unreachable

	// Binary operations.
	token.KwAdd:  true, // add
	token.KwFadd: true, // fadd
	token.KwSub:  true, // sub
	token.KwFsub: true, // fsub
	token.KwMul:  true, // mul
	token.KwFmul: true, // fmul
	token.KwUdiv: true, // udiv
	token.KwSdiv: true, // sdiv
	token.KwFdiv: true, // fdiv
	token.KwUrem: true, // urem
	token.KwSrem: true, // srem
	token.KwFrem: true, // frem

	// Bitwise binary operations.
	token.KwShl:  true, // shl
	token.KwLshr: true, // lshr
	token.KwAshr: true, // ashr
	token.KwAnd:  true, // and
	token.KwOr:   true, // or
	token.KwXor:  true, // xor

	// Vector operations.
	token.KwExtractelement: true, // extractelement
	token.KwInsertelement:  true, // insertelement
	token.KwShufflevector:  true, // shufflevector

	// Aggregate operations.
	token.KwExtractvalue: true, // extractvalue
	token.KwInsertvalue:  true, // insertvalue

	// Memory access and addressing operations.
	token.KwAlloca:        true, // alloca
	token.KwLoad:          true, // load
	token.KwStore:         true, // store
	token.KwFence:         true, // fence
	token.KwCmpxchg:       true, // cmpxchg
	token.KwAtomicrmw:     true, // atomicrmw
	token.KwGetelementptr: true, // getelementptr

	// Conversion operations.
	token.KwTo:            true, // to
	token.KwTrunc:         true, // trunc
	token.KwZext:          true, // zext
	token.KwSext:          true, // sext
	token.KwFptrunc:       true, // fptrunc
	token.KwFpext:         true, // fpext
	token.KwFptoui:        true, // fptoui
	token.KwFptosi:        true, // fptosi
	token.KwUitofp:        true, // uitofp
	token.KwSitofp:        true, // sitofp
	token.KwPtrtoint:      true, // ptrtoint
	token.KwInttoptr:      true, // inttoptr
	token.KwBitcast:       true, // bitcast
	token.KwAddrspacecast: true, // addrspacecast

	// Other operations.
	token.KwIcmp:       true, // icmp
	token.KwFcmp:       true, // fcmp
	token.KwPhi:        true, // phi
	token.KwSelect:     true, // select
	token.KwCall:       true, // call
	token.KwVaArg:      true, // va_arg
	token.KwLandingpad: true, // landingpad
}
