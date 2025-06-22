package main

import (
	"flag" // standard package to read from command-line arguments. (https://pkg.go.dev/flag)
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang,
		"lang",
		"en",
		"The required language, e.g en, ur, fr and more")
	flag.Parse()

	greeting := greet(language(lang))
	fmt.Println(greeting)
}

// language represents the language's code
type language string

// phrasebook holds greeting for each supported language
var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",     // Greek
	"en": "Hello world",       // English
	"fr": "Bonjour le monde",  // French
	"he": "שלום עולם",         // Hebrew
	"ur": "ہیلو دنیا",         // Urdu
	"vi": "Xin chào Thế Giới", // Vietnamese
}

// greet says hello to the world in various languages
func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting
}
