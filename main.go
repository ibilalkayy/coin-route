package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/safepay/safepay-test/backend"
	"github.com/safepay/safepay-test/structs"
)

func ExchangeRouting(w http.ResponseWriter, r *http.Request) {
	btcAmount := backend.TakeBTCAmount(w, r)
	usdAmount := backend.GiveUSDAmount(w, btcAmount)
	response := structs.Response{
		BTCAmount: btcAmount,
		USDAmount: usdAmount,
		Exchange:  "coinbase",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Execute() error {
	http.HandleFunc("/exchange-routing", ExchangeRouting)
	fmt.Println("Listening the server at port: 4000")
	return http.ListenAndServe(":4000", nil)
}

func main() {
	err := Execute()
	if err != nil {
		log.Fatal(err)
	}
}
