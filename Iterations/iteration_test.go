package iterations

import "testing"

func TestRepeater(t *testing.T) {
	got := Repeat("a")
	expected := "aaaaa"
	if got != expected {
		t.Errorf("expected %q, but got %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
