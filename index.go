package main

import (
	"errors"
	"fmt"
	"log"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	result, err := Hello("Robert")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	//messages.Hello("Me")
}

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
