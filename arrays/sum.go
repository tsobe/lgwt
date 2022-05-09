package arrays

func Sum(numbers []int) int {
	return reduce(numbers, 0, func(num, sum int) int {
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
	var initSums []int
	return reduce(matrix, initSums, func(numbers, tailSums []int) []int {
		if len(numbers) == 0 {
			tailSums = append(tailSums, 0)
		} else {
			tail := numbers[1:]
			tailSums = append(tailSums, Sum(tail))
		}
		return tailSums
	})
}

func reduce[I any](input []I, initVal I, accumulator func(I, I) I) I {
	result := initVal
	for _, val := range input {
		result = accumulator(val, result)
	}
	return result
}
