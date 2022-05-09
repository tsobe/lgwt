package arrays

import (
	"strings"
	"testing"
)

type Person struct {
	Name string
}

func TestFind(t *testing.T) {
	t.Run("Find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})
		assertTrue(t, found)
		assertEqual(t, firstEvenNumber, 2)
	})
	t.Run("Find the best programmer", func(t *testing.T) {
		people := []Person{
			{Name: "Kent Beck"},
			{Name: "Martin Fowler"},
			{Name: "Chris James"},
		}

		king, found := Find(people, func(p Person) bool {
			return strings.Contains(p.Name, "Chris")
		})

		assertTrue(t, found)
		assertEqual(t, king, Person{Name: "Chris James"})
	})
}

func assertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}
