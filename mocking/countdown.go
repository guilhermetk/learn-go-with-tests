package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const outputMessage = "Go!"

func Countdown(w io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		time.Sleep(time.Second)
	}
	fmt.Fprint(w, outputMessage)
}

func main() {
	Countdown(os.Stdout)
}
