package main

import "fmt"

const englishHelloPrefix = "Hello, "
const englishHelloSuffix = "!"
const englishHelloDefaultName = "World"

func Hello(name string) string {
	if name == "" {
		name = englishHelloDefaultName
	}
	return fmt.Sprintf("%s%s%s", englishHelloPrefix, name, englishHelloSuffix)
}

func main() {
	fmt.Println(Hello("Chris"))
}
