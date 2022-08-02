package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5

	if (got != want) {
		t.Errorf("got %d, but want %d", got, want)
	}
}

func ExampleAdd() {
	sum := Add(3, 5)
	fmt.Println(sum)
	// Output: 8
}