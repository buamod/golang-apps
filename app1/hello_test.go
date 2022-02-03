package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("Ibrahim")
	want := "Hello, Ibrahim"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestFailure(t *testing.T) {
	got := hello("Ibrahim")
	want := ""

	if got == want {
		t.Errorf("got %q want %q", got, want)
	}
	t.Logf("got %q want %q", got, want)
}
