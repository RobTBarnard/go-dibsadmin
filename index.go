package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"rsc.io/quote"
)

func main() {
	router := httprouter.New()
	router.GET("/", landingPage)
	router.GET("/GetProductLines", getProductLines)
	log.Fatal(http.ListenAndServe(":8080", router))

	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	result, err := Hello("Robert")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	//messages.Hello("Me")

	//handleRequests()
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

func landingPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Fprint(w, "Welcome to the landing page!")
	fmt.Println("Request received!")
}

type productLine struct {
	productline string `json:"productline"`
	description string `json:"description"`
}

func getProductLines(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	productLines := &productLine{
		productline: "4001",
		description: "a product line",
	}

	result, err := json.Marshal(productLines)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, string(result))
}

// func handleRequests() {
// 	http.HandleFunc("/", landingPage)
// 	http.HandleFunc("/GetProductLines", getProductLines)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
