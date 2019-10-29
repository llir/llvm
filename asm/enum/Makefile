TOOLS_DIR=../../tools
ENUM_DIR=../../ir/enum
TYPES_DIR=../../ir/types

all: gen

${TOOLS_DIR}/string2enum:
	make -C ${TOOLS_DIR}

gen: ${TOOLS_DIR}/string2enum
	${TOOLS_DIR}/string2enum -linecomment -type AtomicOp ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type AtomicOrdering ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type CallingConv ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type ChecksumKind ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type ClauseType ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type DIFlag ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type DISPFlag ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type DLLStorageClass ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type DwarfAttEncoding ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type DwarfCC ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type DwarfLang ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type DwarfMacinfo ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type DwarfOp ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type DwarfTag ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type DwarfVirtuality ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type EmissionKind ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type FastMathFlag ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type FPred ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type FuncAttr ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type IPred ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type Linkage ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type NameTableKind ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type OverflowFlag ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type ParamAttr ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type Preemption ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type ReturnAttr ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type SelectionKind ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type Tail ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type TLSModel ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type UnnamedAddr ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type Visibility ${ENUM_DIR}
	${TOOLS_DIR}/string2enum -linecomment -type FloatKind -pkg enum ${TYPES_DIR}
	goimports -w *.go

clean:
	$(RM) *_string.go

.PHONY: all gen clean
