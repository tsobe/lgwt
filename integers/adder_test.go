package integers

import (
	"testing"
)

func TestAdd(t *testing.T) {
	want := 7

	got := Add(2, 5)

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
