package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Victor")
	want := "Hello, Victor"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
