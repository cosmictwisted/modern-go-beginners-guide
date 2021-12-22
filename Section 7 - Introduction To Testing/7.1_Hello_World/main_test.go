package main

import "testing"

func TestGreet(t *testing.T) {
	want := "Hello, World!"
	got := greet()
	if want != got {
		t.Error("No Hello Today!")
	}
}

func TestNoGreet(t *testing.T) {
	want := "Hello, World!"
	got := noGreet()
	if want != got {
		t.Error("No Hello Today!")
	}
}
