package main

import (
	"testing"
)

//Testing Hello
func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Abdul Learns Go", func(t *testing.T) {
		got := Hello("Abdul", "")
		want := "Abdul Learns Go"
		assertCorrectMessage(t, got, want)

	})

	t.Run("You Learns Go", func(t *testing.T) {
		got := Hello("", "")
		want := "You Learns Go"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Go in Spanish", func(t *testing.T) {
		got := Hello("Abdul", "Spanish")
		want := "Abdul Elodie Go"
		assertCorrectMessage(t, got, want)

	})

	t.Run("Go in French", func(t *testing.T) {
		got := Hello("Abdul", "French")
		want := "Abdul Bonjour Go"
		assertCorrectMessage(t, got, want)

	})

}
