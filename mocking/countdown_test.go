package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spy := &SpySleeper{}
	Countdown(buffer, spy)

	got := buffer.String()
	want := `3
2
1
Go!`

	assertInt(t, spy.Calls, 3)
	assertString(t, got, want)
}

func assertInt(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v exptected %v", got, want)
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Fatalf("got %v exptected %v", got, want)
	}
}
