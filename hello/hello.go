package main

import "fmt"

// Hello returns "Hello, world!"
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	fmt.Println(Hello("world"))
}
