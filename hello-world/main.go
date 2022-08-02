package main

import "fmt"


func greetingPrefix(language string) string {
	helloMap := map[string]string{
		"German": "Hallo",
		"": "Hello",
	}

	return helloMap[language]
}

func Hello(name, language string) string {
	if (name == "") {
		name = "World"
	}

	return fmt.Sprintf("%s, %s", greetingPrefix(language), name)
}

func main() {
	fmt.Println(Hello("Golang", "German"))
}