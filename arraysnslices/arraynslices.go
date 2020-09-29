package arraysnslices

//Sum takes an array and returns the sum of all elements
func Sum(numbers []int) int {
	sum := 0
	for _, numbers := range numbers {
		sum += numbers
	}
	return sum
}

//SumAll takes a list of arrays and returns their sums as an array
func SumAll(numbers ...[]int) (sum []int) {
	length := len(numbers)
	sums := make([]int, length)

	for i, num := range numbers {
		sums[i] = Sum(num)
	}

	return sums
}
