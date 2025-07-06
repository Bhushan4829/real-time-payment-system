package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
	"fmt"
)
import "io"
import (
	"github.com/Bhushan4829/real-time-payment-system/models"
)

type PaymentRequest struct {
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Amount   float64 `json:"amount"`
}

type PaymentResponse struct {
	Status    string  `json:"status"`
	Message   string  `json:"message"`
	Timestamp string  `json:"timestamp"`
	Amount    float64 `json:"amount"`
}
var Transactions []models.Transaction
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Transactions)
}

func SimulatePayment(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	fmt.Println("Received Body:", string(body))

	var req PaymentRequest
	if err := json.Unmarshal(body, &req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// After generating response
	


	// Simulate random success or failure
	status := "failed"
	if rand.Float32() > 0.3 { // 70% chance of success
		status = "success"
	}

	res := PaymentResponse{
		Status:    status,
		Message:   "Payment " + status,
		Timestamp: time.Now().Format(time.RFC3339),
		Amount:    req.Amount,
	}
	transaction := models.Transaction{
		Sender:    req.Sender,
		Receiver:  req.Receiver,
		Amount:    req.Amount,
		Status:    status,
		Timestamp: res.Timestamp,
	}
	Transactions = append(Transactions, transaction)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
