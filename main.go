package main

import "strings"

func main() {
	print("Hello, World!")
}

func cleanInput(text string) []string {
	return strings.Split(strings.ToLower(strings.TrimSpace(text)), " ")
}
