package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	err := Greet(&buffer, "Chris")

	if err != nil {
		t.Fatal("unexpected error encountered")
	}

	got := buffer.String()
	want := "Hello, Chris"
	assertString(t, got, want)

}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
