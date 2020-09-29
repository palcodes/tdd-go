package arraysnslices

//Sum takes an array and returns the sum of all elements
func Sum(numbers []int) int {
	sum := 0
	for _, numbers := range numbers {
		sum += numbers
	}
	return sum
}
