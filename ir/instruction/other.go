package instruction

import "github.com/llir/llvm/ir/types"

// References:
//    http://llvm.org/docs/LangRef.html#other-operations

// TODO: Add support for the remaining other operations:
//    http://llvm.org/docs/LangRef.html#icmp-instruction
//    http://llvm.org/docs/LangRef.html#fcmp-instruction
//    http://llvm.org/docs/LangRef.html#phi-instruction
//    http://llvm.org/docs/LangRef.html#select-instruction
//    http://llvm.org/docs/LangRef.html#call-instruction
//    http://llvm.org/docs/LangRef.html#va-arg-instruction
//    http://llvm.org/docs/LangRef.html#landingpad-instruction

type ICmp struct{}

func (*ICmp) Type() types.Type { panic("ICmp.Type: not yet implemented") }
func (*ICmp) String() string   { panic("ICmp.String: not yet implemented") }

type FCmp struct{}

func (*FCmp) Type() types.Type { panic("FCmp.Type: not yet implemented") }
func (*FCmp) String() string   { panic("FCmp.String: not yet implemented") }

type PHI struct{}

func (*PHI) Type() types.Type { panic("PHI.Type: not yet implemented") }
func (*PHI) String() string   { panic("PHI.String: not yet implemented") }

type Select struct{}

func (*Select) Type() types.Type { panic("Select.Type: not yet implemented") }
func (*Select) String() string   { panic("Select.String: not yet implemented") }

type Call struct{}

func (*Call) Type() types.Type { panic("Call.Type: not yet implemented") }
func (*Call) String() string   { panic("Call.String: not yet implemented") }

type VAArg struct{}

func (*VAArg) Type() types.Type { panic("VAArg.Type: not yet implemented") }
func (*VAArg) String() string   { panic("VAArg.String: not yet implemented") }

type LandingPad struct{}

func (*LandingPad) Type() types.Type { panic("LandingPad.Type: not yet implemented") }
func (*LandingPad) String() string   { panic("LandingPad.String: not yet implemented") }

// isValueInst ensures that only instructions which return values can be
// assigned to the Value interface.
func (*ICmp) isValueInst()       {}
func (*FCmp) isValueInst()       {}
func (*PHI) isValueInst()        {}
func (*Select) isValueInst()     {}
func (*Call) isValueInst()       {}
func (*VAArg) isValueInst()      {}
func (*LandingPad) isValueInst() {}
