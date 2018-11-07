GODIR=`echo ${GOPATH} | awk '{split($$0, arr, ":"); print arr[2]}'`
ENUM_DIR=${GODIR}/src/github.com/llir/llvm/ir/enum

all: gen

gen:
	string2enum -linecomment -type AtomicOp ${ENUM_DIR}
	string2enum -linecomment -type AtomicOrdering ${ENUM_DIR}
	string2enum -linecomment -type CallingConv ${ENUM_DIR}
	string2enum -linecomment -type ChecksumKind ${ENUM_DIR}
	string2enum -linecomment -type ClauseType ${ENUM_DIR}
	string2enum -linecomment -type DIFlag ${ENUM_DIR}
	string2enum -linecomment -type DLLStorageClass ${ENUM_DIR}
	string2enum -linecomment -type DwarfAttEncoding ${ENUM_DIR}
	string2enum -linecomment -type DwarfCC ${ENUM_DIR}
	string2enum -linecomment -type DwarfLang ${ENUM_DIR}
	string2enum -linecomment -type DwarfMacinfo ${ENUM_DIR}
	string2enum -linecomment -type DwarfOp ${ENUM_DIR}
	string2enum -linecomment -type DwarfTag ${ENUM_DIR}
	string2enum -linecomment -type DwarfVirtuality ${ENUM_DIR}
	string2enum -linecomment -type EmissionKind ${ENUM_DIR}
	string2enum -linecomment -type FastMathFlag ${ENUM_DIR}
	string2enum -linecomment -type FPred ${ENUM_DIR}
	string2enum -linecomment -type FuncAttr ${ENUM_DIR}
	string2enum -linecomment -type IPred ${ENUM_DIR}
	string2enum -linecomment -type Linkage ${ENUM_DIR}
	string2enum -linecomment -type NameTableKind ${ENUM_DIR}
	string2enum -linecomment -type OverflowFlag ${ENUM_DIR}
	string2enum -linecomment -type ParamAttr ${ENUM_DIR}
	string2enum -linecomment -type Preemption ${ENUM_DIR}
	string2enum -linecomment -type ReturnAttr ${ENUM_DIR}
	string2enum -linecomment -type SelectionKind ${ENUM_DIR}
	string2enum -linecomment -type Tail ${ENUM_DIR}
	string2enum -linecomment -type TLSModel ${ENUM_DIR}
	string2enum -linecomment -type UnnamedAddr ${ENUM_DIR}
	string2enum -linecomment -type Visibility ${ENUM_DIR}
