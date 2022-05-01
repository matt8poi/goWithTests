package main

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, world"
	if got := Hello(); got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
