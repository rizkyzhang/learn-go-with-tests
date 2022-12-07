package learnstring

import (
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	str := "Go is so cool!"
		
	assertCorrectMessage := func (t testing.TB, got, want int)  {
		t.Helper()

		if (got != want) {
			t.Fatalf("got %d, but want %d", got, want)
		}
	}

	t.Run("Should return 0 given 'wow'", func(t *testing.T) {
		got := strings.Count(str, "wow")
		want := 0

		assertCorrectMessage(t, got, want)
	})

	t.Run("Should return 1 given 'cool'", func(t *testing.T) {
		got := strings.Count(str, "cool")
		want := 1

		assertCorrectMessage(t, got, want)
	})

	t.Run("Should return 2 given 's'", func(t *testing.T) {
		got := strings.Count(str, "s")
		want := 2

		assertCorrectMessage(t, got, want)
	})

	t.Run("Should return 3 given ' '", func(t *testing.T) {
		got := strings.Count(str, " ")
		want := 3

		assertCorrectMessage(t, got, want)
	})
}