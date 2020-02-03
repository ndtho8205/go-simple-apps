package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("say hello to Tho", func(t *testing.T) {
		got := Hello("Tho", "")
		want := "Hello, Tho!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to the world when an empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to Tho in Vietnamese", func(t *testing.T) {
		got := Hello("Tho", "vi")
		want := "Xin chao Tho!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to Tho in French", func(t *testing.T) {
		got := Hello("Tho", "fr")
		want := "Bonjour, Tho!"
		assertCorrectMessage(t, got, want)
	})
}
