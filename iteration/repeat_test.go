package iteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	want := "aaaaa"

	got := Repeat("a")

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}

func TestRepeatParameterized(t *testing.T) {
	want := "aaa"

	got := Repeat("a", 3)

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}
