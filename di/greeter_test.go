package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	want := "Hello John"
	Greet(&buffer, "John")

	got := buffer.String()

	if got != want {
		t.Errorf("Got %q, expected %q", got, want)
	}
}
