package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 10, 4}
	want := 17

	got := Sum(numbers)

	if got != want {
		t.Errorf("Got %d, expected %d. Input: %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	numbers1 := []int{1, 2, 10, 4}
	numbers2 := []int{1, 8}
	want := []int{17, 9}

	got := SumAll(numbers1, numbers2)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, expected %v. Input: %v", got, want, [][]int{numbers1, numbers2})
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("Sum all tails of non-empty slices", func(t *testing.T) {
		numbers1 := []int{1, 2, 10, 4}
		numbers2 := []int{1, 8}
		want := []int{16, 8}

		got := SumAllTails(numbers1, numbers2)

		assertTailSums(t, got, want, numbers1, numbers2)
	})

	t.Run("Sum all tails of empty slices", func(t *testing.T) {
		numbers1 := []int{}
		numbers2 := []int{2, 7}
		want := []int{0, 7}

		got := SumAllTails(numbers1, numbers2)

		if !reflect.DeepEqual(got, want) {
			assertTailSums(t, got, want, numbers1, numbers2)
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("Multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, 1, multiply), 6)
	})

	t.Run("Concatenation of strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, "", concatenate), "abc")
	})
}

func assertTailSums(t *testing.T, got []int, want []int, input ...[]int) {
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, expected %v. Input: %v", got, want, input)
	}
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
