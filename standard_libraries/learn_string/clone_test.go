package learnstring

import (
	"strings"
	"testing"
)

func TestClone(t *testing.T) {
	assertCorrectMessage := func (t testing.TB, got, want string)  {
		t.Helper()

		if (got != want) {
			t.Fatalf("got %s, but want %s", got, want)
		}
	}

	t.Run("Output should be equal to the input", func(t *testing.T) {
			got := strings.Clone("wow")
			want := "wow"

			assertCorrectMessage(t, got, want)
	})
}
