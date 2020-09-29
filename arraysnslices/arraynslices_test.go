package arraysnslices

import "testing"

func TestSum(t *testing.T) {
	t.Run("--- Collection of 5 definite numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("--- Collection of undefined numbers", func(t *testing.T) {
		numbers := []int{12, 241, 64, 76, 89, 51}

		got := Sum(numbers)
		want := 533

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
