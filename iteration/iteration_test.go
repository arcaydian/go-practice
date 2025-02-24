package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeated("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("got %q want %q", repeated, expected)
	}
}

func BenchmarkTest(b *testing.B) {
	for b.Loop() {
		Repeated("a", 5)
	}
}
