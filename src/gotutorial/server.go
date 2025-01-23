package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"io"
	"log"
	"net/http"
)

func main() {
	//Figure out how to hide keys before adding to git
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
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		Country   string `json:"country"`
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculateOrderAmount(req.ProductId)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}

	paymentIntent, err := paymentintent.New(params)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Println(paymentIntent.ClientSecret)

	var response struct {
		ClientSecret string `json:"clientSecret"`
	}

	response.ClientSecret = paymentIntent.ClientSecret

	var buf bytes.Buffer
	err = json.NewEncoder(&buf).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = io.Copy(w, &buf)

	if err != nil {
		fmt.Println(err)
	}
}

func handleHealth(w http.ResponseWriter, r *http.Request) {

	response := []byte("Server is up and running")
	_, err := w.Write(response)

	if err != nil {
		fmt.Println(err)
	}
}

func calculateOrderAmount(productId string) int64 {
	switch productId {
	case "Forever Pants":
		return 26000
	case "Forever Shirt":
		return 15500
	case "Forever Shorts":
		return 15500
	}
	return 0
}
