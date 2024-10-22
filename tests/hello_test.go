package tests

import (
	"internal"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	got := internal.HelloWorld()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
