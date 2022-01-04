package main

import "fmt"

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

func main() {
	fmt.Println(Hello("World", ""))
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefixFor(lang) + name + "!"
}

func helloPrefixFor(lang string) (prefix string) {
	switch lang {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}
