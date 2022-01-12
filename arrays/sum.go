package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(matrix ...[]int) []int {
	var sums []int
	for _, numbers := range matrix {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(matrix ...[]int) []int {
	var tailSums []int
	for _, numbers := range matrix {
		if len(numbers) == 0 {
			tailSums = append(tailSums, 0)
		} else {
			tail := numbers[1:]
			tailSums = append(tailSums, Sum(tail))
		}
	}
	return tailSums
}
