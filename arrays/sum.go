package arrays

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbers ...[]int) []int {
	var sums []int

	for _, num := range numbers {
		sums = append(sums, Sum(num))
	}

	return sums
}

func SumAllTails(numbers ...[]int) []int {
	var sums []int

	for _, num := range numbers {
		if len(num) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(num[1:]))
		}
	}

	return sums
}
