package main

import (
	"fmt"
)

const french = "French"
const spanish = "Spanish"
const englishPrefix = " Learns Go"
const spanishPrefix = " Elodie Go"
const frenchPrefix = " Bonjour Go"

//Language prefix
func langprefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	default:
		prefix = englishPrefix
	}
	return
}

//Hello function
func Hello(name, language string) string {
	if name == "" && language == "" {
		name = "You"
	}
	return name + langprefix(language)
}

func main() {
	fmt.Println(Hello("Abdul", "Spanish"))
}
