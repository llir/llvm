// Code generated by "string2enum -linecomment -type DLLStorageClass /home/u/Desktop/go/src/github.com/llir/llvm/ir/enum"; DO NOT EDIT.

package enum

import "fmt"
import "github.com/llir/llvm/ir/enum"

const _DLLStorageClass_name = "nonedllexportdllimport"

var _DLLStorageClass_index = [...]uint8{0, 4, 13, 22}

func DLLStorageClassFromString(s string) enum.DLLStorageClass {
	if len(s) == 0 {
		return 0
	}
	for i := range _DLLStorageClass_index[:len(_DLLStorageClass_index)-1] {
		if s == _DLLStorageClass_name[_DLLStorageClass_index[i]:_DLLStorageClass_index[i+1]] {
			return enum.DLLStorageClass(i)
		}
	}
	panic(fmt.Errorf("unable to locate DLLStorageClass enum corresponding to %q", s))
}
