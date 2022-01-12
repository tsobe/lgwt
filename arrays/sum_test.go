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
	assertTailSums := func(t *testing.T, got []int, want []int, input ...[]int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v, expected %v. Input: %v", got, want, input)
		}
	}

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
