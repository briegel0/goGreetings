package greetings

import "testing"

func TestGreetings(t *testing.T) {
	want := "Hello, world."
	got, error := Hello(want)

	if error != nil {
		t.Error("error detected")
	}

	if got != "Hi, Hello, world.. Welcome!" {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
