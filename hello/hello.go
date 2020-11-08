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

	if language == "Swedish" {
		return swedishHelloPrefix + name + "!"
	}

	if language == "French" {
		return frenchHelloPrefix + name + "!"
	}

	return englishHelloPrefix + name + "!"
}

func main() {
	fmt.Println(Hello("world", ""))
}
