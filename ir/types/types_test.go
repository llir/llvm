package types

import "testing"

func TestIntTypeEqual(t *testing.T) {
	golden := []struct {
		t    *IntType
		u    *IntType
		want bool
	}{
		{
			t:    &IntType{BitSize: 32},
			u:    &IntType{BitSize: 8},
			want: false,
		},
		{
			t:    &IntType{BitSize: 32},
			u:    I32,
			want: true,
		},
		{
			t:    &IntType{BitSize: 32},
			u:    NewInt(32),
			want: true,
		},
		{
			t:    &IntType{TypeName: "foo", BitSize: 32},
			u:    &IntType{TypeName: "foo", BitSize: 8},
			want: false,
		},
		{
			t:    &IntType{TypeName: "foo", BitSize: 32},
			u:    I32,
			want: true,
		},
		{
			t:    &IntType{TypeName: "foo", BitSize: 32},
			u:    NewInt(32),
			want: true,
		},
	}
	for _, g := range golden {
		got := g.t.Equal(g.u)
		if g.want != got {
			t.Errorf("struct equality mismatch between `%s` and `%s`; expected %t, got %t", g.t.Def(), g.u.Def(), g.want, got)
		}
	}
}

func TestStructTypeEqual(t *testing.T) {
	// Identified (named) struct types are uniqued by type names, not by
	// structural identity.
	golden := []struct {
		t    *StructType
		u    *StructType
		want bool
	}{
		// Unnamed struct types.
		{
			t:    &StructType{Fields: []Type{I32}},
			u:    NewStruct(I8),
			want: false,
		},
		{
			t:    &StructType{Fields: []Type{I32}, Packed: false},
			u:    &StructType{Fields: []Type{I32}, Packed: true},
			want: false,
		},
		{
			t:    &StructType{Fields: []Type{I32}},
			u:    &StructType{Fields: []Type{I32}},
			want: true,
		},
		{
			t:    &StructType{Fields: []Type{I32}, Packed: true},
			u:    &StructType{Fields: []Type{I32}, Packed: true},
			want: true,
		},
		// Identified struct types.
		{
			t:    &StructType{TypeName: "foo", Fields: []Type{I32}},
			u:    &StructType{TypeName: "bar", Fields: []Type{I8}},
			want: false,
		},
		{
			t:    &StructType{TypeName: "foo", Fields: []Type{I32}, Packed: false},
			u:    &StructType{TypeName: "bar", Fields: []Type{I32}, Packed: true},
			want: false,
		},
		{
			t:    &StructType{TypeName: "foo", Fields: []Type{I32}},
			u:    &StructType{TypeName: "foo", Fields: []Type{I32}},
			want: true,
		},
		{
			t:    &StructType{TypeName: "foo", Fields: []Type{I32}, Packed: true},
			u:    &StructType{TypeName: "foo", Fields: []Type{I32}, Packed: true},
			want: true,
		},
		{
			t:    &StructType{TypeName: "foo", Fields: []Type{I32}},
			u:    &StructType{TypeName: "bar", Fields: []Type{I32}},
			want: false,
		},
		{
			t:    &StructType{TypeName: "foo", Fields: []Type{I32}, Packed: true},
			u:    &StructType{TypeName: "bar", Fields: []Type{I32}, Packed: true},
			want: false,
		},
	}
	for _, g := range golden {
		got := g.t.Equal(g.u)
		if g.want != got {
			t.Errorf("struct equality mismatch between `%s` and `%s`; expected %t, got %t", g.t.Def(), g.u.Def(), g.want, got)
		}
	}
}

// Assert that each type implements the types.Type interface.
var (
	_ Type = (*VoidType)(nil)
	_ Type = (*FuncType)(nil)
	_ Type = (*IntType)(nil)
	_ Type = (*FloatType)(nil)
	_ Type = (*MMXType)(nil)
	_ Type = (*PointerType)(nil)
	_ Type = (*VectorType)(nil)
	_ Type = (*LabelType)(nil)
	_ Type = (*TokenType)(nil)
	_ Type = (*MetadataType)(nil)
	_ Type = (*ArrayType)(nil)
	_ Type = (*StructType)(nil)
)
