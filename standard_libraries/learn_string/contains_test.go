package learnstring

import (
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	str := "Go is so cool!"

	assertCorrectMessage := func (t testing.TB, got, want bool)  {
		t.Helper()

		if (got != want) {
			t.Fatalf("got %t, but want %t", got, want)
		}
	}

	t.Run("Should return true if string is equal to substring", func(t *testing.T) {
		got := strings.Contains(str, str)
		want := true

		assertCorrectMessage(t, got, want)
	})

	t.Run("Should return true if string contains substring", func(t *testing.T) {
		got := strings.Contains(str, "cool")
		want := true

		assertCorrectMessage(t, got, want)
	})

	t.Run("Should return false if substring have different case", func(t *testing.T) {
		got := strings.Contains(str, "Cool")
		want := false

		assertCorrectMessage(t, got, want)
	})

	t.Run("Should return false if substring contains spaces", func(t *testing.T) {
		got := strings.Contains(str, "  cool ")
		want := false

		assertCorrectMessage(t, got, want)
	})
}