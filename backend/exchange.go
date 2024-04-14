package backend

import (
	"encoding/json"
	"net/http"

	"github.com/ibilalkayy/coin-route/structs"
)

func ExchangeRouting(w http.ResponseWriter, r *http.Request) {
	btcAmount := TakeBTCAmount(w, r)
	usdAmount := GiveUSDAmount(w, btcAmount)
	response := structs.Response{
		BTCAmount: btcAmount,
		USDAmount: usdAmount,
		Exchange:  "coinbase",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
