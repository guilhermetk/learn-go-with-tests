package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const outputMessage = "Go!"
const sleep = "sleep"
const write = "write"

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

type SpyCountdownOperations struct {
	Calls []string
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(time.Second * 1)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, outputMessage)
}

func main() {
	sleeper := &ConfigurableSleeper{5 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
