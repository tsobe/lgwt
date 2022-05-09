package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("Asserting on integers", func(t *testing.T) {
		assertEqual(t, 1, 1)
		assertNotEqual(t, 1, 2)
	})

	t.Run("Asserting on strings", func(t *testing.T) {
		assertEqual(t, "hello", "hello")
		assertNotEqual(t, "hello", "Grace")
	})
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		stack := new(Stack[int])

		// check stack is empty
		assertTrue(t, stack.IsEmpty())

		// add a thing, then check it's not empty
		stack.Push(123)
		assertFalse(t, stack.IsEmpty())

		// add another thing, pop it back again
		stack.Push(456)
		value, _ := stack.Pop()
		assertEqual(t, value, 456)
		value, _ = stack.Pop()
		assertEqual(t, value, 123)
		assertTrue(t, stack.IsEmpty())
	})
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func assertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %+v", got)
	}
}

func assertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func assertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
