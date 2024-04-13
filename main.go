package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	BTCAmount float64 `json:"btcAmount"`
	USDAmount float64 `json:"usdAmount"`
	Exchange  string  `json:"exchange"`
}

func TakeAmount(w http.ResponseWriter, r *http.Request) float64 {
	amountStr := r.URL.Query().Get("amount")
	if amountStr == "" {
		http.Error(w, "Amount query parameter is missing", http.StatusBadRequest)
		return 0
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "Invalid amount format", http.StatusBadRequest)
		return 0
	}
	return amount
}

func ExchangeRouting(w http.ResponseWriter, r *http.Request) {
	amount := TakeAmount(w, r)

	response := Response{
		BTCAmount: amount,
		USDAmount: 45384.54,
		Exchange:  "coinbase",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Routes() {
	http.HandleFunc("/exchange-routing", ExchangeRouting)
}

func Execute() error {
	Routes()
	fmt.Println("Listening the server at port: 4000")
	return http.ListenAndServe(":4000", nil)
}

func main() {
	err := Execute()
	if err != nil {
		log.Fatal(err)
	}
}
