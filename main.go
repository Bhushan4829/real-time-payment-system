package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Bhushan4829/real-time-payment-system/handlers"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server is up!")
	})

	http.HandleFunc("/simulate-payment", handlers.SimulatePayment)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
