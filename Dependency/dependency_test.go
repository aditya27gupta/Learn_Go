package dependency

import (
	"bytes"
	"testing"
)

func AssertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"
	AssertStrings(t, got, want)
}

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	Countdown(&buffer)
	got := buffer.String()
	want := `3
2
1
Go!`
	AssertStrings(t, got, want)
}
