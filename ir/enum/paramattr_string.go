// Code generated by "stringer -linecomment -type ParamAttr"; DO NOT EDIT.

package enum

import "strconv"

const _ParamAttr_name = "byvalinallocainregnestnoaliasnocapturenonnullreadnonereadonlyreturnedsignextsretswifterrorswiftselfwriteonlyzeroext"

var _ParamAttr_index = [...]uint8{0, 5, 13, 18, 22, 29, 38, 45, 53, 61, 69, 76, 80, 90, 99, 108, 115}

func (i ParamAttr) String() string {
	if i >= ParamAttr(len(_ParamAttr_index)-1) {
		return "ParamAttr(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ParamAttr_name[_ParamAttr_index[i]:_ParamAttr_index[i+1]]
}
