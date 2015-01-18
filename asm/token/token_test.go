package token

import "testing"

func TestKeywords(t *testing.T) {
	m := make(map[Kind]bool)
	for _, kind := range Keywords {
		if kind != Type && !kind.IsKeyword() {
			t.Errorf("%v incorrectly classified as keyword", kind)
			continue
		}
		m[kind] = true
	}
	for kind := keywordStart + 1; kind < keywordEnd; kind++ {
		if kind == instructionStart || kind == instructionEnd {
			continue
		}
		_, ok := m[kind]
		if !ok {
			t.Errorf("Keywords map is missing %v", kind)
		}
	}
}
