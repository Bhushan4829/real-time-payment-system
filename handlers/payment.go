package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
	"fmt"
)
import "io"

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
