package main

import "fmt"


func Hello(name, language string) string {
	if (name == "") {
		name = "World"
	}

	helloMap := map[string]string{
		"German": "Hallo",
		"": "Hello",
	}

	return fmt.Sprintf("%s, %s", helloMap[language], name)
}

func main() {
	fmt.Println(Hello("Golang", "German"))
}