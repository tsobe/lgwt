package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("Asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("Asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "Grace")
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %+v", got)
	}
}
