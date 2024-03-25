package main

import "fmt"

const (
	spanish                 = "Spanish"
	french                  = "French"
	englishHelloPrefix      = "Hello, "
	spanishHelloPrefix      = "Hola, "
	frenchHelloPrefix       = "Bonjour, "
	englishHelloDefaultName = "World"
)

func Hello(name, language string) string {
	if name == "" {
		name = englishHelloDefaultName
	}

	return fmt.Sprintf("%s%s", getPrefix(language), name)
}

func getPrefix(language string) string {
	switch language {
	case spanish:
		return spanishHelloPrefix
	case french:
		return frenchHelloPrefix
	}
	return englishHelloPrefix
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
