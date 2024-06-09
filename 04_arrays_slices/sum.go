package arraysslices

func Sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, slice := range numbersToSum {
		sums = append(sums, Sum(slice))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, slice := range numbersToSum {
		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(slice[1:]))
		}
	}

	return sums
}
