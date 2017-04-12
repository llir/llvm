package metadata

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

// Unmarshal parses the LLVM IR metadata node and stores the result in the value
// pointed to by v.
func Unmarshal(node Node, v interface{}) error {
	d := &decoder{}
	return d.unmarshal(node, v)
}

// A decoder tracks information required to decode LLVM IR metadata.
type decoder struct {
}

// unmarshal parses the LLVM IR metadata node and stores the result in the value
// pointed to by v.
func (d *decoder) unmarshal(node Node, v interface{}) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		t := reflect.TypeOf(v)
		if t == nil {
			return errors.New("metadata: Unmarshal(nil)")
		}
		if t.Kind() != reflect.Ptr {
			return errors.Errorf("metadata: Unmarshal(non-pointer %s)", t)
		}
		return errors.Errorf("metadata: Unmarshal(nil %s)", t)
	}
	if err := d.value(node, rv); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// value decodes a metadata node into the value.
func (d *decoder) value(node Node, v reflect.Value) error {
	if u, ok := v.Interface().(Unmarshaler); ok {
		if err := u.UnmarshalMetadata(node); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}
	panic(fmt.Errorf("support for decoding into %v not yet implemented", v.Type()))
}

// Unmarshaler is the interface implemented by types that can unmarshal an LLVM
// IR metadata description of themselves.
type Unmarshaler interface {
	// UnmarshalMetadata unmarshals the metadata node into the value.
	UnmarshalMetadata(node Node) error
}
