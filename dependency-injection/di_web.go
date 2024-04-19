package main

import (
	"io"
	"log"
	"net/http"
)

func GreetWeb(w io.Writer, str string) {
	greeting := "Hello, " + str
	w.Write([]byte(greeting))
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	GreetWeb(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
