package arraysnslices

import "testing"

func TestSum(t *testing.T) {

	numbers := []int{12, 241, 64, 76, 89, 51}

	got := Sum(numbers)
	want := 533

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}

}
