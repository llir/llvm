// Note, the LLVM IR metadata decoder implementation of this package is heavily
// inspired by encoding/json; which is governed by a BSD license.

package metadata

import (
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

// Unmarshaler is the interface implemented by types that can unmarshal an LLVM
// IR metadata description of themselves.
type Unmarshaler interface {
	// UnmarshalMetadata unmarshals the metadata node into the value.
	UnmarshalMetadata(node Node) error
}

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

// value decodes the metadata node into the value.
func (d *decoder) value(node Node, v reflect.Value) error {
	u, rv := d.indirect(v)
	if u != nil {
		if err := u.UnmarshalMetadata(node); err != nil {
			return errors.WithStack(err)
		}
		return nil
	}
	switch kind := rv.Type().Kind(); kind {
	case reflect.String:
		if err := d.string(node, rv); err != nil {
			return errors.WithStack(err)
		}
		return nil
	default:
		panic(fmt.Errorf("support for decoding into %T (%v) not yet implemented", kind, kind))
	}
}

// string decodes the metadata string into the value.
func (d *decoder) string(node Node, v reflect.Value) error {
	n, ok := node.(*Metadata)
	if !ok {
		return errors.Errorf("invalid metadata node type; expected *metadata.Metadata, got %T", node)
	}
	if len(n.Nodes) != 1 {
		return errors.Errorf("invalid number of metadata nodes; expected 1, got %d", len(n.Nodes))
	}
	s, ok := n.Nodes[0].(*String)
	if !ok {
		return errors.Errorf("invalid metadata string type; expected *metadata.String, got %T", n.Nodes[0])
	}
	v.SetString(s.Val)
	return nil
}

// indirect walks down v allocating pointers as needed, until it gets to a non-
// pointer. if it encounters an Unmarshaler, indirect stops and returns that.
func (d *decoder) indirect(v reflect.Value) (Unmarshaler, reflect.Value) {
	// Note, the indirect method has been copied (with minor modifications) from
	// go/src/encoding/json/decode.go.

	// If v is a named type and is addressable, start with its address, so that
	// if the type has pointer methods, we find them.
	if v.Kind() != reflect.Ptr && v.Type().Name() != "" && v.CanAddr() {
		v = v.Addr()
	}
	for {
		// Load value from interface, but only if the result will be usefully
		// addressable.
		if v.Kind() == reflect.Interface && !v.IsNil() {
			e := v.Elem()
			if e.Kind() == reflect.Ptr && !e.IsNil() && e.Elem().Kind() == reflect.Ptr {
				v = e
				continue
			}
		}

		if v.Kind() != reflect.Ptr {
			break
		}

		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		if v.Type().NumMethod() > 0 {
			if u, ok := v.Interface().(Unmarshaler); ok {
				return u, reflect.Value{}
			}
		}
		v = v.Elem()
	}
	return nil, v
}
