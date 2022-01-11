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
		t.Errorf("Got %d, expected %d. Numbers: %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	numbers1 := []int{1, 2, 10, 4}
	numbers2 := []int{1, 8}
	want := []int{17, 9}

	got := SumAll(numbers1, numbers2)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, expected %v", got, want)
	}
}
