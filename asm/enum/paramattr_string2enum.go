// Code generated by "string2enum -linecomment -type ParamAttr ../../ir/enum"; DO NOT EDIT.

package enum

import (
	"fmt"

	"github.com/llir/llvm/ir/enum"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the string2enum command to generate them again.
	var x [1]struct{}
	_ = x[enum.ParamAttrAllocAlign-0]
	_ = x[enum.ParamAttrAllocPtr-1]
	_ = x[enum.ParamAttrImmArg-2]
	_ = x[enum.ParamAttrInReg-3]
	_ = x[enum.ParamAttrNest-4]
	_ = x[enum.ParamAttrNoAlias-5]
	_ = x[enum.ParamAttrNoCapture-6]
	_ = x[enum.ParamAttrNoFree-7]
	_ = x[enum.ParamAttrNonNull-8]
	_ = x[enum.ParamAttrNoUndef-9]
	_ = x[enum.ParamAttrReadNone-10]
	_ = x[enum.ParamAttrReadOnly-11]
	_ = x[enum.ParamAttrReturned-12]
	_ = x[enum.ParamAttrSignExt-13]
	_ = x[enum.ParamAttrSwiftAsync-14]
	_ = x[enum.ParamAttrSwiftError-15]
	_ = x[enum.ParamAttrSwiftSelf-16]
	_ = x[enum.ParamAttrWriteOnly-17]
	_ = x[enum.ParamAttrZeroExt-18]
}

const _ParamAttr_name = "allocalignallocptrimmarginregnestnoaliasnocapturenofreenonnullnoundefreadnonereadonlyreturnedsignextswiftasyncswifterrorswiftselfwriteonlyzeroext"

var _ParamAttr_index = [...]uint8{0, 10, 18, 24, 29, 33, 40, 49, 55, 62, 69, 77, 85, 93, 100, 110, 120, 129, 138, 145}

// ParamAttrFromString returns the ParamAttr enum corresponding to s.
func ParamAttrFromString(s string) enum.ParamAttr {
	if len(s) == 0 {
		return 0
	}
	for i := range _ParamAttr_index[:len(_ParamAttr_index)-1] {
		if s == _ParamAttr_name[_ParamAttr_index[i]:_ParamAttr_index[i+1]] {
			return enum.ParamAttr(i)
		}
	}
	panic(fmt.Errorf("unable to locate ParamAttr enum corresponding to %q", s))
}

func _(s string) {
	// Check for duplicate string values in type "ParamAttr".
	switch s {
	// 0
	case "allocalign":
	// 1
	case "allocptr":
	// 2
	case "immarg":
	// 3
	case "inreg":
	// 4
	case "nest":
	// 5
	case "noalias":
	// 6
	case "nocapture":
	// 7
	case "nofree":
	// 8
	case "nonnull":
	// 9
	case "noundef":
	// 10
	case "readnone":
	// 11
	case "readonly":
	// 12
	case "returned":
	// 13
	case "signext":
	// 14
	case "swiftasync":
	// 15
	case "swifterror":
	// 16
	case "swiftself":
	// 17
	case "writeonly":
	// 18
	case "zeroext":
	}
}
