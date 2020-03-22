package main

import "testing"

func TestSweep(t *testing.T) {
	got := "hello"
	want := "hello"

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
