package foo

import "testing"

func TestBar(t *testing.T) {
	want := 42
	got := Bar()
	if got != want {
		t.Errorf("expected %v, got %v", want, got)
	}
}

func BenchmarkBar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bar()
	}
}
