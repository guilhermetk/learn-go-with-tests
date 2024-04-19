package main

import (
	"errors"
	"io"
)

var FailedToWriteError = errors.New("Failed to write greeting.")

func Greet(w io.Writer, str string) error {
	greeting := "Hello, " + str
	_, err := w.Write([]byte(greeting))
	if err != nil {
		return FailedToWriteError
	}
	return nil
}

// func main() {
// 	err := Greet(os.Stdout, "Elodie")
// 	if err != nil {
// 		os.Exit(1)
// 	}
// }
