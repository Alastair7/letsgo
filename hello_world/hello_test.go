package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Alex", "");
		want := "Hello, Alex"
	
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T){
		got := Hello("Alex", "Spanish")
		want := "Hola, Alex"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Alex", "French")
		want := "Bonjour, Alex"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Portuguese", func(t *testing.T) {
		got := Hello("Alex", "Portuguese")
		want := "Oi, Alex"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	} }