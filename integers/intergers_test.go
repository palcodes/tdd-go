package integers

import "testing"

func TestIntegers(t *testing.T) {
	sum := Add(3, 5)
	expected := 8

	if sum != expected {
		t.Errorf("\n--- sum of provided integers '%d' but expected '%d'", sum, expected)
	}
}
