package tests

import (
	"internal"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	got := internal.HelloWorld()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
