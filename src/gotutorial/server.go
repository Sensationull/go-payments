package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)

	log.Println("Listening on localhost:4242...")
	var err error = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle Create Payment Intent Request Called")
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle Health Request Called")
}
