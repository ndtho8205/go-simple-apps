package main

import "fmt"

const (
	vietnamese            = "vi"
	vietnameseHelloPrefix = "Xin chao"

	french            = "fr"
	frenchHelloPrefix = "Bonjour,"

	englishHelloPrefix = "Hello,"
)

func main() {
	fmt.Println(Hello("World", ""))
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s %s!", helloPrefix(language), name)
}

func helloPrefix(language string) (prefix string) {
	switch language {
	case vietnamese:
		prefix = vietnameseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}
