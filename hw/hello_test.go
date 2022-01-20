package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying hello to John", func(t *testing.T) {
		want := "Hello John!"

		got := Hello("John", "")

		assertMessage(t, got, want)
	})

	t.Run("Saying hello to empty string", func(t *testing.T) {
		want := "Hello World!"

		got := Hello("", "")

		assertMessage(t, got, want)
	})

	t.Run("Saying hello to John in Spanish", func(t *testing.T) {
		want := "Hola John!"

		got := Hello("John", "Spanish")

		assertMessage(t, got, want)
	})

	t.Run("Saying hello to John in French", func(t *testing.T) {
		want := "Bonjour John!"

		got := Hello("John", "French")

		assertMessage(t, got, want)
	})
}

func assertMessage(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}
