package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Golang")
	want := "Hello, Golang"

	if (got != want) {
		t.Errorf("got %q but want %q", got, want)
	}
}