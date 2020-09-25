package helloworld

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world!"

	if got != want {
		t.Errorf("\ngot %q want %q", got, want)
	}
}
