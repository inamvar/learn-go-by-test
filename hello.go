package main

import "fmt"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bon jour, "
const spanish = "Spanish"
const french = "French"

func hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return  greetingPrefix(language) +  name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main(){
	fmt.Println(hello("", ""))
}