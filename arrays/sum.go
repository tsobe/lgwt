package arrays

func Sum(numbers []int) int {
	return Reduce(numbers, 0, func(num, sum int) int {
		return num + sum
	})
}

func SumAll(matrix ...[]int) []int {
	var sums []int
	for _, numbers := range matrix {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(matrix ...[]int) []int {
	accumulator := func(tailSums, numbers []int) []int {
		if len(numbers) == 0 {
			tailSums = append(tailSums, 0)
		} else {
			tail := numbers[1:]
			tailSums = append(tailSums, Sum(tail))
		}
		return tailSums
	}
	return Reduce(matrix, []int{}, accumulator)
}

func Reduce[A any, B any](input []A, initVal B, accumulator func(B, A) B) B {
	result := initVal
	for _, val := range input {
		result = accumulator(result, val)
	}
	return result
}
