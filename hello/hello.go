package main

import "fmt"

const englishHelloPrefix = "Hello, "
const swedishHelloPrefix = "Hej, "
const frenchHelloPrefix = "Bonjour, "

// Hello returns "Hello, world!"
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "!"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Swedish":
		prefix = swedishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
