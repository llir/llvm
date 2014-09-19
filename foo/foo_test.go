package foo

import "testing"

func TestBar(t *testing.T) {
	want := 42
	got := Bar()
	if got != want {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func TestBaz(t *testing.T) {
	want := "baz"
	got := Baz()
	if got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}

func BenchmarkBar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bar()
	}
}

func BenchmarkBaz(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Baz()
	}
}
