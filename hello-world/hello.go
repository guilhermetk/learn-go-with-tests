package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const englishHelloDefaultName = "World"

func Hello(name, language string) string {
	if name == "" {
		name = englishHelloDefaultName
	}
	if language == "French" {
		return fmt.Sprintf("%s%s", frenchHelloPrefix, name)
	}
	if language == "Spanish" {
		return fmt.Sprintf("%s%s", spanishHelloPrefix, name)
	}
	return fmt.Sprintf("%s%s", englishHelloPrefix, name)
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
