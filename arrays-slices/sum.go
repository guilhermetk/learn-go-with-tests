package arraysslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sums := []int{}
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTail(numbersToSumTail ...[]int) []int {
	sums := []int{}
	for _, numbers := range numbersToSumTail {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return sums

}
