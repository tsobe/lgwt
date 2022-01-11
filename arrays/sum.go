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
