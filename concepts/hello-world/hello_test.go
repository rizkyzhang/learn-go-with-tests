package main

import "testing"

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("John", "")
		want := "Hello, John"

		assertCorrectMessage(t, got, want)
	})

	t.Run("default to 'Hello, World' when name is empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in Indonesian", func(t *testing.T) {
		got := Hello("Wendy", "Indonesian")
		want := "Halo, Wendy"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in French", func(t *testing.T) {
		got := Hello("Wendy", "French")
		want := "Bonjour, Wendy"

		assertCorrectMessage(t, got, want)
	})

	t.Run("default to English when language is empty string", func(t *testing.T) {
		got := Hello("Jack", "")
		want := "Hello, Jack"

		assertCorrectMessage(t, got, want)
	})

	t.Run("default to English when language is not found", func(t *testing.T) {
		got := Hello("Jack", "test")
		want := "Hello, Jack"

		assertCorrectMessage(t, got, want)
	})
}
