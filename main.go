package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {

	text = strings.ToLower(text)
	text = strings.TrimSpace(text)

	return strings.Split(text, " ")
}

func main() {
	fmt.Println("Hello, World!")
}
