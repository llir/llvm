package types

import "testing"

func TestVoidString(t *testing.T) {
	const want = "void"
	typ := NewVoid()
	got := typ.String()
	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestIntString(t *testing.T) {
	golden := []struct {
		want string
		n    int
		err  string
	}{
		{want: "i1", n: 1},
		{want: "i32", n: 32},
		{want: "", n: -1, err: "invalid integer bit width (-1)"},
		{want: "", n: 0, err: "invalid integer bit width (0)"},
		{want: "", n: 1 << 23, err: "invalid integer bit width (8388608)"},
		{want: "i8388607", n: 1<<23 - 1},
	}

	for i, g := range golden {
		typ, err := NewInt(g.n)
		if err != nil {
			if err.Error() != g.err {
				t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			}
			continue
		}
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestFloatString(t *testing.T) {
	golden := []struct {
		want string
		kind FloatKind
	}{
		{want: "half", kind: Float16},
		{want: "float", kind: Float32},
		{want: "double", kind: Float64},
		{want: "fp128", kind: Float128},
		{want: "x86_fp80", kind: X86Float80},
		{want: "ppc_fp128", kind: PPCFloat128},
		{want: "<unknown float type>", kind: -1},
	}

	for i, g := range golden {
		typ := NewFloat(g.kind)
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestMMXString(t *testing.T) {
	const want = "x86_mmx"
	typ := NewMMX()
	got := typ.String()
	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestLabelString(t *testing.T) {
	const want = "label"
	typ := NewLabel()
	got := typ.String()
	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestMetadataString(t *testing.T) {
	const want = "metadata"
	typ := NewMetadata()
	got := typ.String()
	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestFuncString(t *testing.T) {
	voidTyp := NewVoid()
	i8Typ, err := NewInt(8)
	if err != nil {
		t.Fatal(err)
	}
	i32Typ, err := NewInt(32)
	if err != nil {
		t.Fatal(err)
	}
	i8PtrTyp, err := NewPointer(i8Typ)
	if err != nil {
		t.Fatal(err)
	}
	funcTyp, err := NewFunc(i32Typ, []Type{i32Typ}, false) // i32 (i32)
	if err != nil {
		t.Fatal(err)
	}
	labelTyp := NewLabel()

	golden := []struct {
		want     string
		result   Type
		params   []Type
		variadic bool
		err      string
	}{
		// i: 0
		{want: "void ()", result: voidTyp, params: nil},
		// i: 1
		{want: "i32 (i32)", result: i32Typ, params: []Type{i32Typ}},
		// i: 2
		{want: "void (i32, i8)", result: voidTyp, params: []Type{i32Typ, i8Typ}},
		// i: 3
		{want: "i32 (i8*, ...)", result: i32Typ, params: []Type{i8PtrTyp}, variadic: true},
		// i: 4
		// i32 (void)
		{want: "", result: i32Typ, params: []Type{voidTyp}, err: "invalid function parameter type; void type only allowed for function results"},
		// i: 5
		// i32 (i32 (i32))
		{want: "", result: i32Typ, params: []Type{funcTyp}, err: `invalid function parameter type "i32 (i32)"`},
		// i: 6
		// label ()
		{want: "", result: labelTyp, params: nil, err: `invalid result parameter type "label"`},
	}

	for i, g := range golden {
		typ, err := NewFunc(g.result, g.params, g.variadic)
		if err != nil {
			if err.Error() != g.err {
				t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			}
			continue
		}
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestPointerString(t *testing.T) {
	voidTyp := NewVoid()
	i8Typ, err := NewInt(8)
	if err != nil {
		t.Fatal(err)
	}
	i32Typ, err := NewInt(32)
	if err != nil {
		t.Fatal(err)
	}
	i8PtrTyp, err := NewPointer(i8Typ)
	if err != nil {
		t.Fatal(err)
	}
	funcTyp, err := NewFunc(i32Typ, []Type{i32Typ}, false) // i32 (i32)
	if err != nil {
		t.Fatal(err)
	}
	labelTyp := NewLabel()

	golden := []struct {
		want string
		elem Type
		err  string
	}{
		{want: "i32*", elem: i32Typ},
		{want: "i32 (i32)*", elem: funcTyp},
		{want: "i8**", elem: i8PtrTyp},
		{want: "i8**", elem: i8PtrTyp},
		// void*
		{want: "", elem: voidTyp, err: `invalid pointer to "void"; use i8* instead`},
		{want: "", elem: labelTyp, err: `invalid pointer to "label"`},
	}

	for i, g := range golden {
		typ, err := NewPointer(g.elem)
		if err != nil {
			if err.Error() != g.err {
				t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			}
			continue
		}
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestVectorString(t *testing.T) {
	voidTyp := NewVoid()
	i8Typ, err := NewInt(8)
	if err != nil {
		t.Fatal(err)
	}
	i32Typ, err := NewInt(32)
	if err != nil {
		t.Fatal(err)
	}
	i8PtrTyp, err := NewPointer(i8Typ)
	if err != nil {
		t.Fatal(err)
	}
	funcTyp, err := NewFunc(i32Typ, []Type{i32Typ}, false) // i32 (i32)
	if err != nil {
		t.Fatal(err)
	}
	labelTyp := NewLabel()
	mmxTyp := NewMMX()

	golden := []struct {
		want string
		elem Type
		n    int
		err  string
	}{
		{want: "<1 x i32>", elem: i32Typ, n: 1},
		{want: "<5 x i8*>", elem: i8PtrTyp, n: 5},
		{want: "<10 x i8>", elem: i8Typ, n: 10},
		// <-1 x i8>
		{want: "", elem: i8Typ, n: -1, err: "invalid vector length (-1)"},
		// <5 x void>
		{want: "", elem: voidTyp, n: 5, err: "invalid vector element type; void type only allowed for function results"},
		// <3 x label>
		{want: "", elem: labelTyp, n: 3, err: `invalid vector element type "label"`},
		// <7 x label>
		{want: "", elem: mmxTyp, n: 7, err: `invalid vector element type "x86_mmx"`},
		// <2 x i32 (i32)>
		{want: "", elem: funcTyp, n: 2, err: `invalid vector element type "i32 (i32)"`},
	}

	for i, g := range golden {
		typ, err := NewVector(g.elem, g.n)
		if err != nil {
			if err.Error() != g.err {
				t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			}
			continue
		}
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestArrayString(t *testing.T) {
	voidTyp := NewVoid()
	i8Typ, err := NewInt(8)
	if err != nil {
		t.Fatal(err)
	}
	i32Typ, err := NewInt(32)
	if err != nil {
		t.Fatal(err)
	}
	i8PtrTyp, err := NewPointer(i8Typ)
	if err != nil {
		t.Fatal(err)
	}
	funcTyp, err := NewFunc(i32Typ, []Type{i32Typ}, false) // i32 (i32)
	if err != nil {
		t.Fatal(err)
	}
	labelTyp := NewLabel()
	mmxTyp := NewMMX()

	golden := []struct {
		want string
		elem Type
		n    int
		err  string
	}{
		{want: "[1 x i32]", elem: i32Typ, n: 1},
		{want: "[5 x i8*]", elem: i8PtrTyp, n: 5},
		{want: "[10 x i8]", elem: i8Typ, n: 10},
		// [-1 x i8]
		{want: "", elem: i8Typ, n: -1, err: "invalid array length (-1)"},
		// [5 x void]
		{want: "", elem: voidTyp, n: 5, err: "invalid array element type; void type only allowed for function results"},
		// [3 x label]
		{want: "", elem: labelTyp, n: 3, err: `invalid array element type "label"`},
		{want: "[7 x x86_mmx]", elem: mmxTyp, n: 7},
		// [2 x i32 (i32)]
		{want: "", elem: funcTyp, n: 2, err: `invalid array element type "i32 (i32)"`},
	}

	for i, g := range golden {
		typ, err := NewArray(g.elem, g.n)
		if err != nil {
			if err.Error() != g.err {
				t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			}
			continue
		}
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}

func TestStructString(t *testing.T) {
	voidTyp := NewVoid()
	i8Typ, err := NewInt(8)
	if err != nil {
		t.Fatal(err)
	}
	i32Typ, err := NewInt(32)
	if err != nil {
		t.Fatal(err)
	}
	i8PtrTyp, err := NewPointer(i8Typ)
	if err != nil {
		t.Fatal(err)
	}
	funcTyp, err := NewFunc(i32Typ, []Type{i32Typ}, false) // i32 (i32)
	if err != nil {
		t.Fatal(err)
	}
	labelTyp := NewLabel()
	mmxTyp := NewMMX()

	golden := []struct {
		want   string
		fields []Type
		packed bool
		err    string
	}{
		{want: "{i32}", fields: []Type{i32Typ}},
		{want: "{i8, i8}", fields: []Type{i8Typ, i8Typ}},
		// {void}
		{want: "", fields: []Type{voidTyp}, err: "invalid structure field type; void type only allowed for function results"},
		{want: "", fields: []Type{funcTyp}, err: `invalid structure field type "i32 (i32)"`},
		{want: "<{i8*}>", fields: []Type{i8PtrTyp}, packed: true},
		// {label}
		{want: "", fields: []Type{labelTyp}, err: `invalid structure field type "label"`},
		{want: "{x86_mmx}", fields: []Type{mmxTyp}},
	}

	for i, g := range golden {
		typ, err := NewStruct(g.fields, g.packed)
		if err != nil {
			if err.Error() != g.err {
				t.Errorf("i=%d: error mismatch; expected %v, got %v", i, g.err, err)
			}
			continue
		}
		got := typ.String()
		if got != g.want {
			t.Errorf("i=%d: string mismatch; expected %v, got %v", i, g.want, got)
		}
	}
}
