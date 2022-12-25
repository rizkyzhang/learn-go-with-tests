package integers

import (
	"fmt"
	"testing"
)

func assertCorrectMessage(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestAdder(t *testing.T) {
	got := Add(2, 5)
	want := 7

	assertCorrectMessage(t, got, want)
}

func ExampleAdd() {
	sum := Add(3, 9)
	fmt.Println(sum)
	// Output: 12
}
