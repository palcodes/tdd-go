package arraysnslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	numbers := []int{12, 241, 64, 76, 89, 51}

	got := Sum(numbers)
	want := 533

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 4}, []int{2, 8})

	want := []int{5, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
