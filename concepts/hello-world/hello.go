package main

import (
	"fmt"
	"strings"
)

func getHelloPrefix(language string) string {
	helloPrefixes := map[string]string{
		"english":    "Hello, ",
		"french":     "Bonjour, ",
		"indonesian": "Halo, ",
	}

	helloPrefix, ok := helloPrefixes[strings.ToLower(language)]
	if language == "" || !ok {
		return "Hello, "
	}

	return helloPrefix
}

func Hello(name string, language string) string {
	if name == "" {
		return "Hello, World"
	}

	return fmt.Sprintf("%s%s", getHelloPrefix(language), name)
}

func main() {
	fmt.Println(Hello("John", "English"))
}
