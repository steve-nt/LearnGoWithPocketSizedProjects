//Listing 2.6 main.go: Adding a new language

package main

import "fmt"

func main() {
	greeting := greet("fr")
	fmt.Println(greeting)
}

// language represents the language type language string #1
type language string

// greet says hello to the world in the specified language
func greet(l language) string {
	switch l {
	case "en":
		return "Hello world"
	case "fr":
		return "Bonjour le monde"
	default:
		return ""
	}
}
