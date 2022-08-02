package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()

		if (got != want) {
			t.Errorf("got %q but want %q", got, want)
		}
	}

	t.Run("Say hello", func(t *testing.T) {
		got := Hello("Golang", "")
		want := "Hello, Golang"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'Hello, World' when empty name string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in German", func(t *testing.T) {
		got := Hello("William", "German")
		want := "Hallo, William"

		assertCorrectMessage(t, got, want)
	})
}