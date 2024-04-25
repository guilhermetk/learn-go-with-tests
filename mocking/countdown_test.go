package main

import (
	"bytes"
	"slices"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("spy only sleep calls", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy := &SpyCountdownOperations{}
		Countdown(buffer, spy)

		got := buffer.String()
		want := `3
2
1
Go!`

		assertInt(t, len(spy.Calls), 3)
		assertString(t, got, want)

	})

	t.Run("spy sleep and write calls", func(t *testing.T) {
		spyWriteAndSleep := &SpyCountdownOperations{}
		Countdown(spyWriteAndSleep, spyWriteAndSleep)

		want := []string{write, sleep, write, sleep, write, sleep, write}

		if !slices.Equal(want, spyWriteAndSleep.Calls) {
			t.Fatalf("got %v exptected %v", spyWriteAndSleep.Calls, want)
		}
	})

}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
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
