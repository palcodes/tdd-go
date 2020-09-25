package helloworld

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("--- saying hello to people", func(t *testing.T) {
		got := Hello("Aayush")
		want := "Hello, Aayush"

		assertCorrectMessage(t, got, want)
	})

	t.Run("--- say Hello, world! when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
}
